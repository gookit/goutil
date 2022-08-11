package maputil_test

import (
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/testutil/assert"
)

func makMapForSetByPath() map[string]interface{} {
	return map[string]any{
		"key0": "v0",
		"key2": 34,
		"key3": map[string]any{
			"k301": "v301",
			"k303": []string{"v303-1", "v303-2"},
			"k304": map[string]any{
				"k3041": "v3041",
				"k3042": []string{"k3042-1", "k3042-2"},
			},
			"k305": []any{
				map[string]string{
					"k3051": "v3051",
				},
			},
		},
		"key4": map[string]string{
			"k401": "v401",
		},
		"key6": []any{
			map[string]string{
				"k3051": "v3051",
			},
		},
	}
}

func TestSetByPath2_map_add_key(t *testing.T) {
	mp := makMapForSetByPath()
	val := "add-new-key"

	keys1 := []string{"key5"} // ok
	err1 := maputil.SetByKeys2(&mp, keys1, val)
	assert.NoErr(t, err1)
	assert.ContainsKey(t, mp, "key5")
	assert.Eq(t, val, maputil.QuietGet(mp, "key5"))

	// set to map[string]any
	keys2 := []string{"key3", "k302"} // ok
	err2 := maputil.SetByKeys2(&mp, keys2, val)
	assert.NoErr(t, err2)
	assert.Eq(t, val, maputil.QuietGet(mp, "key3.k302"))
	// more deep
	keys3 := []string{"key3", "k304", "k3043"} // ok
	err3 := maputil.SetByKeys2(&mp, keys3, val)
	assert.NoErr(t, err3)
	assert.Eq(t, val, maputil.QuietGet(mp, "key3.k304.k3043"))

	// set to map[string]string
	keys4 := []string{"key4", "k402"} // ok
	err4 := maputil.SetByKeys2(&mp, keys4, val)
	assert.NoErr(t, err4)
	assert.Eq(t, val, maputil.DeepGet(mp, "key4.k402"))
	dump.Println(mp)
}

func TestSetByPath2_map_up_val(t *testing.T) {
	mp := makMapForSetByPath()
	val := "set-new-val"

	keys1 := []string{"key0"} // ok
	err1 := maputil.SetByKeys2(&mp, keys1, val)
	assert.NoErr(t, err1)
	assert.Eq(t, val, maputil.QuietGet(mp, "key0"))

	keys2 := []string{"key3", "k301"} // ok
	err2 := maputil.SetByKeys2(&mp, keys2, val)
	assert.NoErr(t, err2)
	assert.Eq(t, val, maputil.QuietGet(mp, "key3.k301"))

	keys4 := []string{"key4", "k401"} // ok
	err4 := maputil.SetByKeys2(&mp, keys4, val)
	assert.NoErr(t, err4)
	assert.Eq(t, val, maputil.DeepGet(mp, "key4.k401"))
	dump.Println(mp)
}

func TestSetByPath2_slice_val1(t *testing.T) {
	mp := makMapForSetByPath()

	// nVal := "new-value"
	// keys3 := []string{"key3", "k303", "1"} // ok
	// err3 := maputil.SetByKeys2(&mp, keys3, nVal)
	// assert.NoErr(t, err3)
	// assert.Eq(t, nVal, maputil.QuietGet(mp, "key3.k303.1"))

	nItem := "a-new-item"
	keys4 := []string{"key3", "k303", "2"} // ok
	err4 := maputil.SetByKeys2(&mp, keys4, nItem)
	assert.NoErr(t, err4)
	assert.Eq(t, nItem, maputil.QuietGet(mp, "key3.k303.2"))
	dump.Println(mp)
}

func TestSetByPath2_slice_val2(t *testing.T) {
	mp := makMapForSetByPath()
	nVal := "new-value"

	keys2 := []string{"key3", "k303[1]"} // ok
	err2 := maputil.SetByKeys2(&mp, keys2, nVal)
	assert.NoErr(t, err2)
	assert.Eq(t, nVal, maputil.QuietGet(mp, "key3.k303.1"))
	dump.Println(mp)

	// keys3 := []string{"key3", "k303[2]"} // ok
	// err3 := maputil.SetByKeys2(&mp, keys3, "new-item")
	// assert.NoErr(t, err3)
	// assert.Len(t, mp["key3"], 3)
	// dump.Println(mp)
}
