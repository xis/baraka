package baraka

import (
	"testing"
)

type ParserTest struct {
	parser         Parser
	expectedLength int
}

func TestParserParse(t *testing.T) {
	raw := RawMultipartPlainText

	tests := []ParserTest{
		{NewParser(ParserOptions{}), 2},
		{NewParser(ParserOptions{
			Filter: FilterJPEG(),
		}), 0},
	}

	for _, test := range tests {
		req := CreateHTTPRequest(raw)
		processor, err := test.parser.Parse(req)
		if err != nil {
			t.Error(err)
		}
		if len(processor.Content()) != test.expectedLength {
			t.Error("content item length is not equal to expected length")
		}
	}
}
