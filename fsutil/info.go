package fsutil

import (
	"path"
	"path/filepath"
)

// Dir get dir path, without last name.
func Dir(fpath string) string {
	return filepath.Dir(fpath)
}

// PathName get file/dir name from fullpath
func PathName(fpath string) string {
	return path.Base(fpath)
}

// Name get file/dir name from fullpath
func Name(fpath string) string {
	return filepath.Base(fpath)
}

// FileExt get filename ext. alias of path.Ext()
func FileExt(fpath string) string {
	return path.Ext(fpath)
}

// Suffix get filename ext. alias of path.Ext()
func Suffix(fpath string) string {
	return path.Ext(fpath)
}
