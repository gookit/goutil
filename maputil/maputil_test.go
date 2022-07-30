package maputil_test

import (
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
)

// TODO remove
type any = interface{}

func TestMergeStringMap(t *testing.T) {
	ret := maputil.MergeStringMap(map[string]string{"A": "v0"}, map[string]string{"A": "v1"}, false)
	assert.Eq(t, map[string]string{"A": "v0"}, ret)

	ret = maputil.MergeStringMap(map[string]string{"A": "v0"}, map[string]string{"a": "v1"}, true)
	assert.Eq(t, map[string]string{"a": "v0"}, ret)
}

func TestGetByPath(t *testing.T) {
	mp := map[string]interface{}{
		"key0": "val0",
		"key1": map[string]string{"sk0": "sv0"},
		"key2": []string{"sv1", "sv2"},
		"key3": map[string]interface{}{"sk1": "sv1"},
		"key4": []int{1, 2},
		"key5": []interface{}{1, "2", true},
	}

	v, ok := maputil.GetByPath("key0", mp)
	assert.True(t, ok)
	assert.Eq(t, "val0", v)

	v, ok = maputil.GetByPath("key1.sk0", mp)
	assert.True(t, ok)
	assert.Eq(t, "sv0", v)

	v, ok = maputil.GetByPath("key3.sk1", mp)
	assert.True(t, ok)
	assert.Eq(t, "sv1", v)

	// not exists
	v, ok = maputil.GetByPath("not-exits", mp)
	assert.False(t, ok)
	assert.Nil(t, v)
	v, ok = maputil.GetByPath("key2.not-exits", mp)
	assert.False(t, ok)
	assert.Nil(t, v)
	v, ok = maputil.GetByPath("not-exits.subkey", mp)
	assert.False(t, ok)
	assert.Nil(t, v)

	// Slices behaviour
	v, ok = maputil.GetByPath("key2", mp)
	assert.True(t, ok)
	assert.Eq(t, mp["key2"], v)

	v, ok = maputil.GetByPath("key2.0", mp)
	assert.True(t, ok)
	assert.Eq(t, "sv1", v)

	v, ok = maputil.GetByPath("key2.1", mp)
	assert.True(t, ok)
	assert.Eq(t, "sv2", v)

	v, ok = maputil.GetByPath("key4.0", mp)
	assert.True(t, ok)
	assert.Eq(t, 1, v)

	v, ok = maputil.GetByPath("key4.1", mp)
	assert.True(t, ok)
	assert.Eq(t, 2, v)

	v, ok = maputil.GetByPath("key5.0", mp)
	assert.True(t, ok)
	assert.Eq(t, 1, v)

	v, ok = maputil.GetByPath("key5.1", mp)
	assert.True(t, ok)
	assert.Eq(t, "2", v)

	v, ok = maputil.GetByPath("key5.2", mp)
	assert.True(t, ok)
	assert.Eq(t, true, v)

	// Out of bound value
	v, ok = maputil.GetByPath("key2.2", mp)
	assert.False(t, ok)
	assert.Nil(t, v)
}

func TestKeys(t *testing.T) {
	mp := map[string]interface{}{
		"key0": "v0",
		"key1": "v1",
		"key2": 34,
	}

	ln := len(mp)
	ret := maputil.Keys(mp)
	assert.Len(t, ret, ln)
	assert.Contains(t, ret, "key0")
	assert.Contains(t, ret, "key1")
	assert.Contains(t, ret, "key2")

	ret = maputil.Keys(&mp)
	assert.Len(t, ret, ln)
	assert.Contains(t, ret, "key0")
	assert.Contains(t, ret, "key1")

	ret = maputil.Keys(struct {
		a string
	}{"v"})

	assert.Len(t, ret, 0)
}

func TestValues(t *testing.T) {
	mp := map[string]interface{}{
		"key0": "v0",
		"key1": "v1",
		"key2": 34,
	}

	ln := len(mp)
	ret := maputil.Values(mp)

	assert.Len(t, ret, ln)
	assert.Contains(t, ret, "v0")
	assert.Contains(t, ret, "v1")
	assert.Contains(t, ret, 34)

	ret = maputil.Values(struct {
		a string
	}{"v"})

	assert.Len(t, ret, 0)
}

func TestMakeByPath(t *testing.T) {
	mp := maputil.MakeByPath("top.sub", "val")
	assert.NotEmpty(t, mp)
	assert.ContainsKey(t, mp, "top")
	assert.ContainsKey(t, mp, "top")
}

func TestSetByPath2(t *testing.T) {
	mp := map[string]interface{}{
		"key0": "v0",
		"key2": 34,
		"key3": map[string]interface{}{
			"k303": []string{"v303-1", "v303-2"},
		},
		"key4": map[string]string{
			"k401": "v401",
		},
	}

	// reflect.MapOf()

	rv := reflect.ValueOf(mp)
	nv := reflect.ValueOf("new-val")

	// keys := []string{"key5"} // ok
	// keys := []string{"key3", "k302"} // ok
	keys := []string{"key3", "k303[2]"} // ok

	var err error

	for i, key := range keys {
		idx := -1
		isSlice := false
		isPtr := false
		// set value on last key
		isLast := i == len(keys)-1

		// eg: k303[2]
		if pos := strings.IndexRune(key, '['); pos > 0 {
			idx, err = strconv.Atoi(key[pos : len(key)-2])
			key = key[:pos]
		}

		dump.P(isSlice, isPtr)

		isMap := rv.Kind() == reflect.Map
		if isLast {
			if isMap {
				rv.SetMapIndex(reflect.ValueOf(key), nv)
			} else if rv.Kind() == reflect.Slice && idx > -1 {

				if idx > rv.Len() {
					// rv.SetLen(idx+1)
					rv = reflect.Append(rv, reflect.New(nv.Type())) // TODO
				}

				rv.Index(idx).Set(nv)
			} else {
				err = errorx.Rawf("cannot set value for path %q", strings.Join(keys[i:], "."))
			}
			break
		}

		if isMap {
			k := reflect.ValueOf(key)
			if v := rv.MapIndex(k); v.IsValid() {
				// get real type: any -> map
				if v.Kind() == reflect.Interface {
					v = v.Elem()
				}
				if v.Kind() == reflect.Ptr {
					isPtr = true
					v = v.Elem()
				}

				if v.Kind() == reflect.Map {
					rv = v
				} else if v.Kind() == reflect.Slice {
					isSlice = true
					rv = v

					if idx > rv.Len() {
						rv = reflect.Append(rv, reflect.New(nv.Type())) // TODO
					}

					rv.Index(idx).Set(nv)
				}
			}
		} else if rv.Kind() == reflect.Slice && strutil.IsNumeric(key) { // slice
			idx, _ = strconv.Atoi(key)
		} else {
			err = errorx.Rawf("cannot set value for path %q", strings.Join(keys[i:], "."))
			// return
		}
	}

	assert.NoErr(t, err)
	dump.Println(mp)
}

func TestSliceType(t *testing.T) {

	sl := []string{"abc"}
	ty := reflect.TypeOf(sl)

	dump.P(ty.Kind().String())
	if ty.Kind() == reflect.Slice {
		ty = ty.Elem()
		dump.P(ty.Kind().String())
	}
}

func TestSetByPath(t *testing.T) {
	mp := map[string]interface{}{
		"key0": "v0",
		"key1": "v1",
		"key2": 34,
	}

	nmp, err := maputil.SetByPath(mp, "key0", "v00")
	assert.NoErr(t, err)
	assert.ContainsKey(t, nmp, "key0")
	assert.Eq(t, "v00", nmp["key0"])

	nmp, err = maputil.SetByPath(mp, "key3", map[string]interface{}{
		"k301": "v301",
		"k302": 234,
		"k303": []string{"v303-1", "v303-2"},
		"k304": nil,
	})

	assert.NoErr(t, err)
	// dump.P(nmp, mp)
	assert.ContainsKeys(t, nmp, []string{"key3"})
	assert.ContainsKeys(t, nmp["key3"], []string{"k301", "k302", "k303", "k304"})

	nmp, err = maputil.SetByPath(mp, "key4", map[string]string{
		"k401": "v401",
	})
	assert.NoErr(t, err)
	assert.ContainsKey(t, nmp, "key3")

	val, ok := maputil.GetByPath("key4.k401", nmp)
	assert.True(t, ok)
	assert.Eq(t, "v401", val)

	nmp, err = maputil.SetByPath(mp, "key4.k402", "v402")
	assert.NoErr(t, err)
	dump.P(nmp, mp)
}
