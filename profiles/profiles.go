package profiles

import (
	"meteor/filesystem"
	"path/filepath"
	"strings"
)

// Profile represents a profile
type Profile struct {
	Name             string `json:"name"`
	MediaPath        string `json:"mediaPath"`
	ParentalPassword string `json:"parentalPassword"`
}

// ProfileProvider represents a provider of profiles
type ProfileProvider interface {
	GetProfiles() ([]Profile, error)
	GetProfile(profileName string) (Profile, error)
	ProfileExists(profileName string) (bool, error)
}

// Provider represents the IO operations for profiles
type Provider struct {
	path       string
	filesystem filesystem.Filesystem
}

// New creates a new instance of Provider
func New(profilePath string) Provider {
	return Provider{path: profilePath, filesystem: filesystem.New()}
}

// GetProfiles returns back a list of profiles
func (p Provider) GetProfiles() ([]Profile, error) {
	files, err := p.filesystem.GetFiles(p.path)
	if err != nil {
		return []Profile{}, err
	}

	profiles := []Profile{}

	for _, file := range files {
		profile, err := p.GetProfile(stripFileExtension(file.Name))

		if err != nil {
			return []Profile{}, err
		}
		profiles = append(profiles, profile)
	}

	return profiles, nil
}

// GetProfile returns back a Profile of the given name
func (p Provider) GetProfile(profileName string) (Profile, error) {
	profile := Profile{}
	err := p.filesystem.ReadJSONFile(p.path, profileName+".json", &profile)
	if err != nil {
		return profile, err
	}
	profile.Name = profileName
	return profile, err
}

// ProfileExists returns if a profile exists or not
func (p Provider) ProfileExists(profileName string) (bool, error) {
	return p.filesystem.FileExists(p.path, profileName+".json")
}

// StripFileExtension returns a file name with the extension including the final dot removed
func stripFileExtension(fileName string) string {
	fileExt := filepath.Ext(fileName)
	return strings.Replace(fileName, fileExt, "", 1)
}
