package media

import (
	"meteor/filesystem"
)

// Media represents a file based media
type Media struct {
	Name        string `json:"name"`
	IsDirectory bool   `json:"isDirectory"`
	Thumbnail   string `json:"thumbnail"`
}

// Provider represents all media
type Provider struct {
	mediaRoute string
	filesystem filesystem.Filesystem
}

// New returns a new instance of Provider
func New(mediaRoute string, filesystem filesystem.Filesystem) Provider {
	return Provider{mediaRoute: mediaRoute, filesystem: filesystem}
}

// GetLocalMedia returns a list of media files at the designated path
func (m Provider) GetLocalMedia(path string) ([]Media, error) {
	files, err := m.filesystem.GetFiles(path)

	if err != nil {
		return []Media{}, err
	}

	mediaFiles := []Media{}

	for _, file := range files {
		mediaFiles = append(mediaFiles, Media{
			Name:        file.Name,
			IsDirectory: file.IsDirectory,
			Thumbnail:   m.mediaRoute + "/" + file.Name + "/thumbnail",
		})
	}

	return mediaFiles, nil
}

// PathExists return whether or not the path exists in the file system
func (m Provider) PathExists(path string) (bool, error) {
	return m.filesystem.FileExists(path, "")
}
