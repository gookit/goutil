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

// MatchFile only allow file path.
var MatchFile = MatcherFunc(func(el Elem) bool {
	return !el.IsDir()
})

// MatchDir only allow dir path.
var MatchDir = MatcherFunc(func(el Elem) bool {
	return el.IsDir()
})

// StartWithDot match dot file/dir. eg: ".gitignore"
func StartWithDot() MatcherFunc {
	return func(el Elem) bool {
		name := el.Name()
		return len(name) > 0 && name[0] == '.'
	}
}

// MatchDotFile match dot filename. eg: ".idea"
func MatchDotFile() MatcherFunc {
	return func(el Elem) bool {
		return !el.IsDir() && el.Name()[0] == '.'
	}
}

// MatchDotDir match dot dirname. eg: ".idea"
func MatchDotDir() MatcherFunc {
	return func(el Elem) bool {
		return el.IsDir() && el.Name()[0] == '.'
	}
}

// MatchExt match filepath by given file ext.
//
// Usage:
//
//	f := NewFinder('path/to/dir')
//	f.Add(MatchExt(".go"))
//	f.Not(MatchExt(".md"))
func MatchExt(exts ...string) MatcherFunc { return MatchExts(exts) }

// MatchExts filter filepath by given file ext.
func MatchExts(exts []string) MatcherFunc {
	return func(el Elem) bool {
		elExt := path.Ext(el.Name())
		for _, ext := range exts {
			if ext == elExt {
				return true
			}
		}
		return false
	}
}

// MatchName match filepath by given names.
//
// Usage:
//
//	f := NewFinder('path/to/dir')
//	f.Not(MatchName("README.md", "*_test.go"))
func MatchName(names ...string) MatcherFunc { return MatchNames(names) }

// MatchNames match filepath by given names or patterns.
//
// Usage:
//
//	f.Not(MatchNames([]string{"README.md", "*_test.go"}))
func MatchNames(names []string) MatcherFunc {
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

// MatchPrefix match filepath by check given prefixes.
//
// Usage:
//
//	f := NewFinder('path/to/dir')
//	f.Add(finder.MatchPrefix("app_", "README"))
func MatchPrefix(prefixes ...string) MatcherFunc { return MatchPrefixes(prefixes) }

// MatchPrefixes match filepath by check given prefixes.
func MatchPrefixes(prefixes []string) MatcherFunc {
	return func(el Elem) bool {
		for _, pfx := range prefixes {
			if strings.HasPrefix(el.Name(), pfx) {
				return true
			}
		}
		return false
	}
}

// MatchSuffix match filepath by check path has suffixes.
//
// Usage:
//
//	f := NewFinder('path/to/dir')
//	f.Add(finder.MatchSuffix("util.go", "en.md"))
//	f.Not(finder.MatchSuffix("_test.go", ".log"))
func MatchSuffix(suffixes ...string) MatcherFunc { return MatchSuffixes(suffixes) }

// MatchSuffixes match filepath by check path has suffixes.
func MatchSuffixes(suffixes []string) MatcherFunc {
	return func(el Elem) bool {
		for _, sfx := range suffixes {
			if strings.HasSuffix(el.Path(), sfx) {
				return true
			}
		}
		return false
	}
}

// MatchPath match file/dir by given sub paths.
//
// Usage:
//
//	f := NewFinder('path/to/dir')
//	f.Add(MatchPath("need/path"))
func MatchPath(subPaths ...string) MatcherFunc { return MatchPaths(subPaths) }

// MatchPaths match file/dir by given sub paths.
func MatchPaths(subPaths []string) MatcherFunc {
	return func(el Elem) bool {
		for _, subPath := range subPaths {
			if strings.Contains(el.Path(), subPath) {
				return true
			}
		}
		return false
	}
}

// GlobMatch file/dir name by given patterns.
//
// Usage:
//
//	f := NewFinder('path/to/dir')
//	f.AddFilter(GlobMatch("*_test.go"))
func GlobMatch(patterns ...string) MatcherFunc { return GlobMatches(patterns) }

// GlobMatches file/dir name by given patterns.
func GlobMatches(patterns []string) MatcherFunc {
	return func(el Elem) bool {
		for _, pattern := range patterns {
			if ok, _ := path.Match(pattern, el.Name()); ok {
				return true
			}
		}
		return false
	}
}

// RegexMatch match name by given regex pattern
//
// Usage:
//
//	f := NewFinder('path/to/dir')
//	f.AddFilter(RegexMatch(`[A-Z]\w+`))
func RegexMatch(pattern string) MatcherFunc {
	reg := regexp.MustCompile(pattern)

	return func(el Elem) bool {
		return reg.MatchString(el.Name())
	}
}

// NameLike exclude filepath by given name match.
func NameLike(patterns ...string) MatcherFunc { return NameLikes(patterns) }

// NameLikes filter filepath by given name match.
func NameLikes(patterns []string) MatcherFunc {
	return func(el Elem) bool {
		for _, pattern := range patterns {
			if strutil.LikeMatch(pattern, el.Name()) {
				return true
			}
		}
		return false
	}
}

//
// ----------------- built in file info filters -----------------
//

// MatchMtime match file by modify time.
//
// Note: if time is zero, it will be ignored.
//
// Usage:
//
//	f := NewFinder('path/to/dir')
//	// -600 seconds to now(last 10 minutes)
//	f.AddFile(MatchMtime(timex.NowAddSec(-600), timex.ZeroTime))
//	// before 600 seconds(before 10 minutes)
//	f.AddFile(MatchMtime(timex.ZeroTime, timex.NowAddSec(-600)))
func MatchMtime(start, end time.Time) MatcherFunc {
	return MatchModTime(start, end)
}

// MatchModTime filter file by modify time.
func MatchModTime(start, end time.Time) MatcherFunc {
	return func(el Elem) bool {
		if el.IsDir() {
			return false
		}

		fi, err := el.Info()
		if err != nil {
			return false
		}
		return timex.InRange(fi.ModTime(), start, end)
	}
}

var timeNumReg = regexp.MustCompile(`(-?\d+)`)

// HumanModTime filter file by modify time string.
//
// Usage:
//
//	f := finder.NewFinder()
//	f.Include(HumanModTime(">10m")) // before 10 minutes
//	f.Include(HumanModTime("<10m")) // latest 10 minutes, to Now
func HumanModTime(expr string) MatcherFunc {
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

	return MatchModTime(start, end)
}

// FileSize match file by file size. unit: byte
func FileSize(min, max uint64) MatcherFunc { return SizeRange(min, max) }

// SizeRange match file by file size. unit: byte
func SizeRange(min, max uint64) MatcherFunc {
	return func(el Elem) bool {
		if el.IsDir() {
			return false
		}

		fi, err := el.Info()
		if err != nil {
			return false
		}
		return mathutil.InUintRange(uint64(fi.Size()), min, max)
	}
}

// HumanSize match file by file size string. eg: ">1k", "<2m", "1g~3g"
func HumanSize(expr string) MatcherFunc {
	min, max, err := strutil.ParseSizeRange(expr, nil)
	if err != nil {
		panic(err)
	}

	return SizeRange(min, max)
}
