package controllers

import (
	"meteor/db"
	"meteor/mediaplayers"
	"meteor/profiles"
	"path/filepath"
)

type mediaPlayerListener interface {
	Notify(status mediaplayers.MediaPlayerStatus)
}
type profileProvider interface {
	GetProfile(profileName string) (profiles.Profile, error)
}

type mediaStore interface {
	Read(key string) (db.MediaRecord, error)
}

// MediaController represents a controller of a media player
type MediaController struct {
	Controller
	store           mediaStore
	profileProvider profileProvider
	mediaPlayer     mediaplayers.MediaPlayer
}

// NewMediaController returns a new instance of the media controller
func NewMediaController(profileProvider profileProvider, mediaPlayer mediaplayers.MediaPlayer, store mediaStore) MediaController {
	return MediaController{
		store:           store,
		profileProvider: profileProvider,
		mediaPlayer:     mediaPlayer,
	}
}

func (c MediaController) getStartingPosition(mediaFile string) int {
	data, err := c.store.Read(mediaFile)

	if err != nil {
		return 0
	}

	if data.Duration-data.Position <= 10000 {
		return 0
	}

	return data.Position - 3000
}

// LaunchMediaFile plays a media file
func (c MediaController) LaunchMediaFile(profileName, mediafile string) (TextResponse, error) {
	profile, err := c.profileProvider.GetProfile(profileName)

	if err != nil {
		return TextResponse{}, err
	}

	options := mediaplayers.MediaOptions{Position: c.getStartingPosition(filepath.Base(mediafile))}

	playerErr := c.mediaPlayer.Play(filepath.Join(profile.MediaPath, mediafile), profile.MediaArgs, options)

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

// RestartMediaPlayer restarts the media from the beginning
func (c MediaController) RestartMediaPlayer() error {
	return c.mediaPlayer.Restart()
}

func (c MediaController) reportMediaStatus(status mediaplayers.MediaPlayerStatus, listeners []mediaPlayerListener) {
	for _, listener := range listeners {
		listener.Notify(status)
	}
}

//WatchMediaPlayer watches the media player and reports the current status to the listener
func (c MediaController) WatchMediaPlayer(signal chan bool, listeners ...mediaPlayerListener) {
	for terminate := range signal {
		if terminate {
			break
		}

		var status mediaplayers.MediaPlayerStatus

		temp, err := c.mediaPlayer.GetStatus()

		if err != nil {
			status = mediaplayers.MediaPlayerStatus{
				Position: 0,
				Duration: 0,
				State:    0,
			}
		} else {
			status = temp
		}

		c.reportMediaStatus(status, listeners)

		signal <- false
	}
}
