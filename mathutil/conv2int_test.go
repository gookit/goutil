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
		is.Eq(uint(2), mathutil.UintOrDefault("2", 2))
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
		is.Eq(uint64(2), mathutil.Uint64Or("2", 2))
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

func TestToIntXWith(t *testing.T) {
	t.Run("ToIntWith", func(t *testing.T) {
		_, err := mathutil.ToIntWith("2", mathutil.WithStrictMode[int])
		assert.Err(t, err)
		_, err = mathutil.ToIntWith([]string{"2b"}, mathutil.WithStrictMode[int])
		assert.Err(t, err)
	})

	t.Run("ToInt64With", func(t *testing.T) {
		_, err := mathutil.ToInt64With("2", mathutil.WithStrictMode[int64])
		assert.Err(t, err)
		_, err = mathutil.ToInt64With([]string{"2b"}, mathutil.WithStrictMode[int64])
		assert.Err(t, err)
	})
}

func TestToUintXWith(t *testing.T) {
	t.Run("ToUintWith", func(t *testing.T) {
		_, err := mathutil.ToUintWith("2", mathutil.WithStrictMode[uint])
		assert.Err(t, err)
		_, err = mathutil.ToUintWith([]string{"2b"}, mathutil.WithStrictMode[uint])
		assert.Err(t, err)
	})

	t.Run("ToUint64With", func(t *testing.T) {
		_, err := mathutil.ToUint64With("2", mathutil.WithStrictMode[uint64])
		assert.Err(t, err)
		_, err = mathutil.ToUint64With([]string{"2b"}, mathutil.WithStrictMode[uint64])
		assert.Err(t, err)
	})
}

func TestTryStrIntX(t *testing.T) {
	tests := []struct {
		in  string
		out int64
		err bool
	}{
		{in: "2", out: 2},
		{in: "", out: 0},
		{in: "2.3", out: 2},
		{in: "2a", err: true},
		{in: "2.3a", err: true},
		{in: "2.3.4", err: true},
		{in: "2.3.4.5", err: true},
	}

	t.Run("Int", func(t *testing.T) {
		for _, tt := range tests {
			v, err := mathutil.TryStrInt(tt.in)
			if tt.err {
				assert.Err(t, err)
			} else {
				assert.NoErr(t, err)
				assert.Eq(t, int(tt.out), v)
			}
		}
	})

	t.Run("Int64", func(t *testing.T) {
		for _, tt := range tests {
			v, err := mathutil.TryStrInt64(tt.in)
			if tt.err {
				assert.Err(t, err)
			} else {
				assert.NoErr(t, err)
				assert.Eq(t, tt.out, v)
			}
		}
	})

	// TryStrUint64
	t.Run("Uint64", func(t *testing.T) {
		for _, tt := range tests {
			v, err := mathutil.TryStrUint64(tt.in)
			if tt.err {
				assert.Err(t, err)
			} else {
				assert.NoErr(t, err)
				assert.Eq(t, uint64(tt.out), v)
			}
		}
	})
}
