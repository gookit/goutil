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
	sb.WriteMulti('-', 'e', 'f')
	assert.Eq(t, "ab-cd-ef", sb.ResetGet())

	sb.WriteStrings("ab", "-", "cd")
	sb.WriteString("-ef")
	assert.Eq(t, "ab-cd-ef", sb.ResetGet())

	sb.WriteAny("abc")
	sb.WriteAnys(23, "def")
	assert.Eq(t, "abc23def", sb.ResetGet())

	sb.Write([]byte("abc"))
	sb.WriteRune('-')
	sb.Writeln("abc")
	assert.Eq(t, "abc-abc\n", sb.ResetGet())
}
