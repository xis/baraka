package storage

import (
	"os"
	"path/filepath"
)

// FilesystemStorage is the default store to save parts
type FilesystemStorage struct {
	Path string
}

// NewFilesystemStorage creates a new FilesystemStorage
func NewFilesystemStorage(path string) FilesystemStorage {
	return FilesystemStorage{
		Path: path,
	}
}

// Store is a method for storing bytes into disk.
func (s FilesystemStorage) Store(path string, data []byte) error {
	path = filepath.Join(s.Path, path)
	dirName := filepath.Dir(path)

	if !isDir(dirName) {
		err := os.MkdirAll(dirName, os.ModeSticky|os.ModePerm)
		if err != nil {
			return err
		}
	}

	return os.WriteFile(path, data, 0666)
}

// isDir checks if the path is a directory and exists
func isDir(path string) bool {
	f, e := os.Stat(path)

	if e != nil {
		return false
	}

	return f.IsDir()
}
