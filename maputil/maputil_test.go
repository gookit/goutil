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
}

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
