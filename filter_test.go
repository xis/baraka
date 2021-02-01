package baraka

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type FilterTest struct {
	filter         Filter
	part           Part
	expectedResult bool
}

func TestFilter(t *testing.T) {
	tests := []FilterTest{
		{
			filter:         NewExtensionFilter(".jpg"),
			part:           Part{Extension: ".jpg"},
			expectedResult: true,
		},
	}

	for _, test := range tests {
		result := test.filter.Filter(&test.part)
		assert.Equal(t, test.expectedResult, result, "expected result is not equal to real result")
	}
}
