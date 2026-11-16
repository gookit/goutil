package maputil_test

import (
	"strings"
	"testing"

	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/x/assert"
)

func TestEqual(t *testing.T) {
	assert.True(t, maputil.Equal(map[string]int{"a": 1}, map[string]int{"a": 1}))
	assert.False(t, maputil.Equal(map[string]int{"a": 1}, map[string]int{"a": 2}))
	assert.False(t, maputil.Equal(map[string]int{"a": 1}, map[string]int{"a": 1, "b": 2}))
	assert.True(t, maputil.Equal(map[string]int{}, map[string]int(nil)))

	assert.True(t, maputil.EqualFunc(
		map[string]string{"a": "A"},
		map[string]string{"a": "a"},
		func(a, b string) bool { return strings.EqualFold(a, b) },
	))
	assert.False(t, maputil.EqualFunc(
		map[string]string{"a": "A"},
		map[string]string{"b": "a"},
		func(a, b string) bool { return strings.EqualFold(a, b) },
	))
	// the value types may differ
	assert.True(t, maputil.EqualFunc(
		map[string]int{"a": 1},
		map[string]any{"a": 1},
		func(a int, b any) bool { return a == b },
	))
}
