package fmtutil_test

import (
	"testing"
	"time"

	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/x/fmtutil"
)

func TestHowLongAgo(t *testing.T) {
	tests := []struct {
		args int64
		want string
	}{
		{-36, "unknown"},
		{36, "36 secs"},
		{346, "5 mins"},
		{3467, "57 mins"},
		{346778, "4 days"},
		{1200346778, "463 months"},
	}

	for _, tt := range tests {
		assert.Eq(t, tt.want, fmtutil.HowLongAgo(tt.args))
	}
}

func TestFormatDuration(t *testing.T) {
	tests := []struct {
		input    time.Duration
		expected string
	}{
		{30 * time.Second, "00:30"},
		{90 * time.Second, "01:30"},
		{3661 * time.Second, "01:01:01"},
		{7200 * time.Second, "02:00:00"},
		{-1 * time.Second, "00:00"},
	}

	for _, tt := range tests {
		result := fmtutil.FormatDuration(tt.input)
		assert.Eq(t, tt.expected, result)
	}
}
