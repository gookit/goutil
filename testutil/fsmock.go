package testutil

import (
	"github.com/gookit/goutil/testutil/fakeobj"
)

// DirEnt implements the fs.DirEntry
type DirEnt = fakeobj.DirEntry

// NewDirEnt create a fs.DirEntry
func NewDirEnt(fpath string, isDir ...bool) *fakeobj.DirEntry {
	return fakeobj.NewDirEntry(fpath, isDir...)
}
