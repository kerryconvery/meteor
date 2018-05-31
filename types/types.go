package types

// Configuration represents the content of the configuratio file
type Configuration struct {
	ProfilePath string `json:"profile_path"`
}

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
