package reflects_test

import (
	"reflect"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/reflects"
	"github.com/gookit/goutil/testutil/assert"
)

func TestIsFunc(t *testing.T) {
	assert.True(t, reflects.IsFunc(reflects.HasChild))
}

func TestIsEmpty(t *testing.T) {
	is := assert.New(t)

	is.True(reflects.IsEmpty(reflect.ValueOf(nil)))
	is.True(reflects.IsEmpty(reflect.ValueOf("")))

	type T struct{ v interface{} }
	rv := reflect.ValueOf(T{}).Field(0)
	is.True(reflects.IsEmpty(rv))
}

func TestLen(t *testing.T) {
	is := assert.New(t)
	tests := []interface{}{
		"abc",
		123,
		int8(123), int16(123), int32(123), int64(123),
		uint8(123), uint16(123), uint32(123), uint64(123),
		float32(123), float64(123),
		[]int{1, 2, 3}, []string{"a", "b", "c"},
		map[string]string{"k0": "v0", "k1": "v1", "k2": "v2"},
	}

	for _, sample := range tests {
		is.Eq(3, reflects.Len(reflect.ValueOf(sample)))
	}

	ptrArr := &[]string{"a", "b"}
	is.Eq(2, reflects.Len(reflect.ValueOf(ptrArr)))

	is.Eq(-1, reflects.Len(reflect.ValueOf(nil)))
}

func TestSliceSubKind(t *testing.T) {
	noErrTests := []struct {
		val  interface{}
		want reflect.Kind
	}{
		{"invalid", reflect.Invalid},
		{[]int{1, 2}, reflect.Int},
		{[]int8{1, 2}, reflect.Int8},
		{[]int16{1, 2}, reflect.Int16},
		{[]int32{1, 2}, reflect.Int32},
		{[]int64{1, 2}, reflect.Int64},
		{[]uint{1, 2}, reflect.Uint},
		{[]uint8{1, 2}, reflect.Uint8},
		{[]uint16{1, 2}, reflect.Uint16},
		{[]uint32{1, 2}, reflect.Uint32},
		{[]uint64{1, 2}, reflect.Uint64},
		{[]string{"a", "b"}, reflect.String},
		{[]interface{}{"a", "b"}, reflect.Interface},
	}

	for _, item := range noErrTests {
		eleType := reflects.SliceSubKind(reflect.TypeOf(item.val))
		assert.Eq(t, item.want, eleType)
	}
}

func TestSliceItemType(t *testing.T) {
	sl := []string{"abc"}
	ty := reflect.TypeOf(sl)

	assert.Eq(t, reflect.Slice, ty.Kind())
	assert.Eq(t, reflect.String, ty.Elem().Kind())
}

func TestSliceAddItem_fail1(t *testing.T) {
	sl := []string{"abc"}

	rv := reflect.ValueOf(sl)
	ty := reflect.TypeOf(sl)

	assert.Eq(t, reflect.Slice, ty.Kind())
	assert.Eq(t, reflect.String, ty.Elem().Kind())

	// rv = reflect.Append(rv, reflect.New(ty.Elem()).Elem())
	rv.Set(reflect.Append(rv, reflect.New(ty.Elem()).Elem()))

	dump.P(sl, rv.CanAddr(), rv.Interface())
}

func TestSliceAddItem_fail2(t *testing.T) {
	sl := []string{"abc"}

	ty := reflect.TypeOf(sl)
	rv := reflect.ValueOf(&sl)
	ret := rv
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	assert.Eq(t, reflect.Slice, ty.Kind())
	assert.Eq(t, reflect.String, ty.Elem().Kind())

	rv = reflect.Append(rv, reflect.New(ty.Elem()).Elem())

	dump.P(sl, rv.CanAddr(), rv.Interface(), ret.CanAddr())
}

func TestSliceAddItem_3(t *testing.T) {
	var sl interface{}
	sl = []string{"abc"}

	ty := reflect.TypeOf(sl)
	rv := reflect.ValueOf(sl)
	ret := rv
	dump.P(ret.CanAddr())
	// if rv.Kind() == reflect.Ptr {
	// 	rv = rv.Elem()
	// }

	assert.Eq(t, reflect.Slice, ty.Kind())
	assert.Eq(t, reflect.String, ty.Elem().Kind())

	rv = reflect.Append(rv, reflect.New(ty.Elem()).Elem())

	// ret.Set(rv)
	dump.P(sl, rv.CanAddr(), rv.Interface())
}
