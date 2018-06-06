package controllers

import (
	"meteor/media"
	"meteor/profiles"
	"meteor/thumbnails"
	"path/filepath"
)

// ProfilesController is a controller for the /controllerProfiles route
type ProfilesController struct {
	Controller
	profileProvider   profiles.Provider
	mediaProvider     media.Provider
	thumbnailProvider thumbnails.Provider
}

// NewProfilesController returns a new instance of ProfilesController
func NewProfilesController(profileProvider profiles.Provider, mediaProvider media.Provider, thumbnailProvider thumbnails.Provider) ProfilesController {
	return ProfilesController{profileProvider: profileProvider, mediaProvider: mediaProvider, thumbnailProvider: thumbnailProvider}
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

// GetMediaThumbnail returns a thumbnail representing the media
func (c ProfilesController) GetMediaThumbnail(profileName, subPath, filename string) (BinaryResponse, error) {
	profile, err := c.profileProvider.GetProfile(profileName)

	if err != nil {
		return BinaryResponse{}, err
	}

	thumbnail, err := c.thumbnailProvider.GetThumbnail(filepath.Join(profile.MediaPath, subPath), filename)

	if err != nil {
		return BinaryResponse{}, err
	}

	return c.BinaryResponse("image/png", thumbnail), nil
}
