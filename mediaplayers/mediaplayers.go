package mediaplayers

import (
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

type MediaOptions struct {
	Position int
}

// MediaPlayerStatus represents the state of the media player
type MediaPlayerStatus struct {
	NowPlaying string `json:"nowPlaying"`
	State      int    `json:"state"`
	Position   int    `json:"position"`
	Duration   int    `json:"duration"`
}

// MediaPlayer represents a media player
type MediaPlayer interface {
	Play(media string, mediaArgs []string, options MediaOptions) error
	Exit() error
	Pause() error
	Resume() error
	Restart() error
	GetStatus() (MediaPlayerStatus, error)
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

func (m mediaPlayerClassic) buildArguments(filename string, playerArgs, mediaArgs []string, options MediaOptions) []string {
	launchArgs := []string{}

	launchArgs = append(launchArgs, filename)
	launchArgs = append(launchArgs, m.splitArguments(playerArgs)...)
	launchArgs = append(launchArgs, m.splitArguments(mediaArgs)...)

	if options.Position > 0 {
		launchArgs = append(launchArgs, "/start")
		launchArgs = append(launchArgs, strconv.Itoa(options.Position))
	}

	return launchArgs
}

func (m mediaPlayerClassic) Play(media string, mediaArgs []string, options MediaOptions) error {
	cmd := exec.Command(
		m.launchCmd,
		m.buildArguments(media, m.launchArgs, mediaArgs, options)...,
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

func (m mediaPlayerClassic) Restart() error {
	err := m.sendCommand("890")

	if err != nil {
		return err
	}

	return m.Resume()
}

func (m mediaPlayerClassic) GetStatus() (MediaPlayerStatus, error) {
	doc, err := htmlquery.LoadURL(fmt.Sprintf("%s/variables.html", m.apiURL))

	if err != nil {
		return MediaPlayerStatus{}, err
	}

	return m.readVariables(doc), nil
}

func (m mediaPlayerClassic) getFilename(doc *html.Node) string {
	_, file := filepath.Split(m.readVariable(doc, "filepath"))
	return file
}

func (m mediaPlayerClassic) getIntField(doc *html.Node, field string) int {
	value, err := strconv.Atoi(m.readVariable(doc, field))

	if err != nil {
		return 0
	}

	return value
}

func (m mediaPlayerClassic) readVariables(doc *html.Node) MediaPlayerStatus {
	return MediaPlayerStatus{
		NowPlaying: m.getFilename(doc),
		Position:   m.getIntField(doc, "position"),
		Duration:   m.getIntField(doc, "duration"),
		State:      m.getIntField(doc, "state"),
	}
}

func (m mediaPlayerClassic) readVariable(doc *html.Node, id string) string {
	node := htmlquery.FindOne(doc, fmt.Sprintf("//p[@id='%s']", id))
	return htmlquery.InnerText(node)
}
