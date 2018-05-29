package filesystem

import (
	"meteor/tests"
	"testing"
)

type sample struct {
	Name string
}

func TestGetFiles(t *testing.T) {
	files, err := Filesystem{Path: "../test_data/profiles"}.GetFiles()

	tests.ExpectNoError(err, t)

	if len(files) != 2 {
		t.Errorf("Expected two files but got %d", len(files))
	}

	if files[0] != "movies.json" {
		t.Errorf("Expected movies.json but got %s", files[0])
	}
}
func TestGetFilesError(t *testing.T) {
	_, err := Filesystem{Path: "../test_data/does_not_exist"}.GetFiles()

	tests.ExpectError(err, t)
}

func TestReadJsonFile(t *testing.T) {
	content := sample{}
	err := Filesystem{Path: "../test_data"}.ReadJSONFile("sample.json", &content)

	tests.ExpectNoError(err, t)

	if content.Name != "abc" {
		t.Errorf("Expected abc but got %s", content.Name)
	}
}

func TestReadJsonFileError(t *testing.T) {
	err := Filesystem{Path: "../test_data"}.ReadJSONFile("does_not_exist.json", &sample{})

	tests.ExpectError(err, t)
}
