package timex_test

import (
	"testing"
	"time"

	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/timex"
)

func TestUtil_basic(t *testing.T) {
	sec := timex.NowUnix()

	assert.NotEmpty(t, timex.FormatUnix(sec))
	assert.NotEmpty(t, timex.FormatUnixBy(sec, time.RFC3339))

	tt := timex.TodayStart()
	assert.Eq(t, "00:00:00", timex.DateFormat(tt, "H:I:S"))

	tt = timex.TodayEnd()
	assert.Eq(t, "23:59:59", timex.DateFormat(tt, "H:I:S"))

	tt = timex.NowHourStart()
	assert.Eq(t, "00:00", timex.DateFormat(tt, "I:S"))

	tt = timex.NowHourEnd()
	assert.Eq(t, "59:59", timex.DateFormat(tt, "I:S"))
}

func TestNowAddDay(t *testing.T) {
	now := timex.Now()
	tt := timex.NowAddDay(1)
	assert.True(t, tt.Unix() > now.Unix())

	tt = timex.NowAddHour(-3)
	assert.Neq(t, tt.Hour(), now.Hour())
	assert.Eq(t, "3 hours ago", timex.HowLongAgo(int64(now.DiffSec(tt)+2)))

	tt = timex.NowAddMinutes(5)
	assert.Neq(t, tt.Minute(), now.Minute())
}

func TestDateFormat(t *testing.T) {
	now := time.Now()
	assert.Eq(t, now.Format("2006-01-02 15:04:05"), timex.NowDate())

	tests := []struct{ layout, template string }{
		{"20060102 15:04:05", "Ymd H:I:S"},
		{"2006-01-02 15:04:05", "Y-m-d H:I:S"},
		{"2006-01-02 15:04", "Y-m-d H:I"},
		{"01/02 15:04:05", "m/d H:I:S"},
		{"06/01/02 15:04:05", "y/m/d H:I:S"},
		{"06/01/02 15:04:05.000", "y/m/d H:I:S.v"},
	}

	for i, item := range tests {
		date := now.Format(item.layout)
		assert.Eq(t, date, timex.DateFormat(now, item.template))
		if i%2 == 0 {
			assert.Eq(t, date, timex.Date(now, item.template))
		}
	}

	assert.Eq(t, now.Format("01/02 15:04:05.000000"), timex.Date(now, "m/d H:I:S.u"))
}

func TestFormatUnix(t *testing.T) {
	now := time.Now()
	want := now.Format("2006-01-02 15:04:05")

	assert.Eq(t, want, timex.FormatUnix(now.Unix()))
	assert.Eq(t, want, timex.FormatUnixBy(now.Unix(), timex.DefaultLayout))
	assert.Eq(t, want, timex.FormatUnixByTpl(now.Unix(), "Y-m-d H:I:S"))
	// dump.P(want)

	assert.Eq(t, want, timex.Format(now))
	assert.Eq(t, want, timex.FormatBy(now, timex.DefaultLayout))
}

func TestToLayout(t *testing.T) {
	assert.Eq(t, timex.DefaultLayout, timex.ToLayout(""))
	assert.Eq(t, time.RFC3339, timex.ToLayout("c"))
	assert.Eq(t, time.RFC3339, timex.ToLayout("Y-m-dTH:I:SP"))
}

func TestToDur(t *testing.T) {
	tests := []struct {
		in  string
		out time.Duration
		ok  bool
	}{
		{"", time.Duration(0), false},
		{"invalid", time.Duration(0), false},
		{"0", time.Duration(0), true},
		{"now", time.Duration(0), false},
		{"3s", 3 * timex.Second, true},
		{"3sec", 3 * timex.Second, true},
		{"3m", 3 * timex.Minute, true},
		{"3min", 3 * timex.Minute, true},
		{"3h", 3 * timex.Hour, true},
		{"3hours", 3 * timex.Hour, true},
		{"3d", 3 * timex.Day, true},
		{"3day", 3 * timex.Day, true},
		{"1w", 1 * timex.Week, true},
		{"1week", 1 * timex.Week, true},
	}

	for _, item := range tests {
		dur, err := timex.ToDur(item.in)
		if item.ok {
			assert.NoErr(t, err)
		} else {
			assert.Err(t, err)
		}

		assert.Eq(t, item.out, dur)
	}

	dur, err := timex.ToDur("invalid")
	assert.Err(t, err)
	assert.Eq(t, time.Duration(0), dur)

	dur, err = timex.ToDuration("0")
	assert.NoErr(t, err)
	assert.Eq(t, time.Duration(0), dur)
}

func TestIsDuration(t *testing.T) {
	assert.True(t, timex.IsDuration("3s"))
	assert.True(t, timex.IsDuration("3m"))
	assert.True(t, timex.IsDuration("-3h"))
	assert.True(t, timex.IsDuration("0"))

	assert.False(t, timex.IsDuration("3invalid"))
}

func TestInRange(t *testing.T) {
	tests := []struct {
		start, end string
		out        bool
	}{
		{"-5s", "5s", true},
		{"-5s", "", true},
		{"", "-2s", false},
		{"", "", false},
	}

	now := time.Now()

	for _, item := range tests {
		startT, err := timex.TryToTime(item.start, now)
		assert.NoErr(t, err)
		endT, err := timex.TryToTime(item.end, now)
		assert.NoErr(t, err)
		assert.Eq(t, item.out, timex.InRange(now, startT, endT))
	}
}
