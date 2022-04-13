package timex

import "time"

const (
	OneMinSec  = 60
	OneHourSec = 3600
	OneDaySec  = 86400

	OneMin  = time.Minute
	OneHour = time.Hour
	OneDay  = 24 * time.Hour
	OneWeek = 7 * 24 * time.Hour
)

var (
	// DefaultLayout template for format time
	DefaultLayout = "2006-01-02 15:04:05"
)

// TimeX struct
type TimeX struct {
	time.Time
	// DateLayout set the default date format layout
	DateLayout string
}

// Now time
func Now() TimeX {
	return TimeX{
		Time: time.Now(),
	}
}

// Local time for now
func Local() TimeX {
	return TimeX{
		Time: time.Now().In(time.Local),
	}
}

// LocalByName time for now
func LocalByName(tzName string) TimeX {
	loc, err := time.LoadLocation(tzName)
	if err != nil {
		panic(err)
	}

	return TimeX{
		Time: time.Now().In(loc),
	}
}

// Datetime format time use DefaultLayout
func (t *TimeX) Datetime() string {
	if t.DateLayout == "" {
		t.DateLayout = DefaultLayout
	}
	return t.Format(t.DateLayout)
}
