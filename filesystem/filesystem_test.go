package filesystem

import (
	"meteor/tests"
	"testing"
)

type sample struct {
	Name string
}

func TestGetFiles(t *testing.T) {
	files, err := New().GetFiles("../test_data/profiles")

	tests.ExpectNoError(err, t)

	if len(files) != 2 {
		t.Errorf("Expected two files but got %d", len(files))
	}

	if files[0].Name != "movies.json" {
		t.Errorf("Expected movies.json but got %s", files[0].Name)
	}

	if files[0].IsDirectory != false {
		t.Error("Expected file to not be a directory but got true")
	}
}

func TestGetDirectory(t *testing.T) {
	files, err := New().GetFiles("../test_data")

	tests.ExpectNoError(err, t)

	if files[0].Name != "profiles" {
		t.Errorf("Expected profiles but got %s", files[0].Name)
	}

	if files[0].IsDirectory != true {
		t.Error("Expected directory but got false")
	}
}
func TestGetFilesError(t *testing.T) {
	_, err := New().GetFiles("../test_data/does_not_exist")

	tests.ExpectError(err, t)
}

func TestReadJsonFile(t *testing.T) {
	content := sample{}
	err := New().ReadJSONFile("../test_data", "sample.json", &content)

	tests.ExpectNoError(err, t)

	if content.Name != "abc" {
		t.Errorf("Expected abc but got %s", content.Name)
	}
}

func TestReadJsonFileError(t *testing.T) {
	err := New().ReadJSONFile("../test_data", "does_not_exist.json", &sample{})

	tests.ExpectError(err, t)
}

func TestFileExists(t *testing.T) {
	exists, err := New().FileExists("../test_data/profiles", "movies.json")

	tests.ExpectNoError(err, t)

	if exists != true {
		t.Error("Expected true but got false")
	}
}

func TestFileDoesNotExists(t *testing.T) {
	exists, err := New().FileExists("../test_data/profile", "does_not_exist.json")

	tests.ExpectNoError(err, t)

	if exists != false {
		t.Error("Expected false but got true")
	}
}

func TestFileFolderDoesNotExists(t *testing.T) {
	exists, err := New().FileExists("../does_not_exists", "does_exist.json")

	tests.ExpectNoError(err, t)

	if exists != false {
		t.Error("Expected false but got true")
	}
}
