package timex_test

import (
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/timex"
	"github.com/stretchr/testify/assert"
)

func TestFromDate(t *testing.T) {
	tx, err := timex.FromDate("2022-04-20 19:40:34")
	assert.NoError(t, err)
	assert.Equal(t, "2022-04-20 19:40:34", tx.Datetime())

	tx, err = timex.FromDate("2022-04-20 19:40:34", "Y-M-D H:I:S")
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
	assert.Equal(t, tx.Datetime(), tx.DateFormat("Y-M-D H:I:S"))
}

func TestTimeX_SubUnix(t *testing.T) {
	tx := timex.Now()

	after1m := tx.AddMinutes(1)

	assert.Equal(t, timex.OneMinSec, after1m.SubUnix(tx.Time))
}

func TestTimeX_DateFormat(t *testing.T) {
	tx := timex.Now()
	assert.Equal(t, tx.Format(timex.DefaultLayout), tx.DateFormat("Y-M-D H:I:S"))
	assert.Equal(t, tx.Format("2006/01/02 15:04"), tx.TplFormat("Y/M/D H:I"))

	date := tx.Format("06/01/02 15:04")
	assert.Equal(t, date, tx.DateFormat("y/M/D H:I"))
	dump.V(date)
}

func TestTimeX_AddDay(t *testing.T) {
	tx := timex.Now()

	yd := tx.Yesterday()
	yd1 := tx.AddDay(-1)
	assert.Equal(t, yd.Unix(), yd1.Unix())
	assert.Equal(t, yd.Unix(), tx.DayAgo(1).Unix())

	assert.True(t, tx.IsAfter(yd.Time))
	assert.True(t, yd.IsBefore(tx.Time))
	assert.Equal(t, tx.Unix()-yd.Unix(), int64(timex.OneDaySec))
}

func TestTimeX_CustomHMS(t *testing.T) {
	tx := timex.Now()
	assert.Equal(t, "12:23:34", tx.CustomHMS(12, 23, 34).TplFormat("H:I:S"))
}
