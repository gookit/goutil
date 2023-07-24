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
	assert.Eq(t, "", fsutil.Suffix("testdata/testjpg"))
	assert.Eq(t, "", fsutil.Extname("testdata/testjpg"))
	assert.Eq(t, ".jpg", fsutil.FileExt("testdata/test.jpg"))
	assert.Eq(t, ".jpg", fsutil.Suffix("testdata/test.jpg"))
	assert.Eq(t, "jpg", fsutil.Extname("testdata/test.jpg"))

	// IsZipFile
	assert.False(t, fsutil.IsZipFile("testdata/not-exists-file"))
	assert.False(t, fsutil.IsZipFile("testdata/test.jpg"))
	assert.Eq(t, "test.jpg", fsutil.PathName("testdata/test.jpg"))

	assert.Eq(t, "test.jpg", fsutil.Name("path/to/test.jpg"))
	assert.Eq(t, "", fsutil.Name(""))

	assert.NotEmpty(t, fsutil.DirPath("path/to/test.jpg"))
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
	assert.False(t, fsutil.IsEmptyDir("testdata"))
	assert.False(t, fsutil.IsEmptyDir("testdata/not-exist-dir"))
}

func TestIsAbsPath(t *testing.T) {
	assert.True(t, fsutil.IsAbsPath("/data/some.txt"))
	assert.False(t, fsutil.IsAbsPath(""))
	assert.False(t, fsutil.IsAbsPath("some.txt"))
	assert.NoErr(t, fsutil.DeleteIfFileExist("/not-exist"))
}

func TestGlobMatch(t *testing.T) {
	tests := []struct {
		p, s string
		want bool
	}{
		{"a*", "abc", true},
		{"ab.*.ef", "ab.cd.ef", true},
		{"ab.*.*", "ab.cd.ef", true},
		{"ab.cd.*", "ab.cd.ef", true},
		{"ab.*", "ab.cd.ef", true},
		{"a*/b", "a/c/b", false},
		{"a*", "a/c/b", false},
		{"a**", "a/c/b", false},
	}

	for _, tt := range tests {
		assert.Eq(t, tt.want, fsutil.PathMatch(tt.p, tt.s), "case %v", tt)
	}

	assert.False(t, fsutil.PathMatch("ab", "abc"))
	assert.True(t, fsutil.PathMatch("ab*", "abc"))
}
