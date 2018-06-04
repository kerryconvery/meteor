package controllers

import (
	"meteor/media"
	"meteor/profiles"
	"path/filepath"
)

// ProfilesController is a controller for the /controllerProfiles route
type ProfilesController struct {
	Controller
	profileProvider profiles.ProfileProvider
	mediaProvider   media.Provider
}

// NewProfilesController returns a new instance of ProfilesController
func NewProfilesController(profileProvider profiles.ProfileProvider, mediaProvider media.Provider) ProfilesController {
	return ProfilesController{profileProvider: profileProvider, mediaProvider: mediaProvider}
}

// GetAll returns a 200 JSON response containing all controllerProfiles
func (c ProfilesController) GetAll() (JSONResponse, error) {
	controllerProfiles, err := c.profileProvider.GetProfiles()

	if err != nil {
		return JSONResponse{}, err
	}
	return c.JSONResponse(200, controllerProfiles), nil
}

// GetMedia returns response containing the files at the profile path
func (c ProfilesController) GetMedia(profileName, subPath string) (JSONResponse, error) {
	profile, err := c.profileProvider.GetProfile(profileName)

	if err != nil {
		return JSONResponse{}, err
	}

	files, err := c.mediaProvider.GetLocalMedia(filepath.Join(profile.MediaPath, subPath))

	if err != nil {
		return JSONResponse{}, err
	}

	return c.JSONResponse(200, files), nil
}
