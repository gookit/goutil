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

	buf.Writeln("abc", 123)
	assert.Eq(t, "abc123\n", buf.ResetAndGet())

	buf.Println("abc", 123)
	assert.Eq(t, "abc 123\n", buf.ResetAndGet())

	assert.NoErr(t, buf.Close())
	assert.NoErr(t, buf.Flush())
	assert.NoErr(t, buf.Sync())

	buf.WriteStr1Nl("abc")
	assert.Eq(t, "abc\n", buf.ResetAndGet())

	// test WriteStrings
	buf.WriteStrings([]string{"a", "b", "c"})
	assert.Eq(t, "abc", buf.ResetAndGet())

	// test WriteStringNl
	buf.WriteStringNl("abc")
	assert.Eq(t, "abc\n", buf.ResetAndGet())

	// test WriteAnyNl
	buf.WriteAnyNl(1, "abc")
	assert.Eq(t, "1abc\n", buf.ResetAndGet())
}
