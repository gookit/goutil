package finder_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/fsutil/finder"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

func newMockElem(fp string, isDir ...bool) finder.Elem {
	return finder.NewElem(fp, testutil.NewDirEnt(fp, isDir...))
}

func TestFilters_simple(t *testing.T) {
	el := newMockElem("path/some.txt")
	fn := finder.FilterFunc(func(el finder.Elem) bool {
		return false
	})

	assert.False(t, fn(el))

	// with name
	fn = finder.WithName("some.txt")
	assert.True(t, fn(el))
	fn = finder.WithName("not-exist.txt")
	assert.False(t, fn(el))

	// without name
	fn = finder.WithoutName("some.txt")
	assert.False(t, fn(el))
	fn = finder.WithoutName("not-exist.txt")
	assert.True(t, fn(el))

	// with ext
	fn = finder.WithExt(".txt")
	assert.True(t, fn(el))
	fn = finder.WithExt(".js")
	assert.False(t, fn(el))

	// without ext
	fn = finder.WithoutExt(".txt")
	assert.False(t, fn(el))
	fn = finder.WithoutExt(".js")
	assert.True(t, fn(el))

	// with suffix
	fn = finder.WithSuffix("me.txt")
	assert.True(t, fn(el))
	fn = finder.WithSuffix("not-exist.txt")
	assert.False(t, fn(el))
}

func TestExtFilterFunc(t *testing.T) {
	ent := &testutil.DirEnt{}

	fn := finder.WithExt(".log")
	assert.True(t, fn(finder.NewElem("info.log", ent)))
	assert.False(t, fn(finder.NewElem("info.tmp", ent)))

	fn = finder.WithoutExt(".log")
	assert.False(t, fn(finder.NewElem("info.log", ent)))
	assert.True(t, fn(finder.NewElem("info.tmp", ent)))
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

	ent := &testutil.DirEnt{}

	for _, tt := range tests {
		el := finder.NewElem(tt.filePath, ent)
		fn := finder.WithRegexMatch(tt.pattern)
		assert.Eq(t, tt.match, fn(el))
	}

	t.Run("exclude", func(t *testing.T) {
		for _, tt := range tests {
			el := finder.NewElem(tt.filePath, ent)
			fn := finder.WithoutRegexMatch(tt.pattern)
			assert.Eq(t, !tt.match, fn(el))
		}
	})
}

func TestDotDirFilter(t *testing.T) {
	f := finder.EmptyFinder().
		ScanDir("./testdata")

	fmt.Println("no limits:")
	fmt.Println(f)

	dirName := ".dotdir"
	assert.Contains(t, f.FindPaths(), dirName)

	f = finder.EmptyFinder().
		ScanDir("./testdata").
		NoDotDir()

	fmt.Println("NoDotDir limits:")
	fmt.Println(f.Config())
	assert.NotContains(t, f.FindPaths(), dirName)

	f = finder.NewEmpty().
		ScanDir("./testdata").
		WithDir(finder.DotDirFilter(false))

	fmt.Println("DotDirFilter limits:")
	fmt.Println(f)
	assert.NotContains(t, f.FindPaths(), dirName)
}
