package types

// Profile represents a profile
type Profile struct {
	Name     string
	Password string
}

// Profiles represents the IO operations for profiles
type Profiles interface {
	GetProfiles() ([]Profile, error)
	GetProfile(profileName string) (Profile, error)
}
