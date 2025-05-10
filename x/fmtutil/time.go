package fmtutil

import (
	"github.com/gookit/goutil/mathutil"
)

// HowLongAgo format a seconds, got how lang ago
func HowLongAgo(sec int64) string {
	return mathutil.HowLongAgo(sec)
}
