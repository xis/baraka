package baraka

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type RequestTest struct {
	request       *Request
	expectedForms []FormTest
}

type FormTest struct {
	expectedName        string
	expectedPartsLength int
}

func TestRequest(t *testing.T) {
	request := NewRequest(
		map[string][]*Part{
			"test": {
				&PartPlainText,
			},
		},
	)

	tests := []RequestTest{
		{
			request: request,
			expectedForms: []FormTest{
				{
					expectedName:        "test",
					expectedPartsLength: 1,
				},
			},
		},
	}

	for _, test := range tests {
		for _, expectedForm := range test.expectedForms {
			form, err := test.request.GetForm(expectedForm.expectedName)
			if err != nil {
				t.Error("expected form not found in the request")
			}

			assert.Equal(t, expectedForm.expectedPartsLength, len(form), "parts length is not equal to expected parts length")
		}
	}
}
