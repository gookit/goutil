package timex

import (
	"fmt"
	"strings"
	"time"

	"github.com/gookit/goutil/internal/comfunc"
	"github.com/gookit/goutil/strutil"
)

// Elapsed calc elapsed time from start time to end time.
func Elapsed(start, end time.Time) string {
	dur := end.Sub(start)

	switch {
	case dur > time.Hour:
		return fmt.Sprintf("%.2fhrs", dur.Hours())
	case dur > time.Minute:
		return fmt.Sprintf("%.2fmins", dur.Minutes())
	case dur > time.Second:
		return fmt.Sprintf("%.3fs", dur.Seconds())
	case dur > time.Millisecond:
		return fmt.Sprintf("%.2fms", float64(dur.Nanoseconds())/1e6)
	default:
		return fmt.Sprintf("%.2fµs", float64(dur.Nanoseconds())/1e3)
	}
}

// ElapsedNow calc elapsed time from start time to now.
func ElapsedNow(start time.Time) string {
	return Elapsed(start, time.Now())
}

//
// -------- parse time diff to string --------
//

// TimeMessage struct for HowLongAgo2(), FromNowWith()
type TimeMessage struct {
	// Message string or format string
	Message string
	// Seconds time range.
	// first elem is max boundary value, second elem is divisor(unit).
	Seconds []int
}

// TimeMessages time message list.
//
// NOTE: last item.Seconds[0] is min boundary value.
var TimeMessages = []TimeMessage{
	{"< 1 sec ago", []int{0}},
	{"1 sec ago", []int{1}},
	{"%d secs ago", []int{45, 1}},
	{"1 min ago", []int{89}},
	{"%d mins ago", []int{44 * 60, 60}},       // 89s - 44min
	{"1 hour ago", []int{89 * 60}},            // 45min - 89min
	{"%d hours ago", []int{21 * 3600, 3600}},  // 90min - 21hr
	{"1 day ago", []int{35 * 3600}},           // 22 - 35 hours
	{"%d days ago", []int{30 * 86400, 86400}}, // 36hr - 30 day
	// {"1 week ago", []int{10 * 86400}},          // 7 - 10 days
	// {"%d weeks ago", []int{30 * 86400, 604800}},  // 10 - 30 days
	{"1 month ago", []int{45 * 86400}},                 // 31 - 45 days
	{"%d months ago", []int{319 * 86400, 2592000}},     // 44 - 319 days
	{"1 year ago", []int{547 * 86400}},                 // 320 - 547 days
	{"%d years ago", []int{547 * 86400, 12 * 2592000}}, // > 547 days, unit is year
}

// FromNow format time from now, returns like: 1 hour ago, 2 days ago
//
// refer: https://gist.github.com/davidrleonard/259fe449b1ec13bf7d87cde567ca0fde
func FromNow(t time.Time) string {
	return FromNowWith(t, TimeMessages)
}

// FromNowWith format time from now with custom TimeMessage list
func FromNowWith(u time.Time, tms []TimeMessage) string {
	return HowLongAgo2(int64(time.Since(u).Seconds()), tms)
}

// HowLongAgo format diff time seconds to string. alias of HowLongAgo2()
func HowLongAgo(diffSec int64) string {
	return HowLongAgo2(diffSec, TimeMessages)
}

// HowLongAgo2 format diff time seconds with custom TimeMessage list
func HowLongAgo2(diffSec int64, tms []TimeMessage) string {
	length := len(tms)
	diffInt := int(diffSec)

	var msg string
	var secs []int
	for i, item := range tms {
		msg, secs = item.Message, item.Seconds

		// match success: is last elem or diffSec <= secs[0]
		if i+1 == length || diffInt <= secs[0] {
			break
		}
	}

	if len(secs) == 1 {
		return msg
	}
	return fmt.Sprintf(msg, int64(diffInt/secs[1]))
}

//
// -------- parse string to time --------
//

// ToTime parse a datetime string. alias of strutil.ToTime()
func ToTime(s string, layouts ...string) (time.Time, error) {
	return strutil.ToTime(s, layouts...)
}

// ToDur parse a duration string. alias of ToDuration()
func ToDur(s string) (time.Duration, error) { return ToDuration(s) }

// ToDuration parses a duration string. such as "300ms", "-1.5h" or "2h45m".
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
func ToDuration(s string) (time.Duration, error) {
	return comfunc.ToDuration(s)
}

// TryToTime parse a date string or duration string to time.Time.
//
// if s is empty, return zero time.
func TryToTime(s string, bt time.Time) (time.Time, error) {
	if s == "" {
		return ZeroTime, nil
	}
	if s == "now" {
		return time.Now(), nil
	}

	// if s is a duration string, add it to bt(base time)
	if IsDuration(s) {
		dur, err := ToDuration(s)
		if err != nil {
			return ZeroTime, err
		}
		return bt.Add(dur), nil
	}

	// as a date string, parse it to time.Time
	return ToTime(s)
}

//
// -------- parse string to time range --------
//

// ParseRangeOpt is the option for ParseRange
type ParseRangeOpt struct {
	// BaseTime is the base time for relative time string.
	// if is zero, use time.Now() as base time.
	BaseTime time.Time
	// OneAsEnd is the option for one time range.
	//  - False: "-1h" => "-1h,0"; "1h" => "+1h, feature"
	//  - True:  "-1h" => "zero,-1h"; "1h" => "zero,1h"
	OneAsEnd bool
	// AutoSort is the option for sort the time range.
	AutoSort bool
	// SepChar is the separator char for time range string. default is '~'
	SepChar byte
	// BeforeFn hook for before parse time string.
	BeforeFn func(string) string
	// KeywordFn is the function for parse keyword time string.
	KeywordFn func(string) (time.Time, time.Time, error)
}

func ensureOpt(opt *ParseRangeOpt) *ParseRangeOpt {
	if opt == nil {
		opt = &ParseRangeOpt{BaseTime: time.Now(), SepChar: '~'}
	} else {
		if opt.BaseTime.IsZero() {
			opt.BaseTime = time.Now()
		}
		if opt.SepChar == 0 {
			opt.SepChar = '~'
		}
	}

	return opt
}

// ParseRange parse time range expression string to time.Time range.
//
//   - "0" will use opt.BaseTime.
//
// Expression format:
//
//	"-5h~-1h"       	=> 5 hours ago to 1 hour ago
//	"1h~5h"         	=> 1 hour after to 5 hours after
//	"-1h~1h"        	=> 1 hour ago to 1 hour after
//	"-1h"            	=> 1 hour ago to feature. eq "-1h~"
//	"-1h~0"          	=> 1 hour ago to now.
//	"< -1h" OR "~-1h"   => 1 hour ago.
//	"> 1h" OR "1h"     	=> 1 hour after to feature
//	// keyword: now, today, yesterday, tomorrow
//	"today"          => today start to today end
//	"yesterday"      => yesterday start to yesterday end
//	"tomorrow"       => tomorrow start to tomorrow end
//
// Usage:
//
//	start, end, err := ParseRange("-1h~1h", nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(start, end)
func ParseRange(expr string, opt *ParseRangeOpt) (start, end time.Time, err error) {
	opt = ensureOpt(opt)
	expr = strings.TrimSpace(expr)
	if expr == "" {
		err = fmt.Errorf("invalid time range expr %q", expr)
		return
	}

	// parse time range. eg: "5h~1h"
	if strings.IndexByte(expr, opt.SepChar) > -1 {
		s1, s2 := strutil.TrimCut(expr, string(opt.SepChar))
		if s1 == "" && s2 == "" {
			err = fmt.Errorf("invalid time range expr: %s", expr)
			return
		}

		if s1 != "" {
			start, err = TryToTime(s1, opt.BaseTime)
			if err != nil {
				return
			}
		}

		if s2 != "" {
			end, err = TryToTime(s2, opt.BaseTime)
			// auto sort range time
			if opt.AutoSort && err == nil {
				if !start.IsZero() && start.After(end) {
					start, end = end, start
				}
			}
		}

		return
	}

	// single time. eg: "5h", "1h", "-1h"
	if IsDuration(expr) {
		tt, err1 := TryToTime(expr, opt.BaseTime)
		if err1 != nil {
			err = err1
			return
		}

		if opt.OneAsEnd {
			end = tt
		} else {
			start = tt
		}
		return
	}

	// with compare operator. eg: "<1h", ">1h"
	if expr[0] == '<' || expr[0] == '>' {
		tt, err1 := TryToTime(strings.Trim(expr[1:], " ="), opt.BaseTime)
		if err1 != nil {
			err = err1
			return
		}

		if expr[0] == '<' {
			end = tt
		} else {
			start = tt
		}
		return
	}

	// parse keyword time string
	switch expr {
	case "0":
		if opt.OneAsEnd {
			end = opt.BaseTime
		} else {
			start = opt.BaseTime
		}
	case "now":
		if opt.OneAsEnd {
			end = time.Now()
		} else {
			start = time.Now()
		}
	case "today":
		start = DayStart(opt.BaseTime)
		end = DayEnd(opt.BaseTime)
	case "yesterday":
		yd := opt.BaseTime.AddDate(0, 0, -1)
		start = DayStart(yd)
		end = DayEnd(yd)
	case "tomorrow":
		td := opt.BaseTime.AddDate(0, 0, 1)
		start = DayStart(td)
		end = DayEnd(td)
	default:
		// single datetime. eg: "2019-01-01"
		tt, err1 := TryToTime(expr, opt.BaseTime)
		if err1 != nil {
			if opt.KeywordFn == nil {
				err = fmt.Errorf("invalid keyword time string: %s", expr)
				return
			}

			start, end, err = opt.KeywordFn(expr)
			return
		}

		if opt.OneAsEnd {
			end = tt
		} else {
			start = tt
		}
	}

	return
}
