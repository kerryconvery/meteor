package media

import (
	"errors"
	"meteor/filesystem"
	"meteor/tests"
	"testing"
)

type sampleFiles struct {
}

func (f sampleFiles) GetFiles(path string) ([]filesystem.File, error) {
	switch path {
	case "MediaA":
		return []filesystem.File{
				filesystem.File{"movie.avi", false},
				filesystem.File{"music.mp3", false}},
			nil
	case "MediaB":
		return []filesystem.File{
				filesystem.File{"movie.avi", false},
				filesystem.File{"music.mp3", false},
				filesystem.File{"more movies", true}},
			nil
	}
	return []filesystem.File{}, errors.New("unknown path")
}

func (f sampleFiles) FileExists(path, fileName string) (bool, error) {
	if path == "path_exists" {
		return true, nil
	}

	if path == "path_not_found" {
		return false, nil
	}

	return false, errors.New("erorr")
}

func TestGetLocalMedia(t *testing.T) {
	provider := New(sampleFiles{})

	files, err := provider.GetLocalMedia("MediaB")

	tests.ExpectNoError(err, t)

	if len(files) != 3 {
		t.Errorf("Expected 3 files but got %d", len(files))
	}

	if files[0].Name != "movie.avi" {
		t.Errorf("Expected movie.avi but got %s", files[0].Name)
	}

	if files[0].IsDirectory != false {
		t.Error("Expected to be a file but got a directory")
	}

	if files[2].IsDirectory != true {
		t.Error("Expected to be a directory but got a file")
	}
}

func TestGetLocalMediaError(t *testing.T) {
	provider := New(sampleFiles{})

	_, err := provider.GetLocalMedia("MediaC")

	tests.ExpectError(err, t)
}
