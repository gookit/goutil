package timex

import (
	"time"

	"github.com/gookit/goutil/fmtutil"
)

// NowUnix is short of time.Now().Unix()
func NowUnix() int64 {
	return time.Now().Unix()
}

// SetLocalByName set local by tz name. eg: UTC, PRC
func SetLocalByName(tzName string) error {
	location, err := time.LoadLocation(tzName)
	if err != nil {
		return err
	}

	time.Local = location
	return nil
}

// Format use default layout
func Format(t time.Time) string {
	return t.Format(DefaultLayout)
}

// FormatBy given default layout
func FormatBy(t time.Time, layout string) string {
	return t.Format(layout)
}

// Date format time by given date template.
// see ToLayout()
func Date(t time.Time, template string) string {
	return FormatByTpl(t, template)
}

// DateFormat format time by given date template.
// see ToLayout()
func DateFormat(t time.Time, template string) string {
	return FormatByTpl(t, template)
}

// FormatByTpl format time by given date template.
// see ToLayout()
func FormatByTpl(t time.Time, template string) string {
	return t.Format(ToLayout(template))
}

// FormatUnix time seconds use default layout
func FormatUnix(sec int64) string {
	return time.Unix(sec, 0).Format(DefaultLayout)
}

// FormatUnixBy format time seconds use given layout
func FormatUnixBy(sec int64, layout string) string {
	return time.Unix(sec, 0).Format(layout)
}

// FormatUnixByTpl format time seconds use given date template.
// see ToLayout()
func FormatUnixByTpl(sec int64, template string) string {
	return time.Unix(sec, 0).Format(ToLayout(template))
}

// NowAddDay add some day time from now
func NowAddDay(day int) time.Time {
	return time.Now().AddDate(0, 0, day)
}

// NowAddHour add some hour time from now
func NowAddHour(hour int) time.Time {
	return time.Now().Add(time.Duration(hour) * OneHour)
}

// NowAddMinutes add some minutes time from now
func NowAddMinutes(minutes int) time.Time {
	return time.Now().Add(time.Duration(minutes) * OneMin)
}

// NowAddSeconds add some seconds time from now
func NowAddSeconds(seconds int) time.Time {
	return time.Now().Add(time.Duration(seconds) * time.Second)
}

// NowHourStart time
func NowHourStart() time.Time {
	return HourStart(time.Now())
}

// NowHourEnd time
func NowHourEnd() time.Time {
	return HourEnd(time.Now())
}

// AddDay add some day time for given time
func AddDay(t time.Time, day int) time.Time {
	return t.AddDate(0, 0, day)
}

// AddHour add some hour time for given time
func AddHour(t time.Time, hour int) time.Time {
	return t.Add(time.Duration(hour) * OneHour)
}

// AddMinutes add some minutes time for given time
func AddMinutes(t time.Time, minutes int) time.Time {
	return t.Add(time.Duration(minutes) * OneMin)
}

// AddSeconds add some seconds time for given time
func AddSeconds(t time.Time, seconds int) time.Time {
	return t.Add(time.Duration(seconds) * time.Second)
}

// HourStart time for given time
func HourStart(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 0, 0, 0, t.Location())
}

// HourEnd time for given time
func HourEnd(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// DayStart time for given time
func DayStart(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// DayEnd time for given time
func DayEnd(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// TodayStart time
func TodayStart() time.Time {
	return DayStart(time.Now())
}

// TodayEnd time
func TodayEnd() time.Time {
	return DayEnd(time.Now())
}

// HowLongAgo format given timestamp to string.
func HowLongAgo(sec int64) string {
	return fmtutil.HowLongAgo(sec)
}

// ToDuration parses a duration string. such as "300ms", "-1.5h" or "2h45m".
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
func ToDuration(s string) (time.Duration, error) {
	return time.ParseDuration(s)
}
