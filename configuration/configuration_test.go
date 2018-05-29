package configuration

import (
	"testing"
)

func TestGetConfigurationReturnsConfiguration(t *testing.T) {
	var configuration = GetConfiguration("../test_data", "sample_config.json")

	if configuration.ProfilePath != "/media_profiles" {
		t.Errorf("Expected ProfilePath to be /profiles but got %s", configuration.ProfilePath)
	}
}
func TestGetConfigurationReturnsDefaultConfiguration(t *testing.T) {
	var configuration = GetConfiguration("../test_data", "file_does_not_exist.json")

	if configuration.ProfilePath != "/profiles" {
		t.Errorf("GetConfiguration did not return the default configuration. Expected ProfilePath to be /profiles but got %s", configuration.ProfilePath)
	}
}
