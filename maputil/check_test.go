package maputil_test

import (
	"testing"

	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestHasKey(t *testing.T) {
	var mp interface{}

	mp = map[string]string{"key0": "val0"}
	assert.True(t, maputil.HasKey(mp, "key0"))
	assert.False(t, maputil.HasKey(mp, "not-exist"))
	assert.False(t, maputil.HasKey("abc", "not-exist"))
}
