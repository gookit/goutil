package fsutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/stretchr/testify/assert"
)

func TestEmptyFinder(t *testing.T) {
	f := fsutil.EmptyFinder()

	f.
		AddDir("./testdata").
		AddFile("finder.go").
		Each(func(filePath string) {
		fmt.Println(filePath)
	})
}

var testFiles = []string{
	"info.log",
	"error.log",
	"cache.tmp",
	"/some/path/to/info.log",
	"/some/path1/to/cache.tmp",
}

func TestExtFilterFunc(t *testing.T) {
	fn := fsutil.ExtFilterFunc([]string{".log"}, true)
	assert.True(t, fn("info.log"))
	assert.False(t, fn("info.tmp"))

	fn = fsutil.ExtFilterFunc([]string{".log"}, false)
	assert.False(t, fn("info.log"))
	assert.True(t, fn("info.tmp"))
}
