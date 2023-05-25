package finder_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gookit/goutil/fsutil/finder"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestEmptyFinder(t *testing.T) {
	f := finder.EmptyFinder()

	f.
		AddDir("./testdata").
		NoDotFile().
		CacheResult().
		// NoDotDir().
		EachPath(func(filePath string) {
			fmt.Println(filePath)
		})

	assert.NotEmpty(t, f.FindPaths())

	f.Reset()
	assert.Empty(t, f.FindPaths())
}

func TestNewFinder(t *testing.T) {
	finder.NewFinder("./testdata").
		NoDotDir().
		EachStat(func(fi os.FileInfo, filePath string) {
			fmt.Println(filePath, "=>", fi.ModTime())
		})
}

func TestDotFileFilterFunc(t *testing.T) {
	f := finder.NewEmpty().
		AddDir("./testdata")
	assert.NotEmpty(t, f.String())

	fmt.Println("no limits:")
	fmt.Println(f)

	fileName := ".env"
	assert.Contains(t, f.FindPaths(), fileName)

	f = finder.EmptyFinder().
		AddDir("./testdata").
		NoDotFile()

	fmt.Println("NoDotFile limits:")
	fmt.Println(f)
	assert.NotContains(t, f.FindPaths(), fileName)

	f = finder.EmptyFinder().
		AddDir("./testdata").
		WithFilter(finder.DotFileFilter(false))

	fmt.Println("DotFileFilter limits:")
	fmt.Println(f)
	assert.NotContains(t, f.FindPaths(), fileName)
}

func TestDotDirFilterFunc(t *testing.T) {
	f := finder.EmptyFinder().
		AddDir("./testdata")

	fmt.Println("no limits:")
	fmt.Println(f)

	dirName := ".dotdir"
	assert.Contains(t, f.FindPaths(), dirName)

	f = finder.EmptyFinder().
		AddDir("./testdata").
		NoDotDir()

	fmt.Println("NoDotDir limits:")
	fmt.Println(f.Config())
	assert.NotContains(t, f.FindPaths(), dirName)

	f = finder.NewEmpty().
		AddDir("./testdata").
		WithDirFilter(finder.DotDirFilter(false))

	fmt.Println("DotDirFilter limits:")
	fmt.Println(f)
	assert.NotContains(t, f.FindPaths(), dirName)
}

var testFiles = []string{
	"info.log",
	"error.log",
	"cache.tmp",
	"/some/path/to/info.log",
	"/some/path1/to/cache.tmp",
}

func TestExtFilterFunc(t *testing.T) {
	ent := &testutil.DirEnt{}

	fn := finder.ExtFilter(true, ".log")
	assert.True(t, fn(finder.NewElem("info.log", ent)))
	assert.False(t, fn(finder.NewElem("info.tmp", ent)))

	fn = finder.ExtFilter(false, ".log")
	assert.False(t, fn(finder.NewElem("info.log", ent)))
	assert.True(t, fn(finder.NewElem("info.tmp", ent)))

}
