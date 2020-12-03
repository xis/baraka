package baraka

import (
	"testing"
)

func TestInformer(t *testing.T) {
	s, err := NewStorage("./", Options{})
	if err != nil {
		t.Error(err)
	}
	req := CreateRequest(RawMultipartPlainText)
	p, err := s.Parse(req)
	if err != nil {
		t.Error(err)
	}
	if p.Length() != 2 {
		t.Error("files length must be 2")
	}
	filenames := p.Filenames()
	if filenames[0] != "filea.txt" || filenames[1] != "fileb.txt" {
		t.Error("filenames not extracted properly")
	}

	contentTypes := p.ContentTypes()
	if contentTypes[0] != "text/plain" || contentTypes[1] != "text/plain" {
		t.Error("content types not extracted properly")
	}
}
