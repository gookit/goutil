package mathutil_test

import (
	"testing"

	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestMaxFloat(t *testing.T) {
	assert.Eq(t, float64(3), mathutil.MaxFloat(2, 3))
	assert.Eq(t, 3.3, mathutil.Max(2.1, 3.3))
}

func TestMaxI64(t *testing.T) {
	assert.Eq(t, 3, mathutil.MaxInt(2, 3))
	assert.Eq(t, 3, mathutil.MaxInt(3, 2))
	assert.Eq(t, int64(3), mathutil.MaxI64(2, 3))
	assert.Eq(t, int64(3), mathutil.MaxI64(3, 2))

	assert.Eq(t, 3, mathutil.Max[int](3, 2))
	assert.Eq(t, int64(3), mathutil.Max[int64](3, 2))
	assert.Eq(t, int64(3), mathutil.Max(int64(3), int64(2)))
}

func TestSwapMaxInt(t *testing.T) {
	x, y := mathutil.SwapMax(2, 34)
	assert.Eq(t, 34, x)
	assert.Eq(t, 2, y)

	x, y = mathutil.SwapMaxInt(2, 34)
	assert.Eq(t, 34, x)
	assert.Eq(t, 2, y)

	x64, y64 := mathutil.SwapMaxI64(2, 34)
	assert.Eq(t, int64(34), x64)
	assert.Eq(t, int64(2), y64)
}

func TestOrElse(t *testing.T) {
	assert.Eq(t, 23, mathutil.OrElse(23, 21))
	assert.Eq(t, 21.3, mathutil.OrElse[float64](0, 21.3))
}
