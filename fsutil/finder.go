package fsutil

import (
	"os"
	"path"
	"regexp"
	"strings"
)

// PathFilter for filter file path.
type PathFilter interface {
	Filter(filePath string) bool
}

// PathFilterFunc for filter file path.
type PathFilterFunc func(filePath string) bool

// Filter for filter file path.
func (fn PathFilterFunc) Filter(filePath string) bool {
	return fn(filePath)
}

// ContentFilter for filter file path.
type ContentFilter interface {
	Filter(contents, filePath string) bool
}

// TextFilterFunc for filter file contents.
type TextFilterFunc func(contents, filePath string) bool

// Filter for filter file path.
func (fn TextFilterFunc) Filter(contents, filePath string) bool {
	return fn(contents, filePath)
}

// FileFinder struct
type FileFinder struct {
	dirPaths []string // dir paths for find file.

	// include
	includeDirs  []string // include dir names. eg: {"model", ".md"}
	includeExts  []string // include ext names. eg: {".go", ".md"}

	// ignoreExts []string
	excludeDirs  []string // exclude dir names.
	excludeExts  []string // exclude ext names. eg: {".go", ".md"}
	excludeNames []string // exclude file names. eg: {"go.mod"}

	fileFlags int
	// file time limits
	createTime int
	modifyTime int
	updateTime int

	pathFilters []PathFilter
	textFilters []ContentFilter

	// mark has been run find()
	founded bool

	// founded file paths.
	filePaths []string
	// the founded file instances
	osFiles map[string]*os.File
}

// EmptyFinder new empty FileFinder instance
func EmptyFinder() *FileFinder {
	return &FileFinder{}
}

// NewFinder new instance with source dir paths.
func NewFinder(dirPaths []string, filePaths ...string) *FileFinder {
	return &FileFinder{
		dirPaths:  dirPaths,
		filePaths: filePaths,
	}
}

// AddDirPath add source dir for find
func (f *FileFinder) AddDirPath(dirPaths ...string) *FileFinder {
	f.dirPaths = append(f.dirPaths, dirPaths...)
	return f
}

// AddDir add source dir for find. alias of AddDirPath()
func (f *FileFinder) AddDir(dirPaths ...string) *FileFinder {
	f.dirPaths = append(f.dirPaths, dirPaths...)
	return f
}

// ExcludeDir exclude dir names.
func (f *FileFinder) ExcludeDir(dirs ...string) *FileFinder {
	f.excludeDirs = append(f.excludeDirs, dirs...)
	return f
}

// ExcludeName exclude file names.
func (f *FileFinder) ExcludeName(files ...string) *FileFinder {
	f.excludeNames = append(f.excludeNames, files...)
	return f
}

// AddFilter for filter filepath
func (f *FileFinder) AddFilter(filterFuncs ...PathFilter) *FileFinder {
	f.pathFilters = append(f.pathFilters, filterFuncs...)
	return f
}

// WithFilter for filter func for filtering filepath
func (f *FileFinder) WithFilter(filterFuncs ...PathFilter) *FileFinder {
	f.pathFilters = append(f.pathFilters, filterFuncs...)
	return f
}

// AddTextFilter for filter filepath
func (f *FileFinder) AddTextFilter(filterFuncs ...ContentFilter) *FileFinder {
	f.textFilters = append(f.textFilters, filterFuncs...)
	return f
}

// WithTextFilter for filter func for filtering filepath
func (f *FileFinder) WithTextFilter(filterFuncs ...ContentFilter) *FileFinder {
	f.textFilters = append(f.textFilters, filterFuncs...)
	return f
}

// AddFilePaths set founded files
func (f *FileFinder) AddFilePaths(filePaths []string) {
	f.filePaths = append(f.filePaths, filePaths...)
}

// AddFilePath add source file
func (f *FileFinder) AddFilePath(filePaths ...string) *FileFinder {
	f.filePaths = append(f.filePaths, filePaths...)
	return f
}

// AddFile add source file. alias of AddFilePath()
func (f *FileFinder) AddFile(filePaths ...string) *FileFinder {
	f.filePaths = append(f.filePaths, filePaths...)
	return f
}

// FindAll find and return founded file paths.
func (f *FileFinder) FindAll() []string {
	f.find()

	return f.filePaths
}

// Find find file paths.
func (f *FileFinder) Find() *FileFinder {
	f.find()
	return f
}

// do finding
func (f *FileFinder) find() {
	// mark found
	if f.founded {
		return
	}
	f.founded = true

	// TODO do finding

}

// EachContents each file path.
func (f *FileFinder) Each(func(filePath string)) {
}

// EachContents each file os.File
func (f *FileFinder) EachFile(func(file *os.File)) {
}

// EachContents each file os.FileInfo
func (f *FileFinder) EachStat(func(file *os.FileInfo)) {
}

// EachContents each file contents
func (f *FileFinder) EachContents(func(str string)) {
}

func (f *FileFinder) Reset() {
	f.founded = false

	f.filePaths = make([]string, 0)

	f.excludeNames = make([]string, 0)
	f.excludeDirs = make([]string, 0)
}

func (f *FileFinder) String() string {
	return ""
}

//
// built in filter functions
//

// ExtFilterFunc filter filepath by given file ext.
//
// Usage:
//	f := EmptyFiler()
//	f.AddFilter(ExtFilterFunc([]string{".go", ".md"}, true))
//	f.AddFilter(ExtFilterFunc([]string{".log", ".tmp"}, false))
func ExtFilterFunc(exts []string, include bool) PathFilterFunc {
	return func(filePath string) bool {
		fExt := path.Ext(filePath)

		for _, ext := range exts {
			if fExt == ext {
				return include
			}
		}
		return !include
	}
}

// SuffixFilterFunc filter filepath by given file ext.
//
// Usage:
//	f := EmptyFiler()
//	f.AddFilter(SuffixFilterFunc([]string{"util.go", ".md"}, true))
//	f.AddFilter(SuffixFilterFunc([]string{"_test.go", ".log"}, false))
func SuffixFilterFunc(suffixes []string, include bool) PathFilterFunc {
	return func(filePath string) bool {
		for _, sfx := range suffixes {
			if strings.HasSuffix(filePath, sfx) {
				return include
			}
		}
		return !include
	}
}

// PathNameFilterFunc filter filepath by given path names.
func PathNameFilterFunc(names []string, include bool) PathFilterFunc {
	return func(filePath string) bool {
		for _, name := range names {
			if strings.Contains(filePath, name) {
				return include
			}
		}
		return !include
	}
}

// DirNameFilterFunc filter filepath by given dir names.
func DirNameFilterFunc(names []string, include bool) PathFilterFunc {
	return func(filePath string) bool {
		dir := path.Dir(filePath)

		for _, name := range names {
			if strings.Contains(dir, name) {
				return include
			}
		}
		return !include
	}
}

// DotFileFilterFunc filter dot filename. eg: ".gitignore"
func DotFileFilterFunc(include bool) PathFilterFunc {
	return func(filePath string) bool {
		filename := path.Base(filePath)
		if filename[0] == '.' {
			return include
		}

		return !include
	}
}

// GlobFilterFunc filter filepath by given patterns.
//
// Usage:
//	f := EmptyFiler()
//	f.AddFilter(GlobFilterFunc([]string{"*_test.go"}, true))
func GlobFilterFunc(patterns []string, include bool) PathFilterFunc {
	return func(filePath string) bool {
		for _, pattern := range patterns {
			if ok, _ := path.Match(pattern, filePath); ok {
				return include
			}
		}
		return !include
	}
}

// RegexFilterFunc filter filepath by given regex pattern
//
// Usage:
//	f := EmptyFiler()
//	f.AddFilter(RegexFilterFunc(`[A-Z]\w+`, true))
func RegexFilterFunc(pattern string, include bool) PathFilterFunc {
	reg := regexp.MustCompile(pattern)

	return func(filePath string) bool {
		return reg.MatchString(filePath)
	}
}
