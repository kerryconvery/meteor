package controllers

import "meteor/types"

// ProfilesController is a controller for the /profiles route
type ProfilesController struct {
	Controller
	Profiles types.Profiles
}

// NewProfilesController returns a new instance of ProfilesController
func NewProfilesController(profiles types.Profiles) ProfilesController {
	return ProfilesController{Profiles: profiles}
}

// GetAll returns a 200 JSON response containing all profiles
func (c ProfilesController) GetAll() HTTPResponse {
	profiles, err := c.Profiles.GetProfiles()

	if err != nil {
		return c.ErrorResponse(500, err)
	}
	return c.JSONResponse(200, profiles)
}
