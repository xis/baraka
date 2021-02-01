package baraka

import (
	"mime"
	"testing"

	"github.com/stretchr/testify/assert"
)

type InspectorTest struct {
	inspector           Inspector
	data                []byte
	expectedContentType string
}

func TestInspector(t *testing.T) {
	tests := []InspectorTest{
		{
			inspector:           NewDefaultInspector(512),
			data:                []byte("plain text"),
			expectedContentType: "text/plain",
		},
	}

	for _, test := range tests {
		contentType := test.inspector.Inspect(test.data)
		mediaType, _, _ := mime.ParseMediaType(contentType)
		assert.Equal(t, test.expectedContentType, mediaType, "expected extension is not equal to real extension")
	}
}
