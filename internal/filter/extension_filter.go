package filter

import (
	"errors"
	"github.com/xis/baraka/v2"
	"mime"
	"net/http"
)

const (
	sampleSize = 512
)

var (
	ErrExtensionNotFound = errors.New("can't detect file's extension")
)

// ExtensionFilter is a filter which filters the unwanted extensions
// passes the part if the part's extension is in the extensions field
type ExtensionFilter struct {
	extensions []string
}

func NewExtensionFilter(extensions ...string) baraka.Filter {
	return &ExtensionFilter{
		extensions: extensions,
	}
}

func (f *ExtensionFilter) Filter(data []byte) (bool, error) {
	maxSampleSize := sampleSize

	if len(data) < maxSampleSize {
		maxSampleSize = len(data)
	}

	sample := data[:maxSampleSize]
	contentType := http.DetectContentType(sample)

	extensions, err := mime.ExtensionsByType(contentType)
	if err != nil {
		return false, err
	}

	if len(extensions) == 0 {
		return false, ErrExtensionNotFound
	}

	extension := extensions[0]

	for _, validExtension := range f.extensions {
		if extension == validExtension {
			return true, nil
		}
	}

	return false, nil
}
