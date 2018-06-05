package filesystem

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// File represents a file
type File struct {
	Name        string
	IsDirectory bool
}

// Filesystem represents a specific path in the underlying file system
type Filesystem interface {
	GetFiles(path string) ([]File, error)
	ReadJSONFile(path string, fileName string, content interface{}) error
	FileExists(path, fileName string) (bool, error)
}

// localFilesystem represents files on the local disk
type localFilesystem struct {
}

// New returns a new instance of Filesystem
func New() Filesystem {
	return localFilesystem{}
}

// GetFiles returns a list of file names
func (f localFilesystem) GetFiles(path string) ([]File, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return []File{}, err
	}

	fileNames := []File{}
	for _, file := range files {
		fileNames = append(fileNames, File{Name: file.Name(), IsDirectory: file.IsDir()})
	}
	return fileNames, nil
}

// ReadJSONFile reads a json file into content
func (f localFilesystem) ReadJSONFile(path string, fileName string, content interface{}) error {
	raw, err := ioutil.ReadFile(filepath.Join(path, fileName))

	if err != nil {
		return err
	}

	json.Unmarshal(raw, content)

	return nil
}

// FileExists returns true if the given path or file exists otherwise returns false
func (f localFilesystem) FileExists(path, fileName string) (bool, error) {
	_, err := os.Stat(filepath.Join(path, fileName))
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
