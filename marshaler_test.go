package baraka

import (
	"strings"
	"testing"
)

func TestMarshalerJSON(t *testing.T) {
	parsers := []Parser{
		WithMultipartReader{},
		WithParseMultipartForm{},
	}
	tests := []JSONTest{
		{`{"data": 42}`, 1, "application/json"},
		{``, 0, "application/json"},
		{`{"data": 42}`, 0, "@€kds"},
	}

	for _, parser := range parsers {
		for _, ts := range tests {
			s, err := NewStorage("./", parser)
			if err != nil {
				t.Error(err)
			}
			testRaw := strings.ReplaceAll(RawMultipartWithJSON, "@data", ts.raw)
			testRaw = strings.ReplaceAll(testRaw, "@contentType", ts.contentType)
			req := CreateRequest(testRaw)

			p, err := s.Parse(req)
			if err != nil {
				t.Error(err)
			}
			j, err := p.JSON()
			if err != nil && len(j) != ts.expected {
				t.Error(err)
			}
			if len(j) != ts.expected {
				t.Error()
			}
		}
	}
}
