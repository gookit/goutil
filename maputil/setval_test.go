package maputil_test

import (
	"reflect"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/testutil/assert"
)

func makMapForSetByPath() map[string]interface{} {
	return map[string]interface{}{
		"key0": "v0",
		"key2": 34,
		"key3": map[string]interface{}{
			"k301": "v301",
			"k303": []string{"v303-1", "v303-2"},
		},
		"key4": map[string]string{
			"k401": "v401",
		},
	}
}

func TestSetByPath2_map_add_item(t *testing.T) {
	mp := makMapForSetByPath()
	val := "add-new-item"

	keys1 := []string{"key5"} // ok
	err1 := maputil.SetByKeys2(mp, keys1, val)
	assert.NoErr(t, err1)
	assert.ContainsKey(t, mp, "key5")
	assert.Eq(t, val, maputil.QuietGet(mp, "key5"))

	keys2 := []string{"key3", "k302"} // ok
	err2 := maputil.SetByKeys2(mp, keys2, val)
	assert.NoErr(t, err2)
	assert.Eq(t, val, maputil.QuietGet(mp, "key3.k302"))

	keys4 := []string{"key4", "k402"} // ok
	err4 := maputil.SetByKeys2(mp, keys4, val)
	assert.NoErr(t, err4)
	assert.Eq(t, val, maputil.DeepGet(mp, "key4.k402"))
	dump.Println(mp)
}

func TestSetByPath2_map_up_val(t *testing.T) {
	mp := makMapForSetByPath()
	val := "set-new-val"

	keys1 := []string{"key0"} // ok
	err1 := maputil.SetByKeys2(mp, keys1, val)
	assert.NoErr(t, err1)
	assert.Eq(t, val, maputil.QuietGet(mp, "key0"))

	keys2 := []string{"key3", "k301"} // ok
	err2 := maputil.SetByKeys2(mp, keys2, val)
	assert.NoErr(t, err2)
	assert.Eq(t, val, maputil.QuietGet(mp, "key3.k301"))

	keys4 := []string{"key4", "k401"} // ok
	err4 := maputil.SetByKeys2(mp, keys4, val)
	assert.NoErr(t, err4)
	assert.Eq(t, val, maputil.DeepGet(mp, "key4.k401"))
	dump.Println(mp)
}

func TestSetByPath2_slice_val(t *testing.T) {
	mp := makMapForSetByPath()

	keys3 := []string{"key3", "k303[2]"} // ok
	err3 := maputil.SetByKeys2(mp, keys3, "new-value")
	assert.NoErr(t, err3)
	dump.Println(mp)

	keys4 := []string{"key3", "k303.3"} // ok
	err4 := maputil.SetByKeys2(mp, keys4, "new-value")
	assert.NoErr(t, err4)
	dump.Println(mp)
}

func TestSliceItemType(t *testing.T) {
	sl := []string{"abc"}
	ty := reflect.TypeOf(sl)

	assert.Eq(t, reflect.Slice, ty.Kind())
	assert.Eq(t, reflect.String, ty.Elem().Kind())
}
