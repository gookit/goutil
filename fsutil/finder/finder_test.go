package finder_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gookit/goutil/fsutil/finder"
	"github.com/stretchr/testify/assert"
)

func TestEmptyFinder(t *testing.T) {
	f := finder.EmptyFinder()

	f.
		AddDir("./testdata").
		AddFile("finder.go").
		NoDotFile().
		// NoDotDir().
		Find().
		Each(func(filePath string) {
			fmt.Println(filePath)
		})

	assert.NotEmpty(t, f.FilePaths())

	f.Reset()
	assert.Empty(t, f.FilePaths())
}

func TestNewFinder(t *testing.T) {
	finder.NewFinder([]string{"./testdata"}).
		AddFile("finder.go").
		NoDotDir().
		EachStat(func(fi os.FileInfo, filePath string) {
			fmt.Println(filePath, "=>", fi.ModTime())
		})
}

func TestDotFileFilterFunc(t *testing.T) {
	f := finder.EmptyFinder().
		AddDir("./testdata").
		Find()
	fmt.Println("no limits:")
	fmt.Println(f)

	fileName := ".env"
	assert.Contains(t, f.String(), fileName)

	f = finder.EmptyFinder().
		AddDir("./testdata").
		NoDotFile().
		Find()
	fmt.Println("NoDotFile limits:")
	fmt.Println(f)
	assert.NotContains(t, f.String(), fileName)

	f = finder.EmptyFinder().
		AddDir("./testdata").
		WithFilter(finder.DotFileFilterFunc(false)).
		Find()

	fmt.Println("DotFileFilterFunc limits:")
	fmt.Println(f)
	assert.NotContains(t, f.String(), fileName)
}

func TestDotDirFilterFunc(t *testing.T) {
	f := finder.EmptyFinder().
		AddDir("./testdata").
		Find()
	fmt.Println("no limits:")
	fmt.Println(f)

	dirName := ".dotdir"
	assert.Contains(t, f.String(), dirName)

	f = finder.EmptyFinder().
		AddDir("./testdata").
		NoDotDir().
		Find()
	fmt.Println("NoDotDir limits:")
	fmt.Println(f)
	assert.NotContains(t, f.String(), dirName)

	f = finder.EmptyFinder().
		AddDir("./testdata").
		WithDirFilter(finder.DotDirFilterFunc(false)).
		Find()

	fmt.Println("DotDirFilterFunc limits:")
	fmt.Println(f)
	assert.NotContains(t, f.String(), dirName)
}

var testFiles = []string{
	"info.log",
	"error.log",
	"cache.tmp",
	"/some/path/to/info.log",
	"/some/path1/to/cache.tmp",
}

func TestExtFilterFunc(t *testing.T) {
	fn := finder.ExtFilterFunc([]string{".log"}, true)
	assert.True(t, fn("info.log", ""))
	assert.False(t, fn("info.tmp", ""))

	fn = finder.ExtFilterFunc([]string{".log"}, false)
	assert.False(t, fn("info.log", ""))
	assert.True(t, fn("info.tmp", ""))

}
