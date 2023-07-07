package fsutil_test

import (
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestMustCopyFile(t *testing.T) {
	srcPath := "./testdata/cp-file-src.txt"
	dstPath := "./testdata/cp-file-dst.txt"

	assert.NoErr(t, fsutil.RmIfExist(srcPath))
	assert.NoErr(t, fsutil.RmFileIfExist(dstPath))

	_, err := fsutil.PutContents(srcPath, "hello")
	assert.NoErr(t, err)

	fsutil.MustCopyFile(srcPath, dstPath)
	assert.Eq(t, []byte("hello"), fsutil.GetContents(dstPath))
	assert.Eq(t, "hello", fsutil.ReadString(dstPath))

	str, err := fsutil.ReadStringOrErr(dstPath)
	assert.NoErr(t, err)
	assert.Eq(t, "hello", str)

	assert.NoErr(t, fsutil.RmFileIfExist(srcPath))
	assert.NoErr(t, fsutil.RmIfExist(srcPath)) // repeat delete
}

func TestWriteFile(t *testing.T) {
	tempFile := "./testdata/write-file.txt"

	err := fsutil.WriteFile(tempFile, []byte("hello\ngolang"), 0644)
	assert.NoErr(t, err)
	assert.Eq(t, []byte("hello\ngolang"), fsutil.MustReadFile(tempFile))

	// write invalid data
	assert.Panics(t, func() {
		_ = fsutil.WriteFile(tempFile, []string{"hello"}, 0644)
	})
}

func TestMustSave(t *testing.T) {
	opt := fsutil.OpenOptOrNew(nil)
	assert.NotNil(t, opt)
	opt = fsutil.OpenOptOrNew(fsutil.NewOpenOption())
	assert.NotNil(t, opt)

	testFile := "./testdata/must-save.txt"

	fsutil.MustSave(testFile, []byte("hello"),
		fsutil.WithFlag(fsutil.FsCWTFlags),
		fsutil.WithPerm(0644),
	)
	assert.Eq(t, "hello", fsutil.ReadString(testFile))

	// write invalid data
	assert.Panics(t, func() {
		fsutil.MustSave(testFile, []string{"hello"})
	})
}
