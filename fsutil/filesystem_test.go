package fsutil_test

import (
	"bytes"
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/stretchr/testify/assert"
)

func TestPathExists(t *testing.T) {
	assert.False(t, fsutil.FileExists(""))
	assert.False(t, fsutil.PathExists("/not-exist"))
}

func TestIsFile(t *testing.T) {
	assert.False(t, fsutil.IsFile(""))
	assert.False(t, fsutil.IsFile("/not-exist"))
}

func TestIsDir(t *testing.T) {
	assert.False(t, fsutil.IsDir(""))
	assert.False(t, fsutil.IsDir("/not-exist"))
}

func TestIsAbsPath(t *testing.T) {
	assert.True(t, fsutil.IsAbsPath("/data/some.txt"))
}

func TestMimeType(t *testing.T) {
	assert.Equal(t, "", fsutil.MimeType("not-exist"))
	assert.Equal(t,"image/jpeg", fsutil.MimeType("testdata/test.jpg"))

	buf := new(bytes.Buffer)
	buf.Write([]byte("\xFF\xD8\xFF"))
	assert.Equal(t,"image/jpeg", fsutil.ReaderMimeType(buf))
	buf.Reset()

	buf.Write([]byte("text"))
	assert.Equal(t,"text/plain; charset=utf-8", fsutil.ReaderMimeType(buf))
	buf.Reset()

	// buf.Write([]byte(""))
	// assert.Equal(t,"", fsutil.ReaderMimeType(buf))
	// buf.Reset()
}
