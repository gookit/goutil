package maputil_test

import (
	"fmt"
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

	assert.NotPanics(t, func() {
		maputil.FlatWithFunc(nil, nil)
	})
}
