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

type fileSource interface {
	GetFiles(path string) ([]filesystem.File, error)
	FileExists(path, fileName string) (bool, error)
}

// Provider represents all media
type Provider struct {
	source fileSource
}

// New returns a new instance of Provider
func New(source fileSource) Provider {
	return Provider{source: source}
}

// GetLocalMedia returns a list of media files at the designated path
func (m Provider) GetLocalMedia(path string) ([]Media, error) {
	files, err := m.source.GetFiles(path)

	if err != nil {
		return []Media{}, err
	}

	mediaFiles := []Media{}

	for _, file := range files {
		mediaFiles = append(mediaFiles, Media{
			Name:        file.Name,
			IsDirectory: file.IsDirectory,
		})
	}

	return mediaFiles, nil
}

// PathExists return whether or not the path exists in the file system
func (m Provider) PathExists(path string) (bool, error) {
	return m.source.FileExists(path, "")
}
