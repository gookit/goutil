package strutil_test

import (
	"strings"
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestSimilarity(t *testing.T) {
	is := assert.New(t)
	_, ok := strutil.Similarity("hello", "he", 0.3)
	is.True(ok)
}

func TestValid(t *testing.T) {
	is := assert.New(t)

	is.Eq("ab", strutil.Valid("ab", ""))
	is.Eq("ab", strutil.Valid("ab", "cd"))
	is.Eq("cd", strutil.Valid("", "cd"))
	is.Empty(strutil.Valid("", ""))

	is.Eq("cd", strutil.OrElse("", "cd"))
	is.Eq("ab", strutil.OrElse("ab", "cd"))

	var str = "non-empty"
	is.Equal(str, strutil.OrElseNilSafe(&str, "fallback"))
	str = ""
	is.Equal("fallback", strutil.OrElseNilSafe(&str, "fallback"))
	is.Equal("default", strutil.OrElseNilSafe(nil, "default"))

	is.Eq(" ", strutil.ZeroOr(" ", "cd"))
	is.Eq("cd", strutil.ZeroOr("", "cd"))
	is.Eq("ab", strutil.ZeroOr("ab", "cd"))

	is.Eq("cd", strutil.BlankOr("", "cd"))
	is.Eq("cd", strutil.BlankOr(" ", "cd"))
	is.Eq("ab", strutil.BlankOr("ab", "cd"))
	is.Eq("ab", strutil.BlankOr(" ab ", "cd"))

	is.Eq("ab", strutil.OrCond(true, "ab", "cd"))
	is.Eq("cd", strutil.OrCond(false, "ab", "cd"))

	is.Eq("ab", strutil.OrHandle("  ab  ", strings.TrimSpace))
	is.Empty(strutil.OrHandle("", strings.TrimSpace))
}

func TestReplaces(t *testing.T) {
	assert.Eq(t, "tom age is 20", strutil.Replaces(
		"{name} age is {age}",
		map[string]string{
			"{name}": "tom",
			"{age}":  "20",
		}))
}

func TestWrapTag(t *testing.T) {
	assert.Eq(t, "", strutil.WrapTag("", "info"))
	assert.Eq(t, "<info>abc</info>", strutil.WrapTag("abc", "info"))
}

func TestSubstrCount(t *testing.T) {
	s := "I'm fine, thank you, and you"
	substr := "you"
	res, err := strutil.SubstrCount(s, substr)
	assert.NoErr(t, err)
	assert.Eq(t, 2, res)

	res1, err := strutil.SubstrCount(s, substr, 18)
	assert.NoErr(t, err)
	assert.Eq(t, 1, res1)

	res2, err := strutil.SubstrCount(s, substr, 17, 100)
	assert.NoErr(t, err)
	assert.Eq(t, 1, res2)

	res, err = strutil.SubstrCount(s, substr, 16)
	assert.NoErr(t, err)
	assert.Eq(t, 2, res)

	res, err = strutil.SubstrCount(s, substr)
	assert.NoErr(t, err)
	assert.Eq(t, 2, res)

	res, err = strutil.SubstrCount(s, substr, 1, 2, 3)
	assert.Err(t, err)
	assert.Eq(t, 0, res)
}
