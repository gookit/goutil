package structs_test

import (
	"testing"

	"github.com/gookit/goutil/structs"
	"github.com/gookit/goutil/testutil/assert"
)

func TestNewMapData(t *testing.T) {
	md := structs.NewMapData()

	assert.Eq(t, 0, md.Len())

	md.SetData(map[string]interface{}{
		"key0": 234,
	})
	assert.NotEmpty(t, md.Data())
	assert.Eq(t, 234, md.IntVal("key0"))

	md.SetValue("key1", "val1")
	assert.Eq(t, "val1", md.GetVal("key1"))
	assert.Eq(t, "val1", md.StrVal("key1"))
	assert.False(t, md.BoolVal("key1"))
	assert.False(t, md.BoolVal("not-exist"))

	md.SetValue("bol", true)
	assert.True(t, md.BoolVal("bol"))

	md.Clear()
	assert.Eq(t, 0, md.Len())
}

func TestMapDataStore_EnableLock(t *testing.T) {
	md := structs.NewMapData()
	md.EnableLock()

	md.SetData(map[string]interface{}{
		"key0": 234,
		"key1": "abc",
		"key2": true,
	})

	md.Set("key1", "def")
	assert.Eq(t, "def", md.Get("key1"))
	assert.NotEmpty(t, md.String())
}
