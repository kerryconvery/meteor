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
	case "C:\\MediaB":
		return []filesystem.File{
				filesystem.File{"movies\\movie.avi", false},
				filesystem.File{"movies\\more movies", true}},
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

func newMediaProvider() Provider {
	return New(sampleFiles{})
}

func TestGetLocalMedia(t *testing.T) {
	provider := newMediaProvider()

	files, err := provider.GetLocalMedia("C:\\", "MediaB")

	tests.ExpectNoError(err, t)

	if len(files) != 2 {
		t.Errorf("Expected 2 files but got %d", len(files))
	}

	if files[0].Name != "movie.avi" {
		t.Errorf("Expected movie.avi but got %s", files[0].Name)
	}

	if files[0].URI != "MediaB\\movies\\movie.avi" {
		t.Errorf("Expected MediaB\\movies\\movie.avi but got %s", files[0].URI)
	}

	if files[0].IsDirectory != false {
		t.Error("Expected to be a file but got a directory")
	}

	if files[1].Name != "more movies" {
		t.Errorf("Expected more movies but got %s", files[1].Name)
	}

	if files[1].URI != "MediaB\\movies\\more movies" {
		t.Errorf("Expected MediaB\\movies\\more movies but got %s", files[1].URI)
	}

	if files[1].IsDirectory != true {
		t.Error("Expected to be a directory but got a file")
	}
}

func TestGetLocalMediaInvalidMediaPathError(t *testing.T) {
	provider := newMediaProvider()

	_, err := provider.GetLocalMedia("", "MediaC")

	tests.ExpectError(err, t)
}
