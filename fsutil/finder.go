package fsutil

import (
	"os"
	"path"
	"regexp"
	"strings"
)

type PathFilterFunc func(filePath string) bool

// FileFinder struct
type FileFinder struct {
	dirs  []string
	files []string

	// ignoreExts []string
	excludeExts []string
	excludeDirs []string
	excludeFiles []string

	filters []PathFilterFunc

	fileFlags int
}

// EmptyFinder new empty FileFinder instance
func EmptyFinder() *FileFinder {
	return &FileFinder{}
}

// NewFinder new FileFinder instance
func NewFinder(dirs []string, files ...string) *FileFinder {
	return &FileFinder{
		dirs:  dirs,
		files: files,
	}
}

func (f *FileFinder) AddDir(dirs ...string) *FileFinder {
	f.dirs = append(f.dirs, dirs...)
	return f
}

func (f *FileFinder) AddFile(files ...string) *FileFinder {
	f.files = append(f.files, files...)
	return f
}

func (f *FileFinder) AddFilter(filterFuncs ...PathFilterFunc) *FileFinder {
	f.filters = append(f.filters, filterFuncs...)
	return f
}

func (f *FileFinder) WithFilter(filterFuncs ...PathFilterFunc) *FileFinder {
	f.filters = append(f.filters, filterFuncs...)
	return f
}

func (f *FileFinder) ExcludeDir(dirs ...string) *FileFinder {
	f.excludeDirs = append(f.excludeDirs, dirs...)
	return f
}

func (f *FileFinder) ExcludeFile(files ...string) *FileFinder {
	f.excludeFiles = append(f.excludeFiles, files...)
	return f
}

func (f *FileFinder) FindAll() []string {
	return []string{}
}

func (f *FileFinder) Each(func(filePath string)) {
}

func (f *FileFinder) EachStat(func(file *os.FileInfo)) {
}

func (f *FileFinder) EachFile(func(file *os.File)) {
}

func (f *FileFinder) String() string {
	return ""
}

//
// built in filter functions
//

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

func DirNameFilterFunc(names []string, include bool) PathFilterFunc {
	return func(filePath string) bool {
		return !include
	}
}

func PathNameFilterFunc(names []string, include bool) PathFilterFunc {
	return func(filePath string) bool {
		return !include
	}
}
