package reflects

import (
	"errors"
	"reflect"

	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/strutil"
)

// ErrConvertFail error define
var ErrConvertFail = errors.New("convert value type is failure")

// BaseTypeVal convert custom type or intX,uintX,floatX to generic base type.
//
//	intX/unitX 	=> int64
//	floatX      => float64
//	string 	    => string
//
// returns int64,string,float or error
func BaseTypeVal(v reflect.Value) (value interface{}, err error) {
	switch v.Kind() {
	case reflect.String:
		value = v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value = v.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value = int64(v.Uint()) // always return int64
	case reflect.Float32, reflect.Float64:
		value = v.Float()
	default:
		err = ErrConvertFail
	}
	return
}

// ValueByKind reflect value create by give kind
func ValueByKind(val interface{}, kind reflect.Kind) (rv reflect.Value, err error) {
	switch kind {
	case reflect.Int:
		if dstV, err1 := mathutil.ToInt(val); err1 == nil {
			rv = reflect.ValueOf(dstV)
		}
	case reflect.Int8:
		if dstV, err1 := mathutil.ToInt(val); err1 == nil {
			rv = reflect.ValueOf(int8(dstV))
		}
	case reflect.Int16:
		if dstV, err1 := mathutil.ToInt(val); err1 == nil {
			rv = reflect.ValueOf(int16(dstV))
		}
	case reflect.Int32:
		if dstV, err1 := mathutil.ToInt(val); err1 == nil {
			rv = reflect.ValueOf(int32(dstV))
		}
	case reflect.Int64:
		if dstV, err1 := mathutil.ToInt64(val); err1 == nil {
			rv = reflect.ValueOf(dstV)
		}
	case reflect.Uint:
		if dstV, err1 := mathutil.ToUint(val); err1 == nil {
			rv = reflect.ValueOf(uint(dstV))
		}
	case reflect.Uint8:
		if dstV, err1 := mathutil.ToUint(val); err1 == nil {
			rv = reflect.ValueOf(uint8(dstV))
		}
	case reflect.Uint16:
		if dstV, err1 := mathutil.ToUint(val); err1 == nil {
			rv = reflect.ValueOf(uint16(dstV))
		}
	case reflect.Uint32:
		if dstV, err1 := mathutil.ToUint(val); err1 == nil {
			rv = reflect.ValueOf(uint32(dstV))
		}
	case reflect.Uint64:
		if dstV, err1 := mathutil.ToUint(val); err1 == nil {
			rv = reflect.ValueOf(dstV)
		}
	case reflect.String:
		if dstV, err1 := strutil.ToString(val); err1 == nil {
			rv = reflect.ValueOf(dstV)
		}
	case reflect.Bool:
		if bl, ok := val.(bool); ok {
			rv = reflect.ValueOf(bl)
		} else if str, ok := val.(string); ok {
			if dstV, err1 := strutil.ToBool(str); err1 == nil {
				rv = reflect.ValueOf(dstV)
			}
		}
		// TODO ... more kind supports: slice
	}

	if !rv.IsValid() {
		err = ErrConvertFail
	}
	return
}
