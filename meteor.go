package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"meteor/configuration"
	"meteor/controllers"
	"meteor/db"
	"meteor/filesystem"
	"meteor/media"
	"meteor/mediaWatchers"
	"meteor/mediaplayers"
	"meteor/profiles"
	"meteor/thumbnails"
	"meteor/webhook"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func handleJSONResponse(writer http.ResponseWriter, response controllers.JSONResponse) {
	writer.WriteHeader(response.StatusCode)
	writer.Header().Set("Content-Type", response.ContentType)
	json.NewEncoder(writer).Encode(response.Body)
}

func handleBinaryResponse(writer http.ResponseWriter, response controllers.BinaryResponse) {
	writer.WriteHeader(response.StatusCode)
	writer.Header().Set("Content-Type", response.ContentType)
	writer.Write(response.Body.Bytes())
}

func handleTextResponse(writer http.ResponseWriter, response controllers.TextResponse) {
	writer.WriteHeader(response.StatusCode)
	writer.Header().Set("Content-Type", response.ContentType)
	fmt.Fprint(writer, response.Body)
}

func handleError(writer http.ResponseWriter, err error) {
	writer.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(writer, err.Error())
}

func serveUI(webclientPath string) func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Parse the html file
	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		index := template.Must(template.ParseFiles(filepath.Join(webclientPath, "index.html")))
		index.Execute(rw, nil)
	}
}

func main() {
	filesystem := filesystem.New()
	config, err := configuration.GetConfiguration("./", "meteor.json", filesystem)

	if err != nil {
		panic(err)
	}

	profileProvider := profiles.New(config.ProfilePath)
	thumbnailProvider := thumbnails.New(config.ThumbnailPath, config.AssetPath, filesystem)
	profileController := controllers.NewProfilesController(profileProvider, media.New(filesystem), thumbnailProvider)
	webhook := webhook.New()
	store := db.NewMediaStore(config.DatastorePath)

	defer webhook.Stop()
	defer store.Close()

	mediaController := controllers.NewMediaController(
		profileProvider,
		mediaplayers.New(
			config.MediaPlayers[0].LaunchCmd,
			config.MediaPlayers[0].LaunchArgs,
			config.MediaPlayers[0].APIUrl,
		),
		store,
	)

	webhook.Start()

	terminate := make(chan bool, 2)
	terminate <- false

	defer func() { terminate <- true }()
	defer webhook.Stop()

	go mediaController.WatchMediaPlayer(
		terminate,
		mediaWatchers.NewStateNotifier(&webhook),
		mediaWatchers.NewMediaStateRecorder(store),
	)

	router := httprouter.New()

	router.GET("/", serveUI(config.WebClientPath))
	router.GET("/media", serveUI(config.WebClientPath))
	router.ServeFiles("/web/*filepath", http.Dir(config.WebClientPath))

	router.GET("/ws", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		webhook.AddClient(w, r)
	})

	router.GET("/api/profiles", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		response, err := profileController.GetAll()

		if err != nil {
			handleError(w, err)
		} else {
			handleJSONResponse(w, response)
		}
	})

	router.GET("/api/profiles/:profilename/media", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		path := r.URL.Query().Get("uri")

		response, err := profileController.GetMedia(ps.ByName("profilename"), path)

		if err != nil {
			handleError(w, err)
		} else {
			handleJSONResponse(w, response)
		}
	})

	router.GET("/api/profiles/:profilename/media/thumbnail", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		response, err := profileController.GetMediaThumbnail(
			ps.ByName("profilename"),
			r.URL.Query().Get("uri"),
		)

		if err != nil {
			handleError(w, err)
		} else {
			handleBinaryResponse(w, response)
		}
	})

	router.POST("/api/media", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		profileName := r.URL.Query().Get("profile")
		URI := r.URL.Query().Get("uri")

		response, err := mediaController.LaunchMediaFile(profileName, URI)

		if err != nil {
			handleError(w, err)
		} else {
			handleTextResponse(w, response)
		}
	})

	router.DELETE("/api/media", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		err := mediaController.CloseMediaPlayer()
		if err != nil {
			handleError(w, err)
		}
	})

	router.POST("/api/media/pause", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		err := mediaController.PauseMediaPlayer()
		if err != nil {
			handleError(w, err)
		}
	})

	router.POST("/api/media/resume", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		err := mediaController.ResumeMediaPlayer()
		if err != nil {
			handleError(w, err)
		}
	})

	router.POST("/api/media/restart", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		err := mediaController.RestartMediaPlayer()
		if err != nil {
			handleError(w, err)
		}
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ListenPort), router))
}
