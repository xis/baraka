package baraka

import (
	"bytes"
	"io"
	"mime"
	"net/http"

	"github.com/pkg/errors"
)

const defaultMaxFileSize = 10 << 20
const defaultMaxParseCount = 20
const defaultMaxFileCount = 0
const defaultMaxAvailableSlice = 2

var errMaxFileCountExceeded = errors.New("max file count to save exceeded")
var errExtensionNotFound = errors.New("can't detect file's extension")

// Parser contains parsing options and interfaces to do some other actions
type Parser struct {
	Options   ParserOptions
	Filter    Filter
	Inspector Inspector
}

// ParserOptions contains parser's options about parsing.
type ParserOptions struct {
	MaxFileSize       int
	MaxFileCount      int
	MaxParseCount     int
	MaxAvailableSlice int
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
		MaxFileSize:       defaultMaxFileSize,
		MaxFileCount:      defaultMaxFileCount,
		MaxParseCount:     defaultMaxParseCount,
		MaxAvailableSlice: defaultMaxAvailableSlice,
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
// Parse sometimes not creating buf if there is a available buf in availableSlices.
// it gets a buf from availableSlices and use it.
// availableSlices fed by filter function,
// if part can't pass through the filter function, reserved buf for the part gets into availableSlices.
func (parser *Parser) Parse(r *http.Request) (*Request, error) {
	reader, err := r.MultipartReader()
	if err != nil {
		return nil, err
	}

	parts := make(map[string][]*Part)
	request := NewRequest(parts)
	availableSlices := make([][]byte, 0, parser.Options.MaxAvailableSlice)

	for parseCount := 0; parseCount < parser.Options.MaxParseCount; parseCount++ {
		part, err := reader.NextPart()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		var buf []byte
		if len(availableSlices) > 0 {
			buf, availableSlices = availableSlices[len(availableSlices)-1], availableSlices[:len(availableSlices)-1]
		} else {
			buf = make([]byte, 0)
		}

		partBuffer := bytes.NewBuffer(buf)
		partBuffer.ReadFrom(io.LimitReader(part, int64(parser.Options.MaxFileSize)))
		part.Close()

		partBytes := partBuffer.Bytes()

		p := Part{
			Name:    part.FileName(),
			Headers: part.Header,
			Size:    len(partBytes),
			Content: partBytes,
		}

		if parser.Inspector != nil {
			contentType := parser.Inspector.Inspect(partBytes)
			extensions, err := mime.ExtensionsByType(contentType)
			if err != nil {
				return nil, err
			}
			if len(extensions) == 0 {
				return nil, errors.Wrapf(errExtensionNotFound, "filename: %s", part.FileName())
			}
			p.Extension = extensions[0]
		}

		if parser.Filter != nil {
			// execute filter function
			ok := parser.Filter.Filter(&p)
			if !ok {
				if len(availableSlices) <= parser.Options.MaxAvailableSlice {
					// appending []byte to availableSlices to reuse it
					availableSlices = append(availableSlices, partBytes[:0])
				}
				continue
			}
		}
		request.parts[part.FormName()] = append(request.parts[part.FormName()], &p)
		if len := len(request.parts); len == parser.Options.MaxFileCount && len != 0 {
			return nil, errMaxFileCountExceeded
		}
	}
	return request, nil
}
