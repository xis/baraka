package baraka

import (
	"mime/multipart"
	"testing"
)

func TestNewStorage(t *testing.T) {
	s, err := NewStorage("./", WithMultipartReader{})
	if err != nil {
		t.Error(err)
	}
	if s.path == "" && s.parser != nil {
		t.Error("given parameters not set to storage struct")
	}
}

func TestStorageParse(t *testing.T) {
	s, err := NewStorage("./", WithMultipartReader{
		Filter: func(data *multipart.Part) bool {
			return true
		},
	})
	if err != nil {
		t.Error(err)
	}
	req := CreateRequest(RawMultipartPlainText)

	p, err := s.Parse(req)
	if err != nil {
		t.Error(err)
	}
	err = p.saver.Save(s.path)
	if err != nil {
		t.Error(err)
	}

}
