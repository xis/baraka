package baraka

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
)

const (
	defaultMaxFileSize   = 10 << 20
	defaultMaxParseCount = 20
	defaultMaxFileCount  = 0
)

var (
	ErrMaxFileCountExceeded = errors.New("max file count to save exceeded")
	ErrExtensionNotFound    = errors.New("can't detect file's extension")
)

// Parser contains parsing options and interfaces to do some other actions
type Parser struct {
	Options   ParserOptions
	Filter    Filter
	Inspector Inspector
}

// ParserOptions contains parser's options about parsing.
type ParserOptions struct {
	MaxFileSize   int
	MaxFileCount  int
	MaxParseCount int
}

// NewParser creates a new Parser.
func NewParser(options ParserOptions) *Parser {
	return &Parser{
		Options: options,
	}
}

// DefaultParser creates a new parser with the default settings
func DefaultParser() *Parser {
	options := ParserOptions{
		MaxFileSize:   defaultMaxFileSize,
		MaxFileCount:  defaultMaxFileCount,
		MaxParseCount: defaultMaxParseCount,
	}

	return NewParser(options)
}

// SetFilter sets the filter of the parser
func (parser *Parser) SetFilter(filter Filter) *Parser {
	parser.Filter = filter

	return parser
}

// SetInspector sets the inspector of the parser
func (parser *Parser) SetInspector(inspector Inspector) *Parser {
	parser.Inspector = inspector

	return parser
}

// Parse parses the http request with the multipart.Reader.
// reads parts inside the loop which iterates up to MaxParseCount at most.
// creates a []byte (buf) which gonna contain the part data.
func (parser *Parser) Parse(r *http.Request) (*Request, error) {
	reader, err := r.MultipartReader()
	if err != nil {
		return nil, err
	}

	parts := make(map[string][]*Part)
	request := NewRequest(parts)

process:
	for parseCount := 0; parseCount < parser.Options.MaxParseCount; parseCount++ {
		part, err := reader.NextPart()
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		data := bytes.NewBuffer(nil)
		for {
			if data.Len() > parser.Options.MaxFileSize {
				continue process
			}

			buf := make([]byte, 1024)
			n, err := part.Read(buf)
			if err != nil {
				if err == io.EOF {
					buf = buf[:n]
					data.Write(buf)

					break
				}

				return nil, err
			}

			data.Write(buf)
		}
		part.Close()

		p := Part{
			Name:    part.FileName(),
			Headers: part.Header,
			Size:    data.Len(),
			Content: data.Bytes(),
		}

		if parser.Inspector != nil {
			contentType := parser.Inspector.Inspect(data.Bytes())

			extensions, err := mime.ExtensionsByType(contentType)
			if err != nil {
				return nil, err
			}

			if len(extensions) == 0 {
				return nil, fmt.Errorf("%w, filename: %s", ErrExtensionNotFound, part.FileName())
			}

			p.Extension = extensions[0]
		}

		if parser.Filter != nil {
			ok := parser.Filter.Filter(&p)
			if !ok {
				continue
			}
		}

		request.parts[part.FormName()] = append(request.parts[part.FormName()], &p)
		if len := len(request.parts); len == parser.Options.MaxFileCount && len != 0 {
			return nil, ErrMaxFileCountExceeded
		}
	}

	return request, nil
}
