package maputil_test

import (
	"sort"
	"strconv"
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

func TestFilterMap(t *testing.T) {
	mp := map[string]int{"a": 1, "b": 2, "c": 3}

	filtered := maputil.Filter(mp, func(k string, v int) bool { return v%2 == 1 })
	assert.Eq(t, map[string]int{"a": 1, "c": 3}, filtered)
	assert.Nil(t, maputil.Filter(map[string]int{}, func(k string, v int) bool { return true }))

	mapped := maputil.MapValues(mp, func(k string, v int) string { return k + strconv.Itoa(v) })
	assert.Eq(t, map[string]string{"a": "a1", "b": "b2", "c": "c3"}, mapped)
	assert.Nil(t, maputil.MapValues(map[string]int{}, func(k string, v int) string { return "" }))
}

func TestInvert(t *testing.T) {
	assert.Eq(t, map[int]string{1: "a", 2: "b"}, maputil.Invert(map[string]int{"a": 1, "b": 2}))
	assert.Nil(t, maputil.Invert(map[string]int{}))
}

func TestPickOmit(t *testing.T) {
	mp := map[string]int{"a": 1, "b": 2, "c": 3}

	assert.Eq(t, map[string]int{"a": 1, "c": 3}, maputil.Pick(mp, "a", "c"))
	// a key that is not in the map is ignored
	assert.Eq(t, map[string]int{"a": 1}, maputil.Pick(mp, "a", "x"))
	assert.Nil(t, maputil.Pick(mp))
	assert.Nil(t, maputil.Pick(map[string]int{}))

	assert.Eq(t, map[string]int{"a": 1}, maputil.Omit(mp, "b", "c"))
	// omit more keys than the map has
	assert.Eq(t, 0, len(maputil.Omit(mp, "a", "b", "c", "d")))
	assert.Eq(t, mp, maputil.Omit(mp))
	assert.Nil(t, maputil.Omit(map[string]int{}, "a"))
}

func TestMergeFunc(t *testing.T) {
	dst := map[string]int{"a": 1, "b": 2}
	maputil.MergeFunc(dst, map[string]int{"a": 10, "c": 3}, func(dstVal, srcVal int) int {
		return dstVal + srcVal
	})
	assert.Eq(t, map[string]int{"a": 11, "b": 2, "c": 3}, dst)
}

func TestTypedKeysValues(t *testing.T) {
	mp := map[string]int{"a": 1, "b": 2}
	keys := maputil.TypedKeys(mp)
	sort.Strings(keys)
	assert.Eq(t, []string{"a", "b"}, keys)

	values := maputil.TypedValues(mp)
	sort.Ints(values)
	assert.Eq(t, []int{1, 2}, values)

	// empty map gives nil
	assert.Nil(t, maputil.TypedKeys(map[string]int{}))
	assert.Nil(t, maputil.TypedValues(map[string]int{}))
}
