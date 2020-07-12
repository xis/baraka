package baraka

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// Saver is a interface for determine which data type to store
// Parse interface functions returns this interface
type Saver interface {
	Save(path string) error
}

// MultipartForm implements the Saver interface
type MultipartForm struct {
	files multipart.Form
	// MultipartForm filter function runs when saving file
	filter func(data *multipart.File) bool
}

// MultipartParts implements the Saver interface
type MultipartParts struct {
	files []*Header
}

// Save is a method for saving multipart files.
// this method saves multipart.Form data
// will be used when WithParseMultipartForm passed to Storage
func (s MultipartForm) Save(path string) error {
	for _, multiparts := range s.files.File {
		for _, multipart := range multiparts {
			file, err := multipart.Open()
			defer file.Close()
			if err != nil {
				return err
			}
			// execute filter function
			okay := s.filter(&file)
			if !okay {
				continue
			}
			out, err := os.Create(path + multipart.Filename)
			defer out.Close()
			if err != nil {
				return err
			}
			_, err = io.Copy(out, file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Save is a method for saving multipart files.
// this method saves MultipartParts data.
// will be used when WithMultipartReader passed to Storage
func (s MultipartParts) Save(path string) error {
	for _, header := range s.files {
		f := header.Open()
		path, _ := filepath.Abs("./test")
		out, err := os.Create(path + header.Filename)
		defer out.Close()
		if err != nil {
			return err
		}
		_, err = io.Copy(out, f)
		if err != nil {
			return err
		}
	}
	return nil
}
