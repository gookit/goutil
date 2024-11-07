package strutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestBeforeAfter(t *testing.T) {
	// BeforeFirst
	assert.Eq(t, "abc", strutil.BeforeFirst("abc", ":"))
	assert.Eq(t, "abc", strutil.BeforeFirst("abc:123", ":"))
	assert.Eq(t, "abc", strutil.BeforeFirst("abc:123:456", ":"))

	// AfterFirst
	assert.Eq(t, "", strutil.AfterFirst("abc", ":"))
	assert.Eq(t, "123", strutil.AfterFirst("abc:123", ":"))
	assert.Eq(t, "123:456", strutil.AfterFirst("abc:123:456", ":"))

	// BeforeLast
	assert.Eq(t, "abc", strutil.BeforeLast("abc", ":"))
	assert.Eq(t, "abc", strutil.BeforeLast("abc:123", ":"))
	assert.Eq(t, "abc:123", strutil.BeforeLast("abc:123:456", ":"))

	// AfterLast
	assert.Eq(t, "", strutil.AfterLast("abc", ":"))
	assert.Eq(t, "123", strutil.AfterLast("abc:123", ":"))
	assert.Eq(t, "456", strutil.AfterLast("abc:123:456", ":"))
}

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

	b, a = strutil.SplitKV(" name = inhere \n", "=")
	assert.Eq(t, "name", b)
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

	ss = strutil.SplitNValid(" ", ",", 2)
	assert.Empty(t, ss)

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

func TestSplitByWhitespace(t *testing.T) {
	ss := strutil.SplitByWhitespace("a  b c")
	assert.Eq(t, `[]string{"a", "b", "c"}`, fmt.Sprintf("%#v", ss))

	ss = strutil.SplitByWhitespace("a\r\nb\tc d")
	assert.Eq(t, `[]string{"a", "b", "c", "d"}`, fmt.Sprintf("%#v", ss))
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
	tests := []struct {
		str    string
		val    string
		cmt    string
		strict bool
	}{
		{
			str: "value",
			val: "value",
		},
		{
			str: "value // comments at end",
			val: "value",
			cmt: "// comments at end",
		},
		{
			str:    "value // comments at end",
			val:    "value",
			cmt:    "// comments at end",
			strict: true,
		},
		{
			str: "value// comments at end",
			val: "value",
			cmt: "// comments at end",
		},
		{
			str:    "value// comments at end",
			val:    "value// comments at end",
			strict: true,
		},
		{
			str: "value # comments at end",
			val: "value",
			cmt: "# comments at end",
		},
		{
			str:    "value # comments at end",
			val:    "value",
			cmt:    "# comments at end",
			strict: true,
		},
		{
			str: "value# comments at end",
			val: "value",
			cmt: "# comments at end",
		},
		{
			str:    "value# comments at end",
			val:    "value# comments at end",
			strict: true,
		},
	}

	for i, tt := range tests {
		idx := strutil.QuietString(i)
		val, comment := strutil.SplitInlineComment(tt.str+idx, tt.strict)
		if comment == "" {
			assert.Eq(t, tt.val+idx, val)
			assert.Eq(t, tt.cmt, comment)
		} else {
			assert.Eq(t, tt.val, val)
			assert.Eq(t, tt.cmt+idx, comment)
		}
	}
}
