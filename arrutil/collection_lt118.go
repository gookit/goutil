//go:build !go1.18
// +build !go1.18

package arrutil

// Map an object list [object0{},object1{},...] to flatten list [object0.someKey, object1.someKey, ...]
func Map(data any, mapFn func(v any) any) []any {
	panic("please upgrade to go 1.18+")
}
