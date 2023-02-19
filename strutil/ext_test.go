package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestBuilder_WriteAny(t *testing.T) {
	var sb strutil.Builder

	sb.Writef("ab-%s", "c")
	sb.WriteByteNE('d')
	assert.Eq(t, "ab-cd", sb.String())

	sb.Reset()
	sb.WriteStrings("ab", "-", "cd")
	sb.WriteString("-ef")
	assert.Eq(t, "ab-cd-ef", sb.String())

	sb.Reset()
	sb.WriteAny("abc")
	sb.WriteAnys(23, "def")
	assert.Eq(t, "abc23def", sb.String())

	sb.Reset()
	sb.Writeln("abc")
	assert.Eq(t, "abc\n", sb.String())
}
