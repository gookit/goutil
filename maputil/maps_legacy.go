//go:build !go1.21

package maputil

// NOTE: this file has the same functions as maps_go121.go, written by hand for
// the Go versions that do not have the maps package (go1.21+). Keep both in sync.

// Clone copy and return a new map. A nil or empty map gives back nil.
func Clone[M ~map[K]V, K comparable, V any](items M) M {
	if len(items) == 0 {
		return nil
	}

	cloned := make(M, len(items))
	for key, value := range items {
		cloned[key] = value
	}
	return cloned
}

// Copy copies all key/value pairs in src adding them to dst.
// When a key in src is already present in dst,
// the value in dst will be overwritten by the value associated
// with the key in src.
func Copy[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2) {
	for k, v := range src {
		dst[k] = v
	}
}

// DeleteFunc deletes any key/value pairs from m for which del returns true.
func DeleteFunc[M ~map[K]V, K comparable, V any](m M, del func(K, V) bool) {
	for k, v := range m {
		if del(k, v) {
			delete(m, k)
		}
	}
}

// Equal reports whether the two maps have the same key/value pairs.
func Equal[M1 ~map[K]V, M2 ~map[K]V, K, V comparable](m1 M1, m2 M2) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}

// EqualFunc is like Equal, but compares the values with the given function.
func EqualFunc[M1 ~map[K]V1, M2 ~map[K]V2, K comparable, V1, V2 any](m1 M1, m2 M2, eq func(V1, V2) bool) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || !eq(v1, v2) {
			return false
		}
	}
	return true
}
