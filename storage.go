package baraka

import (
	"net/http"
	"path/filepath"
)

// Storage is a struct that determines which path to size and which parsing method to parse.
// Create this struct with NewStorage function.
type Storage struct {
	path   string
	parser Parser
}

// NewStorage creates a new Storage struct.
func NewStorage(path string, parser Parser) (*Storage, error) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	storage := &Storage{
		path:   absolutePath,
		parser: parser,
	}
	return storage, nil
}

// Parse is a function for parsing multipart form files.
// calls Parse from Parser interface and gets Saver,
// creates new process and attaches Saver, *Storage and *http.Request.
// returns *Process for user.
func (s *Storage) Parse(r *http.Request) (*Process, error) {
	saver, err := s.parser.Parse(r)
	if err != nil {
		return nil, err
	}

	p := &Process{
		req:       r,
		storage:   s,
		saver:     saver,
		marshaler: nil,
	}
	return p, nil
}
