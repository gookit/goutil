package fsutil_test

import (
	"io/fs"
	"strings"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestSearchNameUp(t *testing.T) {
	p := fsutil.SearchNameUp("testdata", "finder")
	assert.NotEmpty(t, p)
	assert.True(t, strings.HasSuffix(p, "fsutil"))

	p = fsutil.SearchNameUp("testdata", ".dotdir")
	assert.NotEmpty(t, p)
	assert.True(t, strings.HasSuffix(p, "testdata"))

	p = fsutil.SearchNameUp("testdata", "test.jpg")
	assert.NotEmpty(t, p)
	assert.True(t, strings.HasSuffix(p, "testdata"))

	p = fsutil.SearchNameUp("testdata", "not-exists")
	assert.Empty(t, p)
}

func TestFindInDir(t *testing.T) {
	err := fsutil.FindInDir("path-not-exist", nil)
	assert.NoErr(t, err)

	err = fsutil.FindInDir("testdata/test.jpg", nil)
	assert.NoErr(t, err)

	files := make([]string, 0, 8)
	err = fsutil.FindInDir("testdata", func(fPath string, de fs.DirEntry) error {
		files = append(files, fPath)
		return nil
	})

	dump.P(files)
	assert.NoErr(t, err)
	assert.True(t, len(files) > 0)

	files = files[:0]
	err = fsutil.FindInDir("testdata", func(fPath string, de fs.DirEntry) error {
		files = append(files, fPath)
		return nil
	}, func(fPath string, de fs.DirEntry) bool {
		return !strings.HasPrefix(de.Name(), ".")
	})
	assert.NoErr(t, err)
	assert.True(t, len(files) > 0)

	err = fsutil.FindInDir("testdata", func(fPath string, de fs.DirEntry) error {
		return errorx.Raw("handle error")
	})
	assert.Err(t, err)
}
