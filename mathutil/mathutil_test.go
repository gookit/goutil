package mathutil_test

import (
	"testing"

	"github.com/gookit/goutil/common"
	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestMaxFloat(t *testing.T) {
	assert.Eq(t, float64(3), mathutil.MaxFloat(2, 3))
}

func TestMaxI64(t *testing.T) {
	assert.Eq(t, 3, mathutil.MaxInt(2, 3))
	assert.Eq(t, 3, mathutil.MaxInt(3, 2))
	assert.Eq(t, int64(3), mathutil.MaxI64(2, 3))
	assert.Eq(t, int64(3), mathutil.MaxI64(3, 2))
}

func TestSwapMaxInt(t *testing.T) {
	x, y := mathutil.SwapMaxInt(2, 34)
	assert.Eq(t, 34, x)
	assert.Eq(t, 2, y)

	x64, y64 := mathutil.SwapMaxI64(2, 34)
	assert.Eq(t, int64(34), x64)
	assert.Eq(t, int64(2), y64)
}

func TestCompare(t *testing.T) {
	tests := []struct {
		x, y interface{}
		op   string
	}{
		{2, 2, common.OpEq},
		{2, 3, common.OpNeq},
		{2, 3, common.OpLt},
		{2, 3, common.OpLte},
		{2, 2, common.OpLte},
		{2, 1, common.OpGt},
		{2, 2, common.OpGte},
		{2, 1, common.OpGte},
		{2.2, 2.2, common.OpEq},
		{2.2, 3.1, common.OpNeq},
		{2.3, 3.2, common.OpLt},
		{2.3, 3.3, common.OpLte},
		{2.3, 2.3, common.OpLte},
		{2.3, 1.3, common.OpGt},
		{2.3, 2.3, common.OpGte},
		{2.3, 1.3, common.OpGte},
	}

	for _, test := range tests {
		assert.True(t, mathutil.Compare(test.x, test.y, test.op))
	}

	assert.False(t, mathutil.Compare(2, 3, common.OpGt))
	assert.False(t, mathutil.Compare(nil, 3, common.OpGt))
	assert.False(t, mathutil.Compare(2, nil, common.OpGt))
	assert.False(t, mathutil.Compare("abc", 3, common.OpGt))
	assert.False(t, mathutil.Compare(2, "def", common.OpGt))
}
