// Package goutil Useful utils for the Go: int, string, array/slice, map, format, cli, env, filesystem, testing and more.
package goutil

import (
	"reflect"
	"runtime"
	"strings"
)

// FuncName get func name
func FuncName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// PkgName get current package name
// Usage:
//	funcName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
//	pgkName := goutil.PkgName(funcName)
func PkgName(funcName string) string {
	for {
		lastPeriod := strings.LastIndex(funcName, ".")
		lastSlash := strings.LastIndex(funcName, "/")
		if lastPeriod > lastSlash {
			funcName = funcName[:lastPeriod]
		} else {
			break
		}
	}

	return funcName
}

// GetCallStacks stacks is a wrapper for runtime.
// Stack that attempts to recover the data for all goroutines.
// from glog package
func GetCallStacks(all bool) []byte {
	// We don't know how big the traces are, so grow a few times if they don't fit.
	// Start large, though.
	n := 10000
	if all {
		n = 100000
	}

	var trace []byte
	for i := 0; i < 5; i++ {
		trace = make([]byte, n)
		bts := runtime.Stack(trace, all)
		if bts < len(trace) {
			return trace[:bts]
		}
		n *= 2
	}
	return trace
}

// PanicIfErr if error is not empty
func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
