package fsutil_test

import (
	"os"
	"strings"
	"testing"

	"github.com/gookit/goutil/envutil"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestMkdir(t *testing.T) {
	// TODO windows will error
	if envutil.IsWin() {
		t.Skip("skip mkdir test on Windows")
		return
	}

	err := os.Chmod("testdata", os.ModePerm)

	if assert.NoErr(t, err) {
		assert.NoErr(t, fsutil.Mkdir("testdata/sub/sub21", os.ModePerm))
		assert.NoErr(t, fsutil.Mkdir("testdata/sub/sub22", 0666))
		// 066X will error
		assert.NoErr(t, fsutil.Mkdir("testdata/sub/sub23/sub31", 0774))

		assert.NoErr(t, fsutil.MkParentDir("testdata/sub/sub24/sub32"))
		assert.True(t, fsutil.IsDir("testdata/sub/sub24"))
		fsutil.SafeRemoveAll("testdata/sub")

		assert.NoErr(t, fsutil.MkDirs(0774, "testdata/sub1/sub21", "testdata/sub1/sub22"))
		assert.NoErr(t, fsutil.MkSubDirs(0774, "testdata/sub2", "sub21", "sub22"))
		assert.NoErr(t, fsutil.RemoveSub("testdata", fsutil.ExcludeDotFile, fsutil.ExcludeSuffix(".jpg")))
	}
}

func TestCreateFile(t *testing.T) {
	// TODO windows will error
	// if envutil.IsWin() {
	// 	return
	// }

	file, err := fsutil.CreateFile("testdata/test.txt", 0664, 0666)
	if assert.NoErr(t, err) {
		assert.Eq(t, "testdata/test.txt", file.Name())
		assert.NoErr(t, file.Close())
		assert.NoErr(t, os.Remove(file.Name()))
	}

	file, err = fsutil.CreateFile("testdata/sub/test.txt", 0664, 0777)
	if assert.NoErr(t, err) {
		assert.Eq(t, "testdata/sub/test.txt", file.Name())
		assert.NoErr(t, file.Close())
		assert.NoErr(t, os.RemoveAll("testdata/sub"))
	}

	file, err = fsutil.CreateFile("testdata/sub/sub2/test.txt", 0664, 0777)
	if assert.NoErr(t, err) {
		assert.Eq(t, "testdata/sub/sub2/test.txt", file.Name())
		assert.NoErr(t, file.Close())
		assert.NoErr(t, os.RemoveAll("testdata/sub"))
	}

	fpath := "testdata/sub/sub3/test-must-create.txt"
	file = fsutil.MustCreateFile(fpath, 0, 0766)
	assert.NoErr(t, file.Close())
	assert.NoErr(t, fsutil.RmFileIfExist(fpath))

	err = fsutil.RemoveSub("testdata/sub")
	assert.NoErr(t, err)
}

func TestQuickOpenFile(t *testing.T) {
	fpath := "testdata/quick-open-file.txt"
	assert.NoErr(t, fsutil.RmFileIfExist(fpath))

	file, err := fsutil.QuickOpenFile(fpath)
	assert.NoErr(t, err)
	assert.Eq(t, fpath, file.Name())

	_, err = file.WriteString("hello")
	assert.NoErr(t, err)

	// close
	assert.NoErr(t, file.Close())

	// open for read
	file, err = fsutil.OpenReadFile(fpath)
	assert.NoErr(t, err)
	// var bts [5]byte
	bts := make([]byte, 5)
	_, err = file.Read(bts)
	assert.NoErr(t, err)
	assert.Eq(t, "hello", string(bts))

	// close
	assert.NoErr(t, file.Close())
	assert.NoErr(t, fsutil.Remove(file.Name()))
}

func TestMustOpenFile(t *testing.T) {
	fpath := "testdata/must-open-file.txt"

	assert.Panics(t, func() {
		fsutil.MustOpenFile(fpath, os.O_RDONLY, 0666)
	})

	_, err := fsutil.PutContents(fpath, strings.NewReader("must-open-file"))
	assert.NoErr(t, err)

	of := fsutil.MustOpenFile(fpath, fsutil.FsRWFlags, 0600)
	assert.Eq(t, "must-open-file", fsutil.ReadString(of))
}

func TestOpenAppendFile(t *testing.T) {
	fpath := "./testdata/open-append-file.txt"
	assert.NoErr(t, fsutil.RmFileIfExist(fpath))

	file, err := fsutil.OpenAppendFile(fpath)
	assert.NoErr(t, err)
	assert.Eq(t, fpath, file.Name())

	_, err = file.WriteString("hello")
	assert.NoErr(t, err)
	assert.NoErr(t, file.Close())

	// reopen for write
	file, err = fsutil.OpenAppendFile(fpath)
	assert.NoErr(t, err)

	_, err = file.WriteString(" world")
	assert.NoErr(t, err)
	assert.NoErr(t, file.Close())

	// read all
	s := fsutil.ReadString(fpath)
	assert.Eq(t, "hello world", s)
}

func TestOpenTruncFile(t *testing.T) {
	fpath := "./testdata/open-trunc-file.txt"
	assert.NoErr(t, fsutil.RmFileIfExist(fpath))

	file, err := fsutil.OpenTruncFile(fpath)
	assert.NoErr(t, err)
	assert.Eq(t, fpath, file.Name())

	_, err = file.WriteString("hello")
	assert.NoErr(t, err)
	assert.NoErr(t, file.Close())

	// reopen for write
	file, err = fsutil.OpenTruncFile(fpath)
	assert.NoErr(t, err)

	_, err = file.WriteString(" world")
	assert.NoErr(t, err)

	// read all
	s := fsutil.ReadString(fpath)
	assert.Eq(t, " world", s)
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

func TestUnzip(t *testing.T) {
	assert.Err(t, fsutil.Unzip("/path-not-exists", ""))
}
