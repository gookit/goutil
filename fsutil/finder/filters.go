package finder

import (
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/mathutil"
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

// WithDotFile include dot filename.
func WithDotFile() FilterFunc { return dotFileFilter(true) }

// WithoutDotFile exclude dot filename.
func WithoutDotFile() FilterFunc { return dotFileFilter(false) }

// dotFileFilter filter dot filename. eg: ".gitignore"
func dotFileFilter(include bool) FilterFunc {
	return func(el Elem) bool {
		name := el.Name()
		if len(name) > 0 && name[0] == '.' {
			return include
		}
		return !include
	}
}

// DotDirFilter filter dot dirname. eg: ".idea"
func DotDirFilter(include bool) FilterFunc {
	return func(el Elem) bool {
		if el.IsDir() && el.Name()[0] == '.' {
			return include
		}
		return !include
	}
}

// WithExt filter filepath by given file ext.
//
// Usage:
//
//	f := NewFinder('path/to/dir')
//	f.AddFilter(WithExt(".go"))
//	f.AddExFilter(WithoutExt(".md"))
func WithExt(exts ...string) FilterFunc { return fileExtFilter(true, exts) }

// WithExts filter filepath by given file ext.
func WithExts(exts []string) FilterFunc { return fileExtFilter(true, exts) }

// IncludeExts filter filepath by given file ext.
func IncludeExts(exts []string) FilterFunc { return fileExtFilter(true, exts) }

// WithoutExt filter filepath by given file ext.
func WithoutExt(exts ...string) FilterFunc { return fileExtFilter(false, exts) }

// WithoutExts filter filepath by given file ext.
func WithoutExts(exts []string) FilterFunc { return fileExtFilter(false, exts) }

// ExcludeExts filter filepath by given file ext.
func ExcludeExts(exts []string) FilterFunc { return fileExtFilter(false, exts) }

// fileExtFilter filter filepath by given file ext.
func fileExtFilter(include bool, exts []string) FilterFunc {
	return func(el Elem) bool {
		if len(exts) == 0 {
			return true
		}
		return isContains(path.Ext(el.Name()), exts, include)
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

// WithName filter filepath by given names.
func WithName(names ...string) FilterFunc { return MatchNames(true, names) }

// WithNames filter filepath by given names.
func WithNames(names []string) FilterFunc { return MatchNames(true, names) }

// IncludeNames filter filepath by given names.
func IncludeNames(names []string) FilterFunc { return MatchNames(true, names) }

// WithoutName filter filepath by given names.
func WithoutName(names ...string) FilterFunc { return MatchNames(false, names) }

// WithoutNames filter filepath by given names.
func WithoutNames(names []string) FilterFunc { return MatchNames(false, names) }

// ExcludeNames filter filepath by given names.
func ExcludeNames(names []string) FilterFunc { return MatchNames(false, names) }

// MatchName filter filepath by given names.
func MatchName(names ...string) FilterFunc { return MatchNames(names) }

// MatchNames filter filepath by given names.
func MatchNames(names []string) FilterFunc {
	return func(el Elem) bool {
		elName := el.Name()
		for _, name := range names {
			if name == elName || fsutil.PathMatch(name, elName) {
				return true
			}
		}
		return false
	}
}

// WithPrefix include filepath by check given prefixes.
//
// Usage:
//
//	f := NewFinder('path/to/dir')
//	f.AddFilter(finder.WithPrefix("app_", "README"))
func WithPrefix(prefixes ...string) FilterFunc { return prefixFilter(true, prefixes...) }

// WithoutPrefix exclude filepath by check given prefixes.
func WithoutPrefix(prefixes ...string) FilterFunc { return prefixFilter(false, prefixes...) }

// prefixFilter filter filepath by check name has prefixes.
func prefixFilter(include bool, prefixes ...string) FilterFunc {
	return func(el Elem) bool {
		for _, pfx := range prefixes {
			if strings.HasPrefix(el.Name(), pfx) {
				return include
			}
		}
		return !include
	}
}

// WithSuffix include filepath by check given suffixes.
//
// Usage:
//
//	f := NewFinder('path/to/dir')
//	f.AddFilter(finder.WithSuffix("util.go", "en.md"))
//	f.AddFilter(finder.WithoutSuffix("_test.go", ".log"))
func WithSuffix(suffixes ...string) FilterFunc { return suffixFilter(true, suffixes...) }

// WithoutSuffix exclude filepath by check given suffixes.
func WithoutSuffix(suffixes ...string) FilterFunc { return suffixFilter(false, suffixes...) }

// suffixFilter filter filepath by check path has suffixes.
func suffixFilter(include bool, suffixes ...string) FilterFunc {
	return func(el Elem) bool {
		for _, sfx := range suffixes {
			if strings.HasSuffix(el.Path(), sfx) {
				return include
			}
		}
		return !include
	}
}

// WithPath include file/dir by given sub paths.
//
// Usage:
//
//	f := NewFinder('path/to/dir')
//	f.AddFilter(WithPath("need/path"))
func WithPath(subPaths ...string) FilterFunc { return pathFilter(true, subPaths) }

// WithPaths include file/dir by given sub paths.
func WithPaths(subPaths []string) FilterFunc { return pathFilter(true, subPaths) }

// IncludePaths include file/dir by given sub paths.
func IncludePaths(subPaths []string) FilterFunc { return pathFilter(true, subPaths) }

// WithoutPath exclude file/dir by given sub paths.
func WithoutPath(subPaths ...string) FilterFunc { return pathFilter(false, subPaths) }

// WithoutPaths exclude file/dir by given sub paths.
func WithoutPaths(subPaths []string) FilterFunc { return pathFilter(false, subPaths) }

// ExcludePaths exclude file/dir by given sub paths.
func ExcludePaths(subPaths []string) FilterFunc { return pathFilter(false, subPaths) }

// pathFilter filter file/dir by given sub paths.
func pathFilter(include bool, subPaths []string) FilterFunc {
	return func(el Elem) bool {
		for _, subPath := range subPaths {
			if strings.Contains(el.Path(), subPath) {
				return include
			}
		}
		return !include
	}
}

// WithGlobMatch include filepath by given patterns.
//
// Usage:
//
//	f := NewFinder('path/to/dir')
//	f.AddFilter(WithGlobMatch("*_test.go"))
func WithGlobMatch(patterns ...string) FilterFunc { return globFilter(true, patterns) }

func WithGlobMatches(patterns []string) FilterFunc { return globFilter(true, patterns) }

// WithoutGlobMatch exclude filepath by given patterns.
func WithoutGlobMatch(patterns ...string) FilterFunc { return globFilter(false, patterns) }

// WithoutGlobMatches exclude filepath by given patterns.
func WithoutGlobMatches(patterns []string) FilterFunc { return globFilter(false, patterns) }

// GlobFilter filter filepath by given patterns.
func globFilter(include bool, patterns []string) FilterFunc {
	return func(el Elem) bool {
		for _, pattern := range patterns {
			if ok, _ := path.Match(pattern, el.Name()); ok {
				return include
			}
		}
		return !include
	}
}

// WithRegexMatch include filepath by given regex pattern
//
// Usage:
//
//	f := NewFinder('path/to/dir')
//	f.AddFilter(WithRegexMatch(`[A-Z]\w+`))
func WithRegexMatch(pattern string) FilterFunc { return regexFilter(pattern, true) }

// WithoutRegexMatch exclude filepath by given regex pattern
func WithoutRegexMatch(pattern string) FilterFunc { return regexFilter(pattern, false) }

// regexFilter filter filepath by given regex pattern
func regexFilter(pattern string, include bool) FilterFunc {
	reg := regexp.MustCompile(pattern)

	return func(el Elem) bool {
		if reg.MatchString(el.Path()) {
			return include
		}
		return !include
	}
}

// WithNameLike include filepath by given name match.
func WithNameLike(patterns ...string) FilterFunc { return nameLikeFilter(true, patterns) }

// WithNameLikes include filepath by given name match.
func WithNameLikes(patterns []string) FilterFunc { return nameLikeFilter(true, patterns) }

// WithoutNameLike exclude filepath by given name match.
func WithoutNameLike(patterns ...string) FilterFunc { return nameLikeFilter(false, patterns) }

// WithoutNameLikes exclude filepath by given name match.
func WithoutNameLikes(patterns []string) FilterFunc { return nameLikeFilter(false, patterns) }

// nameLikeFilter filter filepath by given name match.
func nameLikeFilter(include bool, patterns []string) FilterFunc {
	return func(el Elem) bool {
		for _, pattern := range patterns {
			if strutil.LikeMatch(pattern, el.Name()) {
				return include
			}
		}
		return !include
	}
}

//
// ----------------- built in file info filters -----------------
//

// WithModTime filter file by modify time.
//
// Note: if time is zero, it will be ignored.
//
// Usage:
//
//	f := NewFinder('path/to/dir')
//	// -600 seconds to now(last 10 minutes)
//	f.AddFilter(WithModTime(timex.NowAddSec(-600), timex.ZeroTime))
//	// before 600 seconds(before 10 minutes)
//	f.AddFilter(WithModTime(timex.ZeroTime, timex.NowAddSec(-600)))
func WithModTime(start, end time.Time) FilterFunc {
	return modTimeFilter(start, end, true)
}

// WithoutModTime filter file by modify time.
func WithoutModTime(start, end time.Time) FilterFunc {
	return modTimeFilter(start, end, false)
}

// modTimeFilter filter file by modify time.
func modTimeFilter(start, end time.Time, include bool) FilterFunc {
	return func(el Elem) bool {
		fi, err := el.Info()
		if err != nil {
			return !include
		}

		if timex.InRange(fi.ModTime(), start, end) {
			return include
		}
		return !include
	}
}

// WithHumanModTime filter file by modify time string.
//
// Usage:
//
//	f := EmptyFinder()
//	f.AddFilter(WithHumanModTime(">10m")) // before 10 minutes
//	f.AddFilter(WithHumanModTime("<10m")) // latest 10 minutes, to Now
func WithHumanModTime(expr string) FilterFunc { return humanModTimeFilter(expr, true) }

// WithoutHumanModTime filter file by modify time string.
func WithoutHumanModTime(expr string) FilterFunc { return humanModTimeFilter(expr, false) }

var timeNumReg = regexp.MustCompile(`(-?\d+)`)

// humanModTimeFilter filter file by modify time string.
func humanModTimeFilter(expr string, include bool) FilterFunc {
	opt := &timex.ParseRangeOpt{AutoSort: true}
	// convert > to <, < to >
	expr = strutil.Replaces(expr, map[string]string{">": "<", "<": ">"})
	expr = timeNumReg.ReplaceAllStringFunc(expr, func(s string) string {
		if s[0] == '-' {
			return s
		}
		return "-" + s
	})

	start, end, err := timex.ParseRange(expr, opt)
	if err != nil {
		panic(err)
	}

	return func(elem Elem) bool {
		fi, err := elem.Info()
		if err != nil {
			return !include
		}

		if timex.InRange(fi.ModTime(), start, end) {
			return include
		}
		return !include
	}
}

// WithFileSize filter file by file size. unit: byte
func WithFileSize(min, max uint64) FilterFunc { return fileSizeFilter(min, max, true) }

// WithoutFileSize filter file by file size. unit: byte
func WithoutFileSize(min, max uint64) FilterFunc { return fileSizeFilter(min, max, false) }

// fileSizeFilter filter file by file size. unit: byte
func fileSizeFilter(min, max uint64, include bool) FilterFunc {
	return func(el Elem) bool {
		if el.IsDir() {
			return false
		}

		fi, err := el.Info()
		if err != nil {
			return false
		}

		if mathutil.InUintRange(uint64(fi.Size()), min, max) {
			return include
		}
		return !include
	}
}

// WithHumanSize filter file by file size string.
func WithHumanSize(expr string) FilterFunc { return humanSizeFilter(expr, true) }

// WithoutHumanSize filter file by file size string.
func WithoutHumanSize(expr string) FilterFunc { return humanSizeFilter(expr, false) }

// humanSizeFilter filter file by file size string. eg: ">1k", "<2m", "1g~3g"
func humanSizeFilter(expr string, include bool) FilterFunc {
	min, max, err := strutil.ParseSizeRange(expr, nil)
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

		if mathutil.InUintRange(uint64(fi.Size()), min, max) {
			return include
		}
		return !include
	}
}
