package textutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil/textutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestRenderString(t *testing.T) {
	data := map[string]any{
		"name": "inhere",
		"age":  2000,
	}

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
