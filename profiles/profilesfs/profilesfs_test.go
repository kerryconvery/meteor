package profilesfs

import (
	"errors"
	"meteor/tests"
	"meteor/types"
	"reflect"
	"testing"
)

type sampleFiles struct {
	Path string
}

func (f sampleFiles) GetFiles() ([]string, error) {
	switch f.Path {
	case "valid_path":
		return []string{"movies.json", "music.json"}, nil
	case "invalid_file":
		return []string{"movies.json", "invalid_file", "music.json"}, nil
	}
	return []string{}, errors.New("unknown path")
}

func replace(a, b interface{}) {
	ra := reflect.ValueOf(a).Elem()
	rb := reflect.ValueOf(b).Elem()
	ra.Set(rb)
}

func (f sampleFiles) ReadJSONFile(fileName string, content interface{}) error {
	switch fileName {
	case "movies.json":
		replace(content, &types.Profile{Name: "movies", Password: "123"})
		return nil
	case "music.json":
		replace(content, &types.Profile{Name: "music", Password: ""})
		return nil
	}
	return errors.New("File not found")
}

func NewProfilesValidPath() ProfilesFS {
	return ProfilesFS{FS: sampleFiles{Path: "valid_path"}}
}

func NewProfilesInvalidPath() ProfilesFS {
	return ProfilesFS{FS: sampleFiles{Path: "invalid_path"}}
}

func NewProfilesWithInvalidFile() ProfilesFS {
	return ProfilesFS{FS: sampleFiles{Path: "invalid_file"}}
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

	if profiles[0].Password != "123" {
		t.Errorf("Expected a profile password 123 but got %s", profiles[0].Password)
	}

	if profiles[1].Password != "" {
		t.Errorf("Expected no profile password but got %s", profiles[1].Password)
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

	if profile.Password != "123" {
		t.Errorf("Expected profile password to be 123 but got %s", profile.Password)
	}
}

func TestGetProfileNoOptionalFields(t *testing.T) {
	profile, err := NewProfilesValidPath().GetProfile("music")

	tests.ExpectNoError(err, t)

	if profile.Password != "" {
		t.Errorf("Expected profile password to be empty but got %s", profile.Password)
	}
}

func TestGetProfileError(t *testing.T) {
	_, err := NewProfilesValidPath().GetProfile("profile_not_found")

	tests.ExpectError(err, t)
}

func TestStripFileExtension(t *testing.T) {
	fileName := stripFileExtension("sample.json")

	if fileName != "sample" {
		t.Errorf("StripFileExtension - Expected %s but got %s", "sample", fileName)
	}
}
