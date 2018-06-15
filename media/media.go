package media

import (
	"meteor/filesystem"
	"path/filepath"
)

// Media represents a file based media
type Media struct {
	Name        string `json:"name"`
	URI         string `json:"uri"`
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
func (m Provider) GetLocalMedia(mediaPath, subpath string) ([]Media, error) {
	files, err := m.source.GetFiles(filepath.Join(mediaPath, subpath))

	if err != nil {
		return []Media{}, err
	}

	mediaFiles := []Media{}

	for _, file := range files {
		mediaFiles = append(mediaFiles, Media{
			Name:        filepath.Base(file.Name),
			URI:         filepath.Join(subpath, file.Name),
			IsDirectory: file.IsDirectory,
		})
	}

	return mediaFiles, nil
}
