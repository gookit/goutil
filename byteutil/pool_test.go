package byteutil_test

import (
	"testing"

	"github.com/gookit/goutil/byteutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestNewChanPool(t *testing.T) {
	p := byteutil.NewChanPool(10, 8, 8)

	assert.Equal(t, 8, p.Width())
	assert.Equal(t, 8, p.WidthCap())

	p.Put([]byte("abc"))
	assert.Equal(t, []byte("abc"), p.Get())
}
