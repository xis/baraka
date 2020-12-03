package baraka

import (
	"net/http"
)

// Storage is a struct that determines which path to save and which parsing method to parse.
// Create this struct with NewStorage function.
type Storage struct {
	path   string
	parser Parser
}

// defaultMaxMemory 10 mb per file
const defaultMaxMemory = 10 << 20

// NewStorage creates a new Storage struct.
func NewStorage(path string, parser Parser) (*Storage, error) {
	storage := &Storage{
		path:   path,
		parser: parser,
	}
	return storage, nil
}

// Parse same with ParseButMax but uses defaultMaxMemory 10MB, and unlimited file count.
func (s *Storage) Parse(r *http.Request) (*Process, error) {
	p, err := s.ParseButMax(defaultMaxMemory, 0, r)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// ParseButMax is a function for parsing multipart form files.
// calls Parse from Parser interface and gets Saver,
// creates new process and attaches Saver, *Storage and *http.Request to it.
// returns *Process for user.
func (s *Storage) ParseButMax(maxFileSize int64, maxFileCount int64, r *http.Request) (*Process, error) {
	saver, err := s.parser.Parse(maxFileSize, maxFileCount, r)
	if err != nil {
		return nil, err
	}
	p := &Process{
		req:     r,
		storage: s,
		saver:   saver,
	}
	return p, nil
}
