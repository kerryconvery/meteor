package mediaplayers

import (
	"strings"
	"testing"

	"github.com/antchfx/htmlquery"
)

func TestBuildArguments(t *testing.T) {
	mediaplayer := mediaPlayerClassic{}

	playerArgs := []string{"/play", "/fullscreen", "/autoclose"}
	mediaArgs := []string{"/audiorenderer 2", "/someotherarg"}
	options := MediaOptions{Position: 10}

	launchArgs := mediaplayer.buildArguments("media.avi", playerArgs, mediaArgs, options)

	if len(launchArgs) != 9 {
		t.Errorf("Expected 9 arguments but got %d", len(launchArgs))
	}

	if launchArgs[7] != "/start" {
		t.Errorf("Expected start switch but got %s", launchArgs[7])
	}

	if launchArgs[8] != "10" {
		t.Errorf("Expected start switch value 10 but got %s", launchArgs[8])
	}
}

func TestSplitArguments(t *testing.T) {
	arguments := []string{"/audiorenderer 2", "/play", "/file abc"}

	mediaPlayer := mediaPlayerClassic{}

	splitArguments := mediaPlayer.splitArguments(arguments)

	if len(splitArguments) != 5 {
		t.Errorf("Expected 5 arguments but got %d", len(splitArguments))
	}
}

func TestReadVariables(t *testing.T) {
	html := "<p id='filepath'>c:\\movies\\movie.avi</p><p id='position'>100</p><p id='duration'>1000</p><p id='state'>1</p>"

	mediaPlayer := mediaPlayerClassic{}

	doc, err := htmlquery.Parse(strings.NewReader(html))

	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	info := mediaPlayer.readVariables(doc)

	if info.NowPlaying != "movie.avi" {
		t.Errorf("Expected movie.avi but got %s", info.NowPlaying)
	}

	if info.Position != 100 {
		t.Errorf("Expected 100 but got %d", info.Position)
	}

	if info.Duration != 1000 {
		t.Errorf("Expected 1000 but got %d", info.Duration)
	}

	if info.State != 1 {
		t.Errorf("Expected 1 but got %d", info.State)
	}
}
