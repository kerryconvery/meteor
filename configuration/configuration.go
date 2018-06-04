package configuration

import (
	"meteor/filesystem"
)

const defaultProfilePath = "/profiles"

// Configuration represents the content of the configuratio file
type Configuration struct {
	ProfilePath string `json:"profilePath"`
}

// GetConfiguration reads a configuration file and returns back a configuration object.
// In the event the file could not be read the a default configuration is returned instread.
func GetConfiguration(path, configurationFile string, filesystem filesystem.Filesystem) Configuration {
	var configuration Configuration

	err := filesystem.ReadJSONFile(path, configurationFile, &configuration)

	if err != nil {
		return Configuration{defaultProfilePath}
	}
	return configuration
}
