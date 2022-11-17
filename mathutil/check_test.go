package mathutil_test

import (
	"testing"

	"github.com/gookit/goutil/comdef"
	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/testutil/assert"
)

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
}
