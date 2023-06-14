# finder

[![GoDoc](https://godoc.org/github.com/goutil/fsutil/finder?status.svg)](https://godoc.org/github.com/goutil/fsutil/finder)

`finder` Provides a simple and convenient filedir lookup function, 
supports filtering, excluding, matching, ignoring, etc.
and with some commonly built-in matchers.

## Usage

```go
package main

import (
	"github.com/gookit/goutil/dump"
	"github.com/goutil/fsutil/finder"
)

func main() {
	ff := finder.NewFinder()
	ff.AddScan("/tmp", "/usr/local", "/usr/local/share")
	ff.ExcludeDir("abc", "def").ExcludeFile("*.log", "*.tmp")

	ss := ff.FindPaths()
	dump.P(ss)
}
```

## Built-in Matchers

```go
func FileSize(min, max uint64) MatcherFunc
func GlobMatch(patterns ...string) MatcherFunc
func GlobMatches(patterns []string) MatcherFunc
func HumanModTime(expr string) MatcherFunc
func HumanSize(expr string) MatcherFunc
func MatchDotDir() MatcherFunc
func MatchDotFile() MatcherFunc
func MatchExt(exts ...string) MatcherFunc
func MatchExts(exts []string) MatcherFunc
func MatchModTime(start, end time.Time) MatcherFunc
func MatchMtime(start, end time.Time) MatcherFunc
func MatchName(names ...string) MatcherFunc
func MatchNames(names []string) MatcherFunc
func MatchPath(subPaths []string) MatcherFunc
func MatchPaths(subPaths []string) MatcherFunc
func MatchPrefix(prefixes ...string) MatcherFunc
func MatchPrefixes(prefixes []string) MatcherFunc
func MatchSuffix(suffixes ...string) MatcherFunc
func MatchSuffixes(suffixes []string) MatcherFunc
func NameLike(patterns ...string) MatcherFunc
func NameLikes(patterns []string) MatcherFunc
func RegexMatch(pattern string) MatcherFunc
func SizeRange(min, max uint64) MatcherFunc
func StartWithDot() MatcherFunc
```

