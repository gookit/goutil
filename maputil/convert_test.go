package maputil_test

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestKeyToLower(t *testing.T) {
	src := map[string]string{"A": "v0"}
	ret := maputil.KeyToLower(src)

	assert.Contains(t, ret, "a")
	assert.NotContains(t, ret, "A")
}

func TestToStringMap(t *testing.T) {
	src := map[string]any{"a": "v0", "b": 23}
	ret := maputil.ToStringMap(src)

	assert.Eq(t, ret["a"], "v0")
	assert.Eq(t, ret["b"], "23")

	keys := []string{"key0", "key1"}

	mp := maputil.CombineToSMap(keys, []string{"val0", "val1"})
	assert.Len(t, mp, 2)
	assert.Eq(t, "val0", mp.Str("key0"))
}

func TestToAnyMap(t *testing.T) {
	src := map[string]string{"a": "v0", "b": "23"}

	mp := maputil.ToAnyMap(src)
	assert.Len(t, mp, 2)
	assert.Eq(t, "v0", mp["a"])

	src1 := map[string]any{"a": "v0", "b": "23"}
	mp = maputil.ToAnyMap(src1)
	assert.Len(t, mp, 2)
	assert.Eq(t, "v0", mp["a"])

	_, err := maputil.TryAnyMap(123)
	assert.Err(t, err)
}

func TestHTTPQueryString(t *testing.T) {
	src := map[string]any{"a": "v0", "b": 23}
	str := maputil.HTTPQueryString(src)

	fmt.Println(str)
	assert.Contains(t, str, "b=23")
	assert.Contains(t, str, "a=v0")
}

func TestToString2(t *testing.T) {
	src := map[string]any{"a": "v0", "b": 23}

	s := maputil.ToString2(src)
	assert.Contains(t, s, "b:23")
	assert.Contains(t, s, "a:v0")
}

func TestToString(t *testing.T) {
	src := map[string]any{"a": "v0", "b": 23}

	s := maputil.ToString(src)
	dump.P(s)
	assert.Contains(t, s, "b:23")
	assert.Contains(t, s, "a:v0")

	s = maputil.ToString(nil)
	assert.Eq(t, "", s)

	s = maputil.ToString(map[string]any{})
	assert.Eq(t, "{}", s)

	s = maputil.ToString(map[string]any{"": nil})
	assert.Eq(t, "{:}", s)
}

func TestFlatten(t *testing.T) {
	data := map[string]any{
		"name": "inhere",
		"age":  234,
		"top": map[string]any{
			"sub0": "val0",
			"sub1": []string{"val1-0", "val1-1"},
		},
	}

	mp := maputil.Flatten(data)
	assert.ContainsKeys(t, mp, []string{"age", "name", "top.sub0", "top.sub1[0]", "top.sub1[1]"})
	assert.Nil(t, maputil.Flatten(nil))

	fmp := make(map[string]string)
	maputil.FlatWithFunc(data, func(path string, val reflect.Value) {
		fmp[path] = fmt.Sprintf("%v", val.Interface())
	})
	dump.P(fmp)
	assert.Eq(t, "inhere", fmp["name"])
	assert.Eq(t, "234", fmp["age"])
	assert.Eq(t, "val0", fmp["top.sub0"])
	assert.Eq(t, "val1-0", fmp["top.sub1[0]"])

	assert.NotPanics(t, func() {
		maputil.FlatWithFunc(nil, nil)
	})
}

func TestStringsMapToAnyMap(t *testing.T) {
	assert.Nil(t, maputil.StringsMapToAnyMap(nil))

	hh := http.Header{
		"key0": []string{"val0", "val1"},
		"key1": []string{"val2"},
	}

	mp := maputil.StringsMapToAnyMap(hh)
	assert.Contains(t, mp, "key0")
	assert.Contains(t, mp, "key1")
	assert.Len(t, mp["key0"], 2)

	dm := maputil.Data(mp)
	assert.Eq(t, "val0", dm.Str("key0.0"))
	assert.Eq(t, "val2", dm.Str("key1"))
}
