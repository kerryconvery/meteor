package configuration

import (
	"meteor/filesystem"
	"meteor/types"
)

const defaultProfilePath = "/profiles"

// GetConfiguration reads a configuration file and returns back a configuration object.
// In the event the file could not be read the a default configuration is returned instread.
func GetConfiguration(path, configurationFile string) types.Configuration {
	var configuration types.Configuration

	err := filesystem.New(path).ReadJSONFile(configurationFile, &configuration)

	if err != nil {
		return types.Configuration{defaultProfilePath}
	}
	return configuration
}
