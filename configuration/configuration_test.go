package configuration

import (
	"testing"
)

func TestGetConfigurationReturnsConfiguration(t *testing.T) {
	configuration := GetConfiguration("../test_data", "sample_config.json")

	if configuration.ProfilePath != "/media_profiles" {
		t.Errorf("Expected ProfilePath to be /profiles but got %s", configuration.ProfilePath)
	}
}
func TestGetConfigurationReturnsDefaultConfiguration(t *testing.T) {
	configuration := GetConfiguration("../test_data", "file_does_not_exist.json")

	if configuration.ProfilePath != "/profiles" {
		t.Errorf("Expected ProfilePath to be /profiles but got %s", configuration.ProfilePath)
	}
}
