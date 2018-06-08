package configuration

import (
	"meteor/filesystem"
	"meteor/tests"
	"testing"
)

func TestGetConfigurationReturnsConfiguration(t *testing.T) {
	configuration, err := GetConfiguration("../test_data", "sample_config.json", filesystem.New())

	tests.ExpectNoError(err, t)

	if configuration.ProfilePath != "/media_profiles" {
		t.Errorf("Expected ProfilePath to be /nedia_profiles but got %s", configuration.ProfilePath)
	}
}
func TestGetConfigurationReturnsError(t *testing.T) {
	_, err := GetConfiguration("../test_data", "file_does_not_exist.json", filesystem.New())

	tests.ExpectError(err, t)
}
