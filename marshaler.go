package baraka

import (
	"mime"
)

// Marshaler is an interface that contains json marshal function
type Marshaler interface {
	JSON() ([][]byte, error)
}

// JSON @
func (m MultipartForm) JSON() ([][]byte, error) {
	jsons := [][]byte{}
	for _, multiparts := range m.files.File {
		for _, file := range multiparts {
			if file.Size == 0 {
				continue
			}
			media, _, err := mime.ParseMediaType(file.Header.Get("Content-Type"))
			if err != nil {
				return nil, err
			}
			if media == "application/json" {
				f, err := file.Open()
				defer f.Close()
				if err != nil {
					return nil, err
				}
				b := make([]byte, file.Size)
				_, err = f.Read(b)
				if err != nil {
					return nil, err
				}
				jsons = append(jsons, b)
			}
		}
	}
	return jsons, nil
}

// JSON @
func (m MultipartParts) JSON() ([][]byte, error) {
	jsons := [][]byte{}
	for _, file := range m.files {
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
