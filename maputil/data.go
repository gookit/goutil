package maputil

import (
	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/strutil"
)

// Data an map data type
type Data map[string]interface{}

// Get value from the data map
func (d Data) Get(key string) interface{} {
	return d[key]
}

// Set value to the data map
func (d Data) Set(key string, val interface{}) {
	d[key] = val
}

// Has value on the data map
func (d Data) Has(key string) bool {
	_, ok := d[key]
	return ok
}

// Int value get
func (d Data) Int(key string) int {
	val, ok := d[key]
	if !ok {
		return 0
	}

	return mathutil.MustInt(val)
}

// Int64 value get
func (d Data) Int64(key string) int64 {
	val, ok := d[key]
	if !ok {
		return 0
	}

	return mathutil.MustInt64(val)
}

// Str value get by key
func (d Data) Str(key string) string {
	val, ok := d[key]
	if !ok {
		return ""
	}

	return strutil.MustString(val)
}

// Default get value from the data map with default value
func (d Data) Default(key string, def interface{}) interface{} {
	val, ok := d[key]
	if ok {
		return val
	}

	return def
}

// StringMap convert to map[string]string
func (d Data) StringMap() map[string]string {
	return ToStringMap(d)
}

// String data to string
func (d Data) String() string {
	// var buf []byte TODO
	// for k, v := range d {
	//
	// }

	return ""
}
