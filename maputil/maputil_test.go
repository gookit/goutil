package maputil_test

import (
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/testutil/assert"
)

// TODO remove
type any = interface{}

func TestMergeStringMap(t *testing.T) {
	ret := maputil.MergeSMap(map[string]string{"A": "v0"}, map[string]string{"A": "v1"}, false)
	assert.Eq(t, map[string]string{"A": "v0"}, ret)

	ret = maputil.MergeSMap(map[string]string{"A": "v0"}, map[string]string{"a": "v1"}, true)
	assert.Eq(t, map[string]string{"a": "v0"}, ret)
}

func TestMakeByPath(t *testing.T) {
	mp := maputil.MakeByPath("top.sub", "val")
	assert.NotEmpty(t, mp)
	assert.ContainsKey(t, mp, "top")
	assert.ContainsKey(t, mp, "top")
}

func TestSetByPath(t *testing.T) {
	mp := map[string]interface{}{
		"key0": "v0",
		"key1": "v1",
		"key2": 34,
	}

	err := maputil.SetByPath(&mp, "key0", "v00")
	assert.NoErr(t, err)
	assert.ContainsKey(t, mp, "key0")
	assert.Eq(t, "v00", mp["key0"])

	err = maputil.SetByPath(&mp, "key3", map[string]interface{}{
		"k301": "v301",
		"k302": 234,
		"k303": []string{"v303-1", "v303-2"},
		"k304": nil,
	})

	// dump.P(mp)
	assert.NoErr(t, err)
	assert.ContainsKeys(t, mp, []string{"key3"})
	assert.ContainsKeys(t, mp["key3"], []string{"k301", "k302", "k303", "k304"})

	err = maputil.SetByPath(&mp, "key4", map[string]string{
		"k401": "v401",
	})
	assert.NoErr(t, err)
	assert.ContainsKey(t, mp, "key3")

	val, ok := maputil.GetByPath("key4.k401", mp)
	assert.True(t, ok)
	assert.Eq(t, "v401", val)

	err = maputil.SetByPath(&mp, "key4.k402", "v402")
	assert.NoErr(t, err)

	dump.P(mp)
}
