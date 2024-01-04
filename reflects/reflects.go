// Package reflects Provide extends reflect util functions.
package reflects

import (
	"reflect"
)

var emptyValue = reflect.Value{}

var (
	anyType   = reflect.TypeOf((*any)(nil)).Elem()
	errorType = reflect.TypeOf((*error)(nil)).Elem()

	// fmtStringerType  = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	reflectValueType = reflect.TypeOf((*reflect.Value)(nil)).Elem()
)
