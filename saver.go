package baraka

import (
	"io"
	"os"
	"strconv"
)

// Saver is a interface for determine which data type to store
// Parse interface functions returns this interface
type Saver interface {
	Save(prefix string, path string) error
}

// Parts implements the Saver interface
type Parts struct {
	files []*Header
	len   int
}

// Save is a method for saving multipart files.
// this method saves Parts data.
func (s *Parts) Save(prefix string, path string) error {
	for i, header := range s.files {
		f := header.Open()
		out, err := os.Create(path + prefix + strconv.Itoa(i))
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
