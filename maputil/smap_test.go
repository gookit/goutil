package maputil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestSMap_usage(t *testing.T) {
	mp := maputil.SMap{
		"k1": "23",
		"k2": "ab",
		"k3": "true",
		"k4": "1,2",
	}

	assert.True(t, mp.Has("k1"))
	assert.True(t, mp.HasValue("true"))
	assert.True(t, mp.Bool("k3"))
	assert.False(t, mp.IsEmpty())
	assert.False(t, mp.HasValue("not-exist"))
	assert.Len(t, mp.Keys(), 4)
	assert.Len(t, mp.Values(), 4)

	val, ok := mp.Value("k2")
	assert.True(t, ok)
	assert.Eq(t, "ab", val)

	// int
	assert.Eq(t, 23, mp.Int("k1"))
	assert.Eq(t, int64(23), mp.Int64("k1"))

	// str
	assert.Eq(t, "23", mp.Str("k1"))
	assert.Eq(t, "ab", mp.Get("k2"))

	// slice
	assert.Eq(t, []int{1, 2}, mp.Ints("k4"))
	assert.Eq(t, []string{"1", "2"}, mp.Strings("k4"))
	assert.Nil(t, mp.Strings("not-exist"))

	// Default
	assert.Eq(t, "ab", mp.Default("k2", "abc"))
	assert.Eq(t, "abc", mp.Default("notExists", "abc"))

	// not exists
	assert.False(t, mp.Bool("notExists"))
	assert.Eq(t, 0, mp.Int("notExists"))
	assert.Eq(t, int64(0), mp.Int64("notExists"))
	assert.Eq(t, "", mp.Str("notExists"))
	assert.Empty(t, mp.Ints("notExists"))
}

func TestSMap_ToKVPairs(t *testing.T) {
	mp := maputil.SMap{
		"k1": "23",
		"k2": "ab",
	}
	arr := mp.ToKVPairs()
	assert.Len(t, arr, 4)
	str := fmt.Sprint(arr)
	assert.StrContains(t, str, "k1 23")
	assert.StrContains(t, str, "k2 ab")

	mp.Set("k3", "true")
	assert.Eq(t, "true", mp.Get("k3"))

	mp.Load(map[string]string{
		"k4": "1,2",
	})
	assert.Eq(t, "1,2", mp.Get("k4"))
}
