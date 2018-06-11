package mediaplayers

import (
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
)

// MediaPlayer represents a media player
type MediaPlayer interface {
	Play(media string, mediaArgs []string) error
	Exit() error
	Pause() error
	Resume() error
}

type mediaPlayerClassic struct {
	launchCmd  string
	launchArgs []string
	apiURL     string
}

// New returns new instance of MediaPlayers
func New(launchCmd string, launchArgs []string, apiURL string) MediaPlayer {
	return mediaPlayerClassic{launchCmd, launchArgs, apiURL}
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

	return cmd.Start()
}

func (m mediaPlayerClassic) sendCommand(command string) error {
	_, err := http.PostForm(
		fmt.Sprintf("%s/command.html", m.apiURL),
		url.Values{"wm_command": {command}},
	)
	return err
}

func (m mediaPlayerClassic) Exit() error {
	return m.sendCommand("816")
}

func (m mediaPlayerClassic) Pause() error {
	return m.sendCommand("888")
}

func (m mediaPlayerClassic) Resume() error {
	return m.sendCommand("887")
}
