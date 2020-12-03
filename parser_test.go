package baraka

import (
	"testing"
)

func TestParserParse(t *testing.T) {
	raw := RawMultipartPlainText

	parsers := []Parser{
		Options{},
		Options{
			Filter: FilterJPEG(),
		},
	}
	for _, parser := range parsers {
		s, err := NewStorage("./", parser)
		if err != nil {
			t.Error(err)
		}
		req := CreateRequest(raw)

		_, err = s.Parse(req)
		if err != nil {
			t.Error(err)
		}
		req = CreateRequest(raw)
		_, err = s.ParseButMax(32<<20, 1, req)
		if err != nil {
			t.Error(err)
		}
	}
}
