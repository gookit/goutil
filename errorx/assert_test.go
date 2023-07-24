package errorx_test

import (
	"testing"

	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/testutil/assert"
)

func TestAssert_methods(t *testing.T) {
	// IsFalse
	assert.NoErr(t, errorx.IsFalse(false))
	assert.Err(t, errorx.IsFalse(true))
	assert.Err(t, errorx.IsFalse(true, "error msg"))

	// IsTrue
	assert.NoErr(t, errorx.IsTrue(true))
	assert.Err(t, errorx.IsTrue(false))

	// IsIn
	assert.NoErr(t, errorx.IsIn(1, []int{1, 2, 3}))
	assert.Err(t, errorx.IsIn(4, []int{1, 2, 3}))
	assert.Err(t, errorx.IsIn(4, []int{1, 2, 3}, "error msg"))
	assert.Err(t, errorx.IsIn(4, []int{1, 2, 3}, "error %s", "msg"))

	// NotIn
	assert.NoErr(t, errorx.NotIn(4, []int{1, 2, 3}))
	assert.Err(t, errorx.NotIn(1, []int{1, 2, 3}))
	assert.Err(t, errorx.NotIn(1, []int{1, 2, 3}, "error msg"))
}
