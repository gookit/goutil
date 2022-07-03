package stdutil

import (
	"reflect"

	"github.com/gookit/goutil/reflects"
)

// TODO IsEqual(s, d interface{}) bool
// IsContains(v, sub interface{}) bool

// IsNil value check
func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}
	return reflects.IsNil(reflect.ValueOf(v))
}

// IsEmpty value check
func IsEmpty(v interface{}) bool {
	if v == nil {
		return true
	}
	return reflects.IsEmpty(reflect.ValueOf(v))
}

// IsFunc value
func IsFunc(val interface{}) bool {
	if val == nil {
		return false
	}
	return reflect.TypeOf(val).Kind() == reflect.Func
}

// ValueIsEmpty reflect value check.
//
// Deprecated: please use reflects.IsEmpty()
func ValueIsEmpty(v reflect.Value) bool {
	return reflects.IsEmpty(v)
}

// ValueLen get reflect value length
//
// Deprecated: please use reflects.Len()
func ValueLen(v reflect.Value) int {
	return reflects.Len(v)
}
