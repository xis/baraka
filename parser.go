package baraka

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
)

// Parser is an interface which determine which method to use when parsing multipart files
type Parser interface {
	Parse(maxFileSize int64, maxFileCount int64, r *http.Request) (Saver, error)
}

// Options implements the Parser interface
// filter function runs when parsing the file
type Options struct {
	Filter func(data *multipart.Part) bool
}

// Parse @ parses with multipart.Reader()
func (options Options) Parse(maxFileSize int64, maxFileCount int64, r *http.Request) (Saver, error) {
	reader, err := r.MultipartReader()
	if err != nil {
		return nil, err
	}
	// reserve an additional 2 MB for non-file parts.
	maxFileSize += int64(2 << 20)
	var parts Parts
	for {
		part, err := reader.NextPart()
		if err != nil || err == io.EOF {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if options.Filter != nil {
			// execute filter function
			ok := options.Filter(part)
			if !ok {
				continue
			}
		}

		fileName := part.FileName()
		var b bytes.Buffer
		fh := &Header{
			Filename: fileName,
			Header:   part.Header,
		}
		n, err := io.CopyN(&b, part, maxFileSize+1)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n > maxFileSize {
			continue
		}
		fh.content = b.Bytes()
		fh.Size = int64(len(fh.content))
		parts.files = append(parts.files, fh)
		parts.len++
		if len := int64(parts.len); len == maxFileCount && len != 0 {
			break
		}
	}
	return &parts, nil
}
