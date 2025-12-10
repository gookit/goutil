package fakeobj_test

import (
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/timex"
	"github.com/gookit/goutil/x/fakeobj"
)

func TestNewClock(t *testing.T) {
	tc := fakeobj.NewClock("2025-11-02 16:32:05")

	// add time
	tc.Add(timex.Day)

	assert.Eq(t, "2025-11-03 16:32:05", tc.Datetime())
	assert.StrContains(t, tc.Now().String(), "2025-11-03 16:32:05")
}
