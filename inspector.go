package baraka

import "net/http"

// Inspector is a interface for finding the content type of the byte array
type Inspector interface {
	Inspect(data []byte) string
}

// DefaultInspector is the default inspector
// uses http.DetectContentType function to find the content type
type DefaultInspector struct {
	maxSampleSize int
}

// NewDefaultInspector creates a new default inspector
func NewDefaultInspector(maxSampleSize int) Inspector {
	return &DefaultInspector{
		maxSampleSize: maxSampleSize,
	}
}

// Inspect finds the content type of the byte array
func (inspector *DefaultInspector) Inspect(data []byte) string {
	maxSampleSize := inspector.maxSampleSize
	if len(data) < maxSampleSize {
		maxSampleSize = len(data)
	}
	sample := data[:maxSampleSize]
	contentType := http.DetectContentType(sample)
	return contentType
}
