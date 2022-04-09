package strutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/stretchr/testify/assert"
)

func TestCut(t *testing.T) {
	str := "hi,inhere"

	b, a, ok := strutil.Cut(str, ",")
	assert.True(t, ok)
	assert.Equal(t, "hi", b)
	assert.Equal(t, "inhere", a)

	b, a = strutil.MustCut(str, ",")
	assert.Equal(t, "hi", b)
	assert.Equal(t, "inhere", a)

	b, a, ok = strutil.Cut(str, "-")
	assert.False(t, ok)
	assert.Equal(t, str, b)
	assert.Equal(t, "", a)
}

func TestSplit(t *testing.T) {
	ss := strutil.Split("a, , b,c", ",")
	assert.Equal(t, `[]string{"a", "b", "c"}`, fmt.Sprintf("%#v", ss))

	ss = strutil.SplitValid("a, , b,c", ",")
	assert.Equal(t, `[]string{"a", "b", "c"}`, fmt.Sprintf("%#v", ss))

	ss = strutil.SplitN("a, , b,c", ",", 3)
	assert.Equal(t, `[]string{"a", "b", "c"}`, fmt.Sprintf("%#v", ss))

	ss = strutil.SplitN("a, , b,c", ",", 2)
	assert.Equal(t, `[]string{"a", "b,c"}`, fmt.Sprintf("%#v", ss))

	ss = strutil.Split(" ", ",")
	assert.Nil(t, ss)
}

func TestSplitTrimmed(t *testing.T) {
	ss := strutil.SplitTrimmed("a, , b,c", ",")
	assert.Equal(t, `[]string{"a", "", "b", "c"}`, fmt.Sprintf("%#v", ss))

	ss = strutil.SplitNTrimmed("a, , b,c", ",", 2)
	assert.Equal(t, `[]string{"a", ", b,c"}`, fmt.Sprintf("%#v", ss))
}

func TestSubstr(t *testing.T) {
	assert.Equal(t, "abc", strutil.Substr("abcDef", 0, 3))
	assert.Equal(t, "cD", strutil.Substr("abcDef", 2, 2))
	assert.Equal(t, "cDef", strutil.Substr("abcDef", 2, 0))
	assert.Equal(t, "", strutil.Substr("abcDEF", 23, 5))
	assert.Equal(t, "cDEF12", strutil.Substr("abcDEF123", 2, -1))
	assert.Equal(t, "cDEF", strutil.Substr("abcDEF123", 2, -3))
}
