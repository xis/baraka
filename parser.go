package baraka

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

const defaultMaxFileSize = 10 << 20
const defaultMaxParseCount = 20
const defaultMaxFileCount = 0
const defaultMaxAvailableSlice = 2

var errMaxFileCountExceeded = errors.New("max file count to save exceeded")

// Parser is an interface which determine which method to use when parsing multipart files
type Parser interface {
	Parse(r *http.Request) (Processor, error)
}

// parser implements the Parser interface
type parser struct {
	Options ParserOptions
}

// ParserOptions contains parser's options about parsing.
// filter function runs when parsing the file
type ParserOptions struct {
	MaxFileSize       int
	MaxFileCount      int
	MaxParseCount     int
	MaxAvailableSlice int
	Filter            func(b []byte) bool
}

// NewParser creates a new Parser, if you give an empty ParserOptions it uses defaults.
func NewParser(options ParserOptions) Parser {
	if options.MaxFileSize == 0 {
		options.MaxFileSize = defaultMaxFileSize
	}
	if options.MaxParseCount == 0 {
		options.MaxParseCount = defaultMaxParseCount
	}
	if options.MaxFileCount == 0 {
		options.MaxFileCount = defaultMaxFileCount
	}
	if options.MaxAvailableSlice == 0 {
		options.MaxAvailableSlice = defaultMaxAvailableSlice
	}
	return parser{
		options,
	}
}

// Parse parses the http request with the multipart.Reader.
// reads parts inside the loop which iterates up to MaxParseCount at most.
// creates a []byte (buf) which gonna contain the part data.
// Parse sometimes not creating buf if there is a available buf in availableSlices.
// it gets a buf from availableSlices and use it.
// availableSlices fed by filter function,
// if part can't pass through the filter function, reserved buf for the part gets into availableSlices.
func (parser parser) Parse(r *http.Request) (Processor, error) {
	reader, err := r.MultipartReader()
	if err != nil {
		return nil, err
	}

	maxFileSize := parser.Options.MaxFileSize

	parts := make([]*Part, 0, parser.Options.MaxFileCount)
	request := NewRequest(parts...)

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
		partBuffer.ReadFrom(io.LimitReader(part, int64(maxFileSize)))
		part.Close()

		partBytes := partBuffer.Bytes()

		if parser.Options.Filter != nil {
			// execute filter function
			ok := parser.Options.Filter(partBytes)
			if !ok {
				if len(availableSlices) <= parser.Options.MaxAvailableSlice {
					// appending []byte to availableSlices to reuse it
					availableSlices = append(availableSlices, partBytes[:0])
				}
				continue
			}
		}
		p := NewPart(part.FileName(), &part.Header, len(partBytes), partBytes)
		request.parts = append(request.parts, &p)

		if len := len(request.parts); len == parser.Options.MaxFileCount && len != 0 {
			return nil, errMaxFileCountExceeded
		}
	}
	return request, nil
}
