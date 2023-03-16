package fsutil

import (
	"io/fs"
	"os"
	"path/filepath"
)

// SearchNameUp find file/dir name in dirPath or parent dirs,
// return the name of directory path
//
// Usage:
//
//	repoDir := fsutil.SearchNameUp("/path/to/dir", ".git")
func SearchNameUp(dirPath, name string) string {
	dirPath = ToAbsPath(dirPath)

	for {
		namePath := filepath.Join(dirPath, name)
		if PathExists(namePath) {
			return dirPath
		}

		prevLn := len(dirPath)
		dirPath = filepath.Dir(dirPath)
		if prevLn == len(dirPath) {
			return ""
		}
	}
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
