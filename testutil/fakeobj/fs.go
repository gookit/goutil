package fakeobj

import (
	"io"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/gookit/goutil/basefn"
)

// DirEntry implements the fs.DirEntry
type DirEntry struct {
	Dir bool
	Nam string // basename
	Mod fs.FileMode
	Fi  fs.FileInfo
	Err error
}

// NewDirEntry create a fs.DirEntry
func NewDirEntry(fpath string, isDir ...bool) *DirEntry {
	isd := basefn.FirstOr(isDir, false)
	return &DirEntry{Nam: filepath.Base(fpath), Dir: isd, Mod: fs.ModePerm}
}

// Name get
func (d *DirEntry) Name() string {
	return d.Nam
}

// IsDir get
func (d *DirEntry) IsDir() bool {
	return d.Dir
}

// Type get
func (d *DirEntry) Type() fs.FileMode {
	return d.Mod
}

// Info get
func (d *DirEntry) Info() (fs.FileInfo, error) {
	if d.Fi == nil {
		d.Fi = &FileInfo{
			Dir: d.Dir,
			Nam: d.Nam,
			Mod: d.Mod,
		}
	}

	return d.Fi, d.Err
}

// FileInfo implements the fs.FileInfo, fs.File
type FileInfo struct {
	Dir bool
	Nam string // basename
	Mod fs.FileMode
	Mt  time.Time

	// Path full path
	Path string

	Contents string
	CloseErr error
	offset   int
}

// NewFile instance
func NewFile(fpath string) *FileInfo {
	return NewFileInfo(fpath)
}

// NewFileInfo instance
func NewFileInfo(fpath string, isDir ...bool) *FileInfo {
	return &FileInfo{
		Dir:  basefn.FirstOr(isDir, false),
		Nam:  filepath.Base(fpath),
		Mod:  fs.ModePerm,
		Path: fpath,
	}
}

// WithBody set file body contents
func (f *FileInfo) WithBody(s string) *FileInfo {
	f.Contents = s
	return f
}

// WithMtime set file modify time
func (f *FileInfo) WithMtime(mt time.Time) *FileInfo {
	f.Mt = mt
	return f
}

// Reset prepares a FileInfo for reuse.
func (f *FileInfo) Reset() *FileInfo {
	f.offset = 0
	return f
}

// fs.File methods.

// Stat returns the FileInfo structure describing file.
func (f *FileInfo) Stat() (fs.FileInfo, error) {
	return f, nil
}

// Read reads up to len(p) bytes into p.
func (f *FileInfo) Read(p []byte) (int, error) {
	if f.offset >= len(f.Contents) {
		return 0, io.EOF
	}

	n := copy(p, f.Contents[f.offset:])
	f.offset += n
	return n, nil
}

// Close closes the file
func (f *FileInfo) Close() error {
	return f.CloseErr
}

// fs.FileInfo methods.

// Name returns the base name of the file.
func (f *FileInfo) Name() string {
	return f.Nam
}

// Size returns the length in bytes for regular files; system-dependent for others.
func (f *FileInfo) Size() int64 {
	return int64(len(f.Contents))
}

// Mode returns file mode bits.
func (f *FileInfo) Mode() fs.FileMode {
	return f.Mod
}

// ModTime returns the modification time.
func (f *FileInfo) ModTime() time.Time {
	return f.Mt
}

// IsDir returns true if the file is a directory.
func (f *FileInfo) IsDir() bool {
	return f.Dir
}

// Sys returns underlying data source (can return nil).
func (f *FileInfo) Sys() any {
	return nil
}
