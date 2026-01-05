// Package syncs provides synchronization primitives util functions.
package syncs

import (
	"context"
	"sync"
)

// WaitGroup is a wrapper of sync.WaitGroup.
//
// Usage:
//
// 	wg := syncs.WaitGroup{}
//	wg.Go(func() {
// 		time.Sleep(time.Second)
//	})
// 	wg.Wait()
//
type WaitGroup struct {
	sync.WaitGroup
}

// Go runs the given function in a new goroutine. will auto call Add and Done.
func (wg *WaitGroup) Go(fn func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fn()
	}()
}

// ContextValue create a new context with given value
func ContextValue(key, value any) context.Context {
	return context.WithValue(context.Background(), key, value)
}

// SafeMap is a goroutine-safe map.
type SafeMap[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

// NewSafeMap create a new SafeMap instance.
func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		data: make(map[K]V),
	}
}

// Set value to map.
func (m *SafeMap[K, V]) Set(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

// Get value from map.
func (m *SafeMap[K, V]) Get(key K) (value V, ok bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, ok = m.data[key]
	return value, ok
}

// Delete value from map.
func (m *SafeMap[K, V]) Delete(key K) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
}

// Range iterate map values.
func (m *SafeMap[K, V]) Range(fn func(key K, value V)) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for key, value := range m.data {
		fn(key, value)
	}
}

// Clear do clear all map values.
func (m *SafeMap[K, V]) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data = make(map[K]V)
}

// Len get map length.
func (m *SafeMap[K, V]) Len() int {
	return len(m.data)
}
