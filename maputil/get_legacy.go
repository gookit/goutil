//go:build !go1.23

package maputil

import "github.com/gookit/goutil/comdef"

// NOTE: this file has the same functions as get_go123.go, written by hand for the Go
// versions that do not have the maps and slices iterators (go1.23+). Keep both in sync.

// TypedKeys get all keys of the given typed map.
//
// The key order is not fixed, like a range over the map.
func TypedKeys[K comdef.SimpleType, V any](mp map[K]V) (keys []K) {
	for key := range mp {
		keys = append(keys, key)
	}
	return
}

// TypedValues get all values of the given typed map.
//
// The value order is not fixed, like a range over the map.
func TypedValues[K comdef.SimpleType, V any](mp map[K]V) (values []V) {
	for _, val := range mp {
		values = append(values, val)
	}
	return
}
