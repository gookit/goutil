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

	ret = maputil.MergeStrMap(map[string]string{"A": "v0"}, nil)
	assert.Eq(t, map[string]string{"A": "v0"}, ret)

	ret = maputil.AppendSMap(maputil.SM{"A": "v0"}, maputil.SM{"B": "v2"})
	assert.ContainsKeys(t, ret, []string{"A", "B"})

	ret = maputil.MergeSMap(nil, map[string]string{"a": "v1"}, true)
	assert.Eq(t, map[string]string{"a": "v1"}, ret)

	ret = maputil.MergeMultiSMap(maputil.SMap{"a": "v1"}, maputil.SMap{"a": "v2"}, maputil.SMap{"b": "v3"})
	assert.ContainsKeys(t, ret, []string{"a", "b"})

	assert.Eq(t, map[string]string{"a": "v1"}, maputil.FilterSMap(map[string]string{"a": "v1", "b": ""}))
}

func TestMergeL2StrMap(t *testing.T) {
	ret := maputil.MergeL2StrMap(map[string]map[string]string{
		"a": {"a1": "v1", "a2": "v1"},
		"c": {"c1": "c1"},
		"d": {"d2": "v2"},
	}, maputil.L2StrMap{
		"a": {"a2": "v2"},
		"b": {"b1": "v2"},
		"d": {"d2": "v3", "d3": "v4"},
	})

	assert.NotEmpty(t, ret)
	assert.Len(t, ret, 4)
	assert.Equal(t, "v1", ret["a"]["a1"])
	assert.Equal(t, "v2", ret["a"]["a2"])

	// as maputil.L2StrMap test
	r2 := maputil.L2StrMap(ret)
	assert.True(t, r2.Exists("a.a1"))
	assert.False(t, r2.Exists("a.not-exist"))
	assert.Eq(t, "v1", r2.Get("a.a1"))
	assert.NotEmpty(t, r2.StrMap("a"))
	assert.Eq(t, "v2", r2.StrMap("a").Get("a2"))

	r2.Load(map[string]map[string]string{
		"a": {"a1": "v3"},
		"e": {"e1": "v1"},
	})
	assert.Eq(t, "v3", r2.Get("a.a1"))
	assert.Eq(t, "v1", r2.Get("e.e1"))
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

func TestCopy(t *testing.T) {
	dst := map[string]int{"a": 1, "b": 2}
	src := map[string]int{"b": 3, "c": 4}

	maputil.Copy(dst, src)

	assert.Len(t, dst, 3)
	assert.Eq(t, 1, dst["a"])
	assert.Eq(t, 3, dst["b"])
	assert.Eq(t, 4, dst["c"])

	dst2 := map[string]string{"x": "v1"}
	src2 := map[string]string{"y": "v2"}

	maputil.Copy(dst2, src2)

	assert.Len(t, dst2, 2)
	assert.Eq(t, "v1", dst2["x"])
	assert.Eq(t, "v2", dst2["y"])

	dst3 := map[int]string{1: "one"}
	src3 := map[int]string{2: "two", 1: "ONE"}

	maputil.Copy(dst3, src3)

	assert.Len(t, dst3, 2)
	assert.Eq(t, "ONE", dst3[1])
	assert.Eq(t, "two", dst3[2])
}

func TestDeleteFunc(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	maputil.DeleteFunc(m, func(k string, v int) bool {
		return v%2 == 0
	})

	assert.Len(t, m, 2)
	assert.ContainsKey(t, m, "a")
	assert.ContainsKey(t, m, "c")
	assert.NotContainsKey(t, m, "b")
	assert.NotContainsKey(t, m, "d")

	m2 := map[int]string{1: "one", 2: "two", 3: "three"}

	maputil.DeleteFunc(m2, func(k int, v string) bool {
		return k == 2
	})

	assert.Len(t, m2, 2)
	assert.ContainsKey(t, m2, 1)
	assert.ContainsKey(t, m2, 3)
	assert.NotContainsKey(t, m2, 2)

	m3 := map[string]string{"key1": "val1", "key2": "val2"}

	maputil.DeleteFunc(m3, func(k string, v string) bool {
		return k == "key1"
	})

	assert.Len(t, m3, 1)
	assert.NotContainsKey(t, m3, "key1")
	assert.ContainsKey(t, m3, "key2")

	m4 := map[string]int{"a": 1, "b": 2}

	maputil.DeleteFunc(m4, func(k string, v int) bool {
		return false
	})

	assert.Len(t, m4, 2)
}
