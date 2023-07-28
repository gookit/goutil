package clipboard_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/strutil"
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

	cb.WithVerbose(true)
	assert.True(t, cb.Writeable())
	assert.True(t, cb.Readable())

	srcFile := "testdata/testcb.txt"
	srcStr := string(fsutil.MustReadFile(srcFile))
	assert.NotEmpty(t, srcStr)

	err := cb.WriteFromFile(srcFile)
	assert.NoErr(t, err)

	err = cb.WriteFromFile("path/to/not-exists.txt")
	assert.Err(t, err)

	readStr, err := cb.ReadString()
	assert.NoErr(t, err)
	assert.Eq(t, srcStr, strutil.Trim(readStr))

	dstFile := "testdata/read-from-cb.txt"
	assert.NoErr(t, fsutil.RmFileIfExist(dstFile))
	err = cb.ReadToFile(dstFile)
	assert.NoErr(t, err)

	dstStr := string(fsutil.MustReadFile(dstFile))
	assert.Eq(t, srcStr, strutil.Trim(dstStr))

	assert.NoErr(t, cb.Clean())

	_, err = cb.Write([]byte("hello"))
	assert.NoErr(t, err)
	assert.NoErr(t, cb.Flush())
	assert.Eq(t, "hello", cb.SafeString())
	fmt.Println("...end...")
}
