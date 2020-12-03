package baraka

import (
	"bytes"
	"io"
	"net/textproto"
)

// Header is a similar to multipart.FileHeader
// defined it again in here because can't access the unexported field of the multipart.Header
type Header struct {
	Filename string
	Header   textproto.MIMEHeader
	Size     int64
	content  []byte
}

// Open is a similar to multipart.FileHeader.Open function
// opens Header
func (fh *Header) Open() *io.SectionReader {
	r := io.NewSectionReader(bytes.NewReader(fh.content), 0, int64(len(fh.content)))
	return r
}
