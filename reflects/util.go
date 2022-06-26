package reflects

import "reflect"

// Elem returns the value that the interface v contains
// or that the pointer v points to.
func Elem(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		return v.Elem()
	}

	// otherwise, will return self
	return v
}

// HasChild check. eg: array, slice, map, struct
func HasChild(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Struct:
		return true
	}
	return false
}
