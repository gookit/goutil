package comfunc_test

import (
	"testing"

	"github.com/gookit/goutil/internal/comfunc"
	"github.com/gookit/goutil/x/assert"
)

func TestSplitLineToKv(t *testing.T) {
	tests := []struct {
		line string
		k, v string
	}{
		{"key=val", "key", "val"},
		{"key = val ", "key", "val"},
		{"key =val\n", "key", "val"},
		{"key= val\r\n", "key", "val"},
		{" key=val\r", "key", "val"},
		{"key=val\t ", "key", "val"},
		{" key=val\t\n", "key", "val"},
		{"key=val\t\r\n", "key", "val"},
		{"key = val\nue", "key", "val\nue"},
		{" key-one =val ", "key-one", "val"},
		{" key_one = val", "key_one", "val"},
		{" valid=", "valid", ""},
		// invalid input
		{"invalid", "", ""},
		{"=invalid", "", ""},
		{" = invalid", "", ""},
		{"  ", "", ""},
	}

	for _, tt := range tests {
		k, v := comfunc.SplitLineToKv(tt.line, "=")
		assert.Eq(t, tt.k, k)
		assert.Eq(t, tt.v, v)
	}
}

func TestParseEnvLines_quotedInlineComments(t *testing.T) {
	tests := []struct {
		line string
		want string
	}{
		{`KEY="val" # comment`, "val"},
		{`KEY='val' # comment`, "val"},
		// a ' #' inside quotes is part of the value
		{`KEY="pass #1" # comment`, "pass #1"},
		{`KEY='pass #1' # comment`, "pass #1"},
		{`KEY="pass #1"`, "pass #1"},
		{`KEY="a" #c "d"`, "a"},
		// an unclosed quote must not panic
		{`KEY=" # comment`, `"`},
		{`KEY=' # comment`, `'`},
		// unchanged
		{`KEY=val # comment`, "val"},
		{`KEY="a"b"`, `a"b`},
		{`KEY="a" b`, `"a" b`},
	}

	for _, tt := range tests {
		mp, err := comfunc.ParseEnvLines(tt.line, comfunc.ParseEnvLineOption{})
		assert.NoErr(t, err)
		assert.Eq(t, tt.want, mp["KEY"], tt.line)
	}

	// without inline comments, text after the closing quote is kept
	mp, err := comfunc.ParseEnvLines(`KEY="a" # b`, comfunc.ParseEnvLineOption{NotInlineComments: true})
	assert.NoErr(t, err)
	assert.Eq(t, `"a" # b`, mp["KEY"])
}
