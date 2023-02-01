package byteutil_test

import (
	"errors"
	"testing"

	"github.com/gookit/goutil/byteutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestFirstLine(t *testing.T) {
	bs := []byte("hi\ninhere")
	assert.Eq(t, []byte("hi"), byteutil.FirstLine(bs))
	assert.Eq(t, []byte("hi"), byteutil.FirstLine([]byte("hi")))
}

func TestStrOrErr(t *testing.T) {
	bs := []byte("hi, inhere")
	assert.Eq(t, "hi, inhere", byteutil.SafeString(bs, nil))
	assert.Eq(t, "", byteutil.SafeString(bs, errors.New("error")))

	str, err := byteutil.StrOrErr(bs, nil)
	assert.NoErr(t, err)
	assert.Eq(t, "hi, inhere", str)

	str, err = byteutil.StrOrErr(bs, errors.New("error"))
	assert.Err(t, err)
	assert.Eq(t, "", str)
}
