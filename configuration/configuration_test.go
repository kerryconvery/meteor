package configuration

import (
	"testing"
)

func TestGetConfigurationReturnsConfiguration(t *testing.T) {
	const sampleConfiguration = "../test_data/sample_config.json"

	var configuration = GetConfiguration(sampleConfiguration)

	if configuration.ProfilePath != "/media_profiles" {
		t.Errorf("GetConfiguration did not read file %s. Expected ProfilePath to be /profiles but got %s", sampleConfiguration, configuration.ProfilePath)
	}
}
func TestGetConfigurationReturnsDefaultConfiguration(t *testing.T) {
	const sampleConfiguration = "../test_data/file_does_not_exist.json"

	var configuration = GetConfiguration(sampleConfiguration)

	if configuration.ProfilePath != "/profiles" {
		t.Errorf("GetConfiguration did not return the default configuration. Expected ProfilePath to be /profiles but got %s", configuration.ProfilePath)
	}
}
