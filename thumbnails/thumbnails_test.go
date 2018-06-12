package thumbnails

import (
	"bytes"
	"errors"
	"fmt"
	"meteor/tests"
	"testing"
)

type mockImageSource struct {
	file         string
	defaultError error
}

const noImage = "No Image"
const generatedImage = "Generated Image"
const defaultImage = "Default Image"
const existingImage = "Existing Image"

func (p mockImageSource) generateVideoThumbnail(filename string) (*bytes.Buffer, error) {
	fmt.Println(filename)
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
	if filename == "existing_image.jpg" {
		return bytes.NewBufferString(existingImage), nil
	}
	return bytes.NewBufferString(noImage), errors.New("Existing image not found")
}

func (p *mockImageSource) addThumbnail(filename string, image *bytes.Buffer) (int, error) {
	if p.defaultError == nil {
		p.file = filename
		return 100, nil
	}
	return 0, p.defaultError
}
func TestGetThumbnailGenerate(t *testing.T) {
	imageSource := mockImageSource{}
	provider := thumbnailProvider{imageSource: &imageSource}

	thumbnail, err := provider.GetThumbnail("valid_path\\gen_image")

	tests.ExpectNoError(err, t)

	imageStr := thumbnail.String()

	if imageStr != generatedImage {
		t.Errorf("Expected text '%s' but got %s", generatedImage, imageStr)
	}

	if imageSource.file != "gen_image.jpg" {
		t.Errorf("Expected file gen_image.jpg but got %s", imageSource)
	}
}

func TestGetThumbnailAddFileError(t *testing.T) {
	defaultError := errors.New("Error adding new file")
	provider := thumbnailProvider{imageSource: &mockImageSource{defaultError: defaultError}}

	_, err := provider.GetThumbnail("valid_path\\gen_image")

	tests.ExpectNoError(err, t)
}

func TestGetThumbnailExisting(t *testing.T) {
	provider := thumbnailProvider{imageSource: &mockImageSource{}}

	thumbnail, err := provider.GetThumbnail("valid_path\\existing_image")

	tests.ExpectNoError(err, t)

	imageStr := thumbnail.String()

	if imageStr != existingImage {
		t.Errorf("Expected text '%s' but got %s", existingImage, imageStr)
	}
}
func TestGetThumbnailDefault(t *testing.T) {
	provider := thumbnailProvider{imageSource: &mockImageSource{}}

	thumbnail, err := provider.GetThumbnail("valid_path\\gen_failed")

	tests.ExpectNoError(err, t)

	imageStr := thumbnail.String()

	if imageStr != defaultImage {
		t.Errorf("Expected text '%s' but got %s", defaultImage, imageStr)
	}
}
func TestGetThumbnailError(t *testing.T) {
	provider := thumbnailProvider{
		imageSource: &mockImageSource{
			defaultError: errors.New("Could not read default image"),
		},
	}

	_, err := provider.GetThumbnail("valid_path\\gen_failed")

	tests.ExpectError(err, t)
}
