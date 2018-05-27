package profiles

// ProfileIO represents the IO operations for profiles
type profileIO interface {
	ReadProfileNames() ([]string, error)
	ReadProfileProperties(profileName string) (ProfileProperties, error)
}

// Profile contains the profile properties as well as the name
type Profile struct {
	Name   string
	Locked bool
}

// Profiles reads profiles
type Profiles struct {
	io profileIO
}

// New returns a new instance of Profiles
func New(io profileIO) Profiles {
	return Profiles{io}
}

// GetProfile returns back a profile
func (p Profiles) GetProfile(profileName string) (Profile, error) {
	var profileProperties, err = p.io.ReadProfileProperties(profileName)

	if err != nil {
		return Profile{}, err
	}

	var profile = Profile{Name: profileName, Locked: profileProperties.Password != ""}

	return profile, nil
}

// GetProfiles returns a list of profiles located at the profile path otherwise it returns an empty list
func (p Profiles) GetProfiles() ([]Profile, error) {
	var profiles []Profile

	var profileNames, err = p.io.ReadProfileNames()

	if err != nil {
		return profiles, err
	}

	for _, profileName := range profileNames {
		var profile, err = p.GetProfile(profileName)

		if err == nil {
			profiles = append(profiles, profile)
		} else {
			return profiles, err
		}
	}

	return profiles, nil
}
