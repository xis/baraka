package baraka

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
)

// Parser is an interface which determine which method to use when parsing multipart files
type Parser interface {
	Parse(r *http.Request) (Saver, error)
}

// WithParseMultipartForm implements the Parser interface
// filter function runs when saving the file
type WithParseMultipartForm struct {
	Filter func(data *multipart.File) bool
}

// WithMultipartReader implements the Parser interface
// filter function runs when parsing the file
type WithMultipartReader struct {
	Filter func(data *multipart.Part) bool
}

// Parse @ parses with request.ParseMultiparmForm()
func (parser WithParseMultipartForm) Parse(r *http.Request) (Saver, error) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return nil, err
	}
	var files MultipartForm
	files.files = *r.MultipartForm
	files.filter = parser.Filter
	return files, nil
}

// Parse @ parses with multipart.Reader()
func (parser WithMultipartReader) Parse(r *http.Request) (Saver, error) {
	reader, _ := r.MultipartReader()
	var mp MultipartParts
	var maxMemory int64 = 32 << 20
	// Reserve an additional 10 MB for non-file parts.
	for {
		part, err := reader.NextPart()
		if err != nil {
			break
		}
		fileName := part.FileName()

		if parser.Filter != nil {
			// execute filter function
			ok := parser.Filter(part)
			if !ok {
				continue
			}
		}

		var b bytes.Buffer
		fh := &Header{
			Filename: fileName,
			Header:   part.Header,
		}
		_, err = io.CopyN(&b, part, maxMemory+1)
		if err != nil && err != io.EOF {
			return nil, err
		}
		fh.content = b.Bytes()
		fh.Size = int64(len(fh.content))
		mp.files = append(mp.files, fh)
	}
	return mp, nil
}
