package configuration

import (
	"meteor/filesystem"
	"testing"
)

func TestGetConfigurationReturnsConfiguration(t *testing.T) {
	configuration := GetConfiguration("../test_data", "sample_config.json", filesystem.New())

	if configuration.ProfilePath != "/media_profiles" {
		t.Errorf("Expected ProfilePath to be /nedia_profiles but got %s", configuration.ProfilePath)
	}
}
func TestGetConfigurationReturnsDefaultConfiguration(t *testing.T) {
	configuration := GetConfiguration("../test_data", "file_does_not_exist.json", filesystem.New())

	if configuration.ProfilePath != "/profiles" {
		t.Errorf("Expected ProfilePath to be /profiles but got %s", configuration.ProfilePath)
	}
}
