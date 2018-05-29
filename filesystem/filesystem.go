package filesystem

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

// Filesystem represents a specific path in the underlying file system
type Filesystem struct {
	Path string
}

// New returns a new instance of Filesystem
func New(path string) Filesystem {
	return Filesystem{path}
}

// GetFiles returns a list of file names
func (f Filesystem) GetFiles() ([]string, error) {
	files, err := ioutil.ReadDir(f.Path)
	if err != nil {
		return []string{}, err
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames, nil
}

// ReadJSONFile reads a json file into content
func (f Filesystem) ReadJSONFile(fileName string, content interface{}) error {
	raw, err := ioutil.ReadFile(filepath.Join(f.Path, fileName))

	if err != nil {
		return err
	}

	json.Unmarshal(raw, content)

	return nil
}
