package reflects

import "reflect"

// Value struct
type Value struct {
	reflect.Value
	baseKind BKind
}

// Elem returns the value that the interface v contains
// or that the pointer v points to.
func Elem(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		return v.Elem()
	}

	// otherwise, will return self
	return v
}

// Wrap the give value
func Wrap(rv reflect.Value) Value {
	return Value{
		Value:    rv,
		baseKind: ToBKind(rv.Kind()),
	}
}

// ValueOf the give value
func ValueOf(v interface{}) Value {
	if rv, ok := v.(reflect.Value); ok {
		return Wrap(rv)
	}

	rv := reflect.ValueOf(v)
	return Value{
		Value:    rv,
		baseKind: ToBKind(rv.Kind()),
	}
}

// Indirect value
func (v Value) Indirect() Value {
	return v.Elem()
}

// Elem returns the value that the interface v contains or that the pointer v points to.
//
// TIP: not like reflect.Value.Elem. otherwise, will return self.
func (v Value) Elem() Value {
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		elem := v.Value.Elem()

		return Value{
			Value:    elem,
			baseKind: ToBKind(elem.Kind()),
		}
	}

	// otherwise, will return self
	return v
}

// Type of value.
func (v Value) Type() Type {
	return &xType{
		Type:     v.Value.Type(),
		baseKind: v.baseKind,
	}
}

// BaseKind value
func (v Value) BaseKind() BKind {
	return v.baseKind
}

// Int value. if is uintX will convert to int64
func (v Value) Int() int64 {
	switch v.baseKind {
	case Uint:
		return int64(v.Value.Uint())
	case Int:
		return v.Value.Int()
	}
	panic(&reflect.ValueError{Method: "reflect.Value.Int", Kind: v.Kind()})
}

// Uint value. if is intX will convert to uint64
func (v Value) Uint() uint64 {
	switch v.baseKind {
	case Uint:
		return v.Value.Uint()
	case Int:
		return uint64(v.Value.Int())
	}
	panic(&reflect.ValueError{Method: "reflect.Value.Uint", Kind: v.Kind()})
}
