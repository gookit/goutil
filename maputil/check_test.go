package maputil_test

import (
	"testing"

	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestHasKey(t *testing.T) {
	var mp any = map[string]string{"key0": "val0"}

	assert.True(t, maputil.HasKey(mp, "key0"))
	assert.False(t, maputil.HasKey(mp, "not-exist"))
	assert.False(t, maputil.HasKey("abc", "not-exist"))
}

func TestHasOneKey(t *testing.T) {
	var mp any = map[string]string{"key0": "val0", "key1": "def"}

	ok, key := maputil.HasOneKey(mp, "key0", "not-exist")
	assert.True(t, ok)
	assert.Eq(t, "key0", key)

	ok, key = maputil.HasOneKey("abc", "not-exist")
	assert.Nil(t, key)
	assert.False(t, ok)
}

func TestHasAllKeys(t *testing.T) {
	var mp any = map[string]string{"key0": "val0", "key1": "def"}
	ok, noKey := maputil.HasAllKeys(mp, "key0")
	assert.True(t, ok)
	assert.Nil(t, noKey)

	ok, noKey = maputil.HasAllKeys(mp, "key0", "key1")
	assert.True(t, ok)
	assert.Nil(t, noKey)

	ok, noKey = maputil.HasAllKeys(mp, "key0", "not-exist")
	assert.False(t, ok)
	assert.Eq(t, "not-exist", noKey)

	ok, _ = maputil.HasAllKeys("invalid-map", "not-exist")
	assert.False(t, ok)
}
