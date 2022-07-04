package clipboard_test

import (
	"testing"

	"github.com/gookit/goutil/sysutil/clipboard"
	"github.com/gookit/goutil/testutil/assert"
)

func TestClipboard_read_write(t *testing.T) {
	err := clipboard.WriteString("")
	assert.ErrMsg(t, err, "not write contents")

	if !clipboard.Available() {
		assert.False(t, clipboard.Available())
		return
	}

	err = clipboard.Reset()
	assert.NoErr(t, err)

	str, err := clipboard.ReadString()
	assert.NoErr(t, err)
	assert.Empty(t, str)

	src := "hello, this is clipboard"
	err = clipboard.WriteString(src)
	assert.NoErr(t, err)

	str, err = clipboard.ReadString()
	assert.NoErr(t, err)
	assert.NotEmpty(t, str)
	assert.Eq(t, src, str)
}
