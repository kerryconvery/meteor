package mediaplayers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
)

// MediaPlayer represents a media player
type MediaPlayer interface {
	Play(media string, mediaArgs []string) error
	Exit() error
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

func (m mediaPlayerClassic) sendCommand(command string) (*http.Response, error) {
	response, err := http.PostForm(
		fmt.Sprintf("%s/command.html", m.apiURL),
		url.Values{"wm_command": {command}},
	)

	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 400 {
		body, _ := ioutil.ReadAll(response.Body)
		return nil, errors.New(string(body))
	}

	return response, nil
}

func (m mediaPlayerClassic) Exit() error {
	_, err := m.sendCommand("816")

	if err != nil {
		return err
	}

	return nil
}
