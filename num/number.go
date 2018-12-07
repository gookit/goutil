package num

// Percent returns a values percent of the total
func Percent(val, total int) float64 {
	if total == 0 {
		return float64(0)
	}

	return (float64(val) / float64(total)) * 100
}
