package storage

import (
	"os"
	"syscall"
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

func (suite *FilesystemStorageTestSuite) TestStoreButNestedAndAlreadyExists() {
	err := os.MkdirAll("./testdata/space/star", os.ModeSticky|os.ModePerm)
	suite.NoError(err)

	err = suite.Storage.Store("/space/star/something.txt", []byte("something"))
	suite.NoError(err)
}

func (suite *FilesystemStorageTestSuite) TestStoreButPathIsAFile() {
	err := os.WriteFile("./testdata/space/nebula", []byte("something"), 0666)
	suite.NoError(err)

	err = suite.Storage.Store("/space/nebula/something.txt", []byte("something"))

	suite.ErrorIs(err, syscall.ENOTDIR)
}

func (suite *FilesystemStorageTestSuite) TearDownSuite() {
	err := os.RemoveAll("./testdata")
	suite.NoError(err)
}
