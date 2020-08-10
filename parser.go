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

// Options implements the Parser interface
// filter function runs when parsing the file
type Options struct {
	Filter func(data *multipart.Part) bool
}

// Parse @ parses with multipart.Reader()
func (options Options) Parse(r *http.Request) (Saver, error) {
	reader, err := r.MultipartReader()
	if err != nil {
		return nil, err
	}
	var parts Parts
	var maxMemory int64 = 32 << 20
	// Reserve an additional 10 MB for non-file parts.
	for {
		part, err := reader.NextPart()
		if err != nil || err == io.EOF {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		fileName := part.FileName()

		if options.Filter != nil {
			// execute filter function
			ok := options.Filter(part)
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
		parts.files = append(parts.files, fh)
		parts.len++
	}
	return &parts, nil
}
