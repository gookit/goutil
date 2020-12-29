package fsutil_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/gookit/goutil/envutil"
	"github.com/gookit/goutil/fsutil"
	"github.com/stretchr/testify/assert"
)

func TestPathExists(t *testing.T) {
	assert.False(t, fsutil.PathExists(""))
	assert.False(t, fsutil.PathExists("/not-exist"))
	assert.True(t, fsutil.PathExists("testdata/test.jpg"))
}

func TestIsFile(t *testing.T) {
	assert.False(t, fsutil.FileExists(""))
	assert.False(t, fsutil.IsFile(""))
	assert.False(t, fsutil.IsFile("/not-exist"))
	assert.False(t, fsutil.FileExists("/not-exist"))
	assert.True(t, fsutil.IsFile("testdata/test.jpg"))
}

func TestIsDir(t *testing.T) {
	assert.False(t, fsutil.IsDir(""))
	assert.False(t, fsutil.IsDir("/not-exist"))
	assert.True(t, fsutil.IsDir("testdata"))
}

func TestIsAbsPath(t *testing.T) {
	assert.True(t, fsutil.IsAbsPath("/data/some.txt"))
}

func TestMkdir(t *testing.T) {
	// TODO windows will error
	if envutil.IsWin() {
		return
	}

	err := os.Chmod("./testdata", os.ModePerm)

	if assert.NoError(t, err) {
		assert.NoError(t, fsutil.Mkdir("./testdata/sub/sub21", os.ModePerm))
		assert.NoError(t, fsutil.Mkdir("./testdata/sub/sub22", 0666))
		assert.NoError(t, fsutil.Mkdir("./testdata/sub/sub23/sub31", 0777)) // 066X will error

		assert.NoError(t, os.RemoveAll("./testdata/sub"))
	}
}

func TestCreateFile(t *testing.T) {
	// TODO windows will error
	if envutil.IsWin() {
		return
	}

	file, err := fsutil.CreateFile("./testdata/test.txt", 0664, 0666)
	if assert.NoError(t, err) {
		assert.Equal(t, "./testdata/test.txt", file.Name())
		assert.NoError(t, os.Remove(file.Name()))
	}

	file, err = fsutil.CreateFile("./testdata/sub/test.txt", 0664, 0777)
	if assert.NoError(t, err) {
		assert.Equal(t, "./testdata/sub/test.txt", file.Name())
		assert.NoError(t, os.RemoveAll("./testdata/sub"))
	}

	file, err = fsutil.CreateFile("./testdata/sub/sub2/test.txt", 0664, 0777)
	if assert.NoError(t, err) {
		assert.Equal(t, "./testdata/sub/sub2/test.txt", file.Name())
		assert.NoError(t, os.RemoveAll("./testdata/sub"))
	}
}

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
