package main

import (
	"meteor/configuration"
	"meteor/controllers"
	"meteor/filesystem"
	"meteor/media"
	"meteor/profiles"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

// mapJSONResponseToIris call the appropriate iris methods for each response field
func handleJSONResponse(ctx context.Context, response controllers.JSONResponse) {
	ctx.StatusCode(response.StatusCode)
	ctx.ContentType(response.ContentType)
	ctx.JSON(response.Body)
}

func handleError(ctx context.Context, err error) {
	ctx.StatusCode(500)
	ctx.WriteString(err.Error())
}

func beforeGetProfile(profilePath string) func(ctx context.Context) {
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

func main() {
	filesystem := filesystem.New()
	config := configuration.GetConfiguration("./", "meteor.json", filesystem)
	profileProvider := profiles.New(config.ProfilePath)

	app := iris.New()

	app.Get("/", func(ctx context.Context) {
		ctx.Text("Hello World")
	})

	app.Get("api/profiles", func(ctx context.Context) {
		controller := controllers.NewProfilesController(profileProvider, media.New(ctx.Path(), filesystem))

		response, err := controller.GetAll()

		if err != nil {
			handleError(ctx, err)
		} else {
			handleJSONResponse(ctx, response)
		}
	})

	app.Get("api/profiles/{profilename}/media", beforeGetProfile(config.ProfilePath), func(ctx context.Context) {
		controller := controllers.NewProfilesController(profileProvider, media.New(ctx.Path(), filesystem))

		response, err := controller.GetMedia(ctx.Params().Get("profilename"), "")

		if err != nil {
			handleError(ctx, err)
		} else {
			handleJSONResponse(ctx, response)
		}
	})

	app.Get("api/profiles/{profilename}/media/{path}", beforeGetProfile(config.ProfilePath), func(ctx context.Context) {
		controller := controllers.NewProfilesController(profileProvider, media.New(ctx.Path(), filesystem))

		params := ctx.Params()
		response, err := controller.GetMedia(params.Get("profilename"), params.Get("path"))

		if err != nil {
			handleError(ctx, err)
		} else {
			handleJSONResponse(ctx, response)
		}
	})

	app.Get("api/profiles/{profilename}/media/{media}/thumbnail", func(ctx context.Context) {

	})

	app.Run(iris.Addr(":8080"))
}
