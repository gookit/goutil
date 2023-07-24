// Package reflects Provide extends reflect util functions.
package reflects

import "reflect"

// MakeSliceByElem create a new slice by the element type.
//
// - elType: the type of the element.
// - returns: the new slice.
//
// Usage:
//
//	sl := MakeSliceByElem(reflect.TypeOf(1), 10, 20)
//	sl.Index(0).SetInt(10)
//
//	// Or use reflect.AppendSlice() merge two slice
//	// Or use `for` with `reflect.Append()` add elements
func MakeSliceByElem(elTyp reflect.Type, len, cap int) reflect.Value {
	return reflect.MakeSlice(reflect.SliceOf(elTyp), len, cap)
}
