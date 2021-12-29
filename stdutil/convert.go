package stdutil

import "github.com/gookit/goutil/strutil"

// ToString always convert value to string
func ToString(v interface{}) string {
	s, _ := strutil.AnyToString(v, false)
	return s
}

// MustString convert value(basic type) to string, will panic on convert a complex type.
func MustString(v interface{}) string {
	s, err := strutil.AnyToString(v, true)
	if err != nil {
		panic(err)
	}
	return s
}

// TryString try to convert a value to string
func TryString(v interface{}) (string, error) {
	return strutil.AnyToString(v, true)
}
