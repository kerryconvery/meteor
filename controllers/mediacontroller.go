package controllers

import (
	"meteor/controllers/webhook"
	"meteor/mediaplayers"
	"meteor/profiles"
	"path/filepath"
)

type profileSource interface {
	GetProfile(profileName string) (profiles.Profile, error)
}

// MediaController represents a controller of a media player
type MediaController struct {
	Controller
	profilesProvider profileSource
	mediaPlayer      mediaplayers.MediaPlayer
	webhook          webhook.Webhook
}

// NewMediaController returns a new instance of the media controller
func NewMediaController(profilesProvider profileSource, mediaPlayer mediaplayers.MediaPlayer, webhook webhook.Webhook) MediaController {
	return MediaController{profilesProvider: profilesProvider, mediaPlayer: mediaPlayer, webhook: webhook}
}

func (c MediaController) startReporting() {
	go watchMediaPlayer(c.webhook, c.mediaPlayer)
}

func watchMediaPlayer(webhook webhook.Webhook, mediaPlayer mediaplayers.MediaPlayer) {
	for {
		info, err := mediaPlayer.GetInfo()

		if err != nil {
			webhook.Broadcast("", "", "0")
			return
		}

		webhook.Broadcast(info.NowPlaying, info.Position, info.State)
	}
}

// LaunchMediaFile plays a media file
func (c MediaController) LaunchMediaFile(profileName, mediafile string) (TextResponse, error) {
	profile, err := c.profilesProvider.GetProfile(profileName)

	if err != nil {
		return TextResponse{}, err
	}

	playerErr := c.mediaPlayer.Play(filepath.Join(profile.MediaPath, mediafile), profile.MediaArgs)

	c.startReporting()

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
