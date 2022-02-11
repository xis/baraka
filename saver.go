package baraka

import (
	"os"
	"path/filepath"
)

// Saver is a interface that wraps the Save function
type Saver interface {
	Save(path string, filename string, part *Part) error
}

// FileSystemStore is the default store to save parts
type FileSystemStore struct {
	Path string
}

// NewFileSystemStore creates a new FileSystemStore
func NewFileSystemStore(path string) FileSystemStore {
	return FileSystemStore{
		Path: path,
	}
}

// Save is a method for saving parts into disk.
func (s FileSystemStore) Save(path string, filename string, part *Part) error {
	if !isDir(path) {
		err := os.MkdirAll(path, os.ModeSticky|os.ModePerm)
		if err != nil {
			return err
		}
	}

	extension := part.Extension
	if part.Extension == "" {
		extension = filepath.Ext(part.Name)
	}

	out, err := os.Create(filepath.Join(path, filename+extension))
	if err != nil {
		return err
	}

	_, err = out.Write(part.Content)
	if err != nil {
		return err
	}

	err = out.Close()
	if err != nil {
		return err
	}

	return nil
}

// helper functions

// isDir checks if the path is a directory and exists
func isDir(path string) bool {
	f, e := os.Stat(path)

	if e != nil {
		return false
	}

	return f.IsDir()
}
