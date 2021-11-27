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