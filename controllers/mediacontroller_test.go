package controllers

import (
	"errors"
	"fmt"
	"meteor/db"
	"meteor/mediaplayers"
	"meteor/profiles"
	"meteor/tests"
	"testing"

	"github.com/stretchr/testify/mock"
)

type profilesProvider struct {
}

type mockStore struct {
}

type mockMediaPlayerListener struct {
	mock.Mock
	signalChan chan bool
}

func (m mockStore) Read(key string) (db.MediaRecord, error) {
	return db.MediaRecord{
		Position: 0,
		Duration: 0,
	}, nil
}

func (p profilesProvider) GetProfile(profileName string) (profiles.Profile, error) {
	if profileName == "valid_profile" {
		return profiles.Profile{MediaPath: "valid_path", MediaArgs: []string{"/arg"}}, nil
	}
	return profiles.Profile{}, errors.New("Could not find profile")
}

type mockMediaPlayer struct {
	mock.Mock
}

func (m mockMediaPlayer) Play(media string, mediaArgs []string, options mediaplayers.MediaOptions) error {
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

func (m mockMediaPlayer) Restart() error {
	return nil
}

func (m *mockMediaPlayer) GetStatus() (mediaplayers.MediaPlayerStatus, error) {
	args := m.Called()
	return args.Get(0).(mediaplayers.MediaPlayerStatus), args.Error(1)
}

func (m *mockMediaPlayerListener) Notify(status mediaplayers.MediaPlayerStatus) {
	m.Called(status)
	m.signalChan <- true
}

func getMediaController() MediaController {
	mediaPlayer := mockMediaPlayer{}
	mediaPlayer.On("GetStatus").Return(mediaplayers.MediaPlayerStatus{}, nil)

	return NewMediaController(profilesProvider{}, &mediaPlayer, mockStore{})
}

func TestLaunchMediaFile(t *testing.T) {
	controller := getMediaController()

	response, err := controller.LaunchMediaFile("valid_profile", "valid_media_file")

	tests.ExpectNoError(err, t)

	if response.StatusCode != 201 {
		t.Errorf("Expected response code 201 but got %d", response.StatusCode)
	}
}
func TestLaunchMediaFileError(t *testing.T) {
	controller := getMediaController()

	_, err := controller.LaunchMediaFile("valid_profile", "invalid_media_file")

	tests.ExpectError(err, t)
}

func TestWatchMediaPlayer(t *testing.T) {
	signalChan := make(chan bool, 2)
	signalChan <- false

	defer func() { signalChan <- true }()

	status := mediaplayers.MediaPlayerStatus{}
	mediaPlayer := mockMediaPlayer{}

	controller := NewMediaController(profilesProvider{}, &mediaPlayer, mockStore{})

	listener := &mockMediaPlayerListener{}
	listener.signalChan = make(chan bool, 1)

	mediaPlayer.On("GetStatus").Return(status, nil)
	listener.On("Notify", status)

	go controller.WatchMediaPlayer(signalChan, listener)

	<-listener.signalChan

	listener.AssertExpectations(t)
}
