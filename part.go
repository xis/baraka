package baraka

import (
	"net/textproto"
)

// Part is a struct that you can access the all contents of the multipart.Part
type Part struct {
	Name      string
	Headers   textproto.MIMEHeader
	Size      int
	Content   []byte
	Extension string
}

// GetHeader returns the value of the given header key
func (p *Part) GetHeader(key string) string {
	return p.Headers.Get(key)
}
