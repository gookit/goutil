package mathutil_test

import (
	"testing"

	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestMaxFloat(t *testing.T) {
	assert.Eq(t, float64(3), mathutil.MaxFloat(2, 3))
	assert.Eq(t, 3.3, mathutil.MaxFloat(2.1, 3.3))

	assert.Eq(t, 3.3, mathutil.Max(2.1, 3.3))
	assert.Eq(t, 3.3, mathutil.Max(3.3, 2.1))

	assert.Eq(t, 2.1, mathutil.Min(2.1, 3.3))
	assert.Eq(t, 2.1, mathutil.Min(3.3, 2.1))
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

	x, y = mathutil.SwapMax(34, 2)
	assert.Eq(t, 34, x)
	assert.Eq(t, 2, y)

	x, y = mathutil.SwapMin(2, 34)
	assert.Eq(t, 2, x)
	assert.Eq(t, 34, y)

	x, y = mathutil.SwapMin(34, 2)
	assert.Eq(t, 2, x)
	assert.Eq(t, 34, y)

	x, y = mathutil.SwapMaxInt(2, 34)
	assert.Eq(t, 34, x)
	assert.Eq(t, 2, y)
	x, y = mathutil.SwapMaxInt(34, 2)
	assert.Eq(t, 34, x)
	assert.Eq(t, 2, y)

	x64, y64 := mathutil.SwapMaxI64(2, 34)
	assert.Eq(t, int64(34), x64)
	assert.Eq(t, int64(2), y64)
	x64, y64 = mathutil.SwapMaxI64(34, 2)
	assert.Eq(t, int64(34), x64)
	assert.Eq(t, int64(2), y64)
}
