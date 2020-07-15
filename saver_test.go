package baraka

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
)

var raw string = `
--MyBoundary
Content-Disposition: form-data; name="filea"; filename="filea.txt"
Content-Type: text/plain

test file
--MyBoundary--
`

func TestSave_MultipartForm(t *testing.T) {
	path, _ := filepath.Abs("./test_form_")
	b := strings.NewReader(strings.ReplaceAll(raw, "\n", "\r\n"))
	reader := multipart.NewReader(b, "MyBoundary")
	createdForm, err := reader.ReadForm(32 << 20)
	if err != nil {
		t.Error(err)
	}

	form := MultipartForm{
		files: *createdForm,
		filter: func(data multipart.File) bool {
			return true
		},
	}
	err = form.Save(path)
	if err != nil && err != io.EOF {
		t.Error(err)
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
	path, _ := filepath.Abs("./test_part_")
	b := strings.NewReader(strings.ReplaceAll(raw, "\n", "\r\n"))
	reader := multipart.NewReader(b, "MyBoundary")
	var mp MultipartParts
	for {
		part, err := reader.NextPart()
		if err != nil {
			break
		}
		var b bytes.Buffer
		fh := &Header{
			Filename: part.FileName(),
			Header:   part.Header,
		}
		var maxMemory int64 = 32 << 20
		_, err = io.CopyN(&b, part, maxMemory+1)
		if err != nil && err != io.EOF {
			t.Error(err)
		}
		fh.content = b.Bytes()
		fh.Size = int64(len(fh.content))
		mp.files = append(mp.files, fh)
	}
	err := mp.Save(path)
	if err != nil {
		t.Error(err)
	}

	mp.files[0].Filename = ""
	err = mp.Save("")
	if err != nil && err.(*os.PathError).Err != syscall.ENOENT {
		t.Error(err)
	}
}
