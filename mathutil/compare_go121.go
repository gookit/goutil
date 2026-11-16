//go:build go1.21

package mathutil

import "github.com/gookit/goutil/comdef"

// Min compare two value and return min value.
//
// NOTE: uses the min builtin, so a NaN float gives NaN back.
func Min[T comdef.Number](x, y T) T {
	return min(x, y)
}

// Max compare two value and return max value.
//
// NOTE: uses the max builtin, so a NaN float gives NaN back.
func Max[T comdef.Number](x, y T) T {
	return max(x, y)
}
