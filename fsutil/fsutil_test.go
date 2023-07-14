package fsutil_test

import (
	"io/fs"
	"testing"

	"github.com/gookit/goutil/basefn"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestMain(m *testing.M) {
	err := fsutil.RemoveSub("./testdata", func(fPath string, ent fs.DirEntry) bool {
		return fsutil.PathMatch(ent.Name(), "*.txt")
	})
	basefn.MustOK(err)

	m.Run()
}

func TestTempDir(t *testing.T) {
	dir, err := fsutil.TempDir("testdata", "temp.*")
	assert.NoErr(t, err)
	assert.True(t, fsutil.IsDir(dir))
	assert.NoErr(t, fsutil.Remove(dir))
}

func TestSplitPath(t *testing.T) {
	dir, file := fsutil.SplitPath("/path/to/dir/some.txt")
	assert.Eq(t, "/path/to/dir/", dir)
	assert.Eq(t, "some.txt", file)
}

func TestToAbsPath(t *testing.T) {
	assert.Eq(t, "", fsutil.ToAbsPath(""))
	assert.Eq(t, "/path/to/dir/", fsutil.ToAbsPath("/path/to/dir/"))
	assert.Neq(t, "~/path/to/dir", fsutil.ToAbsPath("~/path/to/dir"))
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
