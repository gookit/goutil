package mathutil_test

import (
	"encoding/json"
	"math"
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
		"2a",
		[]int{1},
	}

	overTests := []any{
		// case for overflow
		int64(math.MaxInt32 + 1),
		uint(math.MaxInt32 + 1),
		uint32(math.MaxInt32 + 1),
		uint64(math.MaxInt32 + 1),
		time.Duration(math.MaxInt32 + 1),
		json.Number("2147483648"),
	}

	// To int
	t.Run("To int", func(t *testing.T) {
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

		is.Eq(0, mathutil.SafeInt(nil))
		is.Eq(-2, mathutil.MustInt("-2"))
		is.Eq(-2, mathutil.IntOrPanic("-2"))
		for _, in := range tests {
			is.Eq(2, mathutil.MustInt(in))
			is.Eq(2, mathutil.QuietInt(in))
		}
		for _, in := range errTests {
			is.Eq(0, mathutil.SafeInt(in))
		}
		for _, in := range overTests {
			intVal, err = mathutil.ToInt(in)
			is.Err(err, "input: %v", in)
			is.Eq(0, intVal)
		}

		is.Panics(func() {
			mathutil.MustInt([]int{23})
		})
		is.Panics(func() {
			mathutil.IntOrPanic([]int{23})
		})
	})

	// To uint
	t.Run("To uint", func(t *testing.T) {
		uintVal, err := mathutil.Uint("2")
		is.Nil(err)
		is.Eq(uint64(2), uintVal)

		uintVal, err = mathutil.UintOrErr("2")
		is.Nil(err)
		is.Eq(uint64(2), uintVal)

		_, err = mathutil.ToUint("-2")
		is.Err(err)

		is.Eq(uint64(0), mathutil.QuietUint("-2"))
		for _, in := range tests {
			is.Eq(uint64(2), mathutil.SafeUint(in))
		}
		for _, in := range errTests {
			is.Eq(uint64(0), mathutil.QuietUint(in))
		}

		is.Eq(uint64(0), mathutil.QuietUint(nil))
		is.Eq(uint64(2), mathutil.MustUint("2"))
		is.Panics(func() {
			mathutil.MustUint([]int{23})
		})
	})

	// To int64
	t.Run("To int64", func(t *testing.T) {
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
			is.Eq(int64(0), mathutil.QuietInt64(in))
			is.Eq(int64(0), mathutil.SafeInt64(in))
		}

		is.Eq(int64(0), mathutil.QuietInt64(nil))
		is.Panics(func() {
			mathutil.MustInt64([]int{23})
		})
	})
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

	is.Eq("[1]", mathutil.QuietString([]int{1}))

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
	is.Eq(123.5, mathutil.FloatOrPanic("123.5"))
	is.Eq(123.5, mathutil.QuietFloat("123.5"))
	is.Eq(float64(0), mathutil.QuietFloat(nil))
	is.Eq(float64(0), mathutil.QuietFloat("invalid"))
	is.Eq(float64(0), mathutil.QuietFloat([]int{23}))

	is.Panics(func() {
		mathutil.MustFloat("invalid")
	})
	is.Panics(func() {
		mathutil.FloatOrPanic("invalid")
	})

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
