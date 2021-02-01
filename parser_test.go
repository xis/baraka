package baraka

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ParserTest struct {
	parser         *Parser
	expectedLength int
}

func TestParserParse(t *testing.T) {
	raw := RawMultipartPlainText

	tests := []ParserTest{
		{
			DefaultParser(),
			2,
		},
		{
			DefaultParser().SetFilter(NewExtensionFilter(".jpg", ".png")).SetInspector(NewDefaultInspector(512)),
			0,
		},
	}

	for _, test := range tests {
		req := CreateHTTPRequest(raw)
		request, err := test.parser.Parse(req)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, test.expectedLength, len(request.parts), "parts length is not equal to expected parts length")
	}
}
