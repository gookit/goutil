package textscan_test

import (
	"testing"

	"github.com/gookit/goutil/strutil/textscan"
	"github.com/gookit/goutil/testutil/assert"
)

func TestAddKind(t *testing.T) {
	nk := textscan.Kind(99)
	assert.False(t, textscan.HasKind(nk))
	textscan.AddKind(nk, "Kind99")
	assert.True(t, textscan.HasKind(nk))
	assert.Eq(t, "Kind99", textscan.KindString(nk))

	assert.Panics(t, func() {
		textscan.AddKind(nk, "Kind99")
	})
}

func TestNewStringToken(t *testing.T) {
	tok := textscan.NewStringToken(textscan.TokKey, "key2")
	assert.False(t, tok.HasMore())
	assert.NoErr(t, tok.ScanMore(nil))
	assert.Err(t, tok.MergeSame(nil))
	assert.True(t, tok.IsValid())
	assert.True(t, textscan.IsKindToken(textscan.TokKey, tok))
	assert.Eq(t, textscan.TokKey.String(), tok.Kind().String())
	assert.Eq(t, "key2", tok.Value())
	assert.Eq(t, "key2", tok.String())
}

func TestNewEmptyToken(t *testing.T) {
	tok := textscan.NewEmptyToken()
	assert.False(t, tok.HasMore())
	assert.NoErr(t, tok.ScanMore(nil))
	assert.Err(t, tok.MergeSame(nil))
	assert.False(t, tok.IsValid())
	assert.True(t, textscan.IsKindToken(textscan.TokInvalid, tok))
	assert.Eq(t, textscan.TokInvalid.String(), tok.Kind().String())
	assert.Eq(t, "<Invalid>", tok.String())
}
