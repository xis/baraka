package baraka

import (
	"os"
	"syscall"
	"testing"
)

func TestSaver_Parts(t *testing.T) {
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
	files, err := mp.SaveToBytes()
	if err != nil {
		t.Error(err)
	}
	if string(files[0]) != "test file a" && string(files[1]) != "test files b" {
		t.Error("files not read properly")
	}
	if err != nil {
		t.Error(err)
	}
	mp.files[0].Filename = ""
	err = mp.Save("test", "")
	if err != nil && err.(*os.PathError).Err != syscall.ENOENT {
		t.Error(err)
	}
	err = mp.Save("test", "./", "text/plain")
	if err != nil {
		t.Error(err)
	}
}
