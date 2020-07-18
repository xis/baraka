package baraka

import (
	"io"
	"mime/multipart"
	"os"
	"syscall"
	"testing"
)

func TestSave_MultipartForm(t *testing.T) {
	s, err := NewStorage("./", WithParseMultipartForm{})
	if err != nil {
		t.Error(err)
	}
	req := CreateRequest(RawMultipartPlainText)
	p, err := s.Parse(req)
	if err != nil {
		t.Error(err)
	}

	form := p.saver.(MultipartForm)

	form.filter = func(data multipart.File) bool {
		return false
	}

	path := "./"
	err = form.Save(path)
	if err != nil && err != io.EOF {
		t.Error(err)
	}
	form.filter = func(data multipart.File) bool {
		return true
	}
	path = ""
	for _, form := range form.files.File {
		for _, file := range form {
			file.Filename = ""
		}
	}
	err = form.Save(path)
	if err != nil && err.(*os.PathError).Err != syscall.ENOENT {
		t.Error(err)
	}
}

func TestSave_MultipartParts(t *testing.T) {
	s, err := NewStorage("./", WithMultipartReader{})
	if err != nil {
		t.Error(err)
	}
	req := CreateRequest(RawMultipartPlainText)
	p, err := s.Parse(req)
	if err != nil {
		t.Error(err)
	}
	mp := p.saver.(MultipartParts)
	err = mp.Save("./")
	if err != nil {
		t.Error(err)
	}
	mp.files[0].Filename = ""
	err = mp.Save("")
	if err != nil && err.(*os.PathError).Err != syscall.ENOENT {
		t.Error(err)
	}
}
