package textscan_test

import (
	"testing"

	"github.com/gookit/goutil/strutil/textscan"
	"github.com/gookit/goutil/testutil/assert"
)

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
