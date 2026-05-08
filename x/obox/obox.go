// Package obox provides a lightweight dependency management container
// with generic support, named registration, and multiple lifetime modes.
package obox

import (
	"errors"
	"sync"
)

// Error definitions
var (
	ErrNotFound     = errors.New("obox: dependency not found")
	ErrTypeMismatch = errors.New("obox: type mismatch")
)

// Lifetime defines the lifecycle mode of a dependency
type Lifetime int

const (
	LifetimeSingleton Lifetime = iota // default, single instance
	LifetimeTransient                 // create new instance on each Get
	LifetimeLazy                      // lazy singleton, created on first Get
)

// entry stores the dependency information
type entry struct {
	value    any
	lifetime Lifetime
	factory  func() any
	created  bool // for Lazy mode, mark if already created
}

// Option configures the entry
type Option func(*entry)

// Transient returns an Option that sets Transient lifetime mode
func Transient() Option {
	return func(e *entry) {
		e.lifetime = LifetimeTransient
	}
}

// Lazy returns an Option that sets Lazy lifetime mode with a factory function
func Lazy(factory func() any) Option {
	return func(e *entry) {
		e.lifetime = LifetimeLazy
		e.factory = factory
	}
}

// TransientWithFactory returns an Option that sets Transient lifetime mode with a factory function
func TransientWithFactory(factory func() any) Option {
	return func(e *entry) {
		e.lifetime = LifetimeTransient
		e.factory = factory
	}
}

// container holds all registered dependencies
type container struct {
	mu      sync.RWMutex
	entries map[string]*entry
}

// defaultContainer is the global container for package-level functions
var defaultContainer = &container{
	entries: make(map[string]*entry),
}

// Set registers a dependency with the given name and value.
// Options can be used to configure lifetime mode.
func Set[T any](name string, val T, opts ...Option) {
	defaultContainer.mu.Lock()
	defer defaultContainer.mu.Unlock()

	e := &entry{
		value:    val,
		lifetime: LifetimeSingleton,
	}

	for _, opt := range opts {
		opt(e)
	}

	defaultContainer.entries[name] = e
}

// Get retrieves a dependency by name and type.
// Returns ErrNotFound if not found, ErrTypeMismatch if type doesn't match.
func Get[T any](name string) (T, error) {
	defaultContainer.mu.RLock()
	e, exists := defaultContainer.entries[name]
	defaultContainer.mu.RUnlock()

	if !exists {
		var zero T
		return zero, ErrNotFound
	}

	var val any

	switch e.lifetime {
	case LifetimeSingleton:
		val = e.value
	case LifetimeTransient:
		if e.factory != nil {
			val = e.factory()
		} else {
			val = e.value
		}
	case LifetimeLazy:
		defaultContainer.mu.Lock()
		if !e.created {
			if e.factory != nil {
				e.value = e.factory()
			}
			e.created = true
		}
		val = e.value
		defaultContainer.mu.Unlock()
	}

	result, ok := val.(T)
	if !ok {
		var zero T
		return zero, ErrTypeMismatch
	}

	return result, nil
}

// MustGet retrieves a dependency by name and type.
// Panics if not found or type doesn't match.
func MustGet[T any](name string) T {
	val, err := Get[T](name)
	if err != nil {
		panic(err)
	}
	return val
}

// Has checks if a dependency with the given name exists.
func Has(name string) bool {
	defaultContainer.mu.RLock()
	defer defaultContainer.mu.RUnlock()
	return defaultContainer.entries[name] != nil
}

// Delete removes a dependency by name.
func Delete(name string) {
	defaultContainer.mu.Lock()
	defer defaultContainer.mu.Unlock()
	delete(defaultContainer.entries, name)
}

// Reset clears all registered dependencies.
func Reset() {
	defaultContainer.mu.Lock()
	defer defaultContainer.mu.Unlock()
	defaultContainer.entries = make(map[string]*entry)
}
