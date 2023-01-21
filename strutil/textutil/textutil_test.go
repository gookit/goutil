package textutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/strutil/textutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestReplaceVars(t *testing.T) {
	format := ""
	tplVars := map[string]any{
		"name":   "inhere",
		"key_01": "inhere",
		"key-02": "inhere",
		"info":   map[string]any{"age": 230, "sex": "man"},
	}

	tests := []struct {
		tplText string
		want    string
	}{
		{"hi inhere", "hi inhere"},
		{"hi {{name}}", "hi inhere"},
		{"hi {{ name}}", "hi inhere"},
		{"hi {{name }}", "hi inhere"},
		{"hi {{ name }}", "hi inhere"},
		{"hi {{ key_01 }}", "hi inhere"},
		{"hi {{ key-02 }}", "hi inhere"},
		{"hi {{ info.age }}", "hi 230"},
	}

	for i, tt := range tests {
		t.Run(strutil.JoinAny(" ", "case", i), func(t *testing.T) {
			if got := textutil.ReplaceVars(tt.tplText, tplVars, format); got != tt.want {
				t.Errorf("ReplaceVars() = %v, want = %v", got, tt.want)
			}
		})
	}

	// custom format
	assert.Equal(t, "hi inhere", textutil.ReplaceVars("hi {$name}", tplVars, "{$,}"))
	assert.Equal(t, "hi {$name}", textutil.ReplaceVars("hi {$name}", nil, "{$,}"))
}

func TestIsMatchAll(t *testing.T) {
	str := "hi inhere, age is 120"
	assert.True(t, textutil.IsMatchAll(str, []string{"hi", "inhere"}))
	assert.False(t, textutil.IsMatchAll(str, []string{"hi", "^inhere"}))
}
