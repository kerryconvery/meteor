package thumbnails

import (
	"bytes"
	"meteor/filesystem"
	"os/exec"
	"path/filepath"
)

type imageProvider interface {
	getDefaultImage() (*bytes.Buffer, error)
	getExistingThumbnail(fileName string) (*bytes.Buffer, error)
	generateVideoThumbnail(source string) (*bytes.Buffer, error)
	addThumbnail(filename string, image *bytes.Buffer) (int, error)
}

type fileSource interface {
	ReadFile(path, fileName string) (*bytes.Buffer, error)
	WriteFile(path, fileName string, buffer *bytes.Buffer) (int, error)
}

type imageSource struct {
	thumbnailPath    string
	defaultImagePath string
	files            fileSource
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
	source := imageSource{
		thumbnailPath:    thumbnailPath,
		defaultImagePath: defaultImagePath,
		files:            filesystem,
	}

	return thumbnailProvider{source}
}

// GetThumbnail returns a thumbnail for the supplied video file
func (p thumbnailProvider) GetThumbnail(path, filename string) (*bytes.Buffer, error) {
	existingImage, err := p.imageSource.getExistingThumbnail(filename)

	if err == nil {
		return existingImage, nil
	}

	generatedImage, err := p.imageSource.generateVideoThumbnail(filepath.Join(path, filename))

	if err == nil {
		p.imageSource.addThumbnail(filename, generatedImage)

		return generatedImage, nil
	}

	return p.imageSource.getDefaultImage()
}

func (p imageSource) getDefaultImage() (*bytes.Buffer, error) {
	return p.files.ReadFile(p.defaultImagePath, "file.png")
}

func (p imageSource) getExistingThumbnail(fileName string) (*bytes.Buffer, error) {
	return p.files.ReadFile(p.thumbnailPath, fileName+".jpg")
}

func (p imageSource) addThumbnail(filename string, image *bytes.Buffer) (int, error) {
	return p.files.WriteFile(p.thumbnailPath, filename+".jpg", image)
}

// Generate extracts a thumbnail from a video file and returns a buffer containing the image
func (p imageSource) generateVideoThumbnail(source string) (*bytes.Buffer, error) {
	cmd := exec.Command(
		"ffmpeg",
		"-ss",
		"00:05:00.00",
		"-i",
		source,
		"-vframes",
		"1",
		"-s",
		"64x64",
		"-f",
		"singlejpeg",
		"-")
	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	err := cmd.Run()
	if err != nil {
		return &buffer, err
	}

	return &buffer, nil
}
