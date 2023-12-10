package mathutil_test

import (
	"testing"

	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestOrElse(t *testing.T) {
	assert.Eq(t, 23, mathutil.OrElse(23, 21))
	assert.Eq(t, 21.3, mathutil.OrElse[float64](0, 21.3))
}

func TestLessOr(t *testing.T) {
	assert.Eq(t, 23, mathutil.LessOr(23, 25, 0))
	assert.Eq(t, 11, mathutil.LessOr(23, 21, 11))
	assert.Eq(t, 11, mathutil.LessOr(21, 21, 11))

	// LteOr
	assert.Eq(t, 23, mathutil.LteOr(23, 25, 0))
	assert.Eq(t, 11, mathutil.LteOr(23, 21, 11))
	assert.Eq(t, 21, mathutil.LteOr(21, 21, 11))
}

func TestGreaterOr(t *testing.T) {
	assert.Eq(t, 23, mathutil.GreaterOr(23, 21, 0))
	assert.Eq(t, 21, mathutil.GreaterOr(23, 25, 21))
	assert.Eq(t, 11, mathutil.GreaterOr(21, 21, 11))

	// GteOr
	assert.Eq(t, 23, mathutil.GteOr(23, 21, 0))
	assert.Eq(t, 21, mathutil.GteOr(23, 25, 21))
	assert.Eq(t, 21, mathutil.GteOr(21, 21, 11))
}

func TestDiv(t *testing.T) {
	assert.Eq(t, float64(20), mathutil.Div(27, 1.35))
	assert.Eq(t, 14, mathutil.DivInt(27, 2))
	assert.Eq(t, 20, mathutil.DivF2i(27, 1.35))
}

func TestMul(t *testing.T) {
	assert.Eq(t, float64(5), mathutil.Mul(2, 2.35))
	assert.Eq(t, 36, mathutil.MulF2i(27, 1.35))
}
