package fsutil_test

import (
	"os"
	"strings"
	"testing"

	"github.com/gookit/goutil/cliutil"
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
		// 066X will error
		assert.NoError(t, fsutil.Mkdir("./testdata/sub/sub23/sub31", 0777))

		assert.NoError(t, fsutil.MkParentDir("./testdata/sub/sub24/sub32"))
		assert.True(t, fsutil.IsDir("./testdata/sub/sub24"))

		assert.NoError(t, os.RemoveAll("./testdata/sub"))
	} else {
		cliutil.Redln("chmod dir ./testdata fail")
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

	fpath := "./testdata/sub/sub3/test-must-create.txt"
	assert.NoError(t, fsutil.RmFileIfExist(fpath))
	file = fsutil.MustCreateFile(fpath, 0, 0766)
	assert.NoError(t, file.Close())
}

func TestQuickOpenFile(t *testing.T) {
	fpath := "./testdata/quick-open-file.txt"
	assert.NoError(t, fsutil.RmFileIfExist(fpath))

	file, err := fsutil.QuickOpenFile(fpath)
	assert.NoError(t, err)
	assert.Equal(t, fpath, file.Name())

	_, err = file.WriteString("hello")
	assert.NoError(t, err)

	// close
	assert.NoError(t, file.Close())

	// open for read
	file, err = fsutil.OpenReadFile(fpath)
	assert.NoError(t, err)
	// var bts [5]byte
	bts := make([]byte, 5)
	_, err = file.Read(bts)
	assert.NoError(t, err)
	assert.Equal(t, "hello", string(bts))

	// close
	assert.NoError(t, file.Close())
	assert.NoError(t, fsutil.Remove(file.Name()))
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

func TestDiscardReader(t *testing.T) {
	sr := strings.NewReader("hello")
	fsutil.DiscardReader(sr)

	assert.Empty(t, fsutil.MustReadReader(sr))
	assert.Empty(t, fsutil.GetContents(sr))
}

func TestGetContents(t *testing.T) {
	fpath := "./testdata/get-contents.txt"
	assert.NoError(t, fsutil.RmFileIfExist(fpath))

	_, err := fsutil.PutContents(fpath, "hello")
	assert.NoError(t, err)

	assert.Nil(t, fsutil.ReadExistFile("/path-not-exist"))
	assert.Equal(t, []byte("hello"), fsutil.ReadExistFile(fpath))

	assert.Panics(t, func() {
		fsutil.GetContents(45)
	})
}

func TestMustCopyFile(t *testing.T) {
	srcPath := "./testdata/cp-file-src.txt"
	dstPath := "./testdata/cp-file-dst.txt"

	assert.NoError(t, fsutil.RmIfExist(srcPath))
	assert.NoError(t, fsutil.RmFileIfExist(dstPath))

	_, err := fsutil.PutContents(srcPath, "hello")
	assert.NoError(t, err)

	fsutil.MustCopyFile(srcPath, dstPath)
	assert.Equal(t, []byte("hello"), fsutil.GetContents(dstPath))
}

func TestUnzip(t *testing.T) {
	assert.Error(t, fsutil.Unzip("/path-not-exists", ""))
}
