package profiles

import "testing"

func TestReadProfileNames(t *testing.T) {
	var profileNames, _ = filesystem.ReadProfileNames()

	if len(profileNames) != 2 {
		t.Errorf("Expected length 2 but got %d", len(profileNames))
	}

	if profileNames[0] != "movies" {
		t.Errorf("Expected profile name movies but got %s", profileNames[0])
	}
}
func TestReadProfileNamesFolderDoesNotExist(t *testing.T) {
	var _, err = ProfileFileSystem{ProfilePath: "../test_data/does_not_exist"}.ReadProfileNames()

	if err == nil {
		t.Error("ReadProfileNames - Expected and error but got nil")
	}
}
func TestReadPropertiesProfile(t *testing.T) {
	var properties, err = filesystem.ReadProfileProperties("movies")

	if err != nil {
		t.Errorf("ReadProfile - Expected movies profile but got error %s", err)
	}

	if properties.Password != "123" {
		t.Errorf("ReadProfile - Expected movies profile pssword to be 123 but got %s", properties.Password)
	}
}
func TestReadProfilePropertiesOptionalProperties(t *testing.T) {
	var properties, _ = filesystem.ReadProfileProperties("music")

	if properties.Password != "" {
		t.Errorf("ReadProfile - Expected music profile pssword to be empty but got %s", properties.Password)
	}
}
func TestReadProfilePropertiesNotFound(t *testing.T) {
	var _, err = filesystem.ReadProfileProperties("not_found")

	if err == nil {
		t.Error("ReadProfile - Expected error but got nil error")
	}
}
func TestStripFileExtension(t *testing.T) {
	var fileName = stripFileExtension("sample.json")

	if fileName != "sample" {
		t.Errorf("StripFileExtension - Expected %s but got %s", "sample", fileName)
	}
}
