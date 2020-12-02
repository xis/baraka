package baraka

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
)

// defaultMaxFileSize 10 mb per file
const defaultMaxFileSize = 10 << 20

// Parser is an interface which determine which method to use when parsing multipart files
type Parser interface {
	Parse(r *http.Request) (Processor, error)
}

type parser struct {
	Options ParserOptions
}

// Options implements the Parser interface
// filter function runs when parsing the file
type ParserOptions struct {
	MaxFileSize  int
	MaxFileCount int
	Filter       func(data *multipart.Part) bool
}

func NewParser(options ParserOptions) Parser {
	if options.MaxFileSize == 0 {
		options.MaxFileSize = defaultMaxFileSize
	}
	return parser{
		options,
	}
}

// Parse @ parses with multipart.Reader()
func (parser parser) Parse(r *http.Request) (Processor, error) {
	reader, err := r.MultipartReader()
	if err != nil {
		return nil, err
	}

	maxFileSize := parser.Options.MaxFileSize
	// reserve an additional 2 MB for non-file parts.
	maxFileSize += 2 << 20

	parts := make([]*Part, 0, parser.Options.MaxFileCount)
	request := NewRequest(parts...)
	for {
		part, err := reader.NextPart()
		if err != nil || err == io.EOF {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if parser.Options.Filter != nil {
			// execute filter function
			ok := parser.Options.Filter(part)
			if !ok {
				continue
			}
		}
		var b bytes.Buffer
		n, err := io.CopyN(&b, part, int64(maxFileSize+1))
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n > int64(maxFileSize) {
			continue
		}
		p := NewPart(part.FileName(), &part.Header, b.Len(), b.Bytes())
		request.parts = append(request.parts, &p)
		if len := len(request.parts); len == parser.Options.MaxFileCount && len != 0 {
			break
		}
	}
	return request, nil
}
