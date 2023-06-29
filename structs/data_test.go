package structs_test

import (
	"testing"

	"github.com/gookit/goutil/structs"
	"github.com/gookit/goutil/testutil/assert"
)

func TestLiteData_Data(t *testing.T) {
	d := structs.NewLiteData(nil)
	d.SetData(map[string]any{
		"key0": 234,
		"key1": "abc",
		"key2": true,
	})

	assert.NotEmpty(t, d.Data())
	v, ok := d.Value("key0")
	assert.True(t, ok)
	assert.Eq(t, 234, v)

	d.Merge(map[string]any{
		"key1": "def",
		"key4": "value4",
	})
	assert.Eq(t, "def", d.StrVal("key1"))
	assert.Eq(t, "value4", d.GetVal("key4"))

	d.ResetData()
	assert.Empty(t, d.Data())
}

func TestNewData(t *testing.T) {
	md := structs.NewData()
	assert.Eq(t, 0, md.DataLen())

	md.SetData(map[string]any{
		"key0": 234,
		"sub": map[string]any{
			"skey1": "abc",
			"skey2": true,
		},
	})
	assert.NotEmpty(t, md.Data())
	assert.Eq(t, 234, md.IntVal("key0"))

	v, ok := md.Value("key0")
	assert.True(t, ok)
	assert.Eq(t, 234, v)

	md.SetValue("key1", "val1")
	assert.Eq(t, "val1", md.GetVal("key1"))
	assert.Eq(t, "val1", md.StrVal("key1"))
	assert.False(t, md.BoolVal("key1"))
	assert.False(t, md.BoolVal("not-exist"))

	md.SetValue("bol", true)
	assert.True(t, md.BoolVal("bol"))

	md.ResetData()
	assert.Eq(t, 0, md.DataLen())
}

func TestDataStore_EnableLock(t *testing.T) {
	md := structs.NewData()
	md.EnableLock()

	md.SetData(map[string]any{
		"key0": 234,
		"key1": "abc",
		"key2": true,
	})

	md.Set("key1", "def")
	assert.Eq(t, "def", md.Get("key1"))
	assert.NotEmpty(t, md.String())
}

func TestNewOrderedData(t *testing.T) {
	od := structs.NewOrderedData(10)
	od.Set("key0", 234)
	od.Load(map[string]any{
		"key1": "abc",
		"key2": true,
	})

	assert.NotEmpty(t, od.Keys())
}
