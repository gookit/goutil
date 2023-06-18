package byteutil_test

import (
	"testing"

	"github.com/gookit/goutil/byteutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestBuffer_WriteAny(t *testing.T) {
	buf := byteutil.NewBuffer()

	buf.Printf("ab-%s", "c")
	buf.PrintByte('d')
	assert.Eq(t, "ab-cd", buf.ResetAndGet())

	buf.WriteStr("ab", "-", "cd")
	buf.WriteStr1("-ef")
	assert.Eq(t, "ab-cd-ef", buf.ResetAndGet())

	buf.WriteAny(23, "abc")
	assert.Eq(t, "23abc", buf.ResetGet())

	buf.Writeln("abc")
	assert.Eq(t, "abc\n", buf.ResetAndGet())

	assert.NoErr(t, buf.Close())
	assert.NoErr(t, buf.Flush())
}
