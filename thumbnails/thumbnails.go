package thumbnails

import (
	"bytes"
	"errors"
	"meteor/filesystem"
	"os/exec"
	"path/filepath"
)

type imageProvider interface {
	getDefaultImage() (*bytes.Buffer, error)
	getExistingThumbnail(fileName string) (*bytes.Buffer, error)
	generateVideoThumbnail(source string) (*bytes.Buffer, error)
}

type imageSource struct {
	thumbnailPath    string
	defaultImagePath string
	filesystem       filesystem.Filesystem
}

// Provider represents a thumbnail provider
type Provider interface {
	GetThumbnail(path, filename string) (*bytes.Buffer, error)
}

type thumbnailProvider struct {
	imageSource imageProvider
}

// New returns a new instance of the thumbnail provider
func New(thumbnailPath, defaultImagePath string, filesystem filesystem.Filesystem) Provider {
	return thumbnailProvider{
		imageSource{
			thumbnailPath:    thumbnailPath,
			defaultImagePath: defaultImagePath,
			filesystem:       filesystem,
		},
	}
}

// GetThumbnail returns a thumbnail for the supplied video file
func (p thumbnailProvider) GetThumbnail(path, filename string) (*bytes.Buffer, error) {
	existingImage, err := p.imageSource.getExistingThumbnail(filename)

	if err == nil {
		return existingImage, nil
	}

	generatedImage, err := p.imageSource.generateVideoThumbnail(filepath.Join(path, filename))

	if err == nil {
		return generatedImage, nil
	}

	return p.imageSource.getDefaultImage()
}

func (p imageSource) getDefaultImage() (*bytes.Buffer, error) {
	return p.filesystem.ReadImageFile(p.defaultImagePath, "file.png")
}

func (p imageSource) getExistingThumbnail(fileName string) (*bytes.Buffer, error) {
	return p.filesystem.ReadImageFile(p.thumbnailPath, fileName)
}

// Generate extracts a thumbnail from a video file and returns a buffer containing the image
func (p imageSource) generateVideoThumbnail(source string) (*bytes.Buffer, error) {
	// fmt.Sprintf("ffmpeg -ss 00:05:00.00 -i '%s' -filter:v 'scale=64:64:force_original_aspect_ratio=decrease' -vframes 1 -q:v 2 '%s'")

	cmd := exec.Command(
		"ffmpeg",
		"-ss 00:05:00.00",
		"-i",
		source,
		"-filter:v 'scale=64:64:force_original_aspect_ratio=decrease'",
		"-vframes",
		"1",
		"-q:v 2",
	)

	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	if cmd.Run() != nil {
		return &buffer, errors.New("could not generate thumbnail")
	}
	return &buffer, nil
}
