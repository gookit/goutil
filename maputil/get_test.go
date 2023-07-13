package maputil_test

import (
	"reflect"
	"testing"

	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/jsonutil"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestGetByPath(t *testing.T) {
	mp := map[string]any{
		"key0": "val0",
		"key1": map[string]string{"sk0": "sv0"},
		"key2": []string{"sv1", "sv2"},
		"key3": map[string]any{"sk1": "sv1"},
		"key4": []int{1, 2},
		"key5": []any{1, "2", true},
		"mlMp": []map[string]any{
			{
				"code":  "001",
				"names": []string{"John", "abc"},
			},
			{
				"code":  "002",
				"names": []string{"Tom", "def"},
			},
		},
	}

	tests := []struct {
		path string
		want any
		ok   bool
	}{
		{"key0", "val0", true},
		{"key1.sk0", "sv0", true},
		{"key3.sk1", "sv1", true},
		// not exists
		{"not-exits", nil, false},
		{"key2.not-exits", nil, false},
		{"not-exits.subkey", nil, false},
		// slices behaviour
		{"key2", mp["key2"], true},
		{"key2.0", "sv1", true},
		{"key2.1", "sv2", true},
		{"key4.0", 1, true},
		{"key4.1", 2, true},
		{"key5.0", 1, true},
		{"key5.1", "2", true},
		{"key5.2", true, true},
		// out of bound
		{"key4.3", nil, false},
		// deep sub map
		{"mlMp.*.code", []any{"001", "002"}, true},
		{"mlMp.*.names", []any{
			[]string{"John", "abc"},
			[]string{"Tom", "def"},
		}, true},
		{"mlMp.*.names.1", []any{"abc", "def"}, true},
	}

	for _, tt := range tests {
		v, ok := maputil.GetByPath(tt.path, mp)
		assert.Eq(t, tt.ok, ok, tt.path)
		assert.Eq(t, tt.want, v, tt.path)
	}

	// v, ok := maputil.GetByPath("mlMp.*.names.1", mp)
	// assert.True(t, ok)
	// assert.Eq(t, []any{"abc", "def"}, v)
}

var mlMp = map[string]any{
	"names": []string{"John", "Jane", "abc"},
	"coding": []map[string]any{
		{
			"details": map[string]any{
				"em": map[string]any{
					"code":              "001-1",
					"encounter_uid":     "1-1",
					"billing_provider":  "Test provider 01-1",
					"resident_provider": "Test Resident Provider-1",
				},
			},
		},
		{
			"details": map[string]any{
				"em": map[string]any{
					"code":              "001",
					"encounter_uid":     "1",
					"billing_provider":  "Test provider 01",
					"resident_provider": "Test Resident Provider",
				},
				"cpt": []map[string]any{
					{
						"code":             "001",
						"encounter_uid":    "2",
						"work_item_uid":    "3",
						"billing_provider": "Test provider 001",
						// "resident_provider": "Test Resident Provider",
					},
					{
						"code":              "OBS01",
						"encounter_uid":     "3",
						"work_item_uid":     "4",
						"billing_provider":  "Test provider OBS01",
						"resident_provider": "Test Resident Provider",
					},
					{
						"code":             "SU002",
						"encounter_uid":    "5",
						"work_item_uid":    "6",
						"billing_provider": "Test provider SU002",
						// "resident_provider": "Test Resident Provider",
					},
				},
			},
		},
	},
}

func TestGetByPath_deepPath(t *testing.T) {
	val, ok := maputil.GetByPath("coding.0.details.em.code", mlMp)
	assert.True(t, ok)
	assert.NotEmpty(t, val)

	val, ok = maputil.GetByPath("coding.*.details", mlMp)
	assert.True(t, ok)
	assert.NotEmpty(t, val)
	// dump.P(ok, val)

	val, ok = maputil.GetByPath("coding.*.details.em", mlMp)
	dump.P(ok, val)
	assert.True(t, ok)

	val, ok = maputil.GetByPath("coding.*.details.em.code", mlMp)
	dump.P(ok, val)
	assert.True(t, ok)
	assert.IsType(t, []any{}, val)

	val, ok = maputil.GetByPath("coding.*.details.cpt.*.encounter_uid", mlMp)
	dump.P(ok, val)
	assert.True(t, ok)
	assert.Len(t, val, 1)
	assert.IsType(t, []any{}, val)

	val, ok = maputil.GetByPath("coding.*.details.cpt.*.work_item_uid", mlMp)
	// dump.P(ok, val)
	assert.True(t, ok)
	assert.IsType(t, []any{}, val)

	val, ok = maputil.GetByPath("coding.*.details.cpt.*.resident_provider", mlMp)
	// dump.P(ok, val)
	assert.True(t, ok)
	assert.IsKind(t, reflect.Slice, val)

	val, ok = maputil.GetByPath("coding.*.details.cpt.*.not-exists", mlMp)
	assert.Nil(t, val)
	assert.False(t, ok)
}

func TestKeys(t *testing.T) {
	mp := map[string]any{
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
	mp := map[string]any{
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

func TestEachAnyMap(t *testing.T) {
	mp := map[string]any{
		"key0": "v0",
		"key1": "v1",
		"key2": 34,
	}

	maputil.EachAnyMap(mp, func(k string, v any) {
		assert.NotEmpty(t, k)
		assert.NotEmpty(t, v)
	})

	assert.Panics(t, func() {
		maputil.EachAnyMap(1, nil)
	})
}

func TestGetByPathKeys(t *testing.T) {
	val, ok := maputil.GetByPathKeys(map[string]any{}, nil)
	assert.True(t, ok)
	assert.Empty(t, val)

	t.Run("sub string-map", func(t *testing.T) {
		mp := map[string]any{
			"top": map[string]string{"key": "value"},
		}

		val, ok := maputil.GetByPathKeys(mp, []string{"top", "key"})
		assert.True(t, ok)
		assert.Eq(t, "value", val)
	})
	t.Run("sub any-map", func(t *testing.T) {
		mp := map[string]any{
			"top": map[any]any{"key": "value"},
		}

		val, ok := maputil.GetByPathKeys(mp, []string{"top", "key"})
		assert.True(t, ok)
		assert.Eq(t, "value", val)
	})

	t.Run("sub []map[string]any", func(t *testing.T) {
		mp := map[string]any{
			"top": []map[string]any{
				{"key": "value"},
				{"key": "value1"},
			},
		}

		val, ok := maputil.GetByPathKeys(mp, []string{"top", "1"})
		assert.True(t, ok)
		assert.NotEmpty(t, val)
		assert.IsKind(t, reflect.Map, val)
		val, ok = maputil.GetByPathKeys(mp, []string{"top", "10"})
		assert.False(t, ok)
		assert.Nil(t, val)
		val, ok = maputil.GetByPathKeys(mp, []string{"top", "invalid"})
		assert.False(t, ok)
		assert.Nil(t, val)

		val, ok = maputil.GetByPathKeys(mp, []string{"top", "*"})
		assert.True(t, ok)
		assert.IsKind(t, reflect.Slice, val)
		assert.NotEmpty(t, val)
		val, ok = maputil.GetByPathKeys(mp, []string{"top", "*", "key"})
		assert.True(t, ok)
		assert.IsKind(t, reflect.Slice, val)
		assert.Len(t, val, 2)
	})
}

// https://github.com/gookit/goutil/issues/109
func TestIssues_109(t *testing.T) {
	mp := make(map[string]any)
	err := jsonutil.DecodeString(`{
  "success": true,
  "result": {
    "total": 2,
    "records": [
      {
        "id": "59fab0fa-8f0a-4065-8863-1dae40166015"
      },
      {
        "id": "7c1bd7f9-2ef4-44c8-9756-2e85156ca58f"
      }
    ]
  }
}`, &mp)
	assert.NoErr(t, err)
	dump.P(mp)

	ids, ok := maputil.GetByPath("result.records.*.id", mp)
	dump.P(ids, arrutil.AnyToStrings(ids))
	assert.True(t, ok)
	assert.Len(t, ids, 2)
}
