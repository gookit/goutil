package maputil

import (
	"strings"

	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/strutil"
)

// Data an map data type
type Data map[string]interface{}

// TryGet value from the data map
func (d Data) TryGet(key string) (interface{}, bool) {
	val, ok := d[key]
	return val, ok
}

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

// Emtpy if the data map
func (d Data) Emtpy() bool {
	return len(d) == 0
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

// Bool value get
func (d Data) Bool(key string) bool {
	val, ok := d[key]
	if !ok {
		return false
	}
	if bl, ok := val.(bool); ok {
		return bl
	}

	str := strutil.MustString(val)
	return strutil.MustBool(str)
}

// Default get value from the data map with default value
func (d Data) Default(key string, def interface{}) interface{} {
	val, ok := d[key]
	if ok {
		return val
	}
	return def
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

	str := strutil.MustString(val)
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

// SMap is alias of map[string]string
type SMap map[string]string

// Has kay on the data map
func (m SMap) Has(key string) bool {
	_, ok := m[key]
	return ok
}

// HasValue on the data map
func (m SMap) HasValue(val string) bool {
	for _, v := range m {
		if v == val {
			return true
		}
	}
	return false
}

// Int value get
func (m SMap) Int(key string) int {
	val, ok := m[key]
	if !ok {
		return 0
	}
	return mathutil.MustInt(val)
}

// Int64 value get
func (m SMap) Int64(key string) int64 {
	val, ok := m[key]
	if !ok {
		return 0
	}
	return mathutil.MustInt64(val)
}

// Str value get
func (m SMap) Str(key string) string {
	return m[key]
}

// Bool value get
func (m SMap) Bool(key string) bool {
	val, ok := m[key]
	if !ok {
		return false
	}
	return strutil.MustBool(val)
}

// Ints value to []int
func (m SMap) Ints(key string) []int {
	val, ok := m[key]
	if !ok {
		return nil
	}
	return strutil.Ints(val, ",")
}

// Strings value to []string
func (m SMap) Strings(key string) (ss []string) {
	val, ok := m[key]
	if !ok {
		return
	}
	return strutil.ToSlice(val, ",")
}

// String data to string
func (m SMap) String() string {
	// return fmt.Sprint(m)
	return ""
}
