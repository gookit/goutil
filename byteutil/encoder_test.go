package byteutil_test

import (
	"testing"

	"github.com/gookit/goutil/byteutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestB64Encoder(t *testing.T) {
	src := []byte("abc1234566")
	dst := byteutil.B64Encoder.Encode(src)
	assert.NotEmpty(t, dst)

	decSrc, err := byteutil.B64Encoder.Decode(dst)
	assert.NoError(t, err)
	assert.Eq(t, src, decSrc)
}

func TestHexEncoder(t *testing.T) {
	src := []byte("abc1234566")
	dst := byteutil.HexEncoder.Encode(src)
	assert.NotEmpty(t, dst)

	decSrc, err := byteutil.HexEncoder.Decode(dst)
	assert.NoError(t, err)
	assert.Eq(t, src, decSrc)
}
