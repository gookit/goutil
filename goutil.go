// Package goutil 💪 Useful utils for the Go: int, string, array/slice, map, error, time, format, CLI, ENV, filesystem,
// system, testing, debug and more.
package goutil

import (
	"fmt"

	"github.com/gookit/goutil/jsonutil"
	"github.com/gookit/goutil/stdutil"
)

// Value alias of stdutil.Value
type Value = stdutil.Value

// Go is a basic promise implementation: it wraps calls a function in a goroutine
// and returns a channel which will later return the function's return value.
func Go(f func() error) error {
	ch := make(chan error)
	go func() {
		ch <- f()
	}()
	return <-ch
}

// Filling src data(map,struct) to dst struct
func Filling(src, dst interface{}) error {
	return jsonutil.Mapping(src, dst)
}

// PanicIfErr if error is not empty
func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Panicf format panic message use fmt.Sprintf
func Panicf(format string, v ...interface{}) {
	panic(fmt.Sprintf(format, v...))
}

// FuncName get func name
func FuncName(f interface{}) string {
	return stdutil.FuncName(f)
}

// PkgName get current package name. alias of stdutil.PkgName()
//
// Usage:
//	funcName := goutil.FuncName(fn)
//	pgkName := goutil.PkgName(funcName)
func PkgName(funcName string) string {
	return stdutil.PkgName(funcName)
}

// GetCallStacks stacks is a wrapper for runtime.
// Stack that attempts to recover the data for all goroutines.
// from glog package
func GetCallStacks(all bool) []byte {
	return stdutil.GetCallStacks(all)
}

// GetCallersInfo returns an array of strings containing the file and line number
// of each stack frame leading
func GetCallersInfo(skip, max int) (callers []string) {
	return stdutil.GetCallersInfo(skip, max)
}
