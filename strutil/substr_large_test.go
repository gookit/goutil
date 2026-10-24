package strutil_test

import (
	"math"
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/x/assert"
)

func TestSubstrLargeLength(t *testing.T) {
	assert.Eq(t, "好世界", strutil.Substr("你好世界", 1, math.MaxInt))
	assert.Eq(t, "bc", strutil.Substr("abcd", 1, -1))
}
