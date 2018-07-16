package controllers

import (
	"meteor/mediaplayers"
	"meteor/profiles"
	"path/filepath"
)

type profileSource interface {
	GetProfile(profileName string) (profiles.Profile, error)
}
type broadcaster interface {
	Broadcast(payload interface{})
}

type datastore interface {
	UpdatePosition(key string, position int) error
	Delete(key string)
}

// MediaController represents a controller of a media player
type MediaController struct {
	Controller
	store            datastore
	broadcaster      broadcaster
	profilesProvider profileSource
	mediaPlayer      mediaplayers.MediaPlayer
}

// NewMediaController returns a new instance of the media controller
func NewMediaController(profilesProvider profileSource, mediaPlayer mediaplayers.MediaPlayer, broadcaster broadcaster, store datastore) MediaController {
	controller := MediaController{
		store:            store,
		profilesProvider: profilesProvider,
		mediaPlayer:      mediaPlayer,
		broadcaster:      broadcaster,
	}

	go controller.watchMediaPlayer()
	return controller
}

func (c MediaController) updateStore(info mediaplayers.MediaPlayerInfo) {
	if info.State == 2 {
		c.store.UpdatePosition(info.NowPlaying, info.Position)
	}

	if info.Position == info.Duration {
		c.store.Delete(info.NowPlaying)
	}
}

func (c MediaController) watchMediaPlayer() {
	for {
		info, err := c.mediaPlayer.GetInfo()

		if err != nil {
			c.broadcaster.Broadcast(mediaplayers.MediaPlayerInfo{Position: 0, Duration: 0, State: 0})
		}

		if info.Position > 0 && info.Position == info.Duration {
			c.mediaPlayer.Exit()
		}

		c.updateStore(info)

		c.broadcaster.Broadcast(info)
	}
}

// LaunchMediaFile plays a media file
func (c MediaController) LaunchMediaFile(profileName, mediafile string) (TextResponse, error) {
	profile, err := c.profilesProvider.GetProfile(profileName)

	if err != nil {
		return TextResponse{}, err
	}

	playerErr := c.mediaPlayer.Play(filepath.Join(profile.MediaPath, mediafile), profile.MediaArgs)

	return c.TextResponse(201, ""), playerErr
}

// CloseMediaPlayer instructs the media player to close
func (c MediaController) CloseMediaPlayer() error {
	return c.mediaPlayer.Exit()
}

// PauseMediaPlayer instructs the media player to close
func (c MediaController) PauseMediaPlayer() error {
	return c.mediaPlayer.Pause()
}

// ResumeMediaPlayer instructs the media player to close
func (c MediaController) ResumeMediaPlayer() error {
	return c.mediaPlayer.Resume()
}
