package fsutil

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/gookit/goutil/strutil"
)

// SearchNameUp find file/dir name in dirPath or parent dirs,
// return the name of directory path
//
// Usage:
//
//	repoDir := fsutil.SearchNameUp("/path/to/dir", ".git")
func SearchNameUp(dirPath, name string) string {
	dir, _ := SearchNameUpx(dirPath, name)
	return dir
}

// SearchNameUpx find file/dir name in dirPath or parent dirs,
// return the name of directory path and dir is changed.
func SearchNameUpx(dirPath, name string) (string, bool) {
	var level int
	dirPath = ToAbsPath(dirPath)

	for {
		namePath := filepath.Join(dirPath, name)
		if PathExists(namePath) {
			return dirPath, level > 0
		}

		level++
		prevLn := len(dirPath)
		dirPath = filepath.Dir(dirPath)
		if prevLn == len(dirPath) {
			return "", false
		}
	}
}

// WalkDir walks the file tree rooted at root, calling fn for each file or
// directory in the tree, including root.
func WalkDir(dir string, fn fs.WalkDirFunc) error {
	return filepath.WalkDir(dir, fn)
}

// GlobWithFunc handle matched file
//
// - TIP: will be not find in subdir.
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
	//
	// - return False will skip handle the file.
	FilterFunc func(fPath string, ent fs.DirEntry) bool
	// HandleFunc type for FindInDir
	HandleFunc func(fPath string, ent fs.DirEntry) error
)

// OnlyFindDir on find
func OnlyFindDir(_ string, ent fs.DirEntry) bool {
	return ent.IsDir()
}

// OnlyFindFile on find
func OnlyFindFile(_ string, ent fs.DirEntry) bool {
	return !ent.IsDir()
}

// IncludeSuffix on find
func IncludeSuffix(ss ...string) FilterFunc {
	return func(_ string, ent fs.DirEntry) bool {
		return strutil.HasOneSuffix(ent.Name(), ss)
	}
}

// ExcludeDotFile on find
func ExcludeDotFile(_ string, ent fs.DirEntry) bool {
	return ent.Name()[0] != '.'
}

// ExcludeSuffix on find
func ExcludeSuffix(ss ...string) FilterFunc {
	return func(_ string, ent fs.DirEntry) bool {
		return !strutil.HasOneSuffix(ent.Name(), ss)
	}
}

// ApplyFilters handle
func ApplyFilters(fPath string, ent fs.DirEntry, filters []FilterFunc) bool {
	for _, filter := range filters {
		if !filter(fPath, ent) {
			return true
		}
	}
	return false
}

// FindInDir code refer the go pkg: path/filepath.glob()
//
// - TIP: will be not find in subdir.
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
		filePath := dir + "/" + ent.Name()

		// apply filters
		if len(filters) > 0 && ApplyFilters(filePath, ent, filters) {
			continue
		}

		if err := handleFn(filePath, ent); err != nil {
			return err
		}
	}
	return nil
}
