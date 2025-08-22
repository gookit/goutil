package comfunc_test

import (
	"testing"
	"time"

	"github.com/gookit/goutil/internal/comfunc"
	"github.com/gookit/goutil/testutil/assert"
)

func TestIsDuration(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{s: "1s", want: true},
		{s: "1m", want: true},
		{s: "1h", want: true},
		{s: "-1h", want: true},
		{s: "-1.4h", want: true},
		{s: "1h35m", want: true},
		{s: "1d", want: true},
		{s: "1.3day", want: true},
		{s: "1day34min", want: true},
		// error cases
		{s: "d", want: false},
		{s: "d2", want: false},
	}

	for _, tt := range tests {
		assert.Eq(t, tt.want, comfunc.IsDuration(tt.s), "input: %s", tt.s)
	}
}

func TestToDuration(t *testing.T) {
	got, err := comfunc.ToDuration("1.5d")
	assert.Nil(t, err)
	assert.Eq(t, 36*time.Hour, got)
	// return

	got, err = comfunc.ToDuration("1h35m")
	assert.Nil(t, err)
	assert.Eq(t, time.Hour+35*time.Minute, got)

	tests := []struct {
		s    string
		want time.Duration
	}{
		{s: "1s", want: time.Second},
		{s: "1m", want: time.Minute},
		{s: "1h", want: time.Hour},
		{s: "1.5h", want: time.Hour + 30*time.Minute},
		{s: "-1h", want: -time.Hour},
		{s: "-1.5h", want: -time.Hour - 30*time.Minute},
		{s: "1h35m", want: time.Hour + 35*time.Minute},
		// extend shorts
		{s: "1d", want: 24 * time.Hour},
		{s: "1.2d", want: 28*time.Hour + 48*time.Minute},
		{s: "1.5d", want: 36 * time.Hour},
		{s: "3d", want: 3 * 24 * time.Hour},
		{s: "2w", want: 2 * 7 * 24 * time.Hour},
		// long unit
		{s: "1hour", want: time.Hour},
		{s: "-21hours", want: -time.Hour * 21},
		{s: "4hour", want: time.Hour * 4},
		{s: "4hours", want: time.Hour * 4},
		{s: "1day", want: time.Hour * 24},
		{s: "3days", want: time.Hour * 24 * 3},
		{s: "2week", want: time.Hour * 24 * 7 * 2},
		{s: "1month2day3h", want: time.Hour*24*32 + time.Hour*3},
		{s: "1hour34min", want: time.Hour + 34*time.Minute},
		{s: "1day34min", want: 24*time.Hour + 34*time.Minute},
		{s: "1day34min5sec", want: 24*time.Hour + 34*time.Minute + 5*time.Second},
		// complex
		{s: "1h34min", want: time.Hour + 34*time.Minute},
		{s: "1d34min5s", want: 24*time.Hour + 34*time.Minute + 5*time.Second},
	}

	for _, tt := range tests {
		got, err := comfunc.ToDuration(tt.s)
		assert.Nil(t, err)
		assert.Eq(t, tt.want, got, "input: %s", tt.s)
	}

	// test error case
	t.Run("error case", func(t *testing.T) {
		_, err := comfunc.ToDuration("")
		assert.Err(t, err)
		_, err = comfunc.ToDuration("-")
		assert.ErrMsg(t, err, "invalid duration string: -")
		_, err = comfunc.ToDuration("d2")
		assert.ErrMsg(t, err, "invalid duration string: d2")
		_, err = comfunc.ToDuration("a3d")
		assert.Err(t, err)
	})

}
