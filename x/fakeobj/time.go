package fakeobj

import (
	"time"

	"github.com/gookit/goutil/x/basefn"
)

// Clock mock time clock for test
type Clock struct {
	tt time.Time
}

// NewClock create a mock clock instance from layout "2006-01-02 15:04:05"
//
// Example:
// 	tc := NewClock("2023-01-01 12:00:00")
// 	tc.Add(time.Second * 15)
//	ds := tc.Datetime() // "2023-01-01 12:00:15"
func NewClock(value string) *Clock {
	nt, err := time.Parse("2006-01-02 15:04:05", value)
	basefn.PanicErr(err)
	return &Clock{tt: nt}
}

// Now get current time.
func (mt *Clock) Now() time.Time {
	return mt.tt
}

// Add progresses time by the given duration.
func (mt *Clock) Add(d time.Duration) {
	mt.tt = mt.tt.Add(d)
}

// Datetime returns the current time in the format "2006-01-02 15:04:05".
func (mt *Clock) Datetime() string {
	return mt.tt.Format("2006-01-02 15:04:05")
}
