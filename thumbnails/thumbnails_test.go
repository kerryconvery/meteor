package thumbnails

import (
	"bytes"
	"errors"
	"meteor/tests"
	"testing"
)

type mockImageSource struct {
	defaultError error
}

const noImage = "No Image"
const generatedImage = "Generated Image"
const defaultImage = "Default Image"
const existingImage = "Existing Image"

func (p mockImageSource) generateVideoThumbnail(filename string) (*bytes.Buffer, error) {
	if filename == "valid_path\\gen_image" {
		return bytes.NewBufferString(generatedImage), nil
	}
	return bytes.NewBufferString(noImage), errors.New("invalid file")
}

func (p mockImageSource) getDefaultImage() (*bytes.Buffer, error) {
	if p.defaultError == nil {
		return bytes.NewBufferString(defaultImage), nil
	}
	return bytes.NewBufferString(noImage), p.defaultError
}

func (p mockImageSource) getExistingThumbnail(filename string) (*bytes.Buffer, error) {
	if filename == "existing_image" {
		return bytes.NewBufferString(existingImage), nil
	}
	return bytes.NewBufferString(noImage), errors.New("Existing image not found")
}
func TestGetThumbnailGenerate(t *testing.T) {
	provider := thumbnailProvider{imageSource: mockImageSource{}}

	thumbnail, err := provider.GetThumbnail("valid_path", "gen_image")

	tests.ExpectNoError(err, t)

	imageStr := thumbnail.String()

	if imageStr != generatedImage {
		t.Errorf("Expected text '%s' but got %s", generatedImage, imageStr)
	}
}

func TestGetThumbnailExisting(t *testing.T) {
	provider := thumbnailProvider{imageSource: mockImageSource{}}

	thumbnail, err := provider.GetThumbnail("valid_path", "existing_image")

	tests.ExpectNoError(err, t)

	imageStr := thumbnail.String()

	if imageStr != existingImage {
		t.Errorf("Expected text '%s' but got %s", existingImage, imageStr)
	}
}

func TestGetThumbnailDefault(t *testing.T) {
	provider := thumbnailProvider{imageSource: mockImageSource{}}

	thumbnail, err := provider.GetThumbnail("valid_path", "gen_failed")

	tests.ExpectNoError(err, t)

	imageStr := thumbnail.String()

	if imageStr != defaultImage {
		t.Errorf("Expected text '%s' but got %s", defaultImage, imageStr)
	}
}
func TestGetThumbnailError(t *testing.T) {
	provider := thumbnailProvider{imageSource: mockImageSource{errors.New("Could not read default image")}}

	_, err := provider.GetThumbnail("valid_path", "gen_failed")

	tests.ExpectError(err, t)
}
