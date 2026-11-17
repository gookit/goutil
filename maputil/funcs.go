package maputil

// Filter returns a new map with the key/value pairs for which fn returns true.
func Filter[K comparable, V any](mp map[K]V, fn func(K, V) bool) map[K]V {
	if len(mp) == 0 {
		return nil
	}

	newMp := make(map[K]V, len(mp))
	for k, v := range mp {
		if fn(k, v) {
			newMp[k] = v
		}
	}
	return newMp
}

// MapValues returns a new map with the values mapped by fn, the keys are kept.
func MapValues[K comparable, V any, R any](mp map[K]V, fn func(K, V) R) map[K]R {
	if len(mp) == 0 {
		return nil
	}

	newMp := make(map[K]R, len(mp))
	for k, v := range mp {
		newMp[k] = fn(k, v)
	}
	return newMp
}

// Invert returns a new map with the keys and values of the given map swapped.
//
// When a value is not unique the last key of the iteration wins, so it is mostly
// useful for a map with unique values.
func Invert[K comparable, V comparable](mp map[K]V) map[V]K {
	if len(mp) == 0 {
		return nil
	}

	newMp := make(map[V]K, len(mp))
	for k, v := range mp {
		newMp[v] = k
	}
	return newMp
}
