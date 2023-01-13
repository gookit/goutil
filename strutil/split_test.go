package strutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestCut(t *testing.T) {
	str := "hi,inhere"

	b, a, ok := strutil.Cut(str, ",")
	assert.True(t, ok)
	assert.Eq(t, "hi", b)
	assert.Eq(t, "inhere", a)

	b, a = strutil.MustCut(str, ",")
	assert.Eq(t, "hi", b)
	assert.Eq(t, "inhere", a)

	assert.Panics(t, func() {
		strutil.MustCut("abc", ",")
	})

	b, a = strutil.QuietCut(str, ",")
	assert.Eq(t, "hi", b)
	assert.Eq(t, "inhere", a)

	b, a = strutil.TrimCut(str, ",")
	assert.Eq(t, "hi", b)
	assert.Eq(t, "inhere", a)

	b, a = strutil.TrimCut(" hi , inhere \n", ",")
	assert.Eq(t, "hi", b)
	assert.Eq(t, "inhere", a)

	b, a, ok = strutil.Cut(str, "-")
	assert.False(t, ok)
	assert.Eq(t, str, b)
	assert.Eq(t, "", a)
}

func TestSplit(t *testing.T) {
	ss := strutil.Split("a, , b,c", ",")
	assert.Eq(t, `[]string{"a", "b", "c"}`, fmt.Sprintf("%#v", ss))

	ss = strutil.SplitValid("a, , b,c", ",")
	assert.Eq(t, `[]string{"a", "b", "c"}`, fmt.Sprintf("%#v", ss))

	ss = strutil.SplitN("a, , b,c", ",", 3)
	assert.Eq(t, `[]string{"a", "b", "c"}`, fmt.Sprintf("%#v", ss))

	ss = strutil.SplitN("a, , b,c", ",", 2)
	assert.Eq(t, `[]string{"a", "b,c"}`, fmt.Sprintf("%#v", ss))

	ss = strutil.SplitN("origin https://github.com/gookit/gitw (push)", " ", 3)
	assert.Len(t, ss, 3)

	ss = strutil.Split(" ", ",")
	assert.Nil(t, ss)
}

func TestSplitTrimmed(t *testing.T) {
	ss := strutil.SplitTrimmed("a, , b,c", ",")
	assert.Eq(t, `[]string{"a", "", "b", "c"}`, fmt.Sprintf("%#v", ss))

	ss = strutil.SplitTrimmed("", ",")
	assert.Empty(t, ss)

	ss = strutil.SplitTrimmed(",", ",")
	assert.NotEmpty(t, ss)

	ss = strutil.SplitNTrimmed("a, , b,c", ",", 2)
	assert.Eq(t, `[]string{"a", ", b,c"}`, fmt.Sprintf("%#v", ss))
}

func TestSubstr(t *testing.T) {
	assert.Eq(t, "abc", strutil.Substr("abcDef", 0, 3))
	assert.Eq(t, "cD", strutil.Substr("abcDef", 2, 2))
	assert.Eq(t, "cDef", strutil.Substr("abcDef", 2, 0))
	assert.Eq(t, "", strutil.Substr("abcDEF", 23, 5))
	assert.Eq(t, "cDEF12", strutil.Substr("abcDEF123", 2, -1))
	assert.Eq(t, "cDEF", strutil.Substr("abcDEF123", 2, -3))
}

func TestSplitInlineComment(t *testing.T) {
	val, comment := strutil.SplitInlineComment("value0")
	assert.Eq(t, "value0", val)
	assert.Eq(t, "", comment)

	val, comment = strutil.SplitInlineComment("value0 // comments at end")
	assert.Eq(t, "value0", val)
	assert.Eq(t, "// comments at end", comment)

	val, comment = strutil.SplitInlineComment("value0 # comments at end")
	assert.Eq(t, "value0", val)
	assert.Eq(t, "# comments at end", comment)
}
