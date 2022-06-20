package maputil

import (
	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/strutil"
)

// SMap is alias of map[string]string
type SMap map[string]string

// IsEmpty of the data map
func (m SMap) IsEmpty() bool {
	return len(m) == 0
}

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

// Value get from the data map
func (m SMap) Value(key string) (string, bool) {
	val, ok := m[key]
	return val, ok
}

// Int value get
func (m SMap) Int(key string) int {
	val, ok := m[key]
	if !ok {
		return 0
	}
	return mathutil.QuietInt(val)
}

// Int64 value get
func (m SMap) Int64(key string) int64 {
	val, ok := m[key]
	if !ok {
		return 0
	}
	return mathutil.QuietInt64(val)
}

// Get value by key
func (m SMap) Get(key string) string {
	return m[key]
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
	return strutil.QuietBool(val)
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
	return ToString2(m)
}
