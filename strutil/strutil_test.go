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
	is.Eq("cd", strutil.OrElse("", "cd"))
	is.Eq("ab", strutil.OrElse("ab", "cd"))

	is.Eq("ab", strutil.OrCond(true, "ab", "cd"))
	is.Eq("cd", strutil.OrCond(false, "ab", "cd"))

	is.Eq("ab", strutil.OrHandle("  ab  ", strings.TrimSpace))
}

func TestRenderTemplate(t *testing.T) {
	tpl := "hi, My name is {{ .name | upFirst }}, age is {{ .age }}"
	assert.Eq(t, "hi, My name is Inhere, age is 2000", strutil.RenderTemplate(tpl, map[string]any{
		"name": "inhere",
		"age":  2000,
	}, nil))
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
	res3, err := strutil.SubstrCount(s, substr, 16)
	assert.NoErr(t, err)
	assert.Eq(t, 2, res3)
}

func TestPrettyJSON(t *testing.T) {
	tests := []any{
		map[string]int{"a": 1},
		struct {
			A int `json:"a"`
		}{1},
	}
	want := `{
    "a": 1
}`
	for _, sample := range tests {
		got, err := strutil.PrettyJSON(sample)
		assert.NoErr(t, err)
		assert.Eq(t, want, got)
	}
}
