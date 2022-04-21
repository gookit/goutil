package timex_test

import (
	"testing"
	"time"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/timex"
	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	sec := timex.NowUnix()

	assert.NotEmpty(t, timex.FormatUnix(sec))
	assert.NotEmpty(t, timex.FormatUnixBy(sec, time.RFC3339))
}

func TestFormatByTpl(t *testing.T) {
	now := time.Now()

	assert.Equal(t, now.Format("20060102 15:04:05"), timex.FormatByTpl(now, "YMD H:I:S"))
	assert.Equal(t, now.Format("2006-01-02 15:04:05"), timex.FormatByTpl(now, "Y-M-D H:I:S"))
	assert.Equal(t, now.Format("01/02 15:04:05"), timex.Date(now, "M/D H:I:S"))
	assert.Equal(t, now.Format("06/01/02 15:04:05"), timex.Date(now, "y/M/D H:I:S"))
	assert.Equal(t, now.Format("2006-01-02 15:04"), timex.DateFormat(now, "Y-M-D H:I"))
}

func TestFormatUnix(t *testing.T) {
	now := time.Now()
	want := now.Format("2006-01-02 15:04:05")

	assert.Equal(t, want, timex.FormatUnix(now.Unix()))
	assert.Equal(t, want, timex.FormatUnixBy(now.Unix(), timex.DefaultLayout))
	assert.Equal(t, want, timex.FormatUnixByTpl(now.Unix(), "Y-M-D H:I:S"))
	dump.P(want)
}

func TestToLayout(t *testing.T) {
	assert.Equal(t, timex.DefaultLayout, timex.ToLayout(""))
	assert.Equal(t, time.RFC3339, timex.ToLayout("Y-M-DTH:I:SZ07:00"))
}
