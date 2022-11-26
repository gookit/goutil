package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestSimilarity(t *testing.T) {
	is := assert.New(t)
	_, ok := strutil.Similarity("hello", "he", 0.3)
	is.True(ok)
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
