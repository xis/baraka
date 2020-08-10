package baraka

import (
	"mime"
)

// Marshaler is an interface that contains json marshal function
type Marshaler interface {
	JSON() ([][]byte, error)
}

// JSON @
func (parts *Parts) JSON() ([][]byte, error) {
	jsons := [][]byte{}
	for _, file := range parts.files {
		if file.Size == 0 {
			continue
		}
		media, _, err := mime.ParseMediaType(file.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		if media == "application/json" {
			jsons = append(jsons, file.content)
		}
	}
	return jsons, nil
}
