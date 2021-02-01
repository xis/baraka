package baraka

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type PartTest struct {
	part            *Part
	expectedHeaders []string
}

func TestPart(t *testing.T) {
	tests := []PartTest{
		{
			part: &PartPlainText,
			expectedHeaders: []string{
				"Content-Type",
			},
		},
	}

	for _, test := range tests {
		for _, header := range test.expectedHeaders {
			value := test.part.GetHeader(header)
			assert.NotEmpty(t, value, "expected header's value is empty")
		}
	}
}
