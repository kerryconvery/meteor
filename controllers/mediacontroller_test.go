package controllers

import (
	"errors"
	"fmt"
	"meteor/profiles"
	"meteor/tests"
	"testing"
)

type profilesProvider struct {
}

func (p profilesProvider) GetProfile(profileName string) (profiles.Profile, error) {
	if profileName == "valid_profile" {
		return profiles.Profile{MediaPath: "valid_path", MediaArgs: []string{"/arg"}}, nil
	}
	return profiles.Profile{}, errors.New("Could not find profile")
}

type mockMediaPlayer struct {
}

func (m mockMediaPlayer) Play(media string, mediaArgs []string) error {
	if media == "valid_path\\valid_media_file" && mediaArgs[0] == "/arg" {
		return nil
	}
	return fmt.Errorf("Invalid media %s or args %s", media, mediaArgs)
}

func (m mockMediaPlayer) Exit() error {
	return nil
}

func (m mockMediaPlayer) Pause() error {
	return nil
}

func (m mockMediaPlayer) Resume() error {
	return nil
}

func GetMediaController() MediaController {
	return NewMediaController(profilesProvider{}, mockMediaPlayer{})
}
func TestLaunchMediaFile(t *testing.T) {
	controller := GetMediaController()

	response, err := controller.LaunchMediaFile("valid_profile", "valid_media_file")

	tests.ExpectNoError(err, t)

	if response.StatusCode != 201 {
		t.Errorf("Expected response code 201 but got %d", response.StatusCode)
	}
}
func TestLaunchMediaFileError(t *testing.T) {
	controller := GetMediaController()

	_, err := controller.LaunchMediaFile("valid_profile", "invalid_media_file")

	tests.ExpectError(err, t)
}
