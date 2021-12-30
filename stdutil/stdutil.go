package stdutil

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

// PanicIfErr if error is not empty
func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

// PanicIf if error is not empty
func PanicIf(err error) {
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
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// PkgName get current package name
//
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
