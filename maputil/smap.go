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

// Has key on the data map
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

// Load data to the map
func (m SMap) Load(data map[string]string) {
	for k, v := range data {
		m[k] = v
	}
}

// Set value to the data map
func (m SMap) Set(key string, val any) {
	m[key] = strutil.MustString(val)
}

// Value get from the data map
func (m SMap) Value(key string) (string, bool) {
	val, ok := m[key]
	return val, ok
}

// Default get value by key. if not found, return defVal
func (m SMap) Default(key, defVal string) string {
	if val, ok := m[key]; ok {
		return val
	}
	return defVal
}

// Get value by key
func (m SMap) Get(key string) string {
	return m[key]
}

// Int value get
func (m SMap) Int(key string) int {
	if val, ok := m[key]; ok {
		return mathutil.QuietInt(val)
	}
	return 0
}

// Int64 value get
func (m SMap) Int64(key string) int64 {
	if val, ok := m[key]; ok {
		return mathutil.QuietInt64(val)
	}
	return 0
}

// Str value get
func (m SMap) Str(key string) string {
	return m[key]
}

// Bool value get
func (m SMap) Bool(key string) bool {
	if val, ok := m[key]; ok {
		return strutil.QuietBool(val)
	}
	return false
}

// Ints value to []int
func (m SMap) Ints(key string) []int {
	if val, ok := m[key]; ok {
		return strutil.Ints(val, ValSepStr)
	}
	return nil
}

// Strings value to []string
func (m SMap) Strings(key string) (ss []string) {
	if val, ok := m[key]; ok {
		return strutil.ToSlice(val, ValSepStr)
	}
	return
}

// IfExist key, then call the fn with value.
func (m SMap) IfExist(key string, fn func(val string)) {
	if val, ok := m[key]; ok {
		fn(val)
	}
}

// IfValid value is not empty, then call the fn
func (m SMap) IfValid(key string, fn func(val string)) {
	if val, ok := m[key]; ok && val != "" {
		fn(val)
	}
}

// Keys of the string-map
func (m SMap) Keys() []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values of the string-map
func (m SMap) Values() []string {
	ss := make([]string, 0, len(m))
	for _, v := range m {
		ss = append(ss, v)
	}
	return ss
}

// ToKVPairs slice convert. eg: {k1:v1,k2:v2} => {k1,v1,k2,v2}
func (m SMap) ToKVPairs() []string {
	pairs := make([]string, 0, len(m)*2)
	for k, v := range m {
		pairs = append(pairs, k, v)
	}
	return pairs
}

// String data to string
func (m SMap) String() string {
	return ToString2(m)
}
