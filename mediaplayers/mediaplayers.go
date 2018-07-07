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

// MediaPlayerInfo represents the state of the media player
type MediaPlayerInfo struct {
	NowPlaying string `json:"nowPlaying"`
	State      int    `json:"state"`
	Position   int    `json:"position"`
	Duration   int    `json:"duration"`
}

// MediaPlayer represents a media player
type MediaPlayer interface {
	Play(media string, mediaArgs []string) error
	Exit() error
	Pause() error
	Resume() error
	GetInfo() (MediaPlayerInfo, error)
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

func (m mediaPlayerClassic) GetInfo() (MediaPlayerInfo, error) {
	doc, err := htmlquery.LoadURL(fmt.Sprintf("%s/variables.html", m.apiURL))

	if err != nil {
		return MediaPlayerInfo{}, err
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

func (m mediaPlayerClassic) readVariables(doc *html.Node) MediaPlayerInfo {
	return MediaPlayerInfo{
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
