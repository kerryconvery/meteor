package mediaplayers

import (
	"os/exec"
	"strings"
)

// MediaPlayer represents a media player
type MediaPlayer interface {
	Play(media string, mediaArgs []string) error
}

type mediaPlayerClassic struct {
	launchCmd  string
	launchArgs []string
}

// New returns new instance of MediaPlayers
func New(launchCmd string, launchArgs []string) MediaPlayer {
	return mediaPlayerClassic{launchCmd: launchCmd, launchArgs: launchArgs}
}

func (m mediaPlayerClassic) splitArguments(arguments []string) []string {
	splitArguments := []string{}

	for _, arg := range arguments {
		splitArguments = append(splitArguments, strings.Split(arg, " ")...)
	}

	return splitArguments
}

func (m mediaPlayerClassic) buildArguments(filename string, playerArgs, mediaArgs []string) []string {
	launchArgs := []string{}

	launchArgs = append(launchArgs, filename)
	launchArgs = append(launchArgs, m.splitArguments(playerArgs)...)
	launchArgs = append(launchArgs, m.splitArguments(mediaArgs)...)

	return launchArgs
}

func (m mediaPlayerClassic) Play(media string, mediaArgs []string) error {
	cmd := exec.Command(
		m.launchCmd,
		m.buildArguments(media, m.launchArgs, mediaArgs)...,
	)
	return cmd.Run()
}
