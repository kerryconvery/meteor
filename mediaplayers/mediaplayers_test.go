package mediaplayers

import "testing"

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
