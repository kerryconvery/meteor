package profiles

import (
	"meteor/filesystem"
	"path/filepath"
	"strings"
)

// Profile represents a profile
type Profile struct {
	Name             string   `json:"name"`
	MediaPath        string   `json:"mediaPath"`
	ParentalPassword string   `json:"parentalPassword"`
	MediaArgs        []string `json:"mediaArgs"`
}

// Provider represents a profileProvider of profiles
type Provider interface {
	GetProfiles() ([]Profile, error)
	GetProfile(profileName string) (Profile, error)
	ProfileExists(profileName string) (bool, error)
}

type fileSource interface {
	GetFiles(path string) ([]filesystem.File, error)
	ReadJSONFile(path string, fileName string, content interface{}) error
	FileExists(path, fileName string) (bool, error)
}

// profileProvider represents the IO operations for profiles
type profileProvider struct {
	path   string
	source fileSource
}

// New creates a new instance of profileProvider
func New(profilePath string) Provider {
	return profileProvider{path: profilePath, source: filesystem.New()}
}

// GetProfiles returns back a list of profiles
func (p profileProvider) GetProfiles() ([]Profile, error) {
	files, err := p.source.GetFiles(p.path)
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
func (p profileProvider) GetProfile(profileName string) (Profile, error) {
	profile := Profile{}
	err := p.source.ReadJSONFile(p.path, profileName+".json", &profile)
	if err != nil {
		return profile, err
	}
	profile.Name = profileName
	return profile, err
}

// ProfileExists returns if a profile exists or not
func (p profileProvider) ProfileExists(profileName string) (bool, error) {
	return p.source.FileExists(p.path, profileName+".json")
}

// StripFileExtension returns a file name with the extension including the final dot removed
func stripFileExtension(fileName string) string {
	fileExt := filepath.Ext(fileName)
	return strings.Replace(fileName, fileExt, "", 1)
}
