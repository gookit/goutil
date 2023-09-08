package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/timex"
)

func TestToTime(t *testing.T) {
	is := assert.New(t)

	testutil.SetTimeLocalUTC()
	defer testutil.RestoreTimeLocal()

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
}

func TestToDuration(t *testing.T) {
	is := assert.New(t)

	dur, err1 := strutil.ToDuration("3s")
	is.NoErr(err1)
	is.Eq(3*timex.Second, dur)

	dur, err1 = strutil.ToDuration("3sec")
	is.NoErr(err1)
	is.Eq(3*timex.Second, dur)

	dur, err1 = strutil.ToDuration("-3sec")
	is.NoErr(err1)
	is.Eq(-3*timex.Second, dur)
}

func TestParseSizeRange(t *testing.T) {
	tests := []struct {
		expr string
		min  uint64
		max  uint64
		ok   bool
	}{
		// invalid
		{"", 0, 0, false},
		// limit min
		{"1", 1, 0, true},
		{"1b", 1, 0, true},
		{"1B", 1, 0, true},
		{"1k", 1024, 0, true},
		{"1KB", 1024, 0, true},
		{"1m", 1024 * 1024, 0, true},
		{"1m~", 1024 * 1024, 0, true},
		{"+1mb", 1024 * 1024, 0, true},
		{"> 1mb", 1024 * 1024, 0, true},
		{">= 1mb", 1024 * 1024, 0, true},
		// limit max
		{"-1M", 0, 1024 * 1024, true},
		{"< 1Mb", 0, 1024 * 1024, true},
		{"<=1Mb", 0, 1024 * 1024, true},
		{"~1Mb", 0, 1024 * 1024, true},
		{"0~1Mb", 0, 1024 * 1024, true},
		// limit range
		{"1kb~1m", 1024, 1024 * 1024, true},
		{"1kb~1mb", 1024, 1024 * 1024, true},
		{"1kb~1kb", 1024, 1024, true},
		// error case
		{"1kb~invalid", 1024, 0, false},
		{"invalid1~invalid2", 0, 0, false},
		{"invalid", 0, 0, false},
		{"1invalid", 0, 0, false},
	}

	is := assert.New(t)
	opt := &strutil.ParseSizeOpt{}
	for _, item := range tests {
		is.WithMsg(item.expr)
		min, max, err := strutil.ParseSizeRange(item.expr, opt)
		is.Equal(item.min, min, "for "+item.expr)
		is.Equal(item.max, max, "for "+item.expr)
		is.Equal(item.ok, err == nil, "for "+item.expr)
	}

	min, max, err := strutil.ParseSizeRange("1kb~1m", nil)
	is.Nil(err)
	is.Equal(uint64(1024), min)
	is.Equal(uint64(1024*1024), max)
}
