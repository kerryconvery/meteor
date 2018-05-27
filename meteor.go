package main

import "meteor/profiles"
import "meteor/configuration"

var config = configuration.GetConfiguration("meteor.json")

func main() {
	profiles.New(profiles.ProfileFileSystem{ProfilePath: config.ProfilePath})
}
