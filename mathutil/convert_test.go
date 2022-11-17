package mathutil_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestToInt(t *testing.T) {
	is := assert.New(t)

	tests := []any{
		2,
		int8(2), int16(2), int32(2), int64(2),
		uint(2), uint8(2), uint16(2), uint32(2), uint64(2),
		float32(2.2), 2.3,
		"2",
		time.Duration(2),
		json.Number("2"),
	}
	errTests := []any{
		nil,
		"2a",
		[]int{1},
	}

	// To int
	intVal, err := mathutil.Int("2")
	is.Nil(err)
	is.Eq(2, intVal)

	intVal, err = mathutil.ToInt("-2")
	is.Nil(err)
	is.Eq(-2, intVal)

	is.Eq(2, mathutil.StrInt("2"))

	intVal, err = mathutil.IntOrErr("-2")
	is.Nil(err)
	is.Eq(-2, intVal)

	is.Eq(-2, mathutil.MustInt("-2"))
	for _, in := range tests {
		is.Eq(2, mathutil.MustInt(in))
		is.Eq(2, mathutil.QuietInt(in))
	}
	for _, in := range errTests {
		is.Eq(0, mathutil.MustInt(in))
	}

	// To uint
	uintVal, err := mathutil.Uint("2")
	is.Nil(err)
	is.Eq(uint64(2), uintVal)

	uintVal, err = mathutil.UintOrErr("2")
	is.Nil(err)
	is.Eq(uint64(2), uintVal)

	_, err = mathutil.ToUint("-2")
	is.Err(err)

	is.Eq(uint64(0), mathutil.MustUint("-2"))
	for _, in := range tests {
		is.Eq(uint64(2), mathutil.MustUint(in))
	}
	for _, in := range errTests {
		is.Eq(uint64(0), mathutil.QuietUint(in))
		is.Eq(uint64(0), mathutil.MustUint(in))
	}

	// To int64
	i64Val, err := mathutil.ToInt64("2")
	is.Nil(err)
	is.Eq(int64(2), i64Val)

	i64Val, err = mathutil.Int64("-2")
	is.Nil(err)
	is.Eq(int64(-2), i64Val)

	i64Val, err = mathutil.Int64OrErr("-2")
	is.Nil(err)
	is.Eq(int64(-2), i64Val)

	for _, in := range tests {
		is.Eq(int64(2), mathutil.MustInt64(in))
	}
	for _, in := range errTests {
		is.Eq(int64(0), mathutil.MustInt64(in))
		is.Eq(int64(0), mathutil.QuietInt64(in))
	}
}

func TestToString(t *testing.T) {
	is := assert.New(t)

	tests := []any{
		2,
		int8(2), int16(2), int32(2), int64(2),
		uint(2), uint8(2), uint16(2), uint32(2), uint64(2),
		float32(2), float64(2),
		// "2",
		time.Duration(2),
		json.Number("2"),
	}

	for _, in := range tests {
		is.Eq("2", mathutil.String(in))
		is.Eq("2", mathutil.QuietString(in))
		is.Eq("2", mathutil.MustString(in))
		val, err := mathutil.ToString(in)
		is.NoErr(err)
		is.Eq("2", val)
	}

	val, err := mathutil.StringOrErr(2)
	is.NoErr(err)
	is.Eq("2", val)

	val, err = mathutil.ToString(nil)
	is.NoErr(err)
	is.Eq("", val)

	is.Panics(func() {
		mathutil.MustString("2")
	})
}

func TestToFloat(t *testing.T) {
	is := assert.New(t)

	tests := []any{
		2,
		int8(2), int16(2), int32(2), int64(2),
		uint(2), uint8(2), uint16(2), uint32(2), uint64(2),
		float32(2), float64(2),
		"2",
		time.Duration(2),
		json.Number("2"),
	}
	for _, in := range tests {
		is.Eq(float64(2), mathutil.MustFloat(in))
	}

	is.Eq(123.5, mathutil.MustFloat("123.5"))
	is.Eq(123.5, mathutil.QuietFloat("123.5"))
	is.Eq(float64(0), mathutil.MustFloat("invalid"))
	is.Eq(float64(0), mathutil.QuietFloat("invalid"))

	fltVal, err := mathutil.ToFloat("123.5")
	is.Nil(err)
	is.Eq(123.5, fltVal)

	fltVal, err = mathutil.Float("-123.5")
	is.Nil(err)
	is.Eq(-123.5, fltVal)

	fltVal, err = mathutil.FloatOrErr("-123.5")
	is.Nil(err)
	is.Eq(-123.5, fltVal)
}
