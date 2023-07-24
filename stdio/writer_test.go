package stdio_test

import (
	"bytes"
	"testing"

	"github.com/gookit/goutil/stdio"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/testutil/fakeobj"
)

func TestNewWriteWrapper(t *testing.T) {
	buf := new(bytes.Buffer)

	w := stdio.NewWriteWrapper(buf)
	_, err := w.WriteString("inhere")
	assert.NoErr(t, err)
	assert.Eq(t, "inhere", w.String())

	err = w.WriteByte(',')
	assert.NoErr(t, err)

	_, err = w.Write([]byte("hi."))
	assert.NoErr(t, err)
	assert.Eq(t, "inhere,hi.", w.String())

	_, err = w.Writef(" ok, %s.", "tom")
	assert.NoErr(t, err)
	assert.Eq(t, "inhere,hi. ok, tom.", w.String())

	b := &fakeobj.IOWriter{}
	w = stdio.WrapW(b)
	n, err := w.WriteString("abc")
	assert.Eq(t, 3, n)
	assert.NoErr(t, err)

	assert.Eq(t, "abc", string(b.Buf))
	assert.Empty(t, w.String())
}
