package mathutil_test

import (
	"testing"

	"github.com/gookit/goutil/mathutil"
	"github.com/stretchr/testify/assert"
)

func TestMaxFloat(t *testing.T) {
	assert.Equal(t, float64(3), mathutil.MaxFloat(2, 3))
}

func TestMaxI64(t *testing.T) {
	assert.Equal(t, 3, mathutil.MaxInt(2, 3))
	assert.Equal(t, 3, mathutil.MaxInt(3, 2))
	assert.Equal(t, int64(3), mathutil.MaxI64(2, 3))
	assert.Equal(t, int64(3), mathutil.MaxI64(3, 2))
}

func TestSwapMaxInt(t *testing.T) {
	x, y := mathutil.SwapMaxInt(2, 34)
	assert.Equal(t, 34, x)
	assert.Equal(t, 2, y)

	x64, y64 := mathutil.SwapMaxI64(2, 34)
	assert.Equal(t, int64(34), x64)
	assert.Equal(t, int64(2), y64)
}
