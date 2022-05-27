package reflects

import "reflect"

// BKind base data kind type
type BKind uint

// base kinds
const (
	Invalid BKind = iota
	Bool
	Int
	Uint
	Float
	String
	Array
	Map
	Struct
	Complex
	Unsupported
)

// ToBaseKind convert reflect.Kind to base kind
func ToBaseKind(kind reflect.Kind) BKind {
	switch kind {
	case reflect.Bool:
		return Bool
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return Int
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return Uint
	case reflect.Float32, reflect.Float64:
		return Float
	case reflect.Complex64, reflect.Complex128:
		return Complex
	case reflect.String:
		return String
	case reflect.Array, reflect.Slice:
		return Array
	case reflect.Map:
		return Map
	case reflect.Struct:
		return Struct
	}

	// like: ptr, func, interface ...
	return Unsupported
}

// Type struct
type Type interface {
	reflect.Type
	// BaseKind value
	BaseKind() BKind
}

type xType struct {
	reflect.Type
	baseKind BKind
}

// TypeOf value
func TypeOf(v interface{}) Type {
	rftTyp := reflect.TypeOf(v)

	return &xType{
		Type:     rftTyp,
		baseKind: ToBaseKind(rftTyp.Kind()),
	}
}

// BaseKind value
func (t *xType) BaseKind() BKind {
	return t.baseKind
}
