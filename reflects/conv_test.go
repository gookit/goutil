package reflects_test

import (
	"math"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/gookit/goutil/reflects"
	"github.com/gookit/goutil/testutil/assert"
)

func TestBaseTypeVal(t *testing.T) {
	tests := []struct {
		want, give any
	}{
		{int64(23), 23},
		{uint64(23), uint(23)},
		{23.4, 23.4},
		// {23.4, float32(23.4)},
		{"abc", "abc"},
	}
	for _, e := range tests {
		val, err := reflects.BaseTypeVal(reflect.ValueOf(e.give))
		assert.NoErr(t, err)
		assert.Eq(t, e.want, val)
	}

	// val = float64(23.399999618530273)
	val, err := reflects.BaseTypeVal(reflect.ValueOf(float32(23.4)))
	assert.NoErr(t, err)
	assert.NotEmpty(t, val)

	val, err = reflects.BaseTypeVal(reflect.ValueOf([]int{23}))
	assert.Err(t, err)
	assert.Nil(t, val)
}

func TestValueByType(t *testing.T) {
	val, err := reflects.ValueByType(true, reflect.TypeOf(false))
	assert.NoErr(t, err)
	assert.True(t, val.Bool())

	val, err = reflects.ValueByType(123, reflect.TypeOf("s"))
	assert.NoErr(t, err)
	assert.Eq(t, "123", val.Interface())

	val, err = reflects.ValueByType("123", reflect.TypeOf(1))
	assert.NoErr(t, err)
	assert.Eq(t, 123, val.Interface())

	// same type
	val, err = reflects.ValueByType("abc", reflect.TypeOf("s"))
	assert.NoErr(t, err)
	assert.Eq(t, "abc", val.Interface())

	// invalid val
	_, err = reflects.ConvToType(nil, reflect.TypeOf(1))
	assert.Err(t, err)
}

func TestValueByType_slice(t *testing.T) {
	val, err := reflects.ValueByType([]int{12, 23}, reflect.TypeOf([]int{}))
	assert.NoErr(t, err)
	assert.Eq(t, []int{12, 23}, val.Interface())

	arr := []string{"val0", "val1"}
	val, err = reflects.ValueByType(arr, reflect.TypeOf([]string{}))
	assert.NoErr(t, err)
	assert.Eq(t, arr, val.Interface())

	// auto conv elem type
	val, err = reflects.ValueByType([]string{"12", "23"}, reflect.TypeOf([]int{}))
	assert.NoErr(t, err)
	assert.Eq(t, []int{12, 23}, val.Interface())

	val, err = reflects.ValueByType([]string{"ab", "cd"}, reflect.TypeOf([]int{}))
	assert.Err(t, err)
	assert.False(t, val.IsValid())
}

func TestValueByType_map(t *testing.T) {
	mp := map[string]string{"key": "val"}
	val, err := reflects.ValueByType(mp, reflect.TypeOf(map[string]string{}))
	assert.NoErr(t, err)
	assert.Eq(t, mp, val.Interface())

	mp = map[string]string{"key": "val"}
	val, err = reflects.ValueByType(mp, reflect.TypeOf(map[int]string{}))
	assert.Err(t, err)
	assert.False(t, val.IsValid())
}

func TestConvSlice(t *testing.T) {
	oldArr := []string{"ab", "cd"}
	newArr, err := reflects.ConvSlice(reflect.ValueOf(oldArr), reflect.TypeOf("s"))
	assert.NoErr(t, err)
	assert.Eq(t, oldArr, newArr.Interface())

	// conv fail
	oldArr = []string{"ab", "cd"}
	newArr, err = reflects.ConvSlice(reflect.ValueOf(oldArr), reflect.TypeOf(1))
	assert.Err(t, err)
	assert.False(t, newArr.IsValid())

	// conv to ints
	oldArr = []string{"12", "23"}
	newArr, err = reflects.ConvSlice(reflect.ValueOf(oldArr), reflect.TypeOf(1))
	assert.NoErr(t, err)
	assert.Eq(t, []int{12, 23}, newArr.Interface())

	assert.Panics(t, func() {
		_, _ = reflects.ConvSlice(reflect.ValueOf("s"), reflect.TypeOf("s"))
	})
}

func TestValueByKind(t *testing.T) {
	tests := []struct {
		want, give any
		// want kind
		kind reflect.Kind
	}{
		{23, "23", reflect.Int},
		{int8(23), 23, reflect.Int8},
		{int16(23), 23, reflect.Int16},
		{int32(23), 23, reflect.Int32},
		{int64(23), 23, reflect.Int64},
		{"23", 23, reflect.String},
		{uint(23), 23, reflect.Uint},
		{uint8(23), 23, reflect.Uint8},
		{uint16(23), 23, reflect.Uint16},
		{uint32(23), 23, reflect.Uint32},
		{uint64(23), 23, reflect.Uint64},
		{float32(23), 23, reflect.Float32},
		{float64(23), 23, reflect.Float64},
	}
	for _, e := range tests {
		val, err := reflects.ValueByKind(e.give, e.kind)
		assert.NoErr(t, err)
		assert.Eq(t, e.want, val.Interface())
	}

	errTests := []struct {
		give any
		kind reflect.Kind
	}{
		{"abc", reflect.Int},
		{true, reflect.Int},
		{23, reflect.Bool},
		// case for overflow
		{143, reflect.Int8},
		{math.MaxInt16 + 1, reflect.Int16},
		{343, reflect.Uint8},
		{int64(math.MaxInt32 + 1), reflect.Int32},
		{int64(math.MaxUint16 + 1), reflect.Uint16},
		{int64(math.MaxUint32 + 1), reflect.Uint32},
	}

	for _, e := range errTests {
		val, err := reflects.ValueByKind(e.give, e.kind)
		assert.Err(t, err, "give: %v, kind: %v", e.give, e.kind)
		assert.False(t, val.IsValid())
	}

	val, err := reflects.ValueByKind("abc", reflect.Int)
	assert.Err(t, err)
	assert.False(t, val.IsValid())

	val, err = reflects.ValueByKind("true", reflect.Bool)
	assert.NoErr(t, err)
	assert.True(t, val.Bool())

	val, err = reflects.ValueByKind(reflect.ValueOf("true"), reflect.Bool)
	assert.NoErr(t, err)
	assert.True(t, val.Bool())
}

func TestToString(t *testing.T) {
	tests := []struct {
		give any
		want string
	}{
		{nil, ""},
		{true, "true"},
		{23, "23"},
		{int8(23), "23"},
		{int16(23), "23"},
		{int32(23), "23"},
		{int64(23), "23"},
		{"23", "23"},
		{uint(23), "23"},
		{uint8(23), "23"},
		{uint16(23), "23"},
		{uint32(23), "23"},
		{uint64(23), "23"},
		{float32(23), "23"},
		{float64(23), "23"},
		{[]int{12, 34}, "[12 34]"},
	}
	for _, e := range tests {
		rv := reflect.ValueOf(e.give)
		assert.Eq(t, e.want, reflects.String(rv))
	}

	rv := reflect.ValueOf([]int{12, 34})
	s, err := reflects.ToString(rv)
	assert.Err(t, err)
	assert.Eq(t, "", s)

	rv = reflect.Value{}
	assert.Eq(t, "", reflects.String(rv))
}

func TestToTimeOrDuration(t *testing.T) {
	tests := []struct {
		str      string
		typ      reflect.Type
		hasErr   bool
		expected any
	}{
		{"3s", reflect.TypeOf(time.Duration(0)), false, 3 * time.Second},
		{"2023-01-01T00:00:00Z", reflect.TypeOf(time.Time{}), false, time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)},
		{"1h2m3s", reflect.TypeOf(time.Duration(0)), false, 3723 * time.Second},
		{"a message", reflect.TypeOf(34), false, "a message"},
		{"2023-01-02T00:00:00Z", reflect.TypeOf(""), false, "2023-01-02T00:00:00Z"},
		{"invalid Time", reflect.TypeOf(time.Time{}), true, nil},
		{"invalid Duration", reflect.TypeOf(time.Duration(0)), true, nil},
	}

	for _, test := range tests {
		result, err := reflects.ToTimeOrDuration(test.str, test.typ)
		if test.hasErr {
			assert.Err(t, err)
		} else {
			assert.NoErr(t, err)
		}
		assert.Eq(t, test.expected, result, "case: "+test.str)
	}

	longStr := strings.Repeat("a long message", 10)
	result, err := reflects.ToTimeOrDuration(longStr, reflect.TypeOf(""))
	assert.NoErr(t, err)
	assert.Eq(t, longStr, result)
}
