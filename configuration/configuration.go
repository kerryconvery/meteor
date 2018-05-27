package configuration

import (
	"encoding/json"
	"io/ioutil"
)

const defaultProfilePath = "/profiles"

// Configuration represents the content of the configuratio file
type Configuration struct {
	ProfilePath string `json:"profile_path"`
}

// GetConfiguration reads a configuration file and returns back a configuration object.
// In the event the file could not be read the a default configuration is returned instread.
func GetConfiguration(configurationFile string) Configuration {
	raw, err := ioutil.ReadFile(configurationFile)

	if err != nil {
		return Configuration{defaultProfilePath}
	}

	var configuration Configuration

	json.Unmarshal(raw, &configuration)

	return configuration
}
