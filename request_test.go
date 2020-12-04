package baraka

import (
	"bytes"
	"net/http"
	"os"
	"testing"
)

type RequestTest struct {
	httpRequest          *http.Request
	parser               Parser
	expectedContent      []string
	expectedLength       int
	expectedFilenames    []string
	expectedContentTypes []string
	expectedJSON         JSONTest
	expectedSave         []SaveTest
}

type JSONTest struct {
	expectedErr         bool
	expectedLength      int
	expectedJSONStrings []string
}

type SaveTest struct {
	expectedExistenceOfFile bool
	expectedFilename        string
	expectedSizeOfFile      int
}

func TestRequest(t *testing.T) {
	raw := RawMultipartPlainText
	parser := NewParser(ParserOptions{})
	httpRequest := CreateHTTPRequest(raw)

	tests := []RequestTest{
		{
			httpRequest:          httpRequest,
			parser:               parser,
			expectedContent:      []string{"test file a", "test file b"},
			expectedLength:       2,
			expectedFilenames:    []string{"filea.txt", "fileb.txt"},
			expectedContentTypes: []string{"text/plain", "text/plain"},
			expectedJSON: JSONTest{
				expectedErr: true, expectedLength: 0, expectedJSONStrings: []string{"", ""},
			},
			expectedSave: []SaveTest{
				{expectedExistenceOfFile: true, expectedFilename: "test_0.txt", expectedSizeOfFile: 11},
				{expectedExistenceOfFile: true, expectedFilename: "test_1.txt", expectedSizeOfFile: 11},
			},
		},
	}

	for _, test := range tests {
		processor, err := test.parser.Parse(test.httpRequest)
		if err != nil {
			t.Error(err)
		}
		testContent(t, test.expectedContent, processor)
		testLength(t, test.expectedLength, processor)
		testFilenames(t, test.expectedFilenames, processor)
		testContentTypes(t, test.expectedContentTypes, processor)
		testJSON(t, test.expectedJSON, processor)
		testSave(t, test.expectedSave, processor)
	}

}

func testContent(t *testing.T, expectedContent []string, request Processor) {
	content := request.Content()
	for key, part := range content {
		if bytes.Compare(part.Content, []byte(expectedContent[key])) != 0 {
			t.Error("part content is not equal to expected content value")
		}
	}
}

func testLength(t *testing.T, expectedLength int, request Processor) {
	length := request.Length()
	if length != expectedLength {
		t.Error("content length of request is not equal to expected content length")
	}
}

func testFilenames(t *testing.T, expectedFilenames []string, request Processor) {
	filenames := request.Filenames()
	for key, filename := range filenames {
		if filename != expectedFilenames[key] {
			t.Error("filename of part is not equal to expected filename")
		}
	}
}

func testContentTypes(t *testing.T, expectedContentTypes []string, request Processor) {
	contentTypes := request.ContentTypes()
	for key, contentType := range contentTypes {
		if contentType != expectedContentTypes[key] {
			t.Error("content type of part is not equal to expected content type", contentType)
		}
	}
}

func testJSON(t *testing.T, expectedJSON JSONTest, request Processor) {
	jsonStrings, err := request.GetJSON()
	if err != nil && !expectedJSON.expectedErr {
		t.Error(err)
	}
	if len(jsonStrings) != expectedJSON.expectedLength {
		t.Error("json bytes array length is not equal to expected length value")
	}
	for key, jsonString := range jsonStrings {
		if jsonString != expectedJSON.expectedJSONStrings[key] {
			t.Error("json string is not equal to expected json string")
		}
	}
}

func testSave(t *testing.T, expectedSave []SaveTest, request Processor) {
	err := request.Save("test_", "./", "application/json")
	if err != nil {
		t.Error(err)
	}

	for _, test := range expectedSave {
		file, err := os.Stat("./" + test.expectedFilename)
		if err != nil {
			t.Error(err)
		}
		if file.Size() != int64(test.expectedSizeOfFile) {
			t.Error("file size is not equal to expected file size")
		}
	}
}
