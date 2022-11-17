package maputil_test

import (
	"testing"

	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestHasKey(t *testing.T) {
	var mp any

	mp = map[string]string{"key0": "val0"}
	assert.True(t, maputil.HasKey(mp, "key0"))
	assert.False(t, maputil.HasKey(mp, "not-exist"))
	assert.False(t, maputil.HasKey("abc", "not-exist"))
}

func TestHasAllKeys(t *testing.T) {
	var mp any

	mp = map[string]string{"key0": "val0", "key1": "def"}
	ok, noKey := maputil.HasAllKeys(mp, "key0")
	assert.True(t, ok)
	assert.Nil(t, noKey)

	ok, noKey = maputil.HasAllKeys(mp, "key0", "key1")
	assert.True(t, ok)
	assert.Nil(t, noKey)

	ok, noKey = maputil.HasAllKeys(mp, "key0", "not-exist")
	assert.False(t, ok)
	assert.Eq(t, "not-exist", noKey)

	ok, _ = maputil.HasAllKeys(mp, "invalid-map", "not-exist")
	assert.False(t, ok)
}
