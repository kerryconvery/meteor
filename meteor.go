package main

import (
	"meteor/configuration"
	"meteor/controllers"
	"meteor/filesystem"
	"meteor/media"
	"meteor/profiles"
	"meteor/thumbnails"
	"path/filepath"

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

func handleError(ctx context.Context, err error) {
	ctx.StatusCode(500)
	ctx.WriteString(err.Error())
}

func checkProfile(profilePath string) func(ctx context.Context) {
	return func(ctx context.Context) {
		profileProvider := profiles.New(profilePath)
		profileName := ctx.Params().Get("profilename")
		exists, err := profileProvider.ProfileExists(profileName)

		if err != nil {
			handleError(ctx, err)
		} else if exists != true {
			ctx.StatusCode(404)
			ctx.WriteString("profile " + profileName + " not found")
		} else {
			ctx.Next()
		}
	}
}

func checkPath(profilePath string, filesystem filesystem.Filesystem) func(ctx context.Context) {
	return func(ctx context.Context) {
		if ctx.URLParamExists("path") {
			path := ctx.URLParam("path")
			profileProvider := profiles.New(profilePath)
			mediaProvider := media.New("", filesystem)
			profileName := ctx.Params().Get("profilename")
			profile, _ := profileProvider.GetProfile(profileName)

			exists, err := mediaProvider.PathExists(filepath.Join(profile.MediaPath, path))

			if err != nil {
				handleError(ctx, err)
			} else if exists != true {
				ctx.StatusCode(404)
				ctx.WriteString("path " + path + " not found")
			} else {
				ctx.Next()
			}
		} else {
			ctx.Next()
		}
	}
}

func main() {
	filesystem := filesystem.New()
	config := configuration.GetConfiguration("./", "meteor.json", filesystem)
	profileProvider := profiles.New(config.ProfilePath)

	app := iris.New()

	app.Get("/", func(ctx context.Context) {
		ctx.Text("Hello World")
	})

	app.Get("api/profiles", func(ctx context.Context) {
		thumbnailProvider := thumbnails.New(config.ThumbnailPath, config.AssetPath, filesystem)

		controller := controllers.NewProfilesController(profileProvider, media.New(ctx.Path(), filesystem), thumbnailProvider)

		response, err := controller.GetAll()

		if err != nil {
			handleError(ctx, err)
		} else {
			handleJSONResponse(ctx, response)
		}
	})

	app.Get("api/profiles/{profilename}/media", func(ctx context.Context) {
		thumbnailProvider := thumbnails.New(config.ThumbnailPath, config.AssetPath, filesystem)

		controller := controllers.NewProfilesController(profileProvider, media.New(ctx.Path(), filesystem), thumbnailProvider)

		path := ""
		if ctx.URLParamExists("path") {
			path = ctx.URLParamTrim("path")
		}
		response, err := controller.GetMedia(ctx.Params().Get("profilename"), path)

		if err != nil {
			handleError(ctx, err)
		} else {
			handleJSONResponse(ctx, response)
		}
	})

	app.Get("api/profiles/{profilename}/media/{media}/thumbnail", func(ctx context.Context) {
		thumbnailProvider := thumbnails.New(config.ThumbnailPath, config.AssetPath, filesystem)

		controller := controllers.NewProfilesController(profileProvider, media.New(ctx.Path(), filesystem), thumbnailProvider)

		path := ""
		if ctx.URLParamExists("path") {
			path = ctx.URLParamTrim("path")
		}
		params := ctx.Params()
		response, err := controller.GetMediaThumbnail(params.Get("profilename"), path, params.Get("media"))

		if err != nil {
			handleError(ctx, err)
		} else {
			handleBinaryResponse(ctx, response)
		}
	})

	app.Run(iris.Addr(":8080"))
}
