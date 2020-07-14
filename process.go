package baraka

import (
	"net/http"
)

// Process @ process is returned after Storage.Parse() function.Process
// stores request for accessing MultipartForm
type Process struct {
	req     *http.Request
	storage *Storage
	saver   Saver
}

// Store calls a function from Saver interface to save files
func (p *Process) Store() error {
	err := p.saver.Save(p.storage.path)
	if err != nil {
		return err
	}
	return nil
}
