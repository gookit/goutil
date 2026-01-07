// Package mathutil provide math(int, number) util functions. eg: convert, math calc, random
package mathutil

import (
	"math"

	"github.com/gookit/goutil/comdef"
)

// Mul computes the `a*b` value, rounding the result.
func Mul[T1, T2 comdef.Number](a T1, b T2) float64 {
	return math.Round(SafeFloat(a) * SafeFloat(b))
}

// MulF2i computes the float64 type a * b value, rounding the result to an integer.
func MulF2i(a, b float64) int {
	return int(math.Round(a * b))
}

// Div computes the `a/b` value, result uses a round handle.
func Div[T1, T2 comdef.Number](a T1, b T2) float64 {
	return math.Round(SafeFloat(a) / SafeFloat(b))
}

// DivInt computes the int type a / b value, rounding the result to an integer.
func DivInt[T comdef.Integer](a, b T) int {
	fv := math.Round(float64(a) / float64(b))
	return int(fv)
}

// DivF2i computes the float64 type a / b value, rounding the result to an integer.
func DivF2i(a, b float64) int {
	return int(math.Round(a / b))
}

// Percent returns a value percentage of the total
func Percent(val, total int) float64 {
	if total == 0 {
		return float64(0)
	}
	return (float64(val) / float64(total)) * 100
}
