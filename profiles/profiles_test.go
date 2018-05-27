package profiles

import (
	"errors"
	"testing"
)

type sampleProfiles struct {
	profilePath string
}

func (p sampleProfiles) ReadProfileNames() ([]string, error) {
	switch p.profilePath {
	case "valid_profiles":
		return []string{"movies", "music"}, nil
	case "with_invalid_profiles":
		return []string{"movies", "cannot_parse", "music"}, nil
	}
	return []string{}, errors.New("profile path does not exist")
}

func (p sampleProfiles) ReadProfileProperties(profileName string) (ProfileProperties, error) {
	switch profileName {
	case "movies":
		return ProfileProperties{Password: "123"}, nil
	case "music":
		return ProfileProperties{}, nil
	case "can_not_parse":
		return ProfileProperties{}, errors.New("Could not parse profile")
	}
	return ProfileProperties{}, errors.New("file does not exist")

}

func validProfiles() Profiles {
	return New(sampleProfiles{profilePath: "valid_profiles"})
}

func withInvalidProfiles() Profiles {
	return New(sampleProfiles{profilePath: "with_invalid_profiles"})
}

var filesystem = ProfileFileSystem{ProfilePath: "../test_data/profiles"}

func TestGetProfiles(t *testing.T) {
	var profiles, err = validProfiles().GetProfiles()

	if err != nil {
		t.Errorf("GetProfiles - Expected no errors but got %s", err)
	}

	if len(profiles) != 2 {
		t.Errorf("GetProfiles did not return a list of profiles. Expected length 2 but got %d", len(profiles))
	}
}
func TestGetProfilesError(t *testing.T) {
	var profiles, err = New(sampleProfiles{profilePath: "does_not_exist"}).GetProfiles()

	if err == nil {
		t.Error("GetProfiles - Expected errors but got nil")
	}

	if len(profiles) != 0 {
		t.Errorf("GetProfiles - Expected profiles of lenght 0 but got %d", len(profiles))
	}
}
func TestGetProfilesInvalidProfile(t *testing.T) {
	var profiles, err = withInvalidProfiles().GetProfiles()

	if err == nil {
		t.Error("Expected an error but got nil")
	}

	if len(profiles) != 1 {
		t.Errorf("Expected two profiles but go %d", len(profiles))
	}
}

func TestGetProfileHasAllFields(t *testing.T) {
	var profile, err = validProfiles().GetProfile("movies")

	if err != nil {
		t.Errorf("GetProfile - Expected no error but got %s", err)
	}

	if profile.Name != "movies" {
		t.Errorf("GetProfile - Expected profile name movies but got %s", profile.Name)
	}

	if profile.Locked != true {
		t.Errorf("GetProfile - Expected profile %s to be locked but got %d", profile.Name, profile.Locked)
	}
}

func TestGetProfileNoOptionalFields(t *testing.T) {
	var profile, err = validProfiles().GetProfile("music")

	if err != nil {
		t.Errorf("GetProfile - Expected no error but got %s", err)
	}

	if profile.Locked == true {
		t.Errorf("GetProfile - Expected profile %s to be unlocked but got %b", profile.Name, profile.Locked)
	}
}

func TestGetProfileHasError(t *testing.T) {
	var _, err = validProfiles().GetProfile("not_found")

	if err == nil {
		t.Error("GetProfile - Expected error but got nil")
	}
}
