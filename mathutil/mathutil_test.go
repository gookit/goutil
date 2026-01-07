package mathutil_test

import (
	"testing"

	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestDiv(t *testing.T) {
	assert.Eq(t, float64(20), mathutil.Div(27, 1.35))
	assert.Eq(t, 14, mathutil.DivInt(27, 2))
	assert.Eq(t, 20, mathutil.DivF2i(27, 1.35))
}

func TestMul(t *testing.T) {
	assert.Eq(t, float64(5), mathutil.Mul(2, 2.35))
	assert.Eq(t, 36, mathutil.MulF2i(27, 1.35))
}
