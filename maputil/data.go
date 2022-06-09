package maputil

import (
	"strings"

	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/strutil"
)

// Data an map data type
type Data map[string]interface{}

// Set value to the data map
func (d Data) Set(key string, val interface{}) {
	d[key] = val
}

// Has value on the data map
func (d Data) Has(key string) bool {
	_, ok := d[key]
	return ok
}

// Emtpy if the data map
func (d Data) Emtpy() bool {
	return len(d) == 0
}

// Get value from the data map
func (d Data) Get(key string) interface{} {
	return d[key]
}

// Value get from the data map
func (d Data) Value(key string) (interface{}, bool) {
	val, ok := d[key]
	return val, ok
}

// GetByPath get value from the data map by path. eg: top.sub
func (d Data) GetByPath(path string) (interface{}, bool) {
	return GetByPath(path, d)
}

// Default get value from the data map with default value
func (d Data) Default(key string, def interface{}) interface{} {
	val, ok := d[key]
	if ok {
		return val
	}
	return def
}

// Int value get
func (d Data) Int(key string) int {
	val, ok := d[key]
	if !ok {
		return 0
	}

	return mathutil.QuietInt(val)
}

// Int64 value get
func (d Data) Int64(key string) int64 {
	val, ok := d[key]
	if !ok {
		return 0
	}

	return mathutil.QuietInt64(val)
}

// Str value get by key
func (d Data) Str(key string) string {
	val, ok := d[key]
	if !ok {
		return ""
	}
	return strutil.QuietString(val)
}

// Bool value get
func (d Data) Bool(key string) bool {
	val, ok := d[key]
	if !ok {
		return false
	}
	if bl, ok := val.(bool); ok {
		return bl
	}

	if str, ok := val.(string); ok {
		return strutil.QuietBool(str)
	}
	return false
}

// Strings get []string value
func (d Data) Strings(key string) []string {
	val, ok := d[key]
	if !ok {
		return nil
	}

	if ss, ok := val.([]string); ok {
		return ss
	}
	return nil
}

// StringsByStr value get by key
func (d Data) StringsByStr(key string) []string {
	val, ok := d[key]
	if !ok {
		return nil
	}

	str := strutil.QuietString(val)
	return strings.Split(str, ",")
}

// StringMap get map[string]string value
func (d Data) StringMap(key string) map[string]string {
	val, ok := d[key]
	if !ok {
		return nil
	}

	if smp, ok := val.(map[string]string); ok {
		return smp
	}
	return nil
}

// ToStringMap convert to map[string]string
func (d Data) ToStringMap() map[string]string {
	return ToStringMap(d)
}

// String data to string
func (d Data) String() string {
	return ToString(d)
}
