package clipboard_test

import (
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/sysutil/clipboard"
	"github.com/gookit/goutil/testutil/assert"
)

func TestClipboard_WriteFromFile(t *testing.T) {
	cb := clipboard.New()
	if ok := cb.Available(); !ok {
		assert.False(t, ok)
		t.Skipf("skip test on program '%s' not found", clipboard.GetReaderBin())
		return
	}

	srcFile := "testdata/testcb.txt"
	srcStr := string(fsutil.MustReadFile(srcFile))
	assert.NotEmpty(t, srcStr)

	err := cb.WriteFromFile(srcFile)
	assert.NoErr(t, err)

	readStr, err := cb.ReadString()
	assert.NoErr(t, err)
	assert.Eq(t, srcStr, readStr)

	dstFile := "testdata/read-from-cb.txt"
	assert.NoErr(t, fsutil.RmFileIfExist(dstFile))
	err = cb.ReadToFile(dstFile)
	assert.NoErr(t, err)

	dstStr := string(fsutil.MustReadFile(dstFile))
	assert.Eq(t, srcStr, dstStr)
}
