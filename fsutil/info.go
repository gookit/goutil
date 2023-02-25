package fsutil

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"github.com/gookit/goutil/internal/comfunc"
)

// Dir get dir path from filepath, without last name.
func Dir(fpath string) string { return filepath.Dir(fpath) }

// PathName get file/dir name from full path
func PathName(fpath string) string { return path.Base(fpath) }

// Name get file/dir name from full path.
//
// eg: path/to/main.go => main.go
func Name(fpath string) string { return filepath.Base(fpath) }

// FileExt get filename ext. alias of path.Ext()
//
// eg: path/to/main.go => ".go"
func FileExt(fpath string) string { return path.Ext(fpath) }

// Extname get filename ext. alias of path.Ext()
//
// eg: path/to/main.go => "go"
func Extname(fpath string) string {
	if ext := path.Ext(fpath); len(ext) > 0 {
		return ext[1:]
	}
	return ""
}

// Suffix get filename ext. alias of path.Ext()
//
// eg: path/to/main.go => ".go"
func Suffix(fpath string) string { return path.Ext(fpath) }

// Expand will parse first `~` as user home dir path.
func Expand(pathStr string) string {
	return comfunc.ExpandPath(pathStr)
}

// ExpandPath will parse `~` as user home dir path.
func ExpandPath(pathStr string) string {
	return comfunc.ExpandPath(pathStr)
}

// SplitPath splits path immediately following the final Separator, separating it into a directory and file name component
func SplitPath(pathStr string) (dir, name string) {
	return filepath.Split(pathStr)
}

// GlobWithFunc handle matched file
func GlobWithFunc(pattern string, fn func(filePath string) error) (err error) {
	files, err := filepath.Glob(pattern)
	if err != nil {
		return err
	}

	for _, filePath := range files {
		err = fn(filePath)
		if err != nil {
			break
		}
	}
	return
}

type (
	// FilterFunc type for FindInDir
	FilterFunc func(fPath string, ent fs.DirEntry) bool
	// HandleFunc type for FindInDir
	HandleFunc func(fPath string, ent fs.DirEntry) error
)

// FindInDir code refer the go pkg: path/filepath.glob()
// - tip: will be not find in subdir.
//
// filters: return false will skip the file.
func FindInDir(dir string, handleFn HandleFunc, filters ...FilterFunc) (e error) {
	fi, err := os.Stat(dir)
	if err != nil || !fi.IsDir() {
		return // ignore I/O error
	}

	// names, _ := d.Readdirnames(-1)
	// sort.Strings(names)

	des, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	for _, ent := range des {
		baseName := ent.Name()
		filePath := dir + "/" + baseName

		// call filters
		if len(filters) > 0 {
			var filtered = false
			for _, filter := range filters {
				if !filter(filePath, ent) {
					filtered = true
					break
				}
			}

			if filtered {
				continue
			}
		}

		if err := handleFn(filePath, ent); err != nil {
			return err
		}
	}
	return nil
}
