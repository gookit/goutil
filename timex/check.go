package timex

import (
	"time"

	"github.com/gookit/goutil/internal/comfunc"
)

// IsDuration check the string is a valid duration string.
// alias of comfunc.IsDuration()
func IsDuration(s string) bool {
	return comfunc.IsDuration(s)
}

// InRange check the dst time is in the range of start and end.
//
// if start is zero, only check dst < end,
// if end is zero, only check dst > start.
func InRange(dst, start, end time.Time) bool {
	if start.IsZero() && end.IsZero() {
		return false
	}

	if start.IsZero() {
		return dst.Before(end)
	}
	if end.IsZero() {
		return dst.After(start)
	}

	return dst.After(start) && dst.Before(end)
}
