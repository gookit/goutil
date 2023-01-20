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
}
