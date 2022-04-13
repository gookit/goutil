package timex

import "time"

// NowUnix is short of time.Now().Unix()
func NowUnix() int64 {
	return time.Now().Unix()
}

// Format use default layout
func Format(t time.Time) string {
	return t.Format(DefaultLayout)
}

// FormatBy given default layout
func FormatBy(t time.Time, layout string) string {
	return t.Format(layout)
}

// NowAddDay add some day time from now
func NowAddDay(day int) time.Time {
	return time.Now().AddDate(0, 0, day)
}

// NowAddMinutes add some minutes time from now
func NowAddMinutes(minutes int) time.Time {
	return time.Now().Add(time.Duration(minutes) * OneMin)
}

// NowAddSeconds add some seconds time from now
func NowAddSeconds(seconds int) time.Time {
	return time.Now().Add(time.Duration(seconds) * time.Second)
}

// AddDay add some day time for given time
func AddDay(t time.Time, day int) time.Time {
	return t.AddDate(0, 0, day)
}

// AddMinutes add some minutes time for given time
func AddMinutes(t time.Time, minutes int) time.Time {
	return t.Add(time.Duration(minutes) * OneMin)
}

// AddSeconds add some seconds time for given time
func AddSeconds(t time.Time, seconds int) time.Time {
	return t.Add(time.Duration(seconds) * time.Second)
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
