package timex_test

import (
	"testing"
	"time"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/timex"
	"github.com/stretchr/testify/assert"
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
	assert.NoError(t, err)
	assert.Equal(t, "2022-04-20 19:40:34", tx.Datetime())

	tx, err = timex.FromDate("2022-04-20 19:40:34", "Y-m-d H:I:S")
	assert.NoError(t, err)
	assert.Equal(t, "2022-04-20 19:40:34", tx.Datetime())
}

func TestTimeX_basic(t *testing.T) {
	tx := timex.Now()
	assert.NotEmpty(t, tx.String())
	assert.NotEmpty(t, tx.Datetime())
}

func TestTimeX_Format(t *testing.T) {
	tx := timex.Now()
	assert.Equal(t, tx.Datetime(), tx.DateFormat("Y-m-d H:I:S"))
}

func TestTimeX_SubUnix(t *testing.T) {
	tx := timex.Now()

	after1m := tx.AddMinutes(1)

	assert.Equal(t, timex.OneMinSec, after1m.SubUnix(tx.Time))
}

func TestTimeX_DateFormat(t *testing.T) {
	tx := timex.Now()
	assert.Equal(t, tx.Format(timex.DefaultLayout), tx.DateFormat("Y-m-d H:I:S"))
	assert.Equal(t, tx.Format(""), tx.DateFormat("Y-m-d H:I:S"))
	assert.Equal(t, tx.Format("2006/01/02 15:04"), tx.TplFormat("Y/m/d H:I"))

	date := tx.Format("06/01/02 15:04")
	dump.V(date)
	assert.Equal(t, date, tx.DateFormat("y/m/d H:I"))

	assert.Equal(t, "23:59:59", tx.DayEnd().DateFormat("H:I:S"))
	assert.Equal(t, "00:00:00", tx.DayStart().DateFormat("H:I:S"))
}

func TestTimeX_AddDay(t *testing.T) {
	tx := timex.Now()

	yd := tx.Yesterday()
	yd1 := tx.AddDay(-1)
	assert.Equal(t, yd.Unix(), yd1.Unix())
	assert.Equal(t, yd.Unix(), tx.DayAgo(1).Unix())

	assert.True(t, tx.IsAfter(yd.Time))
	assert.True(t, tx.IsAfterUnix(yd.Time.Unix()))
	assert.True(t, yd.IsBefore(tx.Time))
	assert.True(t, yd.IsBeforeUnix(tx.T().Unix()))

	assert.Equal(t, tx.Unix()-yd.Unix(), int64(timex.OneDaySec))

	md := tx.Tomorrow()
	yd2 := tx.DayAfter(1)
	assert.Equal(t, md.Unix(), yd2.Unix())
}

func TestTimeX_AddSeconds(t *testing.T) {
	tx := timex.Now()

	h1 := tx.AddHour(1)
	s1 := tx.AddSeconds(timex.OneHourSec)
	assert.Equal(t, h1.Unix(), s1.Unix())

	assert.Equal(t, timex.OneHour, h1.Diff(tx.Time))
	assert.Equal(t, timex.OneHourSec, h1.DiffSec(tx.Time))
}

func TestTimeX_HourStart(t *testing.T) {
	tx := timex.Now()
	hs := tx.HourStart()
	he := tx.HourEnd()

	assert.Equal(t, "00:00", hs.DateFormat("I:S"))
	assert.Equal(t, "59:59", he.DateFormat("I:S"))
}

func TestTimeX_CustomHMS(t *testing.T) {
	tx := timex.Now()
	assert.Equal(t, "12:23:34", tx.CustomHMS(12, 23, 34).TplFormat("H:I:S"))
}
