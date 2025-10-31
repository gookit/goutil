package fmtutil

import (
	"time"

	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/timex"
)

// HowLongAgo format a seconds, got how lang ago
func HowLongAgo(sec int64) string {
	return mathutil.HowLongAgo(sec)
}

// FormatDuration Formatting time consumption is clock format 格式化时间消耗为时钟格式 eg: 90 * time.Second => "01:30"
func FormatDuration(d time.Duration) string {
	return timex.FormatDuration(d)
}
