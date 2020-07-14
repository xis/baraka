package baraka

import (
	"mime/multipart"
	"net/http"
	"strings"
	"testing"
)

func TestStore(t *testing.T) {
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

	err = p.Store()
	if err != nil {
		t.Error(err)
	}
}
