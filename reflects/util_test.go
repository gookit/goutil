package reflects_test

import (
	"reflect"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/reflects"
	"github.com/gookit/goutil/testutil/assert"
)

func TestLen(t *testing.T) {
	is := assert.New(t)
	tests := []any{
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
		val  any
		want reflect.Kind
	}{
		{"invalid", reflect.Invalid},
		{[2]int{1, 2}, reflect.Int},
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
		{[]any{"a", "b"}, reflect.Interface},
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

	ty := reflect.TypeOf(sl)
	rv := reflect.ValueOf(sl)
	assert.False(t, rv.CanAddr())

	assert.Eq(t, reflect.Slice, ty.Kind())
	assert.Eq(t, reflect.String, ty.Elem().Kind())

	rv = reflect.Append(rv, reflect.New(ty.Elem()).Elem())
	// rv.Set(reflect.Append(rv, reflect.New(ty.Elem()).Elem()))
	dump.P(sl, rv.CanAddr(), rv.Interface())

	assert.Eq(t, rv.Len(), 2)
}

func TestSliceAddItem_ok(t *testing.T) {
	sl := []string{"abc"}
	assert.Len(t, sl, 1)

	ty := reflect.TypeOf(sl)
	rv := reflect.ValueOf(&sl)
	assert.False(t, rv.CanAddr())

	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	assert.True(t, rv.CanAddr())
	assert.Eq(t, reflect.Slice, ty.Kind())
	assert.Eq(t, reflect.String, ty.Elem().Kind())

	// rv = reflect.Append(rv, reflect.New(ty.Elem()).Elem())
	rv.Set(reflect.Append(rv, reflect.New(ty.Elem()).Elem()))
	assert.Len(t, sl, 2)

	rv.Index(1).Set(reflect.ValueOf("def"))
	dump.P(sl)
}

func TestSlice_subMap_addItem(t *testing.T) {
	d := []any{
		map[string]string{
			"k3051": "v3051",
		},
	}

	rv := reflect.ValueOf(d)

	vv := rv.Index(0).Elem()
	dump.P(rv.CanAddr(), vv.CanAddr())
	vv.SetMapIndex(reflect.ValueOf("newKey"), reflect.ValueOf("newVal"))
	dump.P(d)
}

func TestMap_subSlice_addItem(t *testing.T) {
	mp := map[string]any{
		"sl": []string{"abc"},
	}

	// ty := reflect.TypeOf(mp)
	rv := reflect.ValueOf(mp)
	dump.P(rv.CanAddr())

	rv.SetMapIndex(reflect.ValueOf("k2"), reflect.ValueOf("v2"))
	dump.P(mp)

	slk := reflect.ValueOf("sl")
	srv := rv.MapIndex(slk)
	if srv.Kind() == reflect.Interface {
		srv = srv.Elem()
	}
	sty := srv.Type()

	dump.P(srv.CanAddr())
	assert.Eq(t, reflect.Slice, sty.Kind())
	assert.Eq(t, reflect.String, sty.Elem().Kind())

	msl := reflect.MakeSlice(sty, 0, srv.Cap())
	dump.P(msl.CanAddr())

	srv = reflect.Append(srv, reflect.New(sty.Elem()).Elem())
	// srv.Set(srv)
	srv.Index(1).Set(reflect.ValueOf("def"))

	rv.SetMapIndex(slk, srv)

	// ret.Set(rv)
	dump.P(mp, srv.Interface())
}

func TestSetValue(t *testing.T) {
	// string
	str := "val"
	rv := reflect.ValueOf(&str)
	err := reflects.SetValue(rv, "new val")
	assert.NoErr(t, err)
	assert.Eq(t, "new val", str)

	// int
	iVal := 234
	rv = reflect.ValueOf(&iVal)
	err = reflects.SetValue(rv, "345")
	assert.NoErr(t, err)
	assert.Eq(t, 345, iVal)

	// panic: reflect: reflect.Value.Set using unaddressable value
	assert.Panics(t, func() {
		rv := reflect.ValueOf("val")
		_ = reflects.SetValue(rv, "new val")
	})

	// test for SetRValue()
	rv = reflect.ValueOf(&iVal)
	reflects.SetRValue(rv, reflect.ValueOf(456))
	assert.Eq(t, 456, iVal)
}

func TestSetValue_map(t *testing.T) {
	// map
	mp := map[string]string{}
	set := map[string]string{"key": "val"}

	rv := reflect.ValueOf(&mp)
	err := reflects.SetValue(rv, set)
	assert.NoErr(t, err)
	assert.Eq(t, set, mp)

	// type error
	err = reflects.SetValue(rv, map[int]string{2: "abc"})
	assert.Err(t, err)
}
