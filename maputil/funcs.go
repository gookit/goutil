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

// Pick returns a new map with only the given keys that exist in the map.
func Pick[K comparable, V any](mp map[K]V, keys ...K) map[K]V {
	if len(mp) == 0 || len(keys) == 0 {
		return nil
	}

	newMp := make(map[K]V, len(keys))
	for _, key := range keys {
		if val, ok := mp[key]; ok {
			newMp[key] = val
		}
	}
	return newMp
}

// Omit returns a new map without the given keys.
func Omit[K comparable, V any](mp map[K]V, keys ...K) map[K]V {
	if len(mp) == 0 {
		return nil
	}
	if len(keys) == 0 {
		return Clone(mp)
	}

	omit := make(map[K]struct{}, len(keys))
	for _, key := range keys {
		omit[key] = struct{}{}
	}

	capNum := len(mp) - len(keys)
	if capNum < 0 {
		capNum = 0
	}

	newMp := make(map[K]V, capNum)
	for k, v := range mp {
		if _, ok := omit[k]; !ok {
			newMp[k] = v
		}
	}
	return newMp
}

// MergeFunc copies the key/value pairs of src into dst, and uses fn to resolve a key
// that is present in both maps. The first argument of fn is the value of dst.
//
// It works like maps.Copy for the keys that only exist in one of the maps.
func MergeFunc[K comparable, V any](dst, src map[K]V, fn func(dstVal, srcVal V) V) {
	for k, srcVal := range src {
		if dstVal, ok := dst[k]; ok {
			dst[k] = fn(dstVal, srcVal)
		} else {
			dst[k] = srcVal
		}
	}
}
