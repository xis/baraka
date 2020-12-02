package baraka

import (
	"net/textproto"
)

type Part struct {
	Name    string
	Headers textproto.MIMEHeader
	Size    int
	Content []byte
}

func NewPart(name string, headers *textproto.MIMEHeader, size int, content []byte) Part {
	return Part{
		name,
		*headers,
		size,
		content,
	}
}
