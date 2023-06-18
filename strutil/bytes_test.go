package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestBuffer_WriteAny(t *testing.T) {
	buf := strutil.NewBuffer()

	buf.Printf("ab-%s", "c")
	buf.PrintByte('d')
	assert.Eq(t, "ab-cd", buf.ResetAndGet())

	buf.WriteStr("ab", "-", "cd")
	buf.WriteStr1("-ef")
	assert.Eq(t, "ab-cd-ef", buf.ResetGet())

	buf.WriteAny(23, "abc")
	assert.Eq(t, "23abc", buf.ResetAndGet())

	buf.Writeln("abc")
	assert.Eq(t, "abc\n", buf.ResetAndGet())
}
