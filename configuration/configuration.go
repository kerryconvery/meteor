package configuration

import (
	"meteor/filesystem"
)

// MediaPlayerConfiguration represents the configuration needed to launch a media player
type MediaPlayerConfiguration struct {
	Name       string   `json:"name"`
	LaunchCmd  string   `json:"launchCmd"`
	LaunchArgs []string `json:"launchArgs"`
	APIUrl     string   `json:"apiUrl"`
}

// Configuration represents the content of the configuratio file
type Configuration struct {
	ProfilePath   string                     `json:"profilePath"`
	ThumbnailPath string                     `json:"thumbnailPath"`
	AssetPath     string                     `json:"assetPath"`
	WebClientPath string                     `json:"webclientPath"`
	MediaPlayers  []MediaPlayerConfiguration `json:"mediaplayers"`
	ListenPort    int                        `json:"listenPort"`
	DatastorePath string                     `json:"datastorePath"`
}

// GetConfiguration reads a configuration file and returns back a configuration object.
// In the event the file could not be read the a default configuration is returned instread.
func GetConfiguration(path, configurationFile string, filesystem filesystem.Filesystem) (Configuration, error) {
	var configuration Configuration

	err := filesystem.ReadJSONFile(path, configurationFile, &configuration)

	if err != nil {
		return Configuration{}, err
	}

	return configuration, nil
}
