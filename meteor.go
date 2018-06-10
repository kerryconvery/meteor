package main

import (
	"meteor/configuration"
	"meteor/controllers"
	"meteor/filesystem"
	"meteor/media"
	"meteor/mediaplayers"
	"meteor/profiles"
	"meteor/thumbnails"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func handleJSONResponse(ctx context.Context, response controllers.JSONResponse) {
	ctx.StatusCode(response.StatusCode)
	ctx.ContentType(response.ContentType)
	ctx.JSON(response.Body)
}

func handleBinaryResponse(ctx context.Context, response controllers.BinaryResponse) {
	ctx.StatusCode(response.StatusCode)
	ctx.ContentType(response.ContentType)
	ctx.Binary(response.Body.Bytes())
}

func handleTextResponse(ctx context.Context, response controllers.TextResponse) {
	ctx.StatusCode(response.StatusCode)
	ctx.ContentType(response.ContentType)
	ctx.Text(response.Body)
}

func handleError(ctx context.Context, err error) {
	ctx.StatusCode(500)
	ctx.WriteString(err.Error())
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

	app := iris.New()

	app.Get("/", func(ctx context.Context) {
		ctx.Text("Hello World")
	})

	app.Get("api/profiles", func(ctx context.Context) {
		response, err := profileController.GetAll()

		if err != nil {
			handleError(ctx, err)
		} else {
			handleJSONResponse(ctx, response)
		}
	})

	app.Get("api/profiles/{profilename}/media", func(ctx context.Context) {
		path := ""
		if ctx.URLParamExists("path") {
			path = ctx.URLParamTrim("path")
		}
		response, err := profileController.GetMedia(ctx.Params().Get("profilename"), path)

		if err != nil {
			handleError(ctx, err)
		} else {
			handleJSONResponse(ctx, response)
		}
	})

	app.Get("api/profiles/{profilename}/media/{media}/thumbnail", func(ctx context.Context) {
		path := ""
		if ctx.URLParamExists("path") {
			path = ctx.URLParamTrim("path")
		}

		response, err := profileController.GetMediaThumbnail(
			ctx.Params().Get("profilename"),
			path,
			ctx.Params().Get("media"),
		)

		if err != nil {
			handleError(ctx, err)
		} else {
			handleBinaryResponse(ctx, response)
		}
	})

	app.Post("api/profiles/{profilename}/media/{media}", func(ctx context.Context) {
		response, err := mediaController.LaunchMediaFile(ctx.Params().Get("profilename"), ctx.Params().Get("media"))
		if err != nil {
			handleError(ctx, err)
		} else {
			handleTextResponse(ctx, response)
		}
	})

	app.Delete("api/profiles/{profilename}/media", func(ctx context.Context) {
		err := mediaController.CloseMediaPlayer()
		if err != nil {
			handleError(ctx, err)
		}
	})

	app.Run(iris.Addr(":8080"))
}
