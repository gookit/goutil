// Package mathutil provide math(int, number) util functions. eg: convert, math calc, random
package mathutil

import (
	"github.com/gookit/goutil/comdef"
)

// OrElse return default value on val is zero, else return val
func OrElse[T comdef.XintOrFloat](val, defVal T) T {
	return ZeroOr(val, defVal)
}

// ZeroOr return default value on val is zero, else return val
func ZeroOr[T comdef.XintOrFloat](val, defVal T) T {
	if val != 0 {
		return val
	}
	return defVal
}

// LessOr return val on val < max, else return default value.
//
// Example:
//
//	LessOr(11, 10, 1) // 1
//	LessOr(2, 10, 1) // 2
//	LessOr(10, 10, 1) // 1
func LessOr[T comdef.XintOrFloat](val, max, devVal T) T {
	if val < max {
		return val
	}
	return devVal
}

// LteOr return val on val <= max, else return default value.
//
// Example:
//
//	LteOr(11, 10, 1) // 11
//	LteOr(2, 10, 1) // 2
//	LteOr(10, 10, 1) // 10
func LteOr[T comdef.XintOrFloat](val, max, devVal T) T {
	if val <= max {
		return val
	}
	return devVal
}

// GreaterOr return val on val > max, else return default value.
//
// Example:
//
//	GreaterOr(23, 0, 2) // 23
//	GreaterOr(0, 0, 2) // 2
func GreaterOr[T comdef.XintOrFloat](val, min, defVal T) T {
	if val > min {
		return val
	}
	return defVal
}

// GteOr return val on val >= max, else return default value.
//
// Example:
//
//	GteOr(23, 0, 2) // 23
//	GteOr(0, 0, 2) // 0
func GteOr[T comdef.XintOrFloat](val, min, defVal T) T {
	if val >= min {
		return val
	}
	return defVal
}
