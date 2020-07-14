package baraka

import (
	"mime/multipart"
	"net/http"
	"strings"
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
	raw := `
--MyBoundary
Content-Disposition: form-data; name="filea"; filename="filea.txt"
Content-Type: text/plain

test file
--MyBoundary--
`
	s, err := NewStorage("./", WithMultipartReader{
		Filter: func(data *multipart.Part) bool {
			return true
		},
	})
	if err != nil {
		t.Error(err)
	}
	b := strings.NewReader(strings.ReplaceAll(raw, "\n", "\r\n"))
	req, _ := http.NewRequest("POST", "http://localhost", b)
	req.Header = http.Header{"Content-Type": {`multipart/form-data; boundary="MyBoundary"`}}

	p, err := s.Parse(req)
	if err != nil {
		t.Error(err)
	}
	err = p.saver.Save(s.path)
	if err != nil {
		t.Error(err)
	}

}
