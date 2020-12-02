package baraka

import (
	"mime"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Request implements the Saver interface
type Request struct {
	parts []*Part
}

func NewRequest(parts ...*Part) *Request {
	return &Request{
		parts,
	}
}

// Save is a method for saving multipart files.
// this method saves Parts data.
func (s *Request) Save(prefix string, path string, excludedContentTypes ...string) error {
	for key := range s.parts {
		file := s.parts[key]
		if len(excludedContentTypes) != 0 {
			excluded := isExcluded(file.Headers.Get("Content-Type"), excludedContentTypes...)
			if excluded {
				continue
			}
		}
		extension := filepath.Ext(file.Name)
		out, err := os.Create(path + prefix + strconv.Itoa(key) + extension)
		defer out.Close()
		if err != nil {
			return err
		}
		_, err = out.Write(file.Content)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Request) Content() []*Part {
	return s.parts
}

// Length returns total count of files
func (request *Request) Length() int {
	return len(request.parts)
}

// Filenames returns names of files
func (request *Request) Filenames() []string {
	filenames := make([]string, len(request.parts))
	for k, v := range request.parts {
		filenames[k] = v.Name
	}
	return filenames
}

// ContentTypes returns content types of files
func (request *Request) ContentTypes() []string {
	contentTypes := make([]string, len(request.parts))
	for k, v := range request.parts {
		contentTypes[k] = v.Headers.Get("Content-Type")
	}
	return contentTypes
}

// JSON returns bytes of json files separately
func (request *Request) GetJSON() ([]string, error) {
	jsons := []string{}
	for _, file := range request.parts {
		if file.Size == 0 {
			continue
		}
		media, _, err := mime.ParseMediaType(file.Headers.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		if media == "application/json" {
			jsons = append(jsons, string(file.Content))
		}
	}
	return jsons, nil
}

// helper functions

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
