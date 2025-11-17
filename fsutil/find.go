package fsutil

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/internal/comfunc"
	"github.com/gookit/goutil/strutil"
)

// FilePathInDirs get full file path in dirs. return empty string if not found.
//
// Params:
//   - file: can be relative path, file name, full path.
//   - dirs: dir paths
func FilePathInDirs(fPath string, dirs ...string) string {
	fPath = comfunc.ExpandHome(fPath)
	if FileExists(fPath) {
		return fPath
	}

	for _, dirPath := range dirs {
		fPath = JoinSubPaths(dirPath, fPath)
		if FileExists(fPath) {
			return fPath
		}
	}
	return "" // not found
}

// FirstExists check multi paths and return first exists path.
func FirstExists(paths ...string) string {
	return MatchFirst(paths, PathExists, "")
}

// FirstExistsDir check multi paths and return first exists dir.
func FirstExistsDir(paths ...string) string {
	return MatchFirst(paths, IsDir, "")
}

// FirstExistsFile check multi paths and return first exists file.
func FirstExistsFile(paths ...string) string {
	return MatchFirst(paths, IsFile, "")
}

// MatchPaths given paths by custom mather func.
func MatchPaths(paths []string, matcher PathMatchFunc) []string {
	var ret []string
	for _, p := range paths {
		if matcher(p) {
			ret = append(ret, p)
		}
	}
	return ret
}

// MatchFirst filter paths by filter func and return first match path.
func MatchFirst(paths []string, matcher PathMatchFunc, defaultPath string) string {
	for _, p := range paths {
		if matcher(p) {
			return p
		}
	}
	return defaultPath
}

// FindParentOption options
type FindParentOption struct {
	MaxLevel int // default: 10
	// NeedDir true: find dirs; false(default): find files
	NeedDir bool
	OnlyOne bool // only find one, default: true
	// Collector func
	Collector func(fullPath string)
	// MatchFunc custom matcher func. return false to stop find.
	MatchFunc func(currentDir string) bool
}

// FindParentOptFn find parent option func
type FindParentOptFn func(opt *FindParentOption)

// FindAllInParentDirs looks for all match file(default)/dir in the current directory and parent directories
func FindAllInParentDirs(dirPath, name string, optFns ...FindParentOptFn) []string {
	var foundPaths []string
	optFns = append(optFns, func(opt *FindParentOption) {
		opt.OnlyOne = false
	})

	FindNameInParentDirs(dirPath, name, func(fullPath string) {
		foundPaths = append(foundPaths, fullPath)
	}, optFns...)
	return foundPaths
}

// FindOneInParentDirs looks for a file(default)/dir in the current directory and parent directories
func FindOneInParentDirs(dirPath, name string, optFns ...FindParentOptFn) string {
	var foundPath string
	FindNameInParentDirs(dirPath, name, func(fullPath string) {
		foundPath = fullPath
	}, optFns...)
	return foundPath
}

// FindNameInParentDirs looks for file(default)/dir in the current directory and parent directories
func FindNameInParentDirs(dirPath, name string, collectFn func(fullPath string), optFns ...FindParentOptFn) {
	opts := &FindParentOption{
		MaxLevel:  10,
		OnlyOne:   true,
		Collector: collectFn,
	}
	for _, fn := range optFns {
		fn(opts)
	}

	FindInParentDirs(dirPath, func(currentDir string) bool {
		filePath := filepath.Join(currentDir, name)
		if fi, err := os.Stat(filePath); err == nil {
			found := false
			if fi.IsDir() {
				found = opts.NeedDir
			} else {
				found = !opts.NeedDir
			}

			if found {
				opts.Collector(filePath)
				return !opts.OnlyOne
			}
		}
		return true
	}, opts.MaxLevel)
}

// FindInParentDirs looks for file/dir in the current directory and parent directories
//  - MatchFunc custom matcher func. return false to stop find.
func FindInParentDirs(dirPath string, matchFunc func(dir string) bool, maxLevel int) {
	currentLv := 1
	currentDir := ToAbsPath(dirPath)

	for {
		// Check if the file exists in the current directory
		if !matchFunc(currentDir) {
			return
		}

		// check find level
		if maxLevel > 0 && currentLv > maxLevel {
			break
		}

		// Get parent directory
		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			// Reached the root, file not found
			return
		}

		// Move to parent directory
		currentLv++
		currentDir = parentDir
	}
}

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
//
// TIP: will recursively found in sub dirs.
func WalkDir(dir string, fn fs.WalkDirFunc) error {
	return filepath.WalkDir(dir, fn)
}

// Glob finds files by glob path pattern. alias of filepath.Glob()
// and support filter matched files by name.
//
// Usage:
//
//	files := fsutil.Glob("/path/to/dir/*.go")
func Glob(pattern string, fls ...NameMatchFunc) []string {
	files, _ := filepath.Glob(pattern)
	if len(fls) == 0 || len(files) == 0 {
		return files
	}

	var matched []string
	for _, file := range files {
		for _, fn := range fls {
			if fn(path.Base(file)) {
				matched = append(matched, file)
				break
			}
		}
	}
	return matched
}

// GlobWithFunc find files by glob path pattern, then handle matched file
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
func OnlyFindDir(_ string, ent fs.DirEntry) bool { return ent.IsDir() }

// OnlyFindFile on find
func OnlyFindFile(_ string, ent fs.DirEntry) bool { return !ent.IsDir() }

// ExcludeNames on find
func ExcludeNames(names ...string) FilterFunc {
	return func(_ string, ent fs.DirEntry) bool {
		return !arrutil.StringsHas(names, ent.Name())
	}
}

// IncludeSuffix on find
func IncludeSuffix(ss ...string) FilterFunc {
	return func(_ string, ent fs.DirEntry) bool {
		return strutil.HasOneSuffix(ent.Name(), ss)
	}
}

// ExcludeDotFile on find
func ExcludeDotFile(_ string, ent fs.DirEntry) bool { return ent.Name()[0] != '.' }

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
// - TIP: default will be not found in sub-dir.
//
// filters: return false will skip the file.
func FindInDir(dir string, handleFn HandleFunc, filters ...FilterFunc) (e error) {
	fi, err := os.Stat(dir)
	if err != nil || !fi.IsDir() {
		return // ignore I/O error
	}

	des, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	// remove the last '/' char
	dirLn := len(dir)
	if dirLn > 1 && dir[dirLn-1] == '/' {
		dir = dir[:dirLn-1]
	}

	for _, ent := range des {
		filePath := dir + "/" + ent.Name()

		// apply filters
		if len(filters) > 0 && ApplyFilters(filePath, ent, filters) {
			continue
		}

		if err1 := handleFn(filePath, ent); err1 != nil {
			return err1
		}
	}
	return nil
}

// FileInDirs returns the first file path in the given dirs.
func FileInDirs(paths []string, names ...string) string {
	for _, pathDir := range paths {
		for _, name := range names {
			file := pathDir + "/" + name
			if IsFile(file) {
				return file
			}
		}
	}
	return ""
}
