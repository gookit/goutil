package checkfn

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gookit/goutil/reflects"
)

// IsNil value check
func IsNil(v any) bool {
	if v == nil {
		return true
	}
	return reflects.IsNil(reflect.ValueOf(v))
}

// IsEmpty value check
func IsEmpty(v any) bool {
	if v == nil {
		return true
	}
	return reflects.IsEmpty(reflect.ValueOf(v))
}

// Contains try loop over the data check if the data includes the element.
//
// data allow types: string, map, array, slice
//
//	map         - check key exists
//	string      - check sub-string exists
//	array,slice - check sub-element exists
//
// Returns:
//   - valid: data is valid
//   - found: element was found
//
// return (false, false) if impossible.
// return (true, false) if element was not found.
// return (true, true) if element was found.
func Contains(data, elem any) (valid, found bool) {
	if data == nil {
		return false, false
	}

	dataRv := reflect.ValueOf(data)
	dataRt := reflect.TypeOf(data)
	dataKind := dataRt.Kind()

	// string
	if dataKind == reflect.String {
		return true, strings.Contains(dataRv.String(), fmt.Sprint(elem))
	}

	// map
	if dataKind == reflect.Map {
		mapKeys := dataRv.MapKeys()
		for i := 0; i < len(mapKeys); i++ {
			if reflects.IsEqual(mapKeys[i].Interface(), elem) {
				return true, true
			}
		}
		return true, false
	}

	// array, slice - other return false
	if dataKind != reflect.Slice && dataKind != reflect.Array {
		return false, false
	}

	for i := 0; i < dataRv.Len(); i++ {
		if reflects.IsEqual(dataRv.Index(i).Interface(), elem) {
			return true, true
		}
	}
	return true, false
}
