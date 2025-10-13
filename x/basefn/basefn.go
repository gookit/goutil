// Package basefn provide some no-dependents util functions
package basefn

import (
	"errors"
	"fmt"

	"github.com/gookit/goutil/internal/comfunc"
)

// Panicf format panic message use fmt.Sprintf
func Panicf(format string, v ...any) { panic(fmt.Sprintf(format, v...)) }

// PanicIf if cond = true, panics with an error message
func PanicIf(cond bool, fmtAndArgs ...any) {
	if cond {
		panic(errors.New(comfunc.FormatWithArgs(fmtAndArgs)))
	}
}

// PanicErr panics if error is not empty
func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

// MustOK if error is not empty, will panic
func MustOK(err error) {
	if err != nil {
		panic(err)
	}
}

// Must return like (v, error). will panic on error, otherwise return v.
//
// Usage:
//
//	// old
//	v, err := fn()
//	if err != nil {
//		panic(err)
//	}
//
//	// new
//	v := goutil.Must(fn())
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// MustIgnore for return like (v, error). Ignore return v and will panic on error.
//
// Useful for io, file operation func: (n int, err error)
//
// Usage:
//
//	// old
//	_, err := fn()
//	if err != nil {
//		panic(err)
//	}
//
//	// new
//	basefn.MustIgnore(fn())
func MustIgnore(_ any, err error) { PanicErr(err) }

// ErrOnFail return input error on cond is false, otherwise return nil
func ErrOnFail(cond bool, err error) error {
	return OrError(cond, err)
}

// OrError return input error on cond is false, otherwise return nil
func OrError(cond bool, err error) error {
	if !cond {
		return err
	}
	return nil
}

// FirstOr get first elem or elseVal
func FirstOr[T any](sl []T, elseVal T) T {
	if len(sl) > 0 {
		return sl[0]
	}
	return elseVal
}

// OrValue get. like: if cond { okVal } else { elVal }
func OrValue[T any](cond bool, okVal, elVal T) T {
	if cond {
		return okVal
	}
	return elVal
}

// OrReturn call okFunc() on condition is true, else call elseFn()
//
// like expr: if cond { okFunc() } else { elseFn() }
func OrReturn[T any](cond bool, okFn, elseFn func() T) T {
	if cond {
		return okFn()
	}
	return elseFn()
}

// ErrFunc type
type ErrFunc func() error

// CallOn call func on condition is true
func CallOn(cond bool, fn ErrFunc) error {
	if cond {
		return fn()
	}
	return nil
}

// CallOrElse call okFunc() on condition is true, else call elseFn()
func CallOrElse(cond bool, okFn, elseFn ErrFunc) error {
	if cond {
		return okFn()
	}
	return elseFn()
}
