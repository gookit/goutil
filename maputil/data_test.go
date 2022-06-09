package maputil_test

import (
	"testing"

	"github.com/gookit/goutil/maputil"
	"github.com/stretchr/testify/assert"
)

func TestData(t *testing.T) {
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
	assert.Equal(t, 23, mp.Get("k1"))

	// int
	assert.Equal(t, 23, mp.Int("k1"))
	assert.Equal(t, int64(23), mp.Int64("k1"))

	// str
	assert.Equal(t, "23", mp.Str("k1"))
	assert.Equal(t, "ab", mp.Str("k2"))

	// set
	mp.Set("new", "val1")
	assert.Equal(t, "val1", mp.Str("new"))

	val, ok := mp.Value("new")
	assert.True(t, ok)
	assert.Equal(t, "val1", val)

	// not exists
	assert.False(t, mp.Bool("notExists"))
	assert.Equal(t, 0, mp.Int("notExists"))
	assert.Equal(t, int64(0), mp.Int64("notExists"))
	assert.Equal(t, "", mp.Str("notExists"))

	// default
	assert.Equal(t, 23, mp.Default("k1", 10))
	assert.Equal(t, 10, mp.Default("notExists", 10))

	assert.Nil(t, mp.StringMap("notExists"))
	assert.Equal(t, map[string]string{"a": "b"}, mp.StringMap("k5"))
}
