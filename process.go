package baraka

import (
	"net/http"
)

// Process @ process is returned after Storage.Parse() function.
// stores request for accessing MultipartForm
type Process struct {
	req     *http.Request
	storage *Storage
	saver   Saver
}

// Store calls a function from Saver interface to save files
func (p *Process) Store(prefix string) error {
	err := p.saver.Save(prefix, p.storage.path)
	if err != nil {
		return err
	}
	return nil
}

// JSON returns bytes of json files separately
func (p *Process) JSON() ([][]byte, error) {
	found, err := p.saver.(Marshaler).JSON()
	if err != nil {
		return nil, err
	}
	return found, nil
}

// Length returns total count of files
func (p *Process) Length() int {
	return p.saver.(Informer).Length()
}

// Filenames returns names of files
func (p *Process) Filenames() []string {
	return p.saver.(Informer).Filenames()
}

// ContentTypes returns content types of files
func (p *Process) ContentTypes() []string {
	return p.saver.(Informer).ContentTypes()
}
