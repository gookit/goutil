package fsutil_test

import (
	"os"
	"strings"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestExpandPath(t *testing.T) {
	path := "~/.kite"

	assert.NotEq(t, path, fsutil.Expand(path))
	assert.NotEq(t, path, fsutil.ExpandPath(path))

	assert.Eq(t, "", fsutil.Expand(""))
	assert.Eq(t, "/path/to", fsutil.Expand("/path/to"))
}

func TestFindInDir(t *testing.T) {
	err := fsutil.FindInDir("path-not-exist", nil)
	assert.NoErr(t, err)

	err = fsutil.FindInDir("testdata/test.jpg", nil)
	assert.NoErr(t, err)

	files := make([]string, 0, 8)
	err = fsutil.FindInDir("testdata", func(fPath string, fi os.FileInfo) error {
		files = append(files, fPath)
		return nil
	})

	dump.P(files)
	assert.NoErr(t, err)
	assert.True(t, len(files) > 0)

	files = files[:0]
	err = fsutil.FindInDir("testdata", func(fPath string, fi os.FileInfo) error {
		files = append(files, fPath)
		return nil
	}, func(fPath string, fi os.FileInfo) bool {
		return !strings.HasPrefix(fi.Name(), ".")
	})
	assert.NoErr(t, err)
	assert.True(t, len(files) > 0)

	err = fsutil.FindInDir("testdata", func(fPath string, fi os.FileInfo) error {
		return errorx.Raw("handle error")
	})
	assert.Err(t, err)
}
