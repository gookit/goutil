# finder

[![GoDoc](https://godoc.org/github.com/goutil/x/finder?status.svg)](https://godoc.org/github.com/goutil/x/finder)

`finder` Provides a simple and convenient file/dir lookup function, 
supports filtering, excluding, matching, ignoring, etc.
and with some commonly built-in matchers.

- Support multiple paths to scan
- Support concurrency scanning
- Support find files and directories
- Support filtering, excluding, matching, ignoring, etc.
- Support built-in matchers, can also customize matchers

## Install

```shell
go get github.com/gookit/goutil/x/finder
```

## Usage

```go
package main

import (
	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/x/finder"
)

func main() {
	ff := finder.NewFinder()
	ff.AddScan("/tmp", "/usr/local", "/usr/local/share")
	ff.ExcludeDir("abc", "def").ExcludeFile("*.log", "*.tmp")
	// add built-in matchers
	ff.Exclude(finder.MatchSuffix("_test.go"), finder.MatchExt(".md"))

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

