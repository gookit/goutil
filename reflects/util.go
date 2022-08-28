package reflects

import (
	"fmt"
	"reflect"
	"strconv"
)

// Elem returns the value that the interface v contains
// or that the pointer v points to.
func Elem(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		return v.Elem()
	}

	// otherwise, will return self
	return v
}

// Len get reflect value length
func Len(v reflect.Value) int {
	v = reflect.Indirect(v)

	// (u)int use width.
	switch v.Kind() {
	case reflect.String:
		return len([]rune(v.String()))
	case reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return v.Len()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return len(strconv.FormatInt(int64(v.Uint()), 10))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return len(strconv.FormatInt(v.Int(), 10))
	case reflect.Float32, reflect.Float64:
		return len(fmt.Sprint(v.Interface()))
	}

	// cannot get length
	return -1
}

// SliceSubKind get sub-elem kind of the array, slice, variadic-var.
//
// Usage:
//
//	SliceSubKind(reflect.TypeOf([]string{"abc"})) // reflect.String
func SliceSubKind(typ reflect.Type) reflect.Kind {
	if typ.Kind() == reflect.Slice {
		return typ.Elem().Kind()
	}
	return reflect.Invalid
}
