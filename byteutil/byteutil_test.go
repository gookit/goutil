package byteutil_test

import (
	"testing"

	"github.com/gookit/goutil/byteutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestFirstLine(t *testing.T) {
	bs := []byte("hi\ninhere")
	assert.Eq(t, []byte("hi"), byteutil.FirstLine(bs))
	assert.Eq(t, []byte("hi"), byteutil.FirstLine([]byte("hi")))
}
