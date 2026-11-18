package maputil_test

import (
	"encoding/json"
	"testing"

	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/x/assert"
)

func TestDeleteByPath(t *testing.T) {
	newMp := func(s string) map[string]any {
		var mp map[string]any
		assert.NoErr(t, json.Unmarshal([]byte(s), &mp))
		return mp
	}
	toJSON := func(mp map[string]any) string {
		b, err := json.Marshal(mp)
		assert.NoErr(t, err)
		return string(b)
	}

	for _, tc := range []struct{ name, input, path, want string }{
		{"top level", `{"a":1,"b":2}`, "a", `{"b":2}`},
		{"nested key", `{"a":{"b":1,"c":2}}`, "a.b", `{"a":{"c":2}}`},
		{"deep nested key", `{"a":{"b":{"c":1,"d":2}}}`, "a.b.c", `{"a":{"b":{"d":2}}}`},
		{"missing key", `{"a":1}`, "x", `{"a":1}`},
		{"missing nested key", `{"a":{"b":1}}`, "a.x", `{"a":{"b":1}}`},
		{"missing parent", `{"a":1}`, "x.y.z", `{"a":1}`},
		{"slice elem index syntax", `{"list":["a","b","c"]}`, "list[1]", `{"list":["a","c"]}`},
		{"slice elem dot syntax", `{"list":["a","b","c"]}`, "list.1", `{"list":["a","c"]}`},
		{"last slice elem", `{"list":["a","b"]}`, "list[1]", `{"list":["a"]}`},
		{"index out of range", `{"list":["a"]}`, "list[3]", `{"list":["a"]}`},
		{"key of a slice elem", `{"list":[{"a":1,"b":2},{"a":3}]}`, "list[0].a", `{"list":[{"b":2},{"a":3}]}`},
		{"nested slice elem", `{"a":{"list":[1,2,3]}}`, "a.list[0]", `{"a":{"list":[2,3]}}`},
		{"nil value of a key", `{"a":{"b":null}}`, "a.b", `{"a":{}}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			mp := newMp(tc.input)
			assert.NoErr(t, maputil.DeleteByPath(&mp, tc.path))
			assert.Eq(t, tc.want, toJSON(mp))
		})
	}

	t.Run("nil map", func(t *testing.T) {
		var mp map[string]any
		assert.NoErr(t, maputil.DeleteByPath(&mp, "a.b"))
		assert.Nil(t, mp)
	})
	t.Run("empty path", func(t *testing.T) {
		mp := newMp(`{"a":1}`)
		assert.NoErr(t, maputil.DeleteByKeys(&mp, nil))
		assert.Eq(t, `{"a":1}`, toJSON(mp))
	})
	t.Run("non-string key map", func(t *testing.T) {
		mp := map[string]any{"a": map[int]string{1: "one"}}
		assert.Err(t, maputil.DeleteByPath(&mp, "a.1"))
	})
	t.Run("delete from a scalar", func(t *testing.T) {
		mp := map[string]any{"a": "text"}
		assert.Err(t, maputil.DeleteByPath(&mp, "a.b"))
	})
}

func TestDeleteByPath_typedValue(t *testing.T) {
	t.Run("typed map value", func(t *testing.T) {
		mp := map[string]any{"group": map[string]string{"a": "1", "b": "2"}}
		assert.NoErr(t, maputil.DeleteByPath(&mp, "group.a"))
		assert.Eq(t, map[string]string{"b": "2"}, mp["group"])
	})
	t.Run("typed slice value", func(t *testing.T) {
		mp := map[string]any{"group": map[string][]int{"list": {1, 2, 3}}}
		assert.NoErr(t, maputil.DeleteByPath(&mp, "group.list[1]"))
		assert.Eq(t, map[string][]int{"list": {1, 3}}, mp["group"])
	})
	t.Run("typed slice of structs", func(t *testing.T) {
		type item struct {
			Name string
			Age  int
		}
		mp := map[string]any{"list": []item{{"a", 1}, {"b", 2}}}
		assert.NoErr(t, maputil.DeleteByPath(&mp, "list[0]"))
		assert.Eq(t, []item{{"b", 2}}, mp["list"])
	})
	t.Run("named string key", func(t *testing.T) {
		type key string
		mp := map[string]any{"group": map[key]int{"a": 1, "b": 2}}
		assert.NoErr(t, maputil.DeleteByPath(&mp, "group.a"))
		assert.Eq(t, map[key]int{"b": 2}, mp["group"])
	})
}

func TestData_DeleteByPath(t *testing.T) {
	d := maputil.Data{"group": map[string]any{"value": "old", "keep": true}}
	assert.NoErr(t, d.DeleteByPath("group.value"))

	group, ok := d["group"].(map[string]any)
	assert.True(t, ok)
	_, has := group["value"]
	assert.False(t, has)
	assert.Eq(t, true, group["keep"])

	// the top level key of the same map is removed too
	assert.NoErr(t, d.DeleteByPath("group"))
	_, has = d["group"]
	assert.False(t, has)
}
