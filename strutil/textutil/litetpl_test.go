package textutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/strutil/textutil"
	"github.com/gookit/goutil/testutil/assert"
)

var data = map[string]any{
	"name": "inhere",
	"age":  2000,
	"subMp": map[string]any{
		"city": "cd",
		"addr": "addr 001",
	},
}

func TestRenderString(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		tpl := "hi, My name is {{ .name | upFirst }}, age is {{ .age }}"
		str := textutil.RenderString(tpl, data)
		assert.Eq(t, "hi, My name is Inhere, age is 2000", str)
	})

	// with default value and alias func
	t.Run("with default value and alias func", func(t *testing.T) {
		tpl := "name: {{ .name | default:guest }}, age: {{ .age | def:18 }}, city: {{ .city | cd }}"
		str := textutil.RenderString(tpl, map[string]any{})
		assert.Eq(t, "name: guest, age: 18, city: cd", str)
	})
}

func TestCustomVarFmt(t *testing.T) {
	lt := textutil.NewLiteTemplate(func(opt *textutil.LiteTemplateOpt) {
		opt.SetVarFmt("{,}")
	})

	tpl := "hi, My name is { name | upFirst }, age is { age }"
	str := lt.RenderString(tpl, data)
	assert.Eq(t, "hi, My name is Inhere, age is 2000", str)

	t.Run("with invalid var format", func(t *testing.T) {
		tpl := "hi, My name is { name | upFirst }, empty {}, age is { age }"
		str := lt.RenderString(tpl, data)
		assert.Eq(t, "hi, My name is Inhere, empty {}, age is 2000", str)
	})
}

func TestRenderFile(t *testing.T) {
	s, err := textutil.RenderFile("testdata/test-lite.tpl", data)
	assert.NoError(t, err)
	fmt.Println(s)

	assert.StrContains(t, s, "hi, My name is Inhere, age is 2000")
	assert.StrContains(t, s, "City: CD")
	assert.StrContains(t, s, "Addr: addr 001")

	// file not exist
	_, err = textutil.RenderFile("testdata/not-exist.tpl", nil)
	assert.Error(t, err)
}
