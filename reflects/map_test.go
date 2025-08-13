package reflects_test

import (
	"reflect"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/reflects"
	"github.com/gookit/goutil/testutil/assert"
)

func TestFlatMap(t *testing.T) {
	assert.Panics(t, func() {
		reflects.FlatMap(reflect.ValueOf("abc"), func(_ string, _ reflect.Value) {
			// nothing
		})
	})

	assert.NotPanics(t, func() {
		reflects.FlatMap(reflect.ValueOf(nil), nil)
	})

	mp := map[string]any{
		"name": "inhere",
		"age":  234,
		"top": map[string]any{
			"sub0": "val0",
			"sub1": []string{"val1-0", "val1-1"},
		},
	}

	flatMp := make(map[string]any, len(mp)*2)
	reflects.FlatMap(reflect.ValueOf(mp), func(path string, val reflect.Value) {
		flatMp[path] = val.Interface()
	})
	dump.P(flatMp)
	assert.Eq(t, "inhere", flatMp["name"])
	assert.Eq(t, "val0", flatMp["top.sub0"])
	assert.Eq(t, "val1-0", flatMp["top.sub1[0]"])
}

func TestTryAnyMap(t *testing.T) {
	smp := map[int]string{
		1: "val1",
		2: "val2",
	}

	mp, err := reflects.TryAnyMap(reflect.ValueOf(smp))
	assert.Nil(t, err)
	assert.Eq(t, "val1", mp["1"])
	assert.Eq(t, "val2", mp["2"])
}

func TestEachStrAnyMap(t *testing.T) {
	smp := map[int]string{
		1: "val1",
		2: "val2",
	}

	mp := make(map[string]any)
	err := reflects.EachStrAnyMap(reflect.ValueOf(smp), func(key string, val any) {
		mp[key] = val
	})

	assert.NoErr(t, err)
	assert.Eq(t, "val1", mp["1"])
	assert.Eq(t, "val2", mp["2"])

	assert.NoErr(t, reflects.EachMap(reflect.ValueOf("abc"), nil))
	assert.Err(t, reflects.EachMap(reflect.ValueOf("abc"), func(key, val reflect.Value) {
		// do nothing
	}))
}
