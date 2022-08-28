package goutil

import (
	"errors"
	"reflect"

	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/reflects"
	"github.com/gookit/goutil/strutil"
)

// ErrConvType error
var ErrConvType = errors.New("convert value type error")

// Bool convert value to bool
func Bool(v interface{}) bool {
	bl, _ := ToBool(v)
	return bl
}

// ToBool try to convert type to bool
func ToBool(v interface{}) (bool, error) {
	if bl, ok := v.(bool); ok {
		return bl, nil
	}

	if str, ok := v.(string); ok {
		return strutil.ToBool(str)
	}
	return false, ErrConvType
}

// String always convert value to string, will ignore error
func String(v interface{}) string {
	s, _ := strutil.AnyToString(v, false)
	return s
}

// ToString convert value to string, will return error on fail.
func ToString(v interface{}) (string, error) {
	return strutil.AnyToString(v, true)
}

// Int convert value to int
func Int(v interface{}) int {
	iv, _ := mathutil.ToInt(v)
	return iv
}

// ToInt try to convert value to int
func ToInt(v interface{}) (int, error) {
	return mathutil.ToInt(v)
}

// Int64 convert value to int64
func Int64(v interface{}) int64 {
	iv, _ := mathutil.ToInt64(v)
	return iv
}

// ToInt64 try to convert value to int64
func ToInt64(v interface{}) (int64, error) {
	return mathutil.ToInt64(v)
}

// Uint convert value to uint64
func Uint(v interface{}) uint64 {
	iv, _ := mathutil.ToUint(v)
	return iv
}

// ToUint try to convert value to uint64
func ToUint(v interface{}) (uint64, error) {
	return mathutil.ToUint(v)
}

// BaseTypeVal convert custom type or intX,uintX,floatX to generic base type.
//
//	intX/unitX 	=> int64
//	floatX      => float64
//	string 	    => string
//
// returns int64,string,float or error
func BaseTypeVal(val interface{}) (value interface{}, err error) {
	return reflects.BaseTypeVal(reflect.ValueOf(val))
}
