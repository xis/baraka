package baraka

import (
	"net/http"
	"strings"
	"testing"
)

func TestParserParse(t *testing.T) {
	raw := `
--MyBoundary
Content-Disposition: form-data; name="filea"; filename="filea.txt"
Content-Type: text/plain

test file
--MyBoundary--
`

	parsers := []Parser{
		WithMultipartReader{},
		WithParseMultipartForm{},
	}

	for _, parser := range parsers {
		s, err := NewStorage("./", parser)
		if err != nil {
			t.Error(err)
		}
		b := strings.NewReader(strings.ReplaceAll(raw, "\n", "\r\n"))
		req, _ := http.NewRequest("POST", "http://localhost", b)
		req.Header = http.Header{"Content-Type": {`multipart/form-data; boundary="MyBoundary"`}}

		_, err = s.Parse(req)
		if err != nil {
			t.Error(err)
		}
	}

}
