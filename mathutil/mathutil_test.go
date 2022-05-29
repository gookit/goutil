package mathutil_test

import (
	"testing"

	"github.com/gookit/goutil/mathutil"
	"github.com/stretchr/testify/assert"
)

func TestMaxI64(t *testing.T) {
	assert.Equal(t, 3, mathutil.MaxInt(2, 3))
	assert.Equal(t, 3, mathutil.MaxInt(3, 2))
	assert.Equal(t, int64(3), mathutil.MaxI64(2, 3))
	assert.Equal(t, int64(3), mathutil.MaxI64(3, 2))
}
