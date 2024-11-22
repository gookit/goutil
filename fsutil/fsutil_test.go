package fsutil_test

import (
	"fmt"
	"io/fs"
	"testing"

	"github.com/gookit/goutil/basefn"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("before test ... clean testdata/*.txt files")
	err := fsutil.RemoveSub("testdata", func(fPath string, ent fs.DirEntry) bool {
		return fsutil.PathMatch("*.txt", ent.Name())
	})
	basefn.MustOK(err)

	m.Run()
}

func TestSplitPath(t *testing.T) {
	dir, file := fsutil.SplitPath("/path/to/dir/some.txt")
	assert.Eq(t, "/path/to/dir/", dir)
	assert.Eq(t, "some.txt", file)

	assert.NotEmpty(t, fsutil.PathSep)
}

func TestToAbsPath(t *testing.T) {
	assert.Eq(t, "/path/to/dir/", fsutil.ToAbsPath("/path/to/dir/"))
	assert.Neq(t, "~/path/to/dir", fsutil.ToAbsPath("~/path/to/dir"))
	assert.NotEmpty(t, fsutil.ToAbsPath(""))
	assert.Neq(t, ".", fsutil.ToAbsPath("."))
	assert.Neq(t, "..", fsutil.ToAbsPath(".."))
	assert.Neq(t, "./", fsutil.ToAbsPath("./"))
	assert.Neq(t, "../", fsutil.ToAbsPath("../"))
}

func TestSlashPath(t *testing.T) {
	assert.Eq(t, "/path/to/dir", fsutil.SlashPath("/path/to/dir"))
	assert.Eq(t, "/path/to/dir", fsutil.UnixPath("/path/to/dir"))
	assert.Eq(t, "/path/to/dir", fsutil.UnixPath("\\path\\to\\dir"))
}
