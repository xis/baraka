package baraka

import (
	"net/http"
	"strings"
	"testing"
)

type HeaderTest struct {
	header http.Header
	parser Parser
}

func TestStore(t *testing.T) {
	tests := []HeaderTest{
		{http.Header{"Content-Type": {`multipart/form-data; boundary="MyBoundary"`}}, WithMultipartReader{}},
		{http.Header{"Content-Type": {`text/plain`}}, WithMultipartReader{}},
		{http.Header{"Content-Type": {`multipart/form-data; boundary="MyBoundary"`}}, WithParseMultipartForm{}},
		{http.Header{"Content-Type": {`text/plain`}}, WithParseMultipartForm{}},
	}

	raw := `
--MyBoundary
Content-Disposition: form-data; name="filea"; filename="filea.txt"
Content-Type: text/plain

test file
--MyBoundary--
`
	for _, test := range tests {
		s, err := NewStorage("./", test.parser)
		if err != nil {
			t.Error(err)
		}
		b := strings.NewReader(strings.ReplaceAll(raw, "\n", "\r\n"))
		req, _ := http.NewRequest("POST", "http://localhost", b)
		req.Header = test.header

		p, err := s.Parse(req)
		if err != nil {
			// expects specified error if content type is text/plain
			if err == http.ErrNotMultipart && test.header["Content-Type"][0] == "text/plain" {
				continue
			}

			t.Fatal(err)
		}

		err = p.Store()
		if err != nil {
			t.Fatal(err)
		}
	}

}
