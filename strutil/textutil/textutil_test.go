package textutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/strutil/textutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestReplaceVars(t *testing.T) {
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
			if got := textutil.ReplaceVars(tt.tplText, tplVars, ""); got != tt.want {
				t.Errorf("ReplaceVars() = %v, want = %v", got, tt.want)
			}
		})
	}

	// custom format
	assert.Equal(t, "hi inhere", textutil.ReplaceVars("hi {$name}", tplVars, "{$,}"))
	assert.Equal(t, "hi inhere age is 230", textutil.ReplaceVars("hi $name age is $info.age", tplVars, "$,"))
	assert.Equal(t, "hi {$name}", textutil.ReplaceVars("hi {$name}", nil, "{$,}"))
}

func TestNewFullReplacer(t *testing.T) {
	vp := textutil.NewFullReplacer("")

	tplVars := map[string]any{
		"name": "inhere",
		"info": map[string]any{"age": 230, "sex": "man"},
	}

	tpl := "hi, {{ name }}, {{ age | 23 }}"
	str := vp.Render(tpl, nil)
	assert.Eq(t, "hi, {{ name }}, 23", str)

	str = vp.Render(tpl, tplVars)
	assert.Eq(t, "hi, inhere, 23", str)

	vp.OnNotFound(func(name string) (val string, ok bool) {
		if name == "name" {
			return "tom", true
		}
		return
	})
	str = vp.Render(tpl, nil)
	assert.Eq(t, "hi, tom, 23", str)
}

func TestRenderSMap(t *testing.T) {
	tplVars := map[string]string{
		"name":   "inhere",
		"age":    "234",
		"key_01": "inhere",
		"key-02": "inhere",
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
	}

	for i, tt := range tests {
		t.Run(strutil.JoinAny(" ", "case", i), func(t *testing.T) {
			if got := textutil.RenderSMap(tt.tplText, tplVars, ""); got != tt.want {
				t.Errorf("RenderSMap() = %v, want = %v", got, tt.want)
			}
		})
	}

	// custom format
	assert.Equal(t, "hi inhere", textutil.RenderSMap("hi {$name}", tplVars, "{$,}"))
	assert.Equal(t, "hi inhere age is 234", textutil.RenderSMap("hi $name age is $age", tplVars, "$"))
	assert.Equal(t, "hi inhere age is 234.", textutil.RenderSMap("hi $name age is $age.", tplVars, "$,"))
	assert.Equal(t, "hi {$name}", textutil.RenderSMap("hi {$name}", nil, "{$,}"))
}

func TestVarReplacer_ParseVars(t *testing.T) {
	vp := textutil.NewVarReplacer("")
	str := "hi {{ name }}, age {{age}}, age {{age }}"
	ss := vp.ParseVars(str)

	assert.NotEmpty(t, ss)
	assert.Len(t, ss, 2)
	assert.Contains(t, ss, "name")
	assert.Contains(t, ss, "age")

	tplVars := map[string]any{
		"name": "inhere",
		"age":  234,
	}
	assert.Equal(t, "hi inhere, age 234, age 234", vp.Render(str, tplVars))
	vp.DisableFlatten()
	assert.Equal(t, "hi inhere, age 234, age 234", vp.Render(str, tplVars))
}

func TestIsMatchAll(t *testing.T) {
	str := "hi inhere, age is 120"
	assert.True(t, textutil.IsMatchAll(str, []string{"hi", "inhere"}))
	assert.False(t, textutil.IsMatchAll(str, []string{"hi", "^inhere"}))
}

func TestParseInlineINI(t *testing.T) {
	mp, err := textutil.ParseInlineINI("")
	assert.NoErr(t, err)
	assert.Empty(t, mp)

	mp, err = textutil.ParseInlineINI("default=inhere")
	assert.NoErr(t, err)
	assert.NotEmpty(t, mp)
	assert.Eq(t, "inhere", mp.Str("default"))

	_, err = textutil.ParseInlineINI("string")
	assert.ErrSubMsg(t, err, "parse inline config error: must")

	_, err = textutil.ParseInlineINI("name=n;default=inhere", "name")
	assert.ErrSubMsg(t, err, "parse inline config error: invalid key name")
}

func TestParseSimpleINI(t *testing.T) {
	t.Run("empty input", func(t *testing.T) {
		mp, err := textutil.ParseSimpleINI("")
		assert.Nil(t, err)
		assert.Empty(t, mp)
	})

	t.Run("only newlines", func(t *testing.T) {
		mp, err := textutil.ParseSimpleINI("\n\n\n")
		assert.Nil(t, err)
		assert.Empty(t, mp)
	})

	t.Run("comment lines only", func(t *testing.T) {
		input := "# comment\n; another comment\n// inline comment"
		mp, err := textutil.ParseSimpleINI(input)
		assert.Nil(t, err)
		assert.Empty(t, mp)
	})

	t.Run("invalid line without equal sign", func(t *testing.T) {
		input := "invalid line"
		mp, err := textutil.ParseSimpleINI(input)
		assert.Err(t, err)
		assert.Contains(t, err.Error(), "invalid line contents")
		assert.Nil(t, mp)
	})

	t.Run("valid key-value pair", func(t *testing.T) {
		input := "key=value"
		mp, err := textutil.ParseSimpleINI(input)
		assert.Nil(t, err)
		assert.Eq(t, "value", mp["key"])
	})

	t.Run("key with inline comment", func(t *testing.T) {
		input := "key=value # this is a comment"
		mp, err := textutil.ParseSimpleINI(input)
		assert.Nil(t, err)
		assert.Eq(t, "value", mp["key"])
	})

	t.Run("multiple valid lines", func(t *testing.T) {
		input := "key1=value1\nkey2=value2\nkey3=value3"
		mp, err := textutil.ParseSimpleINI(input)
		assert.Nil(t, err)
		assert.Eq(t, "value1", mp["key1"])
		assert.Eq(t, "value2", mp["key2"])
		assert.Eq(t, "value3", mp["key3"])
	})

	t.Run("mixed lines", func(t *testing.T) {
		input := "# comment\nkey1=value1\n\nkey2=value2 # inline comment\ninvalidline"
		mp, err := textutil.ParseSimpleINI(input)
		assert.Err(t, err)
		assert.ErrMsgContains(t, err, "invalid line contents")
		assert.Empty(t, mp)
	})
}
