package mathutil

import "math"

// MaxFloat compare and return max value
func MaxFloat(x, y float64) float64 {
	return math.Max(x, y)
}

// MaxInt compare and return max value
func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// SwapMaxInt compare and return max, min value
func SwapMaxInt(x, y int) (int, int) {
	if x > y {
		return x, y
	}
	return y, x
}

// MaxI64 compare and return max value
func MaxI64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

// SwapMaxI64 compare and return max, min value
func SwapMaxI64(x, y int64) (int64, int64) {
	if x > y {
		return x, y
	}
	return y, x
}

// Compare intX,floatX value by given op. returns `srcVal op(=,!=,<,<=,>,>=) dstVal`
//
// Usage:
//	mathutil.Compare(2, 3, ">") // false
//	mathutil.Compare(2, 1.3, ">") // true
//	mathutil.Compare(2.2, 1.3, ">") // true
//	mathutil.Compare(2.1, 2, ">") // true
func Compare(srcVal, dstVal interface{}, op string) (ok bool) {
	if srcVal == nil || dstVal == nil {
		return false
	}

	// float
	if srcFlt, ok := srcVal.(float64); ok {
		if dstFlt, err := ToFloat(dstVal); err == nil {
			return CompFloat(srcFlt, dstFlt, op)
		}
		return false
	}

	if srcFlt, ok := srcVal.(float32); ok {
		if dstFlt, err := ToFloat(dstVal); err == nil {
			return CompFloat(float64(srcFlt), dstFlt, op)
		}
		return false
	}

	// as int64
	srcInt, err := ToInt64(srcVal)
	if err != nil {
		return false
	}

	dstInt, err := ToInt64(dstVal)
	if err != nil {
		return false
	}

	return CompInt64(srcInt, dstInt, op)
}

// CompInt64 compare int64, returns the srcI64 op dstI64
func CompInt64(srcI64, dstI64 int64, op string) (ok bool) {
	switch op {
	case "<":
		ok = srcI64 < dstI64
	case "<=":
		ok = srcI64 <= dstI64
	case ">":
		ok = srcI64 > dstI64
	case ">=":
		ok = srcI64 >= dstI64
	case "=":
		ok = srcI64 == dstI64
	case "!=":
		ok = srcI64 != dstI64
	}
	return
}

// CompFloat compare float64
func CompFloat(srcF64, dstF64 float64, op string) (ok bool) {
	switch op {
	case "<":
		ok = srcF64 < dstF64
	case "<=":
		ok = srcF64 <= dstF64
	case ">":
		ok = srcF64 > dstF64
	case ">=":
		ok = srcF64 >= dstF64
	case "=":
		ok = srcF64 == dstF64
	case "!=":
		ok = srcF64 != dstF64
	}
	return
}
