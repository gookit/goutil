// Package comdef provide some common type or constant definitions
package comdef

import (
	"fmt"
	"io"
)

// ByteStringWriter interface
type ByteStringWriter interface {
	io.Writer
	io.ByteWriter
	io.StringWriter
	fmt.Stringer
}

// StringWriteStringer interface
type StringWriteStringer interface {
	io.StringWriter
	fmt.Stringer
}

type (
	// MarshalFunc define
	MarshalFunc func(v any) ([]byte, error)

	// UnmarshalFunc define
	UnmarshalFunc func(bts []byte, ptr any) error
)

// Int64able interface
type Int64able interface {
	Int64() (int64, error)
}

//
//
// Matcher type
//
//

// Matcher interface
type Matcher[T any] interface {
	Match(s T) bool
}

// MatchFunc definition. implements Matcher interface
type MatchFunc[T any] func(v T) bool

// Match satisfies the Matcher interface
func (fn MatchFunc[T]) Match(v T) bool {
	return fn(v)
}

// StringMatcher interface
type StringMatcher interface {
	Match(s string) bool
}

// StringMatchFunc definition
type StringMatchFunc func(s string) bool

// Match satisfies the StringMatcher interface
func (fn StringMatchFunc) Match(s string) bool {
	return fn(s)
}

// StringHandler interface
type StringHandler interface {
	Handle(s string) string
}

// StringHandleFunc definition
type StringHandleFunc func(s string) string

// Handle satisfies the StringHandler interface
func (fn StringHandleFunc) Handle(s string) string {
	return fn(s)
}

// IntCheckFunc check func
type IntCheckFunc func(val int) error

// StrCheckFunc check func
type StrCheckFunc func(val string) error
