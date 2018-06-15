package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"meteor/configuration"
	"meteor/controllers"
	"meteor/filesystem"
	"meteor/media"
	"meteor/mediaplayers"
	"meteor/profiles"
	"meteor/thumbnails"
	"net/http"

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

func main() {
	filesystem := filesystem.New()
	config, _ := configuration.GetConfiguration("./", "meteor.json", filesystem)
	profileProvider := profiles.New(config.ProfilePath)
	thumbnailProvider := thumbnails.New(config.ThumbnailPath, config.AssetPath, filesystem)
	profileController := controllers.NewProfilesController(profileProvider, media.New(filesystem), thumbnailProvider)
	mediaController := controllers.NewMediaController(
		profileProvider,
		mediaplayers.New(
			config.MediaPlayers[0].LaunchCmd,
			config.MediaPlayers[0].LaunchArgs,
			config.MediaPlayers[0].APIUrl,
		),
	)

	router := httprouter.New()

	router.GET("/", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		//Parse the html file
		index := template.Must(template.ParseFiles("webclient/dist/index.html"))

		index.Execute(rw, nil)
	})

	router.GET("/media", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		//Parse the html file
		index := template.Must(template.ParseFiles("webclient/dist/index.html"))

		index.Execute(rw, nil)
	})

	router.ServeFiles("/webclient/*filepath", http.Dir("webclient"))

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

	log.Fatal(http.ListenAndServe(":8080", router))
}
