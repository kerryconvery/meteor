package profiles

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
)

//ProfileProperties represents the content of a profile
type ProfileProperties struct {
	Password string
}

// ProfileFileSystem represents the underlying file system for accessing profile data
type ProfileFileSystem struct {
	ProfilePath string
}

// ReadProfileNames returns a list of profile names in the specified folder
func (p ProfileFileSystem) ReadProfileNames() ([]string, error) {
	var profiles []string

	files, err := ioutil.ReadDir(p.ProfilePath)

	if err != nil {
		return profiles, err
	}

	for _, file := range files {
		profiles = append(profiles, stripFileExtension(file.Name()))
	}

	return profiles, nil
}

// ReadProfileProperties returns the content of the profile file
func (p ProfileFileSystem) ReadProfileProperties(profileName string) (ProfileProperties, error) {
	raw, err := ioutil.ReadFile(filepath.Join(p.ProfilePath, profileName+".json"))

	if err != nil {
		return ProfileProperties{}, err
	}

	var profileProperties ProfileProperties

	json.Unmarshal(raw, &profileProperties)

	return profileProperties, nil
}

// StripFileExtension returns a file name with the extension including the final dot removed
func stripFileExtension(fileName string) string {
	var fileExt = filepath.Ext(fileName)

	return strings.Replace(fileName, fileExt, "", 1)
}
