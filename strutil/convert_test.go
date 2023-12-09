package strutil_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestStringJoin(t *testing.T) {
	assert.Eq(t, "a:b", strutil.Join(":", "a", "b"))
	assert.Eq(t, "a:b", strutil.Implode(":", "a", "b"))
	assert.Eq(t, "a:b", strutil.JoinList(":", []string{"a", "b"}))

	assert.Eq(t, "ab:23", strutil.JoinAny(":", "ab", 23))
}

func TestStringToBool(t *testing.T) {
	is := assert.New(t)

	tests1 := map[string]bool{
		"1":     true,
		"on":    true,
		"yes":   true,
		"true":  true,
		"false": false,
		"off":   false,
		"no":    false,
		"0":     false,
	}

	for str, want := range tests1 {
		is.Eq(want, strutil.QuietBool(str))
		is.Eq(want, strutil.MustBool(str))
	}

	is.Panics(func() {
		strutil.MustBool("invalid")
	})

	blVal, err := strutil.ToBool("1")
	is.Nil(err)
	is.True(blVal)

	blVal, err = strutil.Bool("10")
	is.Err(err)
	is.False(blVal)
}

func BenchmarkAnyToString_int(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = strutil.AnyToString(3, false)
	}
}

func BenchmarkAnyToString_float(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = strutil.AnyToString(3.4, false)
		// _ = strconv.FormatFloat(3.4, 'f', -1, 64)
	}
}

func BenchmarkAnyToString_string(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = strutil.AnyToString("string", false)
	}
}

func TestToString(t *testing.T) {
	is := assert.New(t)

	tests := []any{
		2,
		int8(2), int16(2), int32(2), int64(2),
		uint(2), uint8(2), uint16(2), uint32(2), uint64(2),
		"2",
		[]byte("2"),
		time.Duration(2),
		json.Number("2"),
	}
	for _, in := range tests {
		is.Eq("2", strutil.QuietString(in))
	}

	is.Eq("2", strutil.SafeString(2))
	is.Eq("err msg", strutil.SafeString(errors.New("err msg")))

	tests1 := []any{
		float32(2.3), 2.3,
	}
	for _, in := range tests1 {
		is.Eq("2.3", strutil.MustString(in))
	}

	str, err := strutil.String(2.3)
	is.NoErr(err)
	is.Eq("2.3", str)

	str, err = strutil.String(true)
	is.NoErr(err)
	is.Eq("true", str)

	str, err = strutil.ToString(true)
	is.NoErr(err)
	is.Eq("true", str)

	str, err = strutil.StringOrErr(true)
	is.NoErr(err)
	is.Eq("true", str)

	_, err = strutil.String(nil)
	is.NoErr(err)

	_, err = strutil.String([]string{"a"})
	is.Err(err)

	str, err = strutil.AnyToString([]string{"a"}, false)
	is.NoErr(err)
	is.Eq("[a]", str)

	str = strutil.QuietString(nil)
	is.Eq("", str)

	str = strutil.StringOrDefault([]string{"a"}, "default")
	is.Eq("default", str)
	str = strutil.StringOrDefault("value", "default")
	is.Eq("value", str)

	is.Panics(func() {
		strutil.StringOrPanic([]string{"a"})
	})
}

func TestByte2string(t *testing.T) {
	is := assert.New(t)

	s := "abc"
	is.Eq(s, strutil.Byte2str([]byte(s)))
	is.Eq(s, strutil.Byte2string([]byte(s)))
	// is.Same(s, strutil.Byte2str([]byte(s)))
	// is.NotSame(s, string([]byte(s)))

	is.Eq([]byte(s), strutil.ToBytes(s))
}

func TestStrToInt(t *testing.T) {
	is := assert.New(t)

	iVal, err := strutil.Int("23")
	is.Nil(err)
	is.Eq(23, iVal)

	iVal, err = strutil.ToInt("-23")
	is.Nil(err)
	is.Eq(-23, iVal)

	is.Eq(-23, strutil.QuietInt("-23"))
	is.Eq(-23, strutil.SafeInt("-23"))

	is.Eq(23, strutil.IntOrDefault("invalid", 23))
	is.Eq(23, strutil.IntOrDefault("23", 25))

	is.Eq(-23, strutil.IntOrPanic("-23"))
	is.Eq(-23, strutil.MustInt("-23"))

	is.PanicsErrMsg(func() {
		strutil.IntOrPanic("abc")
	}, "strconv.Atoi: parsing \"abc\": invalid syntax")
}

func TestStrToInt64(t *testing.T) {
	is := assert.New(t)

	iVal, err := strutil.ToInt64("-23")
	is.Nil(err)
	is.Eq(int64(-23), iVal)

	iVal, err = strutil.Int64OrErr("-23")
	is.Nil(err)
	is.Eq(int64(-23), iVal)

	iVal = strutil.Int64("23")
	is.Nil(err)
	is.Eq(int64(23), iVal)

	is.Eq(int64(23), strutil.Int64OrDefault("invalid", 23))
	is.Eq(int64(23), strutil.Int64OrDefault("23", 25))

	is.Eq(int64(-23), strutil.QuietInt64("-23"))
	is.Eq(int64(-23), strutil.MustInt64("-23"))

	is.Panics(func() {
		strutil.MustInt64("abc")
	})
}

func TestStrToUint(t *testing.T) {
	is := assert.New(t)

	iVal, err := strutil.ToUint("23")
	is.Nil(err)
	is.Eq(uint64(23), iVal)

	iVal, err = strutil.UintOrErr("23")
	is.Nil(err)
	is.Eq(uint64(23), iVal)

	iVal = strutil.Uint("23")
	is.Nil(err)
	is.Eq(uint64(23), iVal)

	is.Eq(uint64(23), strutil.UintOrDefault("invalid", 23))
	is.Eq(uint64(23), strutil.UintOrDefault("23", 25))

	is.Eq(uint64(23), strutil.SafeUint("23"))
	is.Eq(uint64(23), strutil.MustUint("23"))

	is.Panics(func() {
		strutil.MustUint("abc")
	})
}

func TestStrToInts(t *testing.T) {
	is := assert.New(t)

	ints, err := strutil.ToInts("a,b,c")
	is.Err(err)
	is.Len(ints, 0)

	ints = strutil.Ints("a,b,c")
	is.Len(ints, 0)

	ints, err = strutil.ToIntSlice("1,2,3")
	is.Nil(err)
	is.Eq([]int{1, 2, 3}, ints)

	ints = strutil.Ints("1,2,3")
	is.Eq([]int{1, 2, 3}, ints)
}

func TestStr2Array(t *testing.T) {
	is := assert.New(t)

	ss := strutil.Strings("a,b,c", ",")
	is.Len(ss, 3)
	is.Eq(`[]string{"a", "b", "c"}`, fmt.Sprintf("%#v", ss))

	tests := []string{
		// sample
		"a,b,c",
		"a,b,c,",
		",a,b,c",
		"a, b,c",
		"a,,b,c",
		"a, , b,c",
	}

	for _, sample := range tests {
		ss = strutil.ToArray(sample)
		is.Eq(`[]string{"a", "b", "c"}`, fmt.Sprintf("%#v", ss))
	}

	ss = strutil.ToSlice("", ",")
	is.Len(ss, 0)

	ss = strutil.ToStrings(", , ", ",")
	is.Len(ss, 0)
}

//	func TestToOSArgs(t *testing.T) {
//		args := strutil.ToOSArgs(`./app top sub -a ddd --xx "abc
//
// def ghi"`)
//
//		assert.Len(t, args, 7)
//		assert.Eq(t, "abc\ndef ghi", args[6])
//	}

func TestQuote(t *testing.T) {
	is := assert.New(t)

	is.Eq(`"\"it's ok\""`, strutil.Quote(`"it's ok"`))

	is.Eq("", strutil.Unquote("''"))
	is.Eq("a", strutil.Unquote("a"))
	is.Eq("a single-quoted string", strutil.Unquote("'a single-quoted string'"))
	is.Eq("a double-quoted string", strutil.Unquote(`"a double-quoted string"`))
}

func TestToByteSize(t *testing.T) {
	u64 := uint64(0)
	assert.Eq(t, u64, strutil.SafeByteSize("0"))
	assert.Eq(t, u64, strutil.SafeByteSize("0b"))
	assert.Eq(t, u64, strutil.SafeByteSize("0B"))
	assert.Eq(t, u64, strutil.SafeByteSize("0M"))

	tests := []struct {
		bytes uint64
		sizeS string
	}{
		{1, "1"},
		{5, "5"},
		{346, "346"},
		{346, "346B"},
		{3471, "3.39K"},
		{346777, "338.65 Kb"},
		{12341739, "11.77 M"},
		{1202590842, "1.12GB"},
		{1231453023109, "1.12 TB"},
		{1351079888211148, "1.2PB"},
	}

	for _, tt := range tests {
		assert.Eq(t, tt.bytes, strutil.SafeByteSize(tt.sizeS))
	}

	assert.Eq(t, uint64(1), strutil.SafeByteSize("1"))
	assert.Eq(t, uint64(1024*1024), strutil.SafeByteSize("1M"))
	assert.Eq(t, uint64(1024*1024), strutil.SafeByteSize("1MB"))
	assert.Eq(t, uint64(1024*1024), strutil.SafeByteSize("1m"))
	assert.Eq(t, uint64(10485760), strutil.SafeByteSize("10mb"))

	assert.Eq(t, uint64(1024*1024*1024), strutil.SafeByteSize("1G"))
	assert.Eq(t, uint64(1288490188), strutil.SafeByteSize("1.2GB"))
	assert.Eq(t, uint64(1288490188), strutil.SafeByteSize("1.2 GB"))

	size, err := strutil.ToByteSize("invalid")
	assert.Err(t, err)
	assert.Eq(t, uint64(0), size)
}
