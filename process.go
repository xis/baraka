package baraka

import (
	"net/http"
)

// Process @ process is returned after Storage.Parse() function.Process
// stores request for accessing MultipartForm
type Process struct {
	req       *http.Request
	storage   *Storage
	saver     Saver
	marshaler Marshaler
}

// Store calls a function from Saver interface to save files
func (p *Process) Store() error {
	err := p.saver.Save(p.storage.path)
	if err != nil {
		return err
	}
	return nil
}

// JSON @
func (p *Process) JSON() ([][]byte, error) {
	p.marshaler = p.saver.(Marshaler)
	found, err := p.marshaler.JSON()
	if err != nil {
		return nil, err
	}
	return found, nil
}
