package optional

var empty = &optional{v: nil}

// optional struct
type optional struct {
	v interface{}
}

func Of(data interface{}) *optional {
	return &optional{v: data}
}

func OfNillable(data interface{}) *optional {
	if data == nil {
		return empty
	}

	return &optional{v: data}
}

func (o *optional) Map(fn func(v interface{}) interface{}) *optional {
	if o.v == nil {
		return empty
	}

	return OfNillable(fn(o.v))
}

func (o *optional) Get() interface{} {
	if o.v == nil {
		panic("nil value")
	}

	return o.v
}

func (o *optional) OrElse(v interface{}) interface{} {
	if o.v == nil {
		return v
	}

	return o.v
}

func (o *optional) OrElseGet(v interface{}) interface{} {
	if o.v == nil {
		return v
	}

	return o.v
}
