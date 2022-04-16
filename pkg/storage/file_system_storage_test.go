package storage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type FilesystemStorageTestSuite struct {
	suite.Suite

	Storage FilesystemStorage
}

func TestFilesystemStorageTestSuite(t *testing.T) {
	suite.Run(t, new(FilesystemStorageTestSuite))
}

func (suite *FilesystemStorageTestSuite) SetupTest() {
	suite.Storage = NewFilesystemStorage("./testdata")
}

func (suite *FilesystemStorageTestSuite) TestStore() {
	err := suite.Storage.Store("something.txt", []byte("something"))
	suite.NoError(err)
}

func (suite *FilesystemStorageTestSuite) TestStoreButNested() {
	err := suite.Storage.Store("/something/something/something.txt", []byte("something"))
	suite.NoError(err)
}

func (suite *FilesystemStorageTestSuite) TearDownSuite() {
	err := os.RemoveAll("./testdata")
	suite.NoError(err)
}
