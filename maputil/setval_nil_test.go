package maputil_test

import (
	"encoding/json"
	"testing"

	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/x/assert"
)

func TestSetByPath_nilJSONValue(t *testing.T) {
	for _, tc := range []struct{ name, input, path, want string }{
		{"top level", `{"keep":true}`, "value", `{"keep":true,"value":null}`},
		{"new parent", `{"keep":true}`, "new.value", `{"keep":true,"new":{"value":null}}`},
		{"existing nested key", `{"group":{"value":"old"}}`, "group.value", `{"group":{"value":null}}`},
		{"new nested key", `{"group":{}}`, "group.value", `{"group":{"value":null}}`},
		{"new intermediate map", `{"group":{}}`, "group.new.value", `{"group":{"new":{"value":null}}}`},
		{"empty map indexed value", `{}`, "list[1]", `{"list":[null,null]}`},
		{"new indexed value", `{"keep":true}`, "list[1]", `{"keep":true,"list":[null,null]}`},
		{"new nested indexed value", `{"group":{}}`, "group.list[1]", `{"group":{"list":[null,null]}}`},
		{"new intermediate indexed value", `{"group":{}}`, "group.new.list[1]", `{"group":{"new":{"list":[null,null]}}}`},
		{"new indexed map", `{"keep":true}`, "list[1].value", `{"keep":true,"list":[null,{"value":null}]}`},
		{"new parent indexed value", `{"keep":true}`, "new.list[1]", `{"keep":true,"new":{"list":[null,null]}}`},
		{"existing indexed value", `{"list":["old"]}`, "list[0]", `{"list":[null]}`},
		{"grown indexed value", `{"list":["old"]}`, "list.1", `{"list":["old",null]}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var mp map[string]any
			assert.NoErr(t, json.Unmarshal([]byte(tc.input), &mp))
			assert.NoErr(t, maputil.SetByPath(&mp, tc.path, nil))
			out, err := json.Marshal(mp)
			assert.NoErr(t, err)
			assert.Eq(t, tc.want, string(out))
		})
	}
}

func TestSetByPath_nilTypedValue(t *testing.T) {
	for _, tc := range []struct {
		name        string
		value, want any
	}{
		{"pointer", map[string]*int{"value": new(int)}, map[string]*int{"value": nil}},
		{"map", map[string]map[string]int{"value": {"old": 1}}, map[string]map[string]int{"value": nil}},
		{"slice", map[string][]string{"value": {"old"}}, map[string][]string{"value": nil}},
		{"channel", map[string]chan int{"value": make(chan int)}, map[string]chan int{"value": nil}},
		{"function", map[string]func(){"value": func() {}}, map[string]func(){"value": nil}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			mp := map[string]any{"group": tc.value}
			assert.NoErr(t, maputil.SetByPath(&mp, "group.value", nil))
			assert.Eq(t, tc.want, mp["group"])
		})
	}

	t.Run("concrete map value", func(t *testing.T) {
		mp := map[string]any{"group": map[string]string{"value": "old"}}
		assert.ErrMsgContains(t, maputil.SetByPath(&mp, "group.value", nil), "cannot set a nil value to a string map value")
		assert.Eq(t, map[string]string{"value": "old"}, mp["group"])
	})
	t.Run("new nullable slice", func(t *testing.T) {
		mp := map[string]any{"group": map[string][]*int{}}
		assert.NoErr(t, maputil.SetByPath(&mp, "group.list[1]", nil))
		assert.Eq(t, map[string][]*int{"list": {nil, nil}}, mp["group"])
	})
	t.Run("new concrete slice", func(t *testing.T) {
		mp := map[string]any{"group": map[string][]string{}}
		assert.ErrMsgContains(t, maputil.SetByPath(&mp, "group.list[1]", nil), "cannot set a nil value to a string slice elem")
		assert.Eq(t, map[string][]string{}, mp["group"])
	})
	t.Run("incompatible new slice", func(t *testing.T) {
		mp := map[string]any{"group": map[string]string{}}
		assert.ErrMsgContains(t, maputil.SetByPath(&mp, "group.list[1]", nil), "map value")
		assert.Eq(t, map[string]string{}, mp["group"])
	})
	t.Run("incompatible intermediate map", func(t *testing.T) {
		mp := map[string]any{"group": map[string]string{}}
		assert.ErrMsgContains(t, maputil.SetByPath(&mp, "group.new.value", nil), "map value")
		assert.Eq(t, map[string]string{}, mp["group"])
	})
}

func TestMakeByPath_nilIndexedValue(t *testing.T) {
	assert.Eq(t, map[string]any{"list": []any{nil, nil}}, maputil.MakeByPath("list[1]", nil))
	assert.Eq(t, map[string]any{"group": map[string]any{"list": []any{nil, nil}}}, maputil.MakeByKeys([]string{"group", "list[1]"}, nil))
}

func TestData_SetByPath_nilValue(t *testing.T) {
	mp := maputil.Data{"group": map[string]any{"value": "old"}}
	assert.NoErr(t, mp.SetByPath("group.value", nil))
	value, ok := mp["group"].(map[string]any)["value"]
	assert.True(t, ok)
	assert.Nil(t, value)
}
