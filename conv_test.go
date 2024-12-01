package goutil_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/gookit/goutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestToBool(t *testing.T) {
	is := assert.New(t)

	blVal, err := goutil.ToBool("1")
	is.Nil(err)
	is.True(blVal)

	blVal = goutil.Bool("1")
	is.True(blVal)

	is.False(goutil.Bool(false))
	is.False(goutil.Bool(1))
}

func TestToString(t *testing.T) {
	is := assert.New(t)

	str, err := goutil.ToString(23)
	is.Nil(err)
	is.Eq("23", str)

	str = goutil.String(23)
	is.Eq("23", str)
}

func TestToInt(t *testing.T) {
	is := assert.New(t)

	// To int
	iVal, err := goutil.ToInt("2")
	is.Nil(err)
	is.Eq(2, iVal)

	iVal = goutil.Int("-2")
	is.Nil(err)
	is.Eq(-2, iVal)

	// To int64
	i64Val, err := goutil.ToInt64("2")
	is.Nil(err)
	is.Eq(int64(2), i64Val)

	i64Val = goutil.Int64("-2")
	is.Nil(err)
	is.Eq(int64(-2), i64Val)

	// To uint
	uVal, err := goutil.ToUint("2")
	is.Nil(err)
	is.Eq(uint(2), uVal)

	uVal = goutil.Uint("2")
	is.Nil(err)
	is.Eq(uint(2), uVal)

	// To uint64
	u64Val, err := goutil.ToUint64("2")
	is.Nil(err)
	is.Eq(uint64(2), u64Val)

	u64Val = goutil.Uint64("2")
	is.Nil(err)
	is.Eq(uint64(2), u64Val)
}

func TestBaseTypeVal(t *testing.T) {
	is := assert.New(t)

	val, err := goutil.BaseTypeVal(uint64(23))
	is.NoErr(err)
	is.Eq(uint64(23), val)

	val, err = goutil.BaseTypeVal(23)
	is.NoErr(err)
	is.Eq(int64(23), val)

	val, err = goutil.BaseTypeVal(nil)
	is.Err(err)
	is.Nil(val)
}

func TestConvTo(t *testing.T) {
	is := assert.New(t)

	tests := []struct {
		val  any
		kind reflect.Kind
		out  any
		ok   bool
	}{
		// success
		{"23", reflect.Int, 23, true},
		{"23", reflect.Int8, int8(23), true},
		{"23", reflect.Uint8, uint8(23), true},
		{"23", reflect.Int16, int16(23), true},
		{"23", reflect.Uint16, uint16(23), true},
		{"23", reflect.Int32, int32(23), true},
		{"23", reflect.Uint32, uint32(23), true},
		{"23", reflect.Int64, int64(23), true},
		{"23", reflect.Uint64, uint64(23), true},
		{"23", reflect.Float64, float64(23), true},
		{"23", reflect.Float32, float32(23), true},
		{"23", reflect.String, "23", true},
		{"true", reflect.Bool, true, true},
		{nil, reflect.Int, 0, true},
		{nil, reflect.Uint, uint(0), true},
		// failed
		{"23", reflect.Bool, nil, false},
		{"abc", reflect.Float64, nil, false},
		{"abc", reflect.Float32, nil, false},
		{"abc", reflect.Int, nil, false},
		{"abc", reflect.Int8, nil, false},
		{"abc", reflect.Int16, nil, false},
		{"abc", reflect.Int32, nil, false},
		{"abc23", reflect.Int64, nil, false},
		{"abc", reflect.Uint, nil, false},
		{"abc", reflect.Uint8, nil, false},
		{"abc", reflect.Uint16, nil, false},
		{"abc", reflect.Uint32, nil, false},
		{"abc45", reflect.Uint64, nil, false},
		// overflow
		{uint64(math.MaxInt + 2), reflect.Int, nil, false},
		{uint64(math.MaxInt8 + 2), reflect.Int8, nil, false},
		{uint64(math.MaxUint8 + 2), reflect.Uint8, nil, false},
		{uint64(math.MaxInt16 + 2), reflect.Int16, nil, false},
		{uint64(math.MaxUint16 + 2), reflect.Uint16, nil, false},
		{uint64(math.MaxInt32 + 2), reflect.Int32, nil, false},
		{uint64(math.MaxUint32 + 2), reflect.Uint32, nil, false},
		{math.MaxFloat32 * 2.0, reflect.Float32, nil, false},
	}

	// test for goutil.ConvTo
	for _, item := range tests {
		val, err := goutil.ConvTo(item.val, item.kind)
		if item.ok {
			is.NoErr(err)
			is.Eq(item.out, val)
		} else {
			is.Err(err, "val: %v, kind: %v", item.val, item.kind)
		}
	}

	t.Run("extra func", func(t *testing.T) {
		// SafeConv
		is.Eq(23, goutil.SafeConv("23", reflect.Int))
		is.Eq(23, goutil.SafeKind("23", reflect.Int))
		is.Eq(nil, goutil.SafeKind("abc", reflect.Int))

		// ConvOrDefault
		is.Eq(23, goutil.ConvOrDefault("23", reflect.Int, 0))
		is.Eq(23, goutil.ConvOrDefault("abc", reflect.Int, 23))
	})

	// ToKind
	_, err := goutil.ToKind([]int{23}, reflect.Int, nil)
	is.Err(err)
}
