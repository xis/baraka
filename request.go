package baraka

import "errors"

// Request implements the Processor interface.
// contains an array of parts.
// parser.Parse() returns Request as Processor.
type Request struct {
	parts map[string][]*Part
}

// NewRequest creates a new Request with parts inside
func NewRequest(parts map[string][]*Part) *Request {
	return &Request{
		parts,
	}
}

// GetForm returns the parts of the requested form, returns error if form does not exists
func (r *Request) GetForm(formname string) ([]*Part, error) {
	form, ok := r.parts[formname]
	if !ok {
		return nil, errors.New("form not found")
	}
	return form, nil
}
