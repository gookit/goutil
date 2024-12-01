package mathutil_test

import (
	"encoding/json"
	"errors"
	"math"
	"testing"
	"time"

	"github.com/gookit/goutil/internal/comfunc"
	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestWithUserConvFn(t *testing.T) {
	is := assert.New(t)
	in := []int{1}

	// to int
	iv, err := mathutil.ToIntWith(in, mathutil.WithUserConvFn(func(v any) (int, error) {
		return 2, nil
	}))
	is.NoErr(err)
	is.Eq(2, iv)

	// to int64
	i64, err := mathutil.ToInt64With(in, mathutil.WithUserConvFn(func(v any) (int64, error) {
		return 2, nil
	}))
	is.NoErr(err)
	is.Eq(int64(2), i64)

	// to uint
	u, err := mathutil.ToUintWith(in, mathutil.WithUserConvFn(func(v any) (uint, error) {
		return 2, nil
	}))
	is.NoErr(err)
	is.Eq(uint(2), u)

	// to uint64
	u64, err := mathutil.ToUint64With(in, mathutil.WithUserConvFn(func(v any) (uint64, error) {
		return 2, nil
	}))
	is.NoErr(err)
	is.Eq(uint64(2), u64)

	// to float
	f, err := mathutil.ToFloatWith(in, mathutil.WithUserConvFn(func(v any) (float64, error) {
		return 2, nil
	}))
	is.NoErr(err)
	is.Eq(2.0, f)

	// to string
	s, err := mathutil.ToStringWith(in, comfunc.WithUserConvFn(func(v any) (string, error) {
		return "2", nil
	}))
	is.NoErr(err)
	is.Eq("2", s)
}

func TestWithNilAsFail(t *testing.T) {
	var err error
	is := assert.New(t)

	var iv int
	iv, err = mathutil.ToIntWith(nil)
	is.NoErr(err)
	is.Eq(0, iv)
	iv, err = mathutil.ToIntWith(nil, mathutil.WithNilAsFail[int])
	is.Err(err)
	is.Eq(0, iv)

	var i64 int64
	i64, err = mathutil.ToInt64With(nil)
	is.NoErr(err)
	is.Eq(int64(0), i64)
	i64, err = mathutil.ToInt64With(nil, mathutil.WithNilAsFail[int64])
	is.Err(err)
	is.Eq(int64(0), i64)

	// to uint
	_, err = mathutil.ToUint(nil)
	is.NoErr(err)
	_, err = mathutil.ToUintWith(nil, mathutil.WithNilAsFail[uint])
	is.Err(err)

	// to uint64
	_, err = mathutil.ToUint64(nil)
	is.NoErr(err)
	_, err = mathutil.ToUint64With(nil, mathutil.WithNilAsFail[uint64])
	is.Err(err)

	// to float
	_, err = mathutil.Float(nil)
	is.NoErr(err)
	_, err = mathutil.ToFloatWith(nil, mathutil.WithNilAsFail[float64])
	is.Err(err)
}

func TestWithHandlePtr(t *testing.T) {
	var err error
	is := assert.New(t)

	iv1 := 2
	i641 := int64(2)

	// int
	t.Run("to int", func(t *testing.T) {
		var iv int
		iv, err = mathutil.ToIntWith(&iv1)
		is.NoErr(err)
		is.Eq(2, iv)

		_, err = mathutil.ToIntWith(&i641)
		is.Err(err)
		iv, err = mathutil.ToIntWith(&i641, mathutil.WithHandlePtr[int])
		is.NoErr(err)
		is.Eq(2, iv)
	})

	// int64
	t.Run("to int64", func(t *testing.T) {
		var i64 int64
		i64, err = mathutil.ToInt64With(&i641)
		is.NoErr(err)
		is.Eq(int64(2), i64)

		_, err = mathutil.ToInt64With(&iv1)
		is.Err(err)
		i64, err = mathutil.ToInt64With(&iv1, mathutil.WithHandlePtr[int64])
		is.NoErr(err)
		is.Eq(int64(2), i64)
	})

	// uint
	t.Run("to uint", func(t *testing.T) {
		var u uint
		u1 := uint(2)
		u, err = mathutil.ToUintWith(&u1)
		is.NoErr(err)
		is.Eq(uint(2), u)

		_, err = mathutil.ToUintWith(&iv1)
		is.Err(err)
		u, err = mathutil.ToUintWith(&iv1, mathutil.WithHandlePtr[uint])
		is.NoErr(err)
		is.Eq(uint(2), u)
	})

	// uint64
	t.Run("to uint64", func(t *testing.T) {
		var u64 uint64
		u641 := uint64(2)
		u64, err = mathutil.ToUint64With(&u641)
		is.NoErr(err)
		is.Eq(uint64(2), u64)

		_, err = mathutil.ToUint64With(&iv1)
		is.Err(err)
		u64, err = mathutil.ToUint64With(&iv1, mathutil.WithHandlePtr[uint64])
		is.NoErr(err)
		is.Eq(uint64(2), u64)
	})

	// float
	t.Run("to float", func(t *testing.T) {
		var f float64
		f1 := float64(2)
		f, err = mathutil.ToFloatWith(&f1)
		is.NoErr(err)
		is.Eq(float64(2), f)

		_, err = mathutil.ToFloatWith(&iv1)
		is.Err(err)
		f, err = mathutil.ToFloatWith(&iv1, mathutil.WithHandlePtr[float64])
		is.NoErr(err)
		is.Eq(float64(2), f)
	})

	// string
	t.Run("to string", func(t *testing.T) {
		var s string
		s1 := "2"
		s, err = mathutil.ToStringWith(&s1)
		is.NoErr(err)
		is.Eq("2", s)

		_, err = mathutil.ToStringWith(&iv1)
		is.Err(err)
		s, err = mathutil.ToStringWith(&iv1, func(opt *comfunc.ConvOption) {
			opt.HandlePtr = true
		})
		is.NoErr(err)
		is.Eq("2", s)
	})
}

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

		_, err = mathutil.ToInt(nil)
		is.NoErr(err)

		intVal, err = mathutil.IntOrErr("-2")
		is.Nil(err)
		is.Eq(-2, intVal)

		is.Eq(0, mathutil.SafeInt(nil))
		is.Eq(2, mathutil.SafeInt("2.3"))
		is.Eq(-2, mathutil.MustInt("-2"))
		is.Eq(-2, mathutil.IntOrPanic("-2"))
		is.Eq(2, mathutil.IntOrDefault("invalid", 2))

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
		is.Eq(uint(2), uintVal)

		uintVal, err = mathutil.UintOrErr("2")
		is.Nil(err)
		is.Eq(uint(2), uintVal)

		_, err = mathutil.ToUint(nil)
		is.NoErr(err)
		_, err = mathutil.ToUint("-2")
		is.Err(err)

		is.Eq(uint(0), mathutil.QuietUint("-2"))
		for _, in := range tests {
			is.Eq(uint(2), mathutil.SafeUint(in))
		}
		for _, in := range errTests {
			is.Eq(uint(0), mathutil.QuietUint(in))
		}

		is.Eq(uint(0), mathutil.QuietUint(nil))
		is.Eq(uint(2), mathutil.MustUint("2"))
		is.Eq(uint(2), mathutil.UintOrDefault("invalid", 2))
		is.Panics(func() {
			mathutil.MustUint([]int{23})
		})
	})

	// To uint64
	t.Run("To uint64", func(t *testing.T) {
		uintVal, err := mathutil.Uint64("2")
		is.Nil(err)
		is.Eq(uint64(2), uintVal)

		uintVal, err = mathutil.Uint64OrErr("2")
		is.Nil(err)
		is.Eq(uint64(2), uintVal)

		_, err = mathutil.ToUint64(nil)
		is.NoErr(err)
		_, err = mathutil.ToUint64("-2")
		is.Err(err)

		is.Eq(uint64(0), mathutil.QuietUint64("-2"))
		for _, in := range tests {
			is.Eq(uint64(2), mathutil.SafeUint64(in))
		}
		for _, in := range errTests {
			is.Eq(uint64(0), mathutil.QuietUint64(in))
		}

		is.Eq(uint64(0), mathutil.QuietUint64(nil))
		is.Eq(uint64(2), mathutil.MustUint64("2"))
		is.Eq(uint64(2), mathutil.Uint64OrDefault("invalid", 2))
		is.Panics(func() {
			mathutil.MustUint64([]int{23})
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

		_, err = mathutil.ToInt64(nil)
		is.NoErr(err)

		for _, in := range tests {
			is.Eq(int64(2), mathutil.MustInt64(in))
		}
		for _, in := range errTests {
			is.Eq(int64(0), mathutil.QuietInt64(in))
			is.Eq(int64(0), mathutil.SafeInt64(in))
		}

		is.Eq(int64(2), mathutil.SafeInt64("2.3"))
		is.Eq(int64(0), mathutil.QuietInt64(nil))
		is.Eq(int64(2), mathutil.Int64OrDefault("invalid", 2))
		is.Panics(func() {
			mathutil.MustInt64([]int{23})
		})
	})
}

func TestStrInt(t *testing.T) {
	is := assert.New(t)

	is.Eq(2, mathutil.StrInt("2"))
	is.Eq(0, mathutil.StrInt("2a"))
	// StrIntOr
	is.Eq(2, mathutil.StrIntOr("2", 3))
	is.Eq(3, mathutil.StrIntOr("2a", 3))
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

	_, err = mathutil.ToString(nil)
	is.NoErr(err)

	is.Eq("", mathutil.SafeString(nil))
	is.Eq("[1]", mathutil.QuietString([]int{1}))
	is.Eq("23", mathutil.StringOrDefault([]int{1}, "23"))
	is.Eq("23", mathutil.StringOr("23", "2"))

	is.Panics(func() {
		mathutil.StringOrPanic([]int{23})
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

	// FloatOrDefault
	is.Eq(123.5, mathutil.FloatOrDefault("invalid", 123.5))
	is.Eq(123.1, mathutil.FloatOr(123.1, 123.5))

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

	// ToFloatWithFunc
	_, err = mathutil.ToFloatWith([]int{2}, mathutil.WithUserConvFn(func(v any) (float64, error) {
		return 0, errors.New("invalid")
	}))
	is.ErrMsg(err, "invalid")
}

func TestPercent(t *testing.T) {
	assert.Eq(t, float64(34), mathutil.Percent(34, 100))
	assert.Eq(t, float64(0), mathutil.Percent(34, 0))
	assert.Eq(t, float64(-100), mathutil.Percent(34, -34))
}
