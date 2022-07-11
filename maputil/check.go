package maputil

import (
	"reflect"

	"github.com/gookit/goutil/reflects"
)

// HasKey check of the given map.
func HasKey(mp, key interface{}) (ok bool) {
	rftVal := reflect.Indirect(reflect.ValueOf(mp))
	if rftVal.Kind() != reflect.Map {
		return
	}

	for _, keyRv := range rftVal.MapKeys() {
		if reflects.IsEqual(keyRv.Interface(), key) {
			return true
		}
	}
	return
}
