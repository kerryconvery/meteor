package mocks

import (
	"bytes"
	"meteor/filesystem"
)

// Filesystem represents a mock of the filesystem struct
type Filesystem struct {
}

// GetFiles returns a list of file names
func (f Filesystem) GetFiles(path string) ([]filesystem.File, error) {
	return nil, nil
}

// ReadJSONFile reads a json file into content
func (f Filesystem) ReadJSONFile(path string, fileName string, content interface{}) error {
	return nil
}

// FileExists returns true if the given path or file exists otherwise returns false
func (f Filesystem) FileExists(path, fileName string) (bool, error) {
	return false, nil
}

// ReadFile reads a file off disk and returns a buffer
func (f Filesystem) ReadFile(path, fileName string) (*bytes.Buffer, error) {
	return nil, nil
}

// WriteFile write the content of the buffer to a file.  It will create the file if it doesn't exists
func (f Filesystem) WriteFile(path, fileName string, buffer bytes.Buffer) (int, error) {
	return 0, nil
}
