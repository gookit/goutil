// Package stdutil provide some standard util functions for go.
package stdutil

import (
	"fmt"
	"runtime"
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

// GoVersion number get. eg: "1.18.2"
func GoVersion() string {
	return runtime.Version()[2:]
}
