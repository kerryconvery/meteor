package profiles

import (
	"meteor/filesystem"
	"meteor/profiles/profilesfs"
	"meteor/types"
)

// New creates a new instance of Profiles
func New(profilePath string) types.Profiles {
	filesystem := filesystem.Filesystem{Path: profilePath}
	return profilesfs.ProfilesFS{FS: filesystem}
}
