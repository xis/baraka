package baraka

import (
	"os"
	"syscall"
	"testing"
)

func TestSave_MultipartParts(t *testing.T) {
	s, err := NewStorage("./", Options{})
	if err != nil {
		t.Error(err)
	}
	req := CreateRequest(RawMultipartPlainText)
	p, err := s.Parse(req)
	if err != nil {
		t.Error(err)
	}
	mp := p.saver.(*Parts)
	err = mp.Save("test", "./")
	if err != nil {
		t.Error(err)
	}
	mp.files[0].Filename = ""
	err = mp.Save("test", "")
	if err != nil && err.(*os.PathError).Err != syscall.ENOENT {
		t.Error(err)
	}
}
