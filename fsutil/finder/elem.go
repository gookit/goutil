package finder

import (
	"io/fs"

	"github.com/gookit/goutil/strutil"
)

// Elem of find file/dir path result
type Elem interface {
	fs.DirEntry
	// Path get file/dir full path. eg: "/path/to/file.go"
	Path() string
	// Info get file info. like fs.DirEntry.Info(), but will cache result.
	Info() (fs.FileInfo, error)
}

type elem struct {
	fs.DirEntry
	path string
	stat fs.FileInfo
	sErr error
}

// NewElem create a new Elem instance
func NewElem(fPath string, ent fs.DirEntry) Elem {
	return &elem{
		path:     fPath,
		DirEntry: ent,
	}
}

// Path get full file/dir path. eg: "/path/to/file.go"
func (e *elem) Path() string {
	return e.path
}

// Info get file info, will cache result
func (e *elem) Info() (fs.FileInfo, error) {
	if e.stat == nil {
		e.stat, e.sErr = e.DirEntry.Info()
	}
	return e.stat, e.sErr
}

// String get string representation
func (e *elem) String() string {
	return strutil.OrCond(e.IsDir(), "dir: ", "file: ") + e.Path()
}
