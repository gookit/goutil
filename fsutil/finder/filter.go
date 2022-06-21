package finder

import (
	"os"
	"path"
	"regexp"
	"strings"
	"time"
)

// FileFilter for filter file path.
type FileFilter interface {
	FilterFile(filePath, filename string) bool
}

// FileFilterFunc for filter file path.
type FileFilterFunc func(filePath, filename string) bool

// FilterFile Filter for filter file path.
func (fn FileFilterFunc) FilterFile(filePath, filename string) bool {
	return fn(filePath, filename)
}

// DirFilter for filter dir path.
type DirFilter interface {
	FilterDir(dirPath, dirName string) bool
}

// DirFilterFunc for filter file path.
type DirFilterFunc func(dirPath, dirName string) bool

// FilterDir Filter for filter file path.
func (fn DirFilterFunc) FilterDir(dirPath, dirName string) bool {
	return fn(dirPath, dirName)
}

// // BodyFilter for filter file contents.
// type BodyFilter interface {
// 	FilterBody(contents, filePath string) bool
// }
//
// // BodyFilterFunc for filter file contents.
// type BodyFilterFunc func(contents, filePath string) bool
//
// // Filter for filter file path.
// func (fn BodyFilterFunc) FilterBody(contents, filePath string) bool {
// 	return fn(contents, filePath)
// }

// FilterFunc for filter file path.
type FilterFunc func(filePath, filename string) bool

// Filter for filter file path.
func (fn FilterFunc) Filter(filePath, filename string) bool {
	return fn(filePath, filename)
}

//
// ------------------ built in file path filters ------------------
//

// ExtFilterFunc filter filepath by given file ext.
//
// Usage:
//	f := EmptyFinder()
//	f.AddFilter(ExtFilterFunc([]string{".go", ".md"}, true))
//	f.AddFilter(ExtFilterFunc([]string{".log", ".tmp"}, false))
func ExtFilterFunc(exts []string, include bool) FileFilterFunc {
	return func(filePath, _ string) bool {
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
//	f := EmptyFinder()
//	f.AddFilter(SuffixFilterFunc([]string{"util.go", ".md"}, true))
//	f.AddFilter(SuffixFilterFunc([]string{"_test.go", ".log"}, false))
func SuffixFilterFunc(suffixes []string, include bool) FileFilterFunc {
	return func(filePath, _ string) bool {
		for _, sfx := range suffixes {
			if strings.HasSuffix(filePath, sfx) {
				return include
			}
		}
		return !include
	}
}

// PathNameFilterFunc filter filepath by given path names.
func PathNameFilterFunc(names []string, include bool) FileFilterFunc {
	return func(filePath, _ string) bool {
		for _, name := range names {
			if strings.Contains(filePath, name) {
				return include
			}
		}
		return !include
	}
}

// DotFileFilterFunc filter dot filename. eg: ".gitignore"
func DotFileFilterFunc(include bool) FileFilterFunc {
	return func(filePath, filename string) bool {
		// filename := path.Base(filePath)
		if filename[0] == '.' {
			return include
		}
		return !include
	}
}

// ModTimeFilterFunc filter file by modify time.
func ModTimeFilterFunc(limitSec int, op rune, include bool) FileFilterFunc {
	return func(filePath, filename string) bool {
		fi, err := os.Stat(filePath)
		if err != nil {
			return !include
		}

		now := time.Now().Second()
		if op == '>' {
			if now-fi.ModTime().Second() > limitSec {
				return include
			}
			return !include
		}

		// '<'
		if now-fi.ModTime().Second() < limitSec {
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
func GlobFilterFunc(patterns []string, include bool) FileFilterFunc {
	return func(_, filename string) bool {
		for _, pattern := range patterns {
			if ok, _ := path.Match(pattern, filename); ok {
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
func RegexFilterFunc(pattern string, include bool) FileFilterFunc {
	reg := regexp.MustCompile(pattern)

	return func(_, filename string) bool {
		return reg.MatchString(filename)
	}
}

//
// ----------------- built in dir path filters -----------------
//

// DotDirFilterFunc filter dot dirname. eg: ".idea"
func DotDirFilterFunc(include bool) DirFilterFunc {
	return func(_, dirname string) bool {
		if dirname[0] == '.' {
			return include
		}
		return !include
	}
}

// DirNameFilterFunc filter filepath by given dir names.
func DirNameFilterFunc(names []string, include bool) DirFilterFunc {
	return func(_, dirName string) bool {
		for _, name := range names {
			if dirName == name {
				return include
			}
		}
		return !include
	}
}
