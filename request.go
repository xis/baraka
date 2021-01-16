package baraka

import (
	"mime"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Request implements the Processor interface.
// contains an array of parts.
// parser.Parse() returns Request as Processor.
type Request struct {
	parts []*Part
}

// NewRequest creates a new Request with parts inside
func NewRequest(parts ...*Part) *Request {
	return &Request{
		parts,
	}
}

// Save is a method for saving parts into disk.
func (s *Request) Save(prefix string, fileParentDirPath string, excludedContentTypes ...string) error {
	if !isDir(fileParentDirPath) {
		err := os.MkdirAll(fileParentDirPath, os.ModeSticky|os.ModePerm)
		if err != nil {
			return err
		}
	}
	for key := range s.parts {
		file := s.parts[key]
		if len(excludedContentTypes) != 0 {
			excluded := isExcluded(file.Headers.Get("Content-Type"), excludedContentTypes...)
			if excluded {
				continue
			}
		}
		extension := filepath.Ext(file.Name)
		out, err := os.Create(filepath.Join(fileParentDirPath, prefix+strconv.Itoa(key)+extension))
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

// Content method gives you []*Part, so you can access all the data of the parts.
func (s *Request) Content() []*Part {
	return s.parts
}

// Length returns length of the parts
func (request *Request) Length() int {
	return len(request.parts)
}

// Filenames returns filenames of the parts
func (request *Request) Filenames() []string {
	filenames := make([]string, len(request.parts))
	for k, v := range request.parts {
		filenames[k] = v.Name
	}
	return filenames
}

// ContentTypes returns content types of the parts
func (request *Request) ContentTypes() []string {
	contentTypes := make([]string, len(request.parts))
	for k, v := range request.parts {
		contentTypes[k] = v.Headers.Get("Content-Type")
	}
	return contentTypes
}

// GetJSON returns json strings of the application/json parts separately
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

func isDir(path string) bool {
	f, e := os.Stat(path)
	if e != nil {
		return false
	}
	return f.IsDir()
}
