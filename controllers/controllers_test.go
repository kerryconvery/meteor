package controllers

import (
	"errors"
	"meteor/types"
	"testing"
)

type sampleProfiles struct {
	profiles []types.Profile
	err      error
}

func (p sampleProfiles) GetProfiles() ([]types.Profile, error) {
	return p.profiles, p.err
}

func (p sampleProfiles) GetProfile(profileName string) (types.Profile, error) {
	for _, profile := range p.profiles {
		if profile.Name == profileName {
			return profile, nil
		}
	}

	return types.Profile{}, errors.New("Profile not foun")
}

func expectedContentType(expected, actual string, t *testing.T) {
	if expected != actual {
		t.Errorf("Expected content type %s but got %s", expected, actual)
	}
}

func expectedStatusCode(expected, actual int, t *testing.T) {
	if expected != actual {
		t.Errorf("Expected status code %d but got %d", expected, actual)
	}
}

func TestGetProfiles(t *testing.T) {
	profiles := sampleProfiles{[]types.Profile{types.Profile{}}, nil}

	response := NewProfilesController(profiles).GetAll()

	expectedContentType("application/json", response.ContentType, t)
	expectedStatusCode(200, response.StatusCode, t)

	responseBody := response.Body.([]types.Profile)

	if len(responseBody) != 1 {
		t.Errorf("Expected 1 profiles but got %d", len(responseBody))
	}
}

func TestGetProfilesError(t *testing.T) {
	profiles := sampleProfiles{[]types.Profile{}, errors.New("Could not read profiles")}

	response := NewProfilesController(profiles).GetAll()

	expectedStatusCode(500, response.StatusCode, t)
}
