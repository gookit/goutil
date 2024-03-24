package mathutil_test

import (
	"testing"

	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestAbs(t *testing.T) {
	assert.Eq(t, 1, mathutil.Abs(1))
	assert.Eq(t, 1, mathutil.Abs(-1))
}
