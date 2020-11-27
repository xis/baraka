package baraka

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Saver is a interface for determine which data type to store
// Parse interface functions returns this interface
type Saver interface {
	Save(prefix string, path string, excludedContentTypes ...string) error
	SaveToBytes(excludedContentTypes ...string) ([][]byte, error)
}

// Parts implements the Saver interface
type Parts struct {
	files []*Header
	len   int
}

// Save is a method for saving multipart files.
// this method saves Parts data.
func (s *Parts) Save(prefix string, path string, excludedContentTypes ...string) error {
	for key := range s.files {
		fileHeader := s.files[key]
		if len(excludedContentTypes) != 0 {
			excluded := isExcluded(fileHeader.Header.Get("Content-Type"), excludedContentTypes...)
			if excluded {
				continue
			}
		}
		extension := filepath.Ext(fileHeader.Filename)
		file := fileHeader.Open()
		out, err := os.Create(path + prefix + strconv.Itoa(key) + extension)
		defer out.Close()
		if err != nil {
			return err
		}
		_, err = io.Copy(out, file)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Parts) SaveToBytes(excludedContentTypes ...string) ([][]byte, error) {
	files := make([][]byte, len(s.files))
	for key := range s.files {
		fileHeader := s.files[key]
		if len(excludedContentTypes) != 0 {
			excluded := isExcluded(fileHeader.Header.Get("Content-Type"), excludedContentTypes...)
			if excluded {
				continue
			}
		}
		reader := fileHeader.Open()
		file := new(bytes.Buffer)
		file.ReadFrom(reader)
		files[key] = file.Bytes()
	}
	return files, nil
}

// helper functions for saver

// isExcluded returns that if file's content type is on excluded content type list or not.
func isExcluded(contentType string, excludedContentTypes ...string) bool {
	for key := range excludedContentTypes {
		i := strings.Index(contentType, ";")
		if i == -1 {
			i = len(contentType)
		}
		contentType = strings.TrimSpace(strings.ToLower(contentType[0:i]))
		if contentType == excludedContentTypes[key] {
			return true
		}
	}
	return false
}
