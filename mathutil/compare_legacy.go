//go:build !go1.21

package mathutil

import "github.com/gookit/goutil/comdef"

// NOTE: this file has the same functions as compare_go121.go, written by hand for
// the Go versions that do not have the min and max builtins (go1.21+).

// Min compare two value and return min value
func Min[T comdef.Number](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// Max compare two value and return max value
func Max[T comdef.Number](x, y T) T {
	if x > y {
		return x
	}
	return y
}
