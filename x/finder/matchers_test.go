package finder_test

import (
	"testing"
	"time"

	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/x/fakeobj"
	"github.com/gookit/goutil/x/finder"
)

func newMockElem(fp string, isDir ...bool) finder.Elem {
	return finder.NewElem(fp, testutil.NewDirEnt(fp, isDir...))
}

func TestStartWithDot(t *testing.T) {
	tests := []struct {
		elemName string
		expected bool
	}{
		{"abc", false},
		{".gitignore", true},
		{"README.md", false},
		{".env", true},
	}

	for _, test := range tests {
		matcher := finder.StartWithDot()
		result := matcher(newMockElem(test.elemName))
		assert.Equal(t, test.expected, result, "case:"+test.elemName)
	}

	t.Run("MatchDotFile", func(t *testing.T) {
		matcher := finder.MatchDotFile()
		assert.False(t, matcher(newMockElem("some.txt")))
		assert.True(t, matcher(newMockElem(".gitignore")))
		assert.False(t, matcher(newMockElem(".git", true)))
	})

	t.Run("MatchDotDir", func(t *testing.T) {
		matcher := finder.MatchDotDir()
		assert.False(t, matcher(newMockElem("some.txt")))
		assert.False(t, matcher(newMockElem(".gitignore")))
		assert.True(t, matcher(newMockElem(".git", true)))
	})
}

func TestGlobMatch(t *testing.T) {
	matcher := finder.GlobMatch("*.go")
	assert.True(t, matcher(newMockElem("some.go")))
	assert.False(t, matcher(newMockElem("some.txt")))
}

func TestNameLike(t *testing.T) {
	matcher := finder.NameLike("some%")
	assert.True(t, matcher(newMockElem("some_file.go")))
	assert.False(t, matcher(newMockElem("other.txt")))
}

func TestMatchModTime(t *testing.T) {
	now := time.Now()
	start := now.Add(-1 * time.Hour)
	end := now.Add(1 * time.Hour)
	matcher := finder.MatchMtime(start, end)

	// is dir
	assert.False(t, matcher(newMockElem("some-dir", true)))

	// match file mtime
	fPath := "some_file.go"
	fi := fakeobj.NewFileInfo(fPath)
	fi.WithMtime(now)

	ent := testutil.NewDirEnt(fPath)
	ent.Fi = fi

	elem := finder.NewElem(fPath, ent)
	assert.True(t, matcher(elem))

	// info with error
	ent.Err = errorx.Raw("some error for file info")
	assert.False(t, matcher(finder.NewElem(fPath, ent)))
}

func TestHumanModTime(t *testing.T) {
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
