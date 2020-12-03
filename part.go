package baraka

import (
	"net/textproto"
)

// Part is a struct that you can access the all contents of the multipart.Part
type Part struct {
	Name    string
	Headers textproto.MIMEHeader
	Size    int
	Content []byte
}

// NewPart creates a new Part struct
func NewPart(name string, headers *textproto.MIMEHeader, size int, content []byte) Part {
	return Part{
		name,
		*headers,
		size,
		content,
	}
}
