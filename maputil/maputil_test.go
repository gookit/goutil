package maputil_test

import (
	"reflect"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestSimpleMerge(t *testing.T) {
	src := map[string]any{"A": "v0"}
	dst := map[string]any{"A": "v1", "B": "v2"}
	ret := maputil.SimpleMerge(src, dst)
	assert.Len(t, ret, 2)
	assert.Eq(t, "v0", ret["A"])

	dst = map[string]any{"A": "v1", "B": "v2"}
	ret = maputil.SimpleMerge(nil, dst)
	assert.Eq(t, "v1", ret["A"])

	ret = maputil.SimpleMerge(src, nil)
	assert.Eq(t, "v0", ret["A"])

	src = map[string]any{"A": "v0", "B": "v1", "sub": map[string]any{"s1": "sv0"}}
	dst = map[string]any{"A": "v1", "B": "v2", "sub": map[string]any{
		"s1": "sv1",
		"s2": "sv2",
	}}
	ret = maputil.SimpleMerge(src, dst)

	dm := maputil.Data(ret)
	assert.Eq(t, "v0", dm.Str("A"))
	assert.Eq(t, "v1", dm.Str("B"))
	assert.Eq(t, "sv0", dm.Str("sub.s1"))
	assert.Eq(t, "sv2", dm.Str("sub.s2"))
}

func TestMerge1level(t *testing.T) {
	ret := maputil.Merge1level(map[string]any{"A": "v0"}, map[string]any{"A": "v1", "B": "v2"})
	assert.Len(t, ret, 2)
	assert.Eq(t, "v1", ret["A"])

	ret = maputil.Merge1level(map[string]any{"A": "v0"}, nil)
	assert.Eq(t, "v0", ret["A"])

	ret = maputil.Merge1level(nil, map[string]any{"A": "v1", "B": "v2"})
	assert.Eq(t, "v1", ret["A"])
}

func TestMergeStringMap(t *testing.T) {
	ret := maputil.MergeSMap(map[string]string{"A": "v0"}, map[string]string{"A": "v1"}, false)
	assert.Eq(t, map[string]string{"A": "v0"}, ret)

	ret = maputil.MergeSMap(map[string]string{"A": "v0"}, map[string]string{"a": "v1"}, true)
	assert.Eq(t, map[string]string{"a": "v0"}, ret)

	ret = maputil.MergeSMap(map[string]string{"A": "v0"}, nil, false)
	assert.Eq(t, map[string]string{"A": "v0"}, ret)

	ret = maputil.MergeSMap(nil, map[string]string{"a": "v1"}, true)
	assert.Eq(t, map[string]string{"a": "v1"}, ret)

	ret = maputil.MergeMultiSMap(maputil.SMap{"a": "v1"}, maputil.SMap{"a": "v2"}, maputil.SMap{"b": "v3"})
	assert.ContainsKeys(t, ret, []string{"a", "b"})

	assert.Eq(t, map[string]string{"a": "v1"}, maputil.FilterSMap(map[string]string{"a": "v1", "b": ""}))
}

func TestMakeByPath(t *testing.T) {
	mp := maputil.MakeByPath("top.sub", "val")

	assert.NotEmpty(t, mp)
	assert.ContainsKey(t, mp, "top")
	assert.IsKind(t, reflect.Map, mp["top"])
	assert.Eq(t, "val", maputil.DeepGet(mp, "top.sub"))

	mp = maputil.MakeByPath("top.arr[1]", "val")
	dump.P(mp)
	assert.NotEmpty(t, mp)
	assert.ContainsKey(t, mp, "top")
	assert.Eq(t, "{top:map[arr:[ val]]}", maputil.ToString(mp))
	assert.Eq(t, []string{"", "val"}, maputil.DeepGet(mp, "top.arr"))
	assert.Eq(t, "val", maputil.DeepGet(mp, "top.arr.1"))
}

func TestMakeByKeys(t *testing.T) {
	mp := maputil.MakeByKeys([]string{"top"}, "val")
	assert.NotEmpty(t, mp)
	assert.ContainsKey(t, mp, "top")
	assert.Eq(t, "val", mp["top"])

	mp = maputil.MakeByKeys([]string{"top", "sub"}, "val")
	assert.NotEmpty(t, mp)
	assert.ContainsKey(t, mp, "top")
	assert.IsKind(t, reflect.Map, mp["top"])

	mp = maputil.MakeByKeys([]string{"top_arr[]"}, 234)
	// dump.P(mp)
	assert.NotEmpty(t, mp)
	assert.IsKind(t, reflect.Slice, mp["top_arr"])
	assert.Eq(t, 234, maputil.DeepGet(mp, "top_arr.0"))

	mp = maputil.MakeByKeys([]string{"top", "arr[1]"}, "val")
	dump.P(mp)
	assert.NotEmpty(t, mp)
	assert.ContainsKey(t, mp, "top")
	assert.Eq(t, "{top:map[arr:[ val]]}", maputil.ToString(mp))
	assert.Eq(t, []string{"", "val"}, maputil.DeepGet(mp, "top.arr"))
	assert.Eq(t, "val", maputil.DeepGet(mp, "top.arr.1"))
}
