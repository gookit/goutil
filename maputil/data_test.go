package maputil_test

import (
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestData_usage(t *testing.T) {
	mp := maputil.Data{
		"k1": 23,
		"k2": "ab",
		"k3": "true",
		"k4": false,
		"k5": map[string]string{"a": "b"},
	}

	assert.True(t, mp.Has("k1"))
	assert.True(t, mp.Bool("k3"))
	assert.False(t, mp.Bool("k4"))
	assert.False(t, mp.IsEmtpy())
	assert.Eq(t, 23, mp.Get("k1"))

	// int
	assert.Eq(t, 23, mp.Int("k1"))
	assert.Eq(t, int64(23), mp.Int64("k1"))

	// str
	assert.Eq(t, "23", mp.Str("k1"))
	assert.Eq(t, "ab", mp.Str("k2"))

	// set
	mp.Set("new", "val1")
	assert.Eq(t, "val1", mp.Str("new"))

	val, ok := mp.Value("new")
	assert.True(t, ok)
	assert.Eq(t, "val1", val)

	// not exists
	assert.False(t, mp.Bool("notExists"))
	assert.Eq(t, 0, mp.Int("notExists"))
	assert.Eq(t, int64(0), mp.Int64("notExists"))
	assert.Eq(t, "", mp.Str("notExists"))

	// default
	assert.Eq(t, 23, mp.Default("k1", 10))
	assert.Eq(t, 10, mp.Default("notExists", 10))

	assert.Nil(t, mp.StringMap("notExists"))
	assert.Eq(t, map[string]string{"a": "b"}, mp.StringMap("k5"))
}

func TestData_SetByPath(t *testing.T) {
	mp := maputil.Data{
		"k2": "ab",
		"k5": map[string]any{"a": "v0"},
	}

	err := mp.SetByPath("k5.b", "v2")
	assert.NoErr(t, err)

	dump.P(mp)
}
