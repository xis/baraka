/*
	this file contains helper functions and raw values for tests
*/
package baraka

import (
	"net/http"
	"net/textproto"
	"strings"
)

const RawMultipartPlainText = `
--MyBoundary
Content-Disposition: form-data; name="filea"; filename="filea.txt"
Content-Type: text/plain

test file a
--MyBoundary
Content-Disposition: form-data; name="fileb"; filename="fileb.txt"
Content-Type: text/plain

test file b
--MyBoundary--
`

var PartPlainText = Part{
	Name: "Plain",
	Headers: textproto.MIMEHeader{
		"Content-Type": []string{"text/plain"},
	},
	Size:    15,
	Content: []byte("plain text file"),
}

func CreateHTTPRequest(raw string) *http.Request {
	b := strings.NewReader(strings.ReplaceAll(raw, "\n", "\r\n"))
	req, _ := http.NewRequest("POST", "http://test.com", b)
	req.Header = http.Header{"Content-Type": {`multipart/form-data; boundary="MyBoundary"`}}
	return req
}
