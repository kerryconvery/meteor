package controllers

import (
	"errors"
	"meteor/filesystem"
	"meteor/media"
	"meteor/profiles"
	"meteor/tests"
	"testing"
)

type sampleFiles struct {
	filesystem.Filesystem
}

func (f sampleFiles) GetFiles(path string) ([]filesystem.File, error) {
	switch path {
	case "MediaA":
		return []filesystem.File{
				filesystem.File{"movie.avi", false},
				filesystem.File{"music.mp3", false}},
			nil
	case "MediaB":
		return []filesystem.File{
				filesystem.File{"movie.avi", false},
				filesystem.File{"music.mp3", false},
				filesystem.File{"more movies", true}},
			nil
	case "MediaB\\more movies":
		return []filesystem.File{
				filesystem.File{"movie1.avi", false},
				filesystem.File{"movie2.avi", false}},
			nil
	}
	return []filesystem.File{}, errors.New("unknown path")
}

func (f sampleFiles) FileExists(path, filename string) (bool, error) {
	return true, nil
}

type sampleProfiles struct {
	profiles.Provider
	profiles []profiles.Profile
	err      error
}

func (p sampleProfiles) GetProfiles() ([]profiles.Profile, error) {
	return p.profiles, p.err
}

func (p sampleProfiles) GetProfile(profileName string) (profiles.Profile, error) {
	for _, profile := range p.profiles {
		if profile.Name == profileName {
			return profile, nil
		}
	}

	return profiles.Profile{}, errors.New("Profile not found")
}

func NewController(provider profiles.ProfileProvider) ProfilesController {
	return NewProfilesController(provider, media.New("", sampleFiles{}))
}

func TestGetProfiles(t *testing.T) {
	mockProfiles := sampleProfiles{profiles: []profiles.Profile{profiles.Profile{}}, err: nil}

	response, err := NewController(mockProfiles).GetAll()

	tests.ExpectNoError(err, t)

	tests.ExpectContentType("application/json", response.ContentType, t)
	tests.ExpectStatusCode(200, response.StatusCode, t)

	responseBody := response.Body.([]profiles.Profile)

	if len(responseBody) != 1 {
		t.Errorf("Expected 1 profiles but got %d", len(responseBody))
	}
}

func TestGetProfilesError(t *testing.T) {
	profiles := sampleProfiles{profiles: []profiles.Profile{}, err: errors.New("Could not read profiles")}

	_, err := NewProfilesController(profiles, media.New("", sampleFiles{})).GetAll()

	tests.ExpectError(err, t)
}

func TestGetMediaFiles(t *testing.T) {
	profiles := sampleProfiles{profiles: []profiles.Profile{
		profiles.Profile{Name: "ProfileA", MediaPath: "MediaA"},
		profiles.Profile{Name: "ProfileB", MediaPath: "MediaB"},
		profiles.Profile{Name: "ProfileC", MediaPath: "MediaC"}},
		err: nil,
	}

	controller := NewController(profiles)

	response, err := controller.GetMedia("ProfileB", "")

	tests.ExpectNoError(err, t)
	tests.ExpectStatusCode(200, response.StatusCode, t)
	tests.ExpectContentType("application/json", response.ContentType, t)

	mediaFiles := response.Body.([]media.Media)

	if mediaFiles[0].Name != "movie.avi" {
		t.Errorf("Expected movie.avi but got %s", mediaFiles[0].Name)
	}

	if mediaFiles[0].IsDirectory != false {
		t.Error("Expected to be a file but got a directory")
	}

	if mediaFiles[2].IsDirectory != true {
		t.Error("Expected to be a directory but got a file")
	}
}

func TestGetMediaFilesSubPath(t *testing.T) {
	profiles := sampleProfiles{profiles: []profiles.Profile{
		profiles.Profile{Name: "ProfileB", MediaPath: "MediaB"}},
		err: nil,
	}

	controller := NewController(profiles)

	response, err := controller.GetMedia("ProfileB", "more movies")

	tests.ExpectNoError(err, t)

	mediaFiles := response.Body.([]media.Media)

	if len(mediaFiles) != 2 {
		t.Errorf("Expected 2 media files but got %d", len(mediaFiles))
	}

	if mediaFiles[0].Name != "movie1.avi" {
		t.Errorf("Expected movie1.avi but got %s", mediaFiles[0].Name)
	}
}

func TestGetMediaFilesError(t *testing.T) {
	profiles := sampleProfiles{profiles: []profiles.Profile{
		profiles.Profile{Name: "ProfileC", MediaPath: "MediaC"}},
		err: nil,
	}

	controller := NewController(profiles)

	_, err := controller.GetMedia("ProfileC", "")

	tests.ExpectError(err, t)
}
