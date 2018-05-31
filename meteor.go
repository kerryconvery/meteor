package main

import (
	"meteor/configuration"
	"meteor/controllers"
	"meteor/profiles"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

// mapJSONResponseToIris call the appropriate iris methods for each response field
func mapJSONResponseToIris(ctx context.Context, response controllers.HTTPResponse) {
	ctx.StatusCode(response.StatusCode)
	if response.Error != nil {
		ctx.JSON(response.Error)
	} else {
		ctx.ContentType(response.ContentType)
		ctx.JSON(response.Body)
	}
}

func main() {
	config := configuration.GetConfiguration("./", "meteor.json")
	profiles := profiles.New(config.ProfilePath)
	profilesController := controllers.NewProfilesController(profiles)

	app := iris.New()

	app.Get("/", func(ctx context.Context) {
		ctx.Text("Hello World")
	})

	app.Get("/profiles", func(ctx context.Context) {
		mapJSONResponseToIris(ctx, profilesController.GetAll())
	})

	app.Run(iris.Addr(":8080"))
}
