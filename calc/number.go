package calc

import (
	"fmt"
	"time"

	"github.com/gookit/goutil/fmtutil"
)

// Percent returns a values percent of the total
// Deprecated
//	please use mathutil.Percent() instead
func Percent(val, total int) float64 {
	if total == 0 {
		return float64(0)
	}

	return (float64(val) / float64(total)) * 100
}

// ElapsedTime calc elapsed time 计算运行时间消耗 单位 ms(毫秒)
// Deprecated
//	please use mathutil.ElapsedTime() instead
func ElapsedTime(startTime time.Time) string {
	return fmt.Sprintf("%.3f", time.Since(startTime).Seconds()*1000)
}

// DataSize format value.
// Deprecated
//	please use fmtutil.DataSize() instead
func DataSize(size uint64) string {
	return fmtutil.DataSize(size)
}

// HowLongAgo calc time.
// Deprecated
//	please use fmtutil.HowLongAgo() instead
func HowLongAgo(sec int64) string {
	return fmtutil.HowLongAgo(sec)
}
