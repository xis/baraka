/*
	this file contains helper functions and raw values for tests
*/
package baraka

import (
	"mime/multipart"
	"net/http"
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

const RawMultipartWithJSON = `
--MyBoundary
Content-Disposition: form-data; name="filea"; filename="filea.txt"
Content-Type: text/plain

test file
--MyBoundary
Content-Disposition: form-data; name="jsonfile"; filename="jsonfile.json"
Content-Type: @contentType

@data
--MyBoundary--
`

func CreateHTTPRequest(raw string) *http.Request {
	b := strings.NewReader(strings.ReplaceAll(raw, "\n", "\r\n"))
	req, _ := http.NewRequest("POST", "http://localhost", b)
	req.Header = http.Header{"Content-Type": {`multipart/form-data; boundary="MyBoundary"`}}
	return req
}

func FilterJPEG() func(*multipart.Part) bool {
	return func(data *multipart.Part) bool {
		buf := make([]byte, 512)
		data.Read(buf)
		media := http.DetectContentType(buf)
		if media == "image/jpeg" {
			return true
		}
		return false
	}
}
