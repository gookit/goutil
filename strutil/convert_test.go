package strutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/timex"
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

func TestAnyToString(t *testing.T) {
	is := assert.New(t)

	tests := []any{
		2,
		int8(2), int16(2), int32(2), int64(2),
		uint(2), uint8(2), uint16(2), uint32(2), uint64(2),
		"2",
		[]byte("2"),
	}
	for _, in := range tests {
		is.Eq("2", strutil.MustString(in))
		is.Eq("2", strutil.QuietString(in))
	}

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

	str, err = strutil.String(nil)
	is.NoErr(err)
	is.Eq("", str)

	_, err = strutil.String([]string{"a"})
	is.Err(err)

	str, err = strutil.AnyToString([]string{"a"}, false)
	is.NoErr(err)
	is.Eq("[a]", str)
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

	iVal = strutil.QuietInt("-23")
	is.Eq(-23, iVal)

	iVal = strutil.Int2("-23")
	is.Eq(-23, iVal)

	iVal = strutil.IntOrPanic("-23")
	is.Eq(-23, iVal)

	iVal = strutil.MustInt("-23")
	is.Eq(-23, iVal)

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

	iVal = strutil.QuietInt64("-23")
	is.Eq(int64(-23), iVal)

	iVal = strutil.MustInt64("-23")
	is.Eq(int64(-23), iVal)

	is.Panics(func() {
		strutil.MustInt64("abc")
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

func TestToTime(t *testing.T) {
	is := assert.New(t)
	tests := map[string]string{
		"20180927":             "2018-09-27 00:00:00 +0000 UTC",
		"2018-09-27":           "2018-09-27 00:00:00 +0000 UTC",
		"2018-09-27 12":        "2018-09-27 12:00:00 +0000 UTC",
		"2018-09-27T12":        "2018-09-27 12:00:00 +0000 UTC",
		"2018-09-27 12:34":     "2018-09-27 12:34:00 +0000 UTC",
		"2018-09-27T12:34":     "2018-09-27 12:34:00 +0000 UTC",
		"2018-09-27 12:34:45":  "2018-09-27 12:34:45 +0000 UTC",
		"2018-09-27T12:34:45":  "2018-09-27 12:34:45 +0000 UTC",
		"2018/09/27 12:34:45":  "2018-09-27 12:34:45 +0000 UTC",
		"2018/09/27T12:34:45Z": "2018-09-27 12:34:45 +0000 UTC",
		"2018-10-16 12:34:01":  "2018-10-16 12:34:01 +0000 UTC",
	}

	for sample, want := range tests {
		tm, err := strutil.ToTime(sample)
		is.Nil(err, "sample %s => want %s", sample, want)
		is.Eq(want, tm.String())
	}

	tm, err := strutil.ToTime("invalid")
	is.Err(err)
	is.True(tm.IsZero())

	tm, err = strutil.ToTime("invalid", "")
	is.Err(err)
	is.True(tm.IsZero())

	tm, err = strutil.ToTime("2018-09-27T15:34", "2018-09-27 15:34:23")
	is.Err(err)
	is.True(tm.IsZero())

	tm = strutil.MustToTime("2018-09-27T15:34")
	is.Eq("2018-09-27T15:34", timex.FormatByTpl(tm, "Y-m-dTH:I"))

	is.Panics(func() {
		strutil.MustToTime("invalid")
	})

	dur, err1 := strutil.ToDuration("3s")
	is.NoErr(err1)
	is.Eq(3*timex.Second, dur)
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
