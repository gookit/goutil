package maputil_test

import (
	"reflect"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestData_usage(t *testing.T) {
	mp := maputil.Data{
		"k1":    23,
		"k2":    "ab",
		"k3":    "true",
		"k4":    false,
		"k5":    map[string]string{"a": "b"},
		"anyMp": map[string]any{"b": 23},
		"k6":    "23,45",
		"k7":    []string{"ab", "cd"},
	}

	assert.True(t, mp.Has("k1"))
	assert.True(t, mp.Bool("k3"))
	assert.False(t, mp.Bool("k4"))
	assert.False(t, mp.IsEmpty())
	assert.Eq(t, 23, mp.Get("k1"))
	assert.Eq(t, "b", mp.Get("k5.a"))
	assert.Eq(t, 23, mp.Get("anyMp.b"))

	// int
	assert.Eq(t, 23, mp.Int("k1"))
	assert.Eq(t, int64(23), mp.Int64("k1"))

	// str
	assert.Eq(t, "23", mp.Str("k1"))
	assert.Eq(t, "ab", mp.Str("k2"))

	// Strings
	assert.Eq(t, []string{"ab", "cd"}, mp.Strings("k7"))
	assert.Nil(t, mp.Strings("k1"))
	assert.Nil(t, mp.Strings("notExists"))

	// StringsByStr
	assert.Eq(t, []string{"23", "45"}, mp.StringsByStr("k6"))

	// set
	mp.Set("new", "val1")
	assert.Eq(t, "val1", mp.Str("new"))

	val, ok := mp.Value("new")
	assert.True(t, ok)
	assert.Eq(t, "val1", val)

	// not exists
	assert.False(t, mp.Bool("notExists"))
	assert.Eq(t, 0, mp.Int("notExists"))
	assert.Eq(t, int64(0), mp.Int64("notExists"))
	assert.Eq(t, "", mp.Str("notExists"))

	// default
	assert.Eq(t, 23, mp.Default("k1", 10))
	assert.Eq(t, 10, mp.Default("notExists", 10))

	assert.Nil(t, mp.StringMap("k1"))
	assert.Nil(t, mp.StringMap("notExists"))
	assert.Eq(t, map[string]string{"a": "b"}, mp.StrMap("k5"))
	assert.Eq(t, map[string]string{"b": "23"}, mp.StringMap("anyMp"))

	assert.NotEmpty(t, mp.String())
}

func TestData_SetByPath(t *testing.T) {
	mp := maputil.Data{
		"k2": "ab",
		"k5": map[string]any{"a": "v0"},
	}
	assert.Nil(t, mp.Get("k5.b"))
	assert.Len(t, mp.Keys(), 2)
	assert.NotEmpty(t, mp.ToStringMap())

	err := mp.SetByPath("k5.b", "v2")
	assert.NoErr(t, err)
	// dump.P(mp)
	assert.Eq(t, "v2", mp.Get("k5.b"))

	mp.Load(map[string]any{"k2": "val2", "k3": "val3"})
	assert.Eq(t, "val2", mp.Str("k2"))
	assert.Eq(t, "val3", mp.Str("k3"))

	// sub
	assert.Nil(t, mp.Sub("not-exists"))
	sub := mp.Sub("k5")
	assert.Eq(t, "v0", sub.Get("a"))
	assert.Eq(t, "v2", sub.Get("b"))

	// LoadSMap
	mp.LoadSMap(map[string]string{"uk2": "val2", "uk3": "val3"})
	assert.Eq(t, "val2", mp.Str("uk2"))
}

func TestData_SetByPath_case2(t *testing.T) {
	mp := maputil.Data{}
	assert.Eq(t, 0, len(mp))

	err := mp.SetByPath("top2.inline.list.ids", []int{234, 345, 456})
	assert.NoErr(t, err)
	assert.Eq(t, []int{234, 345, 456}, mp.Get("top2.inline.list.ids"))

	err = mp.SetByPath("top2.sub.var-refer", "val1")
	assert.NoErr(t, err)
	assert.Eq(t, "val1", mp.Get("top2.sub.var-refer"))

	err = mp.SetByPath("top2.sub.key2-other", "val2")
	assert.NoErr(t, err)
	assert.Eq(t, "val2", mp.Get("top2.sub.key2-other"))
	// dump.P(mp)
}

func TestData_SetByPath_case3(t *testing.T) {
	mp := maputil.Data{}
	assert.Eq(t, 0, len(mp))

	err := mp.SetByPath("top.sub.key3", "false")
	assert.NoErr(t, err)
	assert.Eq(t, "false", mp.Get("top.sub.key3"))
	assert.False(t, mp.Bool("top.sub.key3"))

	err = mp.SetByPath("top.sub.key4[0]", "abc")
	assert.NoErr(t, err)

	err = mp.SetByPath("top.sub.key4[1]", "def")
	assert.NoErr(t, err)
	sli := mp.Get("top.sub.key4")
	assert.IsKind(t, reflect.Slice, sli)
	assert.Len(t, sli, 2)
	// dump.P(mp, sli)
}

// top.sub.key5[0].f1 = ab
// top.sub.key5[1].f2 = de
func TestData_SetByPath_case4(t *testing.T) {
	mp := maputil.Data{}
	assert.Eq(t, 0, len(mp))

	err := mp.SetByPath("top.sub.key3", "false")
	assert.NoErr(t, err)
	assert.Eq(t, "false", mp.Get("top.sub.key3"))

	err = mp.SetByPath("top.sub.key5[0].f1", "val1")
	assert.NoErr(t, err)
	// dump.P(mp)

	err = mp.SetByPath("top.sub.key5[1].f2", "val2")
	assert.NoErr(t, err)
	dump.P(mp)
	sli := mp.Get("top.sub.key5")
	assert.IsKind(t, reflect.Slice, sli)
	assert.Len(t, sli, 2)
}

func TestData_SetByKeys_emptyData(t *testing.T) {
	// one level
	mp := make(maputil.Data)
	err := mp.SetByKeys([]string{"k3"}, "v3")
	assert.NoErr(t, err)
	dump.P(mp)

	assert.Eq(t, "v3", mp.Str("k3"))

	// two level
	mp1 := make(maputil.Data)
	err = mp1.SetByKeys([]string{"k5", "b"}, "v2")
	assert.NoErr(t, err)
	dump.P(mp1)

	assert.Eq(t, "v2", mp1.Get("k5.b"))
}

func TestData_SetByKeys(t *testing.T) {
	mp := maputil.Data{
		"k2": "ab",
		"k5": map[string]any{"a": "v0"},
	}
	assert.Nil(t, mp.Get("k3"))
	assert.Nil(t, mp.Get("k5.b"))

	assert.NoErr(t, mp.SetByKeys([]string{}, "v3"))

	err := mp.SetByKeys([]string{"k3"}, "v3")
	assert.NoErr(t, err)
	assert.Eq(t, "v3", mp.Str("k3"))

	err = mp.SetByKeys([]string{"k5", "b"}, "v2")
	assert.NoErr(t, err)

	// dump.P(mp)
	assert.Eq(t, "v2", mp.Get("k5.b"))
}
