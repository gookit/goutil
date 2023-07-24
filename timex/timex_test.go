package timex_test

import (
	"testing"
	"time"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/jsonutil"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/timex"
)

func TestWrap(t *testing.T) {
	tx := timex.Wrap(time.Now())
	assert.False(t, tx.IsZero())

	tx = timex.FromTime(time.Now())
	assert.False(t, tx.IsZero())

	tx = timex.Local()
	assert.False(t, tx.IsZero())
	assert.False(t, tx.T().IsZero())

	tx = timex.FromUnix(time.Now().Unix())
	assert.False(t, tx.IsZero())
}

func TestFromDate(t *testing.T) {
	tx, err := timex.FromDate("2022-04-20 19:40:34")
	assert.NoErr(t, err)
	assert.Eq(t, "2022-04-20 19:40:34", tx.Datetime())

	tx, err = timex.FromDate("2022-04-20 19:40:34", "Y-m-d H:I:S")
	assert.NoErr(t, err)
	assert.Eq(t, "2022-04-20 19:40:34", tx.Datetime())
	assert.Eq(t, tx.Unix(), tx.Timestamp())

	_, err = timex.FromDate("invalid")
	assert.Err(t, err)
}

func TestTimeX_basic(t *testing.T) {
	tx := timex.Now()
	assert.NotEmpty(t, tx.String())
	assert.NotEmpty(t, tx.Datetime())
}

func TestTimeX_Format(t *testing.T) {
	tx := timex.Now()
	assert.Eq(t, tx.Datetime(), tx.DateFormat("Y-m-d H:I:S"))
}

func TestTimeX_SubUnix(t *testing.T) {
	tx := timex.Now()

	after1m := tx.AddMinutes(1)

	assert.Eq(t, timex.OneMinSec, after1m.SubUnix(tx.Time))
}

func TestTimeX_DateFormat(t *testing.T) {
	tx := timex.Now()
	assert.Eq(t, tx.Format(timex.DefaultLayout), tx.DateFormat("Y-m-d H:I:S"))
	assert.Eq(t, tx.Format(""), tx.DateFormat("Y-m-d H:I:S"))
	assert.Eq(t, tx.Format("2006/01/02 15:04"), tx.TplFormat("Y/m/d H:I"))

	date := tx.Format("06/01/02 15:04")
	dump.V(date)
	assert.Eq(t, date, tx.DateFormat("y/m/d H:I"))

	assert.Eq(t, "23:59:59", tx.DayEnd().DateFormat("H:I:S"))
	assert.Eq(t, "00:00:00", tx.DayStart().DateFormat("H:I:S"))
}

func TestTimeX_AddDay(t *testing.T) {
	tx := timex.Now()

	yd := tx.Yesterday()
	yd1 := tx.AddDay(-1)
	assert.Eq(t, yd.Unix(), yd1.Unix())
	assert.Eq(t, yd.Unix(), tx.DayAgo(1).Unix())
	assert.Eq(t, yd.Unix(), tx.SubDay(1).Unix())

	assert.Eq(t, "1 day", tx.HowLongAgo(yd.Time))

	assert.True(t, tx.IsAfter(yd.Time))
	assert.True(t, tx.IsAfterUnix(yd.Time.Unix()))
	assert.True(t, yd.IsBefore(tx.Time))
	assert.True(t, yd.IsBeforeUnix(tx.T().Unix()))

	assert.Eq(t, tx.Unix()-yd.Unix(), int64(timex.OneDaySec))

	md := tx.Tomorrow()
	yd2 := tx.DayAfter(1)
	assert.Eq(t, md.Unix(), yd2.Unix())
}

func TestTimeX_AddSeconds(t *testing.T) {
	tx := timex.Now()

	h1 := tx.AddHour(1)
	assert.Eq(t, h1.Unix(), tx.AddString("1h").Unix())

	s1 := tx.AddSeconds(timex.OneHourSec)
	assert.Eq(t, h1.Unix(), s1.Unix())

	assert.Eq(t, timex.OneHour, h1.Diff(tx.Time))
	assert.Eq(t, timex.OneHourSec, h1.DiffSec(tx.Time))
	assert.Eq(t, timex.OneHourSec, h1.DiffUnix(tx.Unix()))

	t2 := s1.SubHour(1)
	assert.Eq(t, tx.Unix(), t2.Unix())
	assert.Eq(t, tx.Unix(), s1.SubMinutes(60).Unix())
}

func TestTimeX_HourStart(t *testing.T) {
	tx := timex.Now()
	hs := tx.HourStart()
	he := tx.HourEnd()

	assert.Eq(t, "00:00", hs.DateFormat("I:S"))
	assert.Eq(t, "59:59", he.DateFormat("I:S"))
}

func TestTimeX_CustomHMS(t *testing.T) {
	tx := timex.Now()
	assert.Eq(t, "12:23:34", tx.CustomHMS(12, 23, 34).TplFormat("H:I:S"))
}

// https://github.com/gookit/goutil/issues/60
func TestTimeX_UnmarshalJSON(t *testing.T) {
	type User struct {
		Time timex.Time `json:"time"`
	}

	req := &User{}
	err := jsonutil.DecodeString(`{
    "time": "2018-10-16 12:34:01"
}`, req)
	assert.NoErr(t, err)
	assert.Eq(t, "2018-10-16 12:34", req.Time.TplFormat("Y-m-d H:i"))

	// UnmarshalText
	tx := &timex.Time{}
	err = tx.UnmarshalText([]byte("2018-10-16 12:34:01"))
	assert.NoErr(t, err)
	assert.Eq(t, "2018-10-16 12:34", tx.TplFormat("Y-m-d H:i"))

	// error
	err = tx.UnmarshalText([]byte("invalid"))
	assert.Err(t, err)

	assert.Nil(t, tx.UnmarshalJSON([]byte("null")))
}
