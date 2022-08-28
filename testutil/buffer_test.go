package testutil_test

import (
	"testing"

	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestNewBuffer(t *testing.T) {
	buf := testutil.NewBuffer()

	buf.WriteString("ab", "-", "cd")
	assert.Eq(t, "ab-cd", buf.ResetAndGet())

	buf.WriteAny(23, "abc")
	assert.Eq(t, "23abc", buf.ResetAndGet())

	buf.Writeln("abc")
	assert.Eq(t, "abc\n", buf.ResetAndGet())
}
