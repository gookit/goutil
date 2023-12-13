package timex_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/timex"
)

func TestElapsedNow(t *testing.T) {
	// hrs
	st := time.Now().Add(-204 * time.Minute)
	assert.Eq(t, "3.40hrs", timex.ElapsedNow(st))

	// min
	st = time.Now().Add(-184 * time.Second)
	assert.Eq(t, "3.07mins", timex.ElapsedNow(st))

	// s
	st = time.Now().Add(-1204 * time.Millisecond)
	assert.Eq(t, "1.204s", timex.ElapsedNow(st))

	// ms
	st = time.Now().Add(-204 * time.Millisecond)
	assert.Eq(t, "204.00ms", timex.ElapsedNow(st))

	// us
	st = time.Now().Add(-2304 * time.Nanosecond)
	assert.StrContains(t, timex.ElapsedNow(st), "2.")
}

func TestFromNow(t *testing.T) {
	lastIdx := len(timex.TimeMessages) - 1
	for i, tm := range timex.TimeMessages {
		sec := tm.Seconds[0]
		if i == lastIdx {
			sec *= 3
		} else if i > 2 {
			sec = sec - i*i
		}

		nt := time.Now().Add(-time.Duration(sec) * time.Second)
		s := timex.FromNow(nt)
		fmt.Println(s, "- from:", tm)
		assert.StrContains(t, s, strings.Replace(tm.Message, "%d", "", -1))
	}
}

func TestHowLongAgo(t *testing.T) {
	assert.Eq(t, "< 1 sec ago", timex.HowLongAgo(-23))
}

func TestTryToTime(t *testing.T) {
	tn := timex.Now()

	// duration string
	durTests := []struct {
		in  string
		out string
		ok  bool
	}{
		{"now", tn.Datetime(), true},
		{"0", tn.Datetime(), true},
		{"3s", tn.AddSeconds(3).Datetime(), true},
		{"3m", tn.AddMinutes(3).Datetime(), true},
	}

	for _, item := range durTests {
		tt, err := timex.TryToTime(item.in, tn.T())
		if item.ok {
			assert.NoErr(t, err)
		} else {
			assert.Err(t, err)
		}
		assert.Eq(t, item.out, timex.Format(tt))
	}

	bt := timex.ZeroTime
	assert.True(t, bt.IsZero())
	assert.Neq(t, 0, bt.Unix())

	noErrTests := []struct {
		in  string
		out string
	}{
		// date string
		{"2020-01-02 15:04:05", "2020-01-02 15:04:05"},
		{"2020-01-02", "2020-01-02 00:00:00"},
		{"2020-01-02 15:04", "2020-01-02 15:04:00"},
		{"2020-01-02 15", "2020-01-02 15:00:00"},
		{"2020-01-02 15:04:05.123", "2020-01-02 15:04:05"},
		{"2020-01-02 15:04:05.123456", "2020-01-02 15:04:05"},
		{"2020-01-02 15:04:05.123456789", "2020-01-02 15:04:05"},
		{"2020-01-02T15:04:05.123456789+08:00", "2020-01-02 15:04:05"},
	}

	for _, item := range noErrTests {
		tt, err := timex.TryToTime(item.in, bt)
		assert.NoErr(t, err)
		assert.Eq(t, item.out, timex.Format(tt))
	}
}

func TestParseRange(t *testing.T) {
	testutil.SetTimeLocalUTC()
	defer testutil.RestoreTimeLocal()

	tests := []struct {
		input string
		start int64
		end   int64
		ok    bool
	}{
		// date string
		{"2020-01-02 15:04:05", 1577977445, timex.ZeroUnix, true},
		{"2020-01-02 15:04:05~2020-01-03 15:04:05", 1577977445, 1578063845, true},
		{"2020-01-02 15:04:06~", 1577977446, timex.ZeroUnix, true},
		{"~2020-01-02 15:04:07", timex.ZeroUnix, 1577977447, true},
		// duration string
		{"-5s", 1672671840, timex.ZeroUnix, true},
		{"> 5s", 1672671850, timex.ZeroUnix, true},
		{"-5s~5s", 1672671840, 1672671850, true},
		{"~5s", timex.ZeroUnix, 1672671850, true},
		{"< 5s", timex.ZeroUnix, 1672671850, true},
		{"1h", 1672675445, timex.ZeroUnix, true},
		{"1hour", 1672675445, timex.ZeroUnix, true},
		// invalid
		{"~", timex.ZeroUnix, timex.ZeroUnix, false},
		{" ", timex.ZeroUnix, timex.ZeroUnix, false},
		{"invalid", timex.ZeroUnix, timex.ZeroUnix, false},
		{"2invalid", timex.ZeroUnix, timex.ZeroUnix, false},
		{"<= 2invalid", timex.ZeroUnix, timex.ZeroUnix, false},
	}

	bt, err := timex.FromDate("2023-01-02 15:04:05")
	assert.NoError(t, err)
	opt := &timex.ParseRangeOpt{
		BaseTime: bt.T(),
	}

	for _, item := range tests {
		start, end, err := timex.ParseRange(item.input, opt)
		assert.Eq(t, item.ok, err == nil, "err for %q", item.input)
		assert.Eq(t, item.start, start.Unix(), "start for %q", item.input)
		assert.Eq(t, item.end, end.Unix(), "end for %q", item.input)
	}

	t.Run("keyword now", func(t *testing.T) {
		start, end, err := timex.ParseRange("now", nil)
		assert.NoError(t, err)
		assert.Eq(t, timex.Now().Unix(), start.Unix())
		assert.Eq(t, timex.ZeroUnix, end.Unix())

		start, end, err = timex.ParseRange("~now", nil)
		assert.NoError(t, err)
		assert.Eq(t, timex.ZeroUnix, start.Unix())
		assert.Eq(t, timex.Now().Unix(), end.Unix())
	})

	t.Run("keyword today", func(t *testing.T) {
		now := timex.Now()
		start, end, err := timex.ParseRange("today", nil)
		assert.NoError(t, err)
		assert.Eq(t, now.DayStart().Unix(), start.Unix())
		assert.Eq(t, now.DayEnd().Unix(), end.Unix())

		start, end, err = timex.ParseRange("~today", nil)
		assert.Error(t, err)
		assert.Eq(t, timex.ZeroUnix, start.Unix())
		assert.Eq(t, timex.ZeroUnix, end.Unix())
	})

	t.Run("keyword yesterday", func(t *testing.T) {
		yd := timex.Now().DayAgo(1)
		start, end, err := timex.ParseRange("yesterday", nil)
		assert.NoError(t, err)
		assert.Eq(t, yd.DayStart().Unix(), start.Unix())
		assert.Eq(t, yd.DayEnd().Unix(), end.Unix())

		start, end, err = timex.ParseRange("~yesterday", nil)
		assert.Error(t, err)
		assert.Eq(t, timex.ZeroUnix, start.Unix())
		assert.Eq(t, timex.ZeroUnix, end.Unix())
	})

	t.Run("keyword tomorrow", func(t *testing.T) {
		td := timex.Now().DayAfter(1)
		start, end, err := timex.ParseRange("tomorrow", nil)
		assert.NoError(t, err)
		assert.Eq(t, td.DayStart().Unix(), start.Unix())
		assert.Eq(t, td.DayEnd().Unix(), end.Unix())

		start, end, err = timex.ParseRange("~tomorrow", nil)
		assert.Error(t, err)
		assert.Eq(t, timex.ZeroUnix, start.Unix())
		assert.Eq(t, timex.ZeroUnix, end.Unix())
	})

	t.Run("auto sort", func(t *testing.T) {
		opt := &timex.ParseRangeOpt{
			AutoSort: true,
		}
		start, end, err := timex.ParseRange("2020-01-02 15:04:05~2020-01-01 15:04:05", opt)
		assert.NoError(t, err)
		assert.Gt(t, end.Unix(), start.Unix())
	})
}
