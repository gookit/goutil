package testutil

import (
	"github.com/gookit/goutil/testutil/fakeobj"
)

// TestWriter struct, useful for testing. alias of fakeobj.Writer
type TestWriter = fakeobj.Writer

// NewTestWriter instance
func NewTestWriter() *TestWriter {
	return &TestWriter{}
}

// DirEnt implements the fs.DirEntry
type DirEnt = fakeobj.DirEntry

// NewDirEnt create a fs.DirEntry
func NewDirEnt(fpath string, isDir ...bool) *fakeobj.DirEntry {
	return fakeobj.NewDirEntry(fpath, isDir...)
}
