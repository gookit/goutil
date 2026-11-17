//go:build go1.23

package maputil

import (
	"maps"
	"slices"

	"github.com/gookit/goutil/comdef"
)

// TypedKeys get all keys of the given typed map.
//
// The key order is not fixed, like a range over the map.
func TypedKeys[K comdef.SimpleType, V any](mp map[K]V) (keys []K) {
	return slices.Collect(maps.Keys(mp))
}

// TypedValues get all values of the given typed map.
//
// The value order is not fixed, like a range over the map.
func TypedValues[K comdef.SimpleType, V any](mp map[K]V) (values []V) {
	return slices.Collect(maps.Values(mp))
}
