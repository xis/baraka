package baraka

import (
	"os"
	"path/filepath"
	"testing"
)

type SaveTest struct {
	part                    Part
	path                    string
	expectedExistenceOfFile bool
	expectedFilename        string
	expectedSizeOfFile      int
}

func TestSaverSave(t *testing.T) {
	tests := []SaveTest{
		{
			part:                    PartPlainText,
			path:                    "test",
			expectedExistenceOfFile: true,
			expectedFilename:        PartPlainText.Name,
			expectedSizeOfFile:      15,
		},
	}
	store := NewFileSystemStore("./")
	for _, test := range tests {
		err := store.Save("test", test.part.Name, &test.part)
		if err != nil {
			t.Error(err)
		}
		file, err := os.Stat(filepath.Join(store.Path, test.path, test.expectedFilename))
		if err != nil {
			t.Error(err)
		}
		if file.Size() != int64(test.expectedSizeOfFile) {
			t.Error("file size is not equal to expected file size")
		}
	}
}
