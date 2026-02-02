// Package lcache provides a simple, thread-safe local cache implementation with TTL support.
//
// Quickly usage:
//
//	lcache.Set("key", "value", 5*time.Minute)
//
//	val, found := lcache.Get("key")
//	if found {
//	    fmt.Println(val)
//	}
//
// Custom configuration:
//
//	cache := lcache.New(
// 		lcache.WithCapacity(10),
//	)
package lcache

import (
	"time"
)

// std 默认的全局缓存实例
var std = New()

// Set value by key with TTL
func Set[T any](key string, val T, ttl time.Duration) {
	std.Set(key, val, ttl)
}

// Get value by key, return zero value if not found
func Get[T any](key string) (T, bool) {
	var zero T // 零值
	val, ok := std.Get(key)
	if !ok {
		return zero, false
	}

	// 类型断言
	res, ok := val.(T)
	if !ok {
		return zero, false
	}
	return res, true
}

// Keys get the keys of the default cache
func Keys() []string {
	return std.Keys()
}

// Len get the number of items in the cache
func Len() int {
	return std.Len()
}

// Clear all items from the default cache
func Clear() {
	std.Clear()
}

// Delete key
func Delete(key string) {
	std.Delete(key)
}

// SaveFile Save the cache data to a file.
func SaveFile(filename string) error {
	return std.SaveFile(filename)
}

// LoadFile Recover cache data from file load
func LoadFile(filename string) error {
	return std.LoadFile(filename)
}
