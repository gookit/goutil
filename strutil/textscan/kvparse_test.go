package textscan_test

import (
	"testing"

	"github.com/gookit/goutil/comdef"
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/strutil/textscan"
	"github.com/gookit/goutil/testutil/assert"
)

func TestKeyValueMatcher_Match(t *testing.T) {
	kv := &textscan.KeyValueMatcher{}

	tok, err := kv.Match(" ", nil)
	assert.NoErr(t, err)
	assert.Nil(t, tok)

	tok, err = kv.Match("invalid format", nil)
	assert.NoErr(t, err)
	assert.Nil(t, tok)

	tok, err = kv.Match("=value0", nil)
	assert.Err(t, err)
	assert.Nil(t, tok)

	tok, err = kv.Match("key0=value0", nil)
	assert.NoErr(t, err)
	assert.IsType(t, &textscan.ValueToken{}, tok)
	val := tok.(*textscan.ValueToken)
	assert.Eq(t, "key0", val.Key())
	assert.Eq(t, "value0", val.Value())
	assert.False(t, val.HasMore())
	assert.Eq(t, "key: key0\nvalue: \"value0\"\ncomments: ", val.String())
	assert.Err(t, tok.MergeSame(textscan.NewEmptyToken()))

	tok, err = kv.Match("key0=value0 // comments at end", nil)
	assert.NoErr(t, err)
	assert.IsType(t, &textscan.ValueToken{}, tok)
	val = tok.(*textscan.ValueToken)
	assert.Eq(t, "key0", val.Key())
	assert.Eq(t, "value0 // comments at end", val.Value())
	assert.Eq(t, "", val.Comment())

	// parse InlineComment
	kv.InlineComment = true
	tok, err = kv.Match("key0=value0 // comments at end", nil)
	assert.NoErr(t, err)
	assert.IsType(t, &textscan.ValueToken{}, tok)
	val = tok.(*textscan.ValueToken)
	assert.Eq(t, "key0", val.Key())
	assert.Eq(t, "value0", val.Value())
	assert.Empty(t, val.Values())
	assert.Eq(t, "// comments at end", val.Comment())

	// change kv-sep
	kv.Separator = comdef.ColonStr
	tok, err = kv.Match("key0: value0 // comments at end", nil)
	assert.NoErr(t, err)
	val = tok.(*textscan.ValueToken)
	assert.Eq(t, "key0", val.Key())
	assert.Eq(t, "value0", val.Value())
	assert.Eq(t, "// comments at end", val.Comment())
}

func TestKeyValueMatcher_KeyCheckFn(t *testing.T) {
	kv := &textscan.KeyValueMatcher{}
	kv.KeyCheckFn = func(key string) error {
		if len(key) < 2 {
			return errorx.Raw("key is invalid")
		}
		return nil
	}

	tok, err := kv.Match("k = value", nil)
	assert.Nil(t, tok)
	assert.ErrMsg(t, err, "key is invalid")

	tok, err = kv.Match("key = value", nil)
	assert.NoErr(t, err)
	assert.True(t, tok.IsValid())
	assert.False(t, tok.HasMore())

	vt := tok.(*textscan.ValueToken)
	assert.Eq(t, "key", vt.Key())
	assert.Eq(t, "value", vt.Value())
}

func TestKeyValueMatcher_MultiLine(t *testing.T) {
	kv := &textscan.KeyValueMatcher{}

	tok, err := kv.Match(`key = """multi line`, nil)
	assert.NoErr(t, err)
	assert.True(t, tok.HasMore())
	assert.True(t, tok.IsValid())

	vt := tok.(*textscan.ValueToken)
	assert.Eq(t, textscan.MultiLineValMarkD, vt.Mark())

	ok, val := kv.DetectEnd(textscan.MultiLineValMarkD, `value"""`)
	assert.True(t, ok)
	assert.Eq(t, "value", val)

	ok, val = kv.DetectEnd(textscan.MultiLineValMarkD, `not end`)
	assert.False(t, ok)
	assert.Eq(t, "not end", val)

	tok, err = kv.Match(`key = '''multi line`, nil)
	assert.NoErr(t, err)
	assert.True(t, tok.HasMore())
	assert.True(t, tok.IsValid())

	vt = tok.(*textscan.ValueToken)
	assert.Eq(t, textscan.MultiLineValMarkS, vt.Mark())

	ok, val = kv.DetectEnd(textscan.MultiLineValMarkS, `value'''`)
	assert.True(t, ok)
	assert.Eq(t, "value", val)
}

func TestKeyValueMatcher_DisableMultiLine(t *testing.T) {
	kv := &textscan.KeyValueMatcher{
		DisableMultiLine: true,
	}

	tok, err := kv.Match(`key = """multi line`, nil)
	assert.NoErr(t, err)
	assert.False(t, tok.HasMore())
	assert.Eq(t, `"""multi line`, tok.Value())

}
