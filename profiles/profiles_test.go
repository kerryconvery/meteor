package profiles

import (
	"errors"
	"meteor/filesystem"
	"meteor/tests"
	"reflect"
	"testing"
)

type sampleFiles struct {
}

func (f sampleFiles) GetFiles(path string) ([]filesystem.File, error) {
	switch path {
	case "valid_path":
		return []filesystem.File{
				filesystem.File{"movies.json", false},
				filesystem.File{"music.json", false}},
			nil
	case "invalid_file":
		return []filesystem.File{
				filesystem.File{"movies.json", false},
				filesystem.File{"invalid_file", false},
				filesystem.File{"music.json", false}},
			nil
	}
	return []filesystem.File{}, errors.New("unknown path")
}

func (f sampleFiles) ReadJSONFile(path, fileName string, content interface{}) error {
	switch fileName {
	case "movies.json":
		replace(content, &Profile{Name: "movies", ParentalPassword: "123"})
		return nil
	case "music.json":
		replace(content, &Profile{Name: "music", ParentalPassword: ""})
		return nil
	}
	return errors.New("File not found")
}

func (f sampleFiles) FileExists(path, fileName string) (bool, error) {
	if path == "valid_path" {
		return fileName == "movies.json", nil
	} else {
		return false, errors.New("Path not found")
	}
}

func replace(a, b interface{}) {
	ra := reflect.ValueOf(a).Elem()
	rb := reflect.ValueOf(b).Elem()
	ra.Set(rb)
}
func NewProfilesValidPath() Provider {
	return profileProvider{"valid_path", sampleFiles{}}
}

func NewProfilesInvalidPath() Provider {
	return profileProvider{"invalid_path", sampleFiles{}}
}

func NewProfilesWithInvalidFile() Provider {
	return profileProvider{"invalid_file", sampleFiles{}}
}

func TestGetProfiles(t *testing.T) {
	profiles, err := NewProfilesValidPath().GetProfiles()

	tests.ExpectNoError(err, t)

	if len(profiles) != 2 {
		t.Errorf("Expected two profiles but got %d", len(profiles))
	}

	if profiles[0].Name != "movies" {
		t.Errorf("Expected a profile called movies but got %s", profiles[0].Name)
	}

	if profiles[0].ParentalPassword != "123" {
		t.Errorf("Expected a profile password 123 but got %s", profiles[0].ParentalPassword)
	}

	if profiles[1].ParentalPassword != "" {
		t.Errorf("Expected no profile password but got %s", profiles[1].ParentalPassword)
	}
}

func TestGetProfilesWithInvalidFile(t *testing.T) {
	_, err := NewProfilesWithInvalidFile().GetProfiles()
	tests.ExpectError(err, t)
}

func TestGetProfilesError(t *testing.T) {
	_, err := NewProfilesInvalidPath().GetProfiles()
	tests.ExpectError(err, t)
}

func TestGetProfile(t *testing.T) {
	profile, err := NewProfilesValidPath().GetProfile("movies")

	tests.ExpectNoError(err, t)

	if profile.Name != "movies" {
		t.Errorf("Expected profile name movies but got %s", profile.Name)
	}

	if profile.ParentalPassword != "123" {
		t.Errorf("Expected profile password to be 123 but got %s", profile.ParentalPassword)
	}
}

func TestProfileExists(t *testing.T) {
	exists, err := NewProfilesValidPath().ProfileExists("movies")

	tests.ExpectNoError(err, t)

	if exists != true {
		t.Error("Expected true but got false")
	}
}

func TestProfileNotExists(t *testing.T) {
	exists, err := NewProfilesValidPath().ProfileExists("does_not_exist")

	tests.ExpectNoError(err, t)

	if exists != false {
		t.Error("Expected false but got true")
	}
}

func TestProfileExistsError(t *testing.T) {
	_, err := NewProfilesInvalidPath().ProfileExists("movies")

	tests.ExpectError(err, t)
}
func TestGetProfileNoOptionalFields(t *testing.T) {
	profile, err := NewProfilesValidPath().GetProfile("music")

	tests.ExpectNoError(err, t)

	if profile.ParentalPassword != "" {
		t.Errorf("Expected profile password to be empty but got %s", profile.ParentalPassword)
	}
}

func TestStripFileExtension(t *testing.T) {
	fileName := stripFileExtension("sample.json")

	if fileName != "sample" {
		t.Errorf("StripFileExtension - Expected %s but got %s", "sample", fileName)
	}
}
