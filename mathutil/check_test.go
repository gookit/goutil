package mathutil_test

import (
	"testing"

	"github.com/gookit/goutil/comdef"
	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestIsNumeric(t *testing.T) {
	assert.True(t, mathutil.IsNumeric('3'))
	assert.False(t, mathutil.IsNumeric('a'))
}

func TestCompare(t *testing.T) {
	tests := []struct {
		x, y any
		op   string
	}{
		{2, 2, comdef.OpEq},
		{2, 3, comdef.OpNeq},
		{2, 3, comdef.OpLt},
		{2, 3, comdef.OpLte},
		{2, 2, comdef.OpLte},
		{2, 1, comdef.OpGt},
		{2, 2, comdef.OpGte},
		{2, 1, comdef.OpGte},
		{2, "1", comdef.OpGte},
		{2.2, 2.2, comdef.OpEq},
		{2.2, 3.1, comdef.OpNeq},
		{2.3, 3.2, comdef.OpLt},
		{2.3, 3.3, comdef.OpLte},
		{2.3, 2.3, comdef.OpLte},
		{2.3, 1.3, comdef.OpGt},
		{2.3, 2.3, comdef.OpGte},
		{2.3, 1.3, comdef.OpGte},
	}

	for _, test := range tests {
		assert.True(t, mathutil.Compare(test.x, test.y, test.op))
	}

	assert.False(t, mathutil.Compare(2, 3, comdef.OpGt))
	assert.False(t, mathutil.Compare(nil, 3, comdef.OpGt))
	assert.False(t, mathutil.Compare(2, nil, comdef.OpGt))
	assert.False(t, mathutil.Compare("abc", 3, comdef.OpGt))
	assert.False(t, mathutil.Compare(2, "def", comdef.OpGt))

	// float64
	assert.False(t, mathutil.Compare(2.4, "def", comdef.OpGt))

	// float32
	assert.True(t, mathutil.Compare(float32(2.3), float32(2.1), comdef.OpGt))
	assert.False(t, mathutil.Compare(float32(2.3), float32(2.1), "<"))
	assert.False(t, mathutil.Compare(float32(2.3), "invalid", "<"))

	assert.True(t, mathutil.CompInt(2, 3, comdef.OpLt))

	// int64
	assert.True(t, mathutil.CompInt64(int64(2), 3, comdef.OpLt))
	assert.True(t, mathutil.CompInt64(int64(22), 3, comdef.OpGt))
	assert.False(t, mathutil.CompInt64(int64(2), 3, comdef.OpGt))
}

func TestInRange(t *testing.T) {
	assert.True(t, mathutil.InRange(1, 1, 2))
	assert.True(t, mathutil.InRange(1, 1, 1))
	assert.False(t, mathutil.InRange(1, 2, 1))
	assert.False(t, mathutil.InRange(1, 2, 2))

	assert.True(t, mathutil.InRange(1.1, 1.1, 2.2))
	assert.True(t, mathutil.InRange(1.1, 1.1, 1.1))
	assert.False(t, mathutil.InRange(1.1, 2.2, 1.1))

	// test for OutRange()
	assert.False(t, mathutil.OutRange(1, 1, 2))
	assert.False(t, mathutil.OutRange(1, 1, 1))
	assert.True(t, mathutil.OutRange(1, 2, 10))

	// test for InUintRange()
	assert.True(t, mathutil.InUintRange[uint](1, 1, 2))
	assert.True(t, mathutil.InUintRange[uint](1, 1, 1))
	assert.True(t, mathutil.InUintRange[uint](1, 1, 0))
	assert.False(t, mathutil.InUintRange[uint](1, 2, 1))
}
