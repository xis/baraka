package baraka

import "net/http"

type Inspector interface {
	Inspect(data []byte) string
}

type DefaultInspector struct {
	maxSampleSize int
}

func NewDefaultInspector(maxSampleSize int) Inspector {
	return &DefaultInspector{
		maxSampleSize: maxSampleSize,
	}
}

func (inspector *DefaultInspector) Inspect(data []byte) string {
	maxSampleSize := inspector.maxSampleSize
	if len(data) < maxSampleSize {
		maxSampleSize = len(data)
	}
	sample := data[:maxSampleSize]
	contentType := http.DetectContentType(sample)
	return contentType
}
