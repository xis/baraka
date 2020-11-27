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

// Store same with StoreWithout but uses default exclude parameter
// excludes json files by default from saving into disk
func (p *Process) Store(prefix string) error {
	return p.StoreWithout(prefix, "application/json")
}

// StoreWithout calls a function from Saver interface to save files
// you can give a variadic parameter for exclude content types
func (p *Process) StoreWithout(prefix string, excludedContentTypes ...string) error {
	return p.saver.Save(prefix, p.storage.path, excludedContentTypes...)
}

func (p *Process) GetBytes(excludedContentTypes ...string) ([][]byte, error) {
	return p.saver.SaveToBytes(excludedContentTypes...)
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
