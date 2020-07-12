package baraka

import (
	"bytes"
	"io"
	"net/textproto"
)

// Header @
type Header struct {
	Filename string
	Header   textproto.MIMEHeader
	Size     int64
	content  []byte
}

// Open @
func (fh *Header) Open() *io.SectionReader {
	r := io.NewSectionReader(bytes.NewReader(fh.content), 0, int64(len(fh.content)))
	return r
}
