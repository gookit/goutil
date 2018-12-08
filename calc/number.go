package calc

import (
	"fmt"
	"time"
)

// Percent returns a values percent of the total
func Percent(val, total int) float64 {
	if total == 0 {
		return float64(0)
	}

	return (float64(val) / float64(total)) * 100
}

// CalcElapsedTime 计算运行时间消耗 单位 ms(毫秒)
func CalcElapsedTime(startTime time.Time) string {
	return fmt.Sprintf("%.3f", time.Since(startTime).Seconds()*1000)
}
