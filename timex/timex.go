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
	// DateLayout set the default date format layout. default use DefaultLayout
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

// SetLocalByName tz name. eg: UTC, PRC
func SetLocalByName(tzName string) error {
	location, err := time.LoadLocation(tzName)
	if err != nil {
		return err
	}

	time.Local = location
	return nil
}

// Datetime format time use DefaultLayout
func (t *TimeX) Datetime() string {
	if t.DateLayout == "" {
		t.DateLayout = DefaultLayout
	}
	return t.Format(t.DateLayout)
}

// AddDay add some day time for the time
func (t *TimeX) AddDay(day int) TimeX {
	return t.AddSeconds(day * OneDaySec)
}

// AddHour add some hour time
func (t *TimeX) AddHour(hours int) TimeX {
	return t.AddSeconds(hours * OneHourSec)
}

// AddMinutes add some minutes time for the time
func (t *TimeX) AddMinutes(minutes int) TimeX {
	return t.AddSeconds(minutes * OneMinSec)
}

// AddSeconds add some seconds time the time
func (t *TimeX) AddSeconds(seconds int) TimeX {
	return TimeX{
		Time: t.Add(time.Duration(seconds) * time.Second),
		// with layout
		DateLayout: DefaultLayout,
	}
}
