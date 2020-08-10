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
	// len is the length of the files in process
	len int
}

// Store calls a function from Saver interface to save files
func (p *Process) Store(prefix string) error {
	err := p.saver.Save(prefix, p.storage.path)
	if err != nil {
		return err
	}
	return nil
}

// JSON @
func (p *Process) JSON() ([][]byte, error) {
	found, err := p.saver.(Marshaler).JSON()
	if err != nil {
		return nil, err
	}
	return found, nil
}

func (p *Process) Length() int {
	return p.saver.(Informer).Length()
}

func (p *Process) Filenames() []string {
	return p.saver.(Informer).Filenames()
}

func (p *Process) ContentTypes() []string {
	return p.saver.(Informer).ContentTypes()
}
