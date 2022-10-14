package textscan_test

import (
	"testing"

	"github.com/gookit/goutil/comdef"
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
