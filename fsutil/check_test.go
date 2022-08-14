package fsutil_test

import (
	"runtime"
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/testutil/assert"
)

//goland:noinspection GoBoolExpressions
func TestFsUtil_common(t *testing.T) {
	assert.Eq(t, "", fsutil.FileExt("testdata/testjpg"))
	assert.Eq(t, ".jpg", fsutil.Suffix("testdata/test.jpg"))

	// IsZipFile
	assert.False(t, fsutil.IsZipFile("testdata/test.jpg"))
	assert.Eq(t, "test.jpg", fsutil.PathName("testdata/test.jpg"))

	assert.Eq(t, "test.jpg", fsutil.Name("path/to/test.jpg"))

	if runtime.GOOS == "windows" {
		assert.Eq(t, "path\\to", fsutil.Dir("path/to/test.jpg"))
	} else {
		assert.Eq(t, "path/to", fsutil.Dir("path/to/test.jpg"))
	}
}

func TestPathExists(t *testing.T) {
	assert.False(t, fsutil.PathExists(""))
	assert.False(t, fsutil.PathExist("/not-exist"))
	assert.False(t, fsutil.PathExists("/not-exist"))
	assert.True(t, fsutil.PathExist("testdata/test.jpg"))
	assert.True(t, fsutil.PathExists("testdata/test.jpg"))
}

func TestIsFile(t *testing.T) {
	assert.False(t, fsutil.FileExists(""))
	assert.False(t, fsutil.IsFile(""))
	assert.False(t, fsutil.IsFile("/not-exist"))
	assert.False(t, fsutil.FileExists("/not-exist"))
	assert.True(t, fsutil.IsFile("testdata/test.jpg"))
	assert.True(t, fsutil.FileExist("testdata/test.jpg"))
}

func TestIsDir(t *testing.T) {
	assert.False(t, fsutil.IsDir(""))
	assert.False(t, fsutil.DirExist(""))
	assert.False(t, fsutil.IsDir("/not-exist"))
	assert.True(t, fsutil.IsDir("testdata"))
	assert.True(t, fsutil.DirExist("testdata"))
}

func TestIsAbsPath(t *testing.T) {
	assert.True(t, fsutil.IsAbsPath("/data/some.txt"))
	assert.NoErr(t, fsutil.DeleteIfFileExist("/not-exist"))
}
