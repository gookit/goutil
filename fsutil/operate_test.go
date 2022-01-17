package fsutil_test

import (
	"os"
	"testing"

	"github.com/gookit/goutil/envutil"
	"github.com/gookit/goutil/fsutil"
	"github.com/stretchr/testify/assert"
)

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
	// if envutil.IsWin() {
	// 	return
	// }

	file, err := fsutil.CreateFile("./testdata/test.txt", 0664, 0666)
	if assert.NoError(t, err) {
		assert.Equal(t, "./testdata/test.txt", file.Name())
		assert.NoError(t, file.Close())
		assert.NoError(t, os.Remove(file.Name()))
	}

	file, err = fsutil.CreateFile("./testdata/sub/test.txt", 0664, 0777)
	if assert.NoError(t, err) {
		assert.Equal(t, "./testdata/sub/test.txt", file.Name())
		assert.NoError(t, file.Close())
		assert.NoError(t, os.RemoveAll("./testdata/sub"))
	}

	file, err = fsutil.CreateFile("./testdata/sub/sub2/test.txt", 0664, 0777)
	if assert.NoError(t, err) {
		assert.Equal(t, "./testdata/sub/sub2/test.txt", file.Name())
		assert.NoError(t, file.Close())
		assert.NoError(t, os.RemoveAll("./testdata/sub"))
	}
}

func TestQuickOpenFile(t *testing.T) {
	fname := "./testdata/quick-open-file.txt"
	file, err := fsutil.QuickOpenFile(fname)
	if assert.NoError(t, err) {
		assert.Equal(t, fname, file.Name())
		assert.NoError(t, file.Close())
		assert.NoError(t, os.Remove(file.Name()))
	}
}

func TestMustRemove(t *testing.T) {
	assert.Panics(t, func() {
		fsutil.MustRm("/path-not-exist")
	})

	assert.Panics(t, func() {
		fsutil.MustRemove("/path-not-exist")
	})
}

func TestQuietRemove(t *testing.T) {
	assert.NotPanics(t, func() {
		fsutil.QuietRm("/path-not-exist")
	})

	assert.NotPanics(t, func() {
		fsutil.QuietRemove("/path-not-exist")
	})
}
