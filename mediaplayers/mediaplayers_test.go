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

	launchArgs := mediaplayer.buildArguments("media.avi", playerArgs, mediaArgs)

	if len(launchArgs) != 7 {
		t.Errorf("Expected 7 arguments but got %d", len(launchArgs))
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
	html := "<p id='filepath'>movie.avi</p><p id='position'>100</p><p id='state'>1</p>"

	mediaPlayer := mediaPlayerClassic{}

	doc, err := htmlquery.Parse(strings.NewReader(html))

	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	info := mediaPlayer.readVariables(doc)

	if info.NowPlaying != "movie.avi" {
		t.Errorf("Expected movie.avi but got %s", info.NowPlaying)
	}

	if info.Position != "100" {
		t.Errorf("Expected 100 but got %s", info.Position)
	}

	if info.State != "1" {
		t.Errorf("Expected 1 but got %s", info.State)
	}
}
