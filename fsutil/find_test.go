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

type dirEnt struct {
	typ   fs.FileMode
	isDir bool
	name  string
}

func (d *dirEnt) Name() string {
	return d.name
}

func (d *dirEnt) IsDir() bool {
	return d.isDir
}

func (d *dirEnt) Type() fs.FileMode {
	return d.typ
}

func (d *dirEnt) Info() (fs.FileInfo, error) {
	panic("implement me")
}

func TestApplyFilters(t *testing.T) {
	e1 := &dirEnt{name: "some-backup"}
	f1 := fsutil.ExcludeSuffix("-backup")

	assert.False(t, f1("", e1))
	assert.True(t, fsutil.ApplyFilters("", e1, []fsutil.FilterFunc{f1}))
	assert.True(t, fsutil.ApplyFilters("", e1, []fsutil.FilterFunc{fsutil.OnlyFindDir}))
	assert.False(t, fsutil.ApplyFilters("", e1, []fsutil.FilterFunc{fsutil.OnlyFindFile}))
	assert.False(t, fsutil.ApplyFilters("", e1, []fsutil.FilterFunc{fsutil.ExcludeDotFile}))
	assert.False(t, fsutil.ApplyFilters("", e1, []fsutil.FilterFunc{fsutil.IncludeSuffix("-backup")}))
	assert.True(t, fsutil.ApplyFilters("", e1, []fsutil.FilterFunc{fsutil.ExcludeNames("some-backup")}))
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
