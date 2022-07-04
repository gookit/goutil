package fsutil_test

import (
	"bytes"
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/stretchr/testify/assert"
)

func TestMimeType(t *testing.T) {
	assert.Equal(t, "", fsutil.MimeType(""))
	assert.Equal(t, "", fsutil.MimeType("not-exist"))
	assert.Equal(t, "image/jpeg", fsutil.MimeType("testdata/test.jpg"))

	buf := new(bytes.Buffer)
	buf.Write([]byte("\xFF\xD8\xFF"))
	assert.Equal(t, "image/jpeg", fsutil.ReaderMimeType(buf))
	buf.Reset()

	buf.Write([]byte("text"))
	assert.Equal(t, "text/plain; charset=utf-8", fsutil.ReaderMimeType(buf))
	buf.Reset()

	buf.Write([]byte(""))
	assert.Equal(t, "", fsutil.ReaderMimeType(buf))
	buf.Reset()

	assert.True(t, fsutil.IsImageFile("testdata/test.jpg"))
}

func TestTempDir(t *testing.T) {
	dir, err := fsutil.TempDir("testdata", "temp.*")
	assert.NoError(t, err)
	assert.DirExists(t, dir)
	assert.NoError(t, fsutil.Remove(dir))
}

func TestRealpath(t *testing.T) {
	assert.Equal(t, "/path/to/dir", fsutil.Realpath("/path/to/some/../dir"))

	dir, file := fsutil.SplitPath("/path/to/dir/some.txt")
	assert.Equal(t, "/path/to/dir/", dir)
	assert.Equal(t, "some.txt", file)
}
