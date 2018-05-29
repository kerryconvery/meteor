package profilesfs

import (
	"meteor/types"
	"path/filepath"
	"strings"
)

type filesystem interface {
	GetFiles() ([]string, error)
	ReadJSONFile(fileName string, content interface{}) error
}

// ProfilesFS represents the profiles file system
type ProfilesFS struct {
	FS filesystem
}

// GetProfiles returns back a list of profiles
func (p ProfilesFS) GetProfiles() ([]types.Profile, error) {
	files, err := p.FS.GetFiles()
	if err != nil {
		return []types.Profile{}, err
	}

	profiles := []types.Profile{}

	for _, file := range files {
		profile, err := p.GetProfile(stripFileExtension(file))

		if err != nil {
			return []types.Profile{}, err
		}
		profiles = append(profiles, profile)
	}

	return profiles, nil
}

// GetProfile returns back a Profile of the given name
func (p ProfilesFS) GetProfile(profileName string) (types.Profile, error) {
	profile := types.Profile{}
	err := p.FS.ReadJSONFile(profileName+".json", &profile)
	if err != nil {
		return profile, err
	}
	profile.Name = profileName
	return profile, err
}

// StripFileExtension returns a file name with the extension including the final dot removed
func stripFileExtension(fileName string) string {
	fileExt := filepath.Ext(fileName)
	return strings.Replace(fileName, fileExt, "", 1)
}
