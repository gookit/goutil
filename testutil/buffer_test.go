package testutil_test

import (
	"testing"

	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestNewBuffer(t *testing.T) {
	buf := testutil.NewBuffer()

	buf.WriteStr("ab", "-", "cd")
	assert.Eq(t, "ab-cd", buf.ResetGet())

	buf.WriteAny(23, "abc")
	assert.Eq(t, "23abc", buf.ResetAndGet())

	buf.Writeln("abc")
	assert.Eq(t, "abc\n", buf.ResetAndGet())
}

func TestNewSafeBuffer(t *testing.T) {
	sb := testutil.NewSafeBuffer()
	_, err := sb.Write([]byte("hello,"))
	assert.NoErr(t, err)
	err = sb.WriteByte('a')
	assert.NoErr(t, err)
	_, err = sb.WriteRune('b')
	assert.NoErr(t, err)

	s := sb.ResetGet()
	assert.Eq(t, "hello,ab", s)

	_, err = sb.WriteString("hello")
	assert.NoErr(t, err)
}
