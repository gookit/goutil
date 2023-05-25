package finder

import (
	"io/fs"
	"path"
	"regexp"
	"strings"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/timex"
)

// ------------------ built in filters ------------------

// OnlyFileFilter only allow file path.
var OnlyFileFilter = FilterFunc(func(el Elem) bool {
	return !el.IsDir()
})

// OnlyDirFilter only allow dir path.
var OnlyDirFilter = FilterFunc(func(el Elem) bool {
	return el.IsDir()
})

// DotDirFilter filter dot dirname. eg: ".idea"
func DotDirFilter(include bool) FilterFunc {
	return func(el Elem) bool {
		if el.IsDir() && el.Path()[0] == '.' {
			return include
		}
		return !include
	}
}

// OnlyFileFilter2 filter only file path.
func OnlyFileFilter2(exts ...string) FilterFunc {
	return func(el Elem) bool {
		if el.IsDir() {
			return false
		}

		if len(exts) == 0 {
			return true
		}
		return isContains(path.Ext(el.Name()), exts, true)
	}
}

func isContains(sub string, list []string, include bool) bool {
	for _, s := range list {
		if s == sub {
			return include
		}
	}
	return !include
}

// ExtFilter filter filepath by given file ext.
//
// Usage:
//
//	f := NewEmpty()
//	f.AddFilter(ExtFilter(".go"))
//	f.AddFilter(ExtFilter(".go", ".php"))
func ExtFilter(include bool, exts ...string) FilterFunc {
	return func(el Elem) bool {
		if len(exts) == 0 {
			return true
		}
		return isContains(path.Ext(el.Path()), exts, include)
	}
}

// NameFilter filter filepath by given names.
func NameFilter(include bool, names ...string) FilterFunc {
	return func(el Elem) bool {
		return isContains(el.Name(), names, include)
	}
}

// SuffixFilter filter filepath by check given suffixes.
//
// Usage:
//
//	f := EmptyFinder()
//	f.AddFilter(finder.SuffixFilter(true, "util.go", "en.md"))
//	f.AddFilter(finder.SuffixFilter(false, "_test.go", ".log"))
func SuffixFilter(include bool, suffixes ...string) FilterFunc {
	return func(el Elem) bool {
		for _, sfx := range suffixes {
			if strings.HasSuffix(el.Path(), sfx) {
				return include
			}
		}
		return !include
	}
}

// PathFilter filter filepath by given sub paths.
//
// Usage:
//
//	f.AddFilter(PathFilter(true, "need/path"))
func PathFilter(include bool, subPaths ...string) FilterFunc {
	return func(el Elem) bool {
		for _, subPath := range subPaths {
			if strings.Contains(el.Path(), subPath) {
				return include
			}
		}
		return !include
	}
}

// DotFileFilter filter dot filename. eg: ".gitignore"
func DotFileFilter(include bool) FilterFunc {
	return func(el Elem) bool {
		name := el.Name()
		if len(name) > 0 && name[0] == '.' {
			return include
		}
		return !include
	}
}

// GlobFilterFunc filter filepath by given patterns.
//
// Usage:
//
//	f := EmptyFiler()
//	f.AddFilter(GlobFilterFunc(true, "*_test.go"))
func GlobFilterFunc(include bool, patterns ...string) FilterFunc {
	return func(el Elem) bool {
		for _, pattern := range patterns {
			if ok, _ := path.Match(pattern, el.Path()); ok {
				return include
			}
		}
		return !include
	}
}

// RegexFilterFunc filter filepath by given regex pattern
//
// Usage:
//
//	f := EmptyFiler()
//	f.AddFilter(RegexFilterFunc(`[A-Z]\w+`, true))
func RegexFilterFunc(pattern string, include bool) FilterFunc {
	reg := regexp.MustCompile(pattern)

	return func(el Elem) bool {
		if reg.MatchString(el.Path()) {
			return include
		}
		return !include
	}
}

//
// ----------------- built in file info filters -----------------
//

// ModTimeFilter filter file by modify time.
//
// Usage:
//
//	f := EmptyFinder()
//	f.AddFilter(ModTimeFilter(600, '>', true)) // 600 seconds to Now(last 10 minutes
//	f.AddFilter(ModTimeFilter(600, '<', false)) // before 600 seconds(before 10 minutes)
func ModTimeFilter(limitSec int, op rune, include bool) FilterFunc {
	return func(el Elem) bool {
		fi, err := el.Info()
		if err != nil {
			return !include
		}

		lt := timex.Now().AddSeconds(-limitSec)
		if op == '>' {
			if lt.After(fi.ModTime()) {
				return include
			}
			return !include
		}

		// '<'
		if lt.Before(fi.ModTime()) {
			return include
		}
		return !include
	}
}

// HumanModTimeFilter filter file by modify time string.
//
// Usage:
//
//	f := EmptyFinder()
//	f.AddFilter(HumanModTimeFilter("10m", '>', true)) // 10 minutes to Now
//	f.AddFilter(HumanModTimeFilter("10m", '<', false)) // before 10 minutes
func HumanModTimeFilter(limit string, op rune, include bool) FilterFunc {
	return func(elem Elem) bool {
		fi, err := elem.Info()
		if err != nil {
			return !include
		}

		lt, err := strutil.ToDuration(limit)
		if err != nil {
			return !include
		}

		if op == '>' {
			if timex.Now().Add(-lt).After(fi.ModTime()) {
				return include
			}
			return !include
		}

		// '<'
		if timex.Now().Add(-lt).Before(fi.ModTime()) {
			return include
		}
		return !include
	}
}

// FileSizeFilter filter file by file size.
func FileSizeFilter(min, max int64, include bool) FilterFunc {
	return func(el Elem) bool {
		if el.IsDir() {
			return false
		}

		fi, err := el.Info()
		if err != nil {
			return false
		}

		return ByteSizeCheck(fi, min, max, include)
	}
}

// HumanSizeFilter filter file by file size string. eg: 1KB, 2MB, 3GB
func HumanSizeFilter(min, max string, include bool) FilterFunc {
	minSize, err := strutil.ToByteSize(min)
	if err != nil {
		panic(err)
	}

	maxSize, err := strutil.ToByteSize(max)
	if err != nil {
		panic(err)
	}

	return func(el Elem) bool {
		if el.IsDir() {
			return false
		}

		fi, err := el.Info()
		if err != nil {
			return false
		}

		return ByteSizeCheck(fi, int64(minSize), int64(maxSize), include)
	}
}

// ByteSizeCheck filter file by file size.
func ByteSizeCheck(fi fs.FileInfo, min, max int64, include bool) bool {
	if min > 0 && fi.Size() < min {
		return !include
	}

	if max > 0 && fi.Size() > max {
		return !include
	}
	return include
}
