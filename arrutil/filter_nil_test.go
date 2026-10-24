package arrutil_test

import (
	"testing"

	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/x/assert"
)

func TestFilterNilInterfaces(t *testing.T) {
	assert.Eq(t, []any{1, "value"}, arrutil.Filter([]any{nil, 0, "", 1, "value"}))
	assert.Eq(t, []any{nil}, arrutil.Filter([]any{nil}, func(any) bool { return true }))
}
