package fsutil_test

import (
	"bytes"
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestMimeType(t *testing.T) {
	assert.Eq(t, "", fsutil.MimeType(""))
	assert.Eq(t, "", fsutil.MimeType("not-exist"))
	assert.Eq(t, "image/jpeg", fsutil.MimeType("testdata/test.jpg"))

	buf := new(bytes.Buffer)
	buf.Write([]byte("\xFF\xD8\xFF"))
	assert.Eq(t, "image/jpeg", fsutil.ReaderMimeType(buf))
	buf.Reset()

	buf.Write([]byte("text"))
	assert.Eq(t, "text/plain; charset=utf-8", fsutil.ReaderMimeType(buf))
	buf.Reset()

	buf.Write([]byte(""))
	assert.Eq(t, "", fsutil.ReaderMimeType(buf))
	buf.Reset()

	assert.True(t, fsutil.IsImageFile("testdata/test.jpg"))
}

func TestTempDir(t *testing.T) {
	dir, err := fsutil.TempDir("testdata", "temp.*")
	assert.NoErr(t, err)
	assert.True(t, fsutil.IsDir(dir))
	assert.NoErr(t, fsutil.Remove(dir))
}

func TestRealpath(t *testing.T) {
	assert.Eq(t, "/path/to/dir", fsutil.Realpath("/path/to/some/../dir"))

	dir, file := fsutil.SplitPath("/path/to/dir/some.txt")
	assert.Eq(t, "/path/to/dir/", dir)
	assert.Eq(t, "some.txt", file)
}
