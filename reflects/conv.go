package reflects

import (
	"reflect"

	"github.com/gookit/goutil/common"
	"github.com/gookit/goutil/internal/comfunc"
	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/strutil"
)

// BaseTypeVal convert custom type or intX,uintX,floatX to generic base type.
//
//	intX/unitX 	=> int64
//	floatX      => float64
//	string 	    => string
//
// returns int64,string,float or error
func BaseTypeVal(v reflect.Value) (value interface{}, err error) {
	v = reflect.Indirect(v)

	switch v.Kind() {
	case reflect.String:
		value = v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value = v.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		value = int64(v.Uint()) // always return int64
	case reflect.Float32, reflect.Float64:
		value = v.Float()
	default:
		err = common.ErrConvType
	}
	return
}

// ValueByType create reflect.Value by give reflect.Type
func ValueByType(val interface{}, typ reflect.Type) (rv reflect.Value, err error) {
	if typ.Kind() <= reflect.Float64 {
		return ValueByKind(val, typ.Kind())
	}

	// check type. like map, slice
	newRv := reflect.ValueOf(val)
	if newRv.Type() == typ {
		return newRv, nil
	}

	err = common.ErrConvType
	return
}

// ValueByKind create reflect.Value by give reflect.Kind
//
// TIPs:
//
//	Only support kind: string, bool, intX, uintX, floatX
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
	case reflect.Float32:
		if dstV, err1 := mathutil.ToFloat(val); err1 == nil {
			rv = reflect.ValueOf(float32(dstV))
		}
	case reflect.Float64:
		if dstV, err1 := mathutil.ToFloat(val); err1 == nil {
			rv = reflect.ValueOf(dstV)
		}
	case reflect.String:
		if dstV, err1 := strutil.ToString(val); err1 == nil {
			rv = reflect.ValueOf(dstV)
		}
	case reflect.Bool:
		if bl, err := comfunc.ToBool(val); err == nil {
			rv = reflect.ValueOf(bl)
		}
	}

	if !rv.IsValid() {
		err = common.ErrConvType
	}
	return
}
