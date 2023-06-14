package finder_test

import (
	"testing"

	"github.com/gookit/goutil/fsutil/finder"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

func newMockElem(fp string, isDir ...bool) finder.Elem {
	return finder.NewElem(fp, testutil.NewDirEnt(fp, isDir...))
}

func TestMultiMatcher(t *testing.T) {
	mf := finder.NewDirMatchers(finder.MatchPath("sub/path"))
	assert.NotEmpty(t, mf)

	mf.Add(finder.MatchName("dir"))

	el := newMockElem("/some/sub/path/dir", true)
	assert.True(t, mf.Apply(el))

	mf = finder.NewFileMatchers(finder.MatchPrefix("some_"))
	assert.NotEmpty(t, mf)
	mf.Add(
		finder.GlobMatch("*_file.go"),
		finder.NameLike("some%"),
	)

	el = newMockElem("/some/sub/path/some_file.go")
	assert.True(t, mf.Apply(el))
}

func TestMatchers_simple(t *testing.T) {
	el := newMockElem("path/some.txt")
	el1 := newMockElem("path/some_file.txt")
	fn := finder.MatcherFunc(func(el finder.Elem) bool {
		return false
	})

	assert.False(t, fn(el))

	// match name
	fn = finder.MatchName("some.txt")
	assert.True(t, fn(el))
	fn = finder.MatchName("not-exist.txt")
	assert.False(t, fn(el))

	// MatchExt
	fn = finder.MatchExt(".txt")
	assert.True(t, fn(el))
	fn = finder.MatchExt(".js")
	assert.False(t, fn(el))

	// MatchSuffix
	fn = finder.MatchSuffix("me.txt")
	assert.True(t, fn(el))
	fn = finder.MatchSuffix("not-exist.txt")
	assert.False(t, fn(el))

	// MatchPrefix
	fn = finder.MatchPrefix("some_")
	assert.False(t, fn(el))
	assert.True(t, fn(el1))

	// MatchPath
	fn = finder.MatchPath("path/some")
	assert.True(t, fn(el))
	fn = finder.MatchPath("not-exist/path")
	assert.False(t, fn(el))
}

func TestRegexMatch(t *testing.T) {
	tests := []struct {
		filePath string
		pattern  string
		match    bool
	}{
		{"path/to/util.go", `\.go$`, true},
		{"path/to/util.go", `\.md$`, false},
		{"path/to/util.md", `\.md$`, true},
		{"path/to/util.md", `\.go$`, false},
	}

	for _, tt := range tests {
		el := newMockElem(tt.filePath)
		fn := finder.RegexMatch(tt.pattern)
		assert.Eq(t, tt.match, fn(el))
	}
}

func TestMatchDotDir(t *testing.T) {
	f := finder.EmptyFinder().
		WithFlags(finder.FlagBoth).
		ScanDir("./testdata")

	dirName := ".dotdir"
	assert.Contains(t, f.FindNames(), dirName)

	t.Run("NoDotDir", func(t *testing.T) {
		f = finder.EmptyFinder().
			ScanDir("./testdata").
			NoDotDir()

		assert.NotContains(t, f.FindNames(), dirName)
	})

	t.Run("Exclude false", func(t *testing.T) {
		f = finder.NewEmpty().
			WithStrFlag("dir").
			ScanDir("./testdata").
			ExcludeDotDir(false)

		assert.Contains(t, f.FindNames(), dirName)
	})
}
