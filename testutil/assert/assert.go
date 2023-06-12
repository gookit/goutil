// Package assert Provides commonly asserts functions for help write Go testing.
//
// inspired the package: github.com/stretchr/testify/assert
package assert

// TestingT is an interface wrapper around *testing.T
type TestingT interface {
	Helper()
	Name() string
	Error(args ...any)
}
