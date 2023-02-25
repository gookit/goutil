package cflag

import (
	"fmt"
	"strconv"

	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/comdef"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/strutil/textutil"
)

/*************************************************************************
 * options: some special flag vars
 * - implemented flag.Value interface
 *************************************************************************/

// Ints The int flag list, implemented flag.Value interface
type Ints []int

// String to string
func (s *Ints) String() string {
	return fmt.Sprintf("%v", *s)
}

// Set new value
func (s *Ints) Set(value string) error {
	intVal, err := strconv.Atoi(value)
	if err == nil {
		*s = append(*s, intVal)
	}

	return err
}

// Strings The string flag list, implemented flag.Value interface
type Strings []string

// String to string
func (s *Strings) String() string {
	return fmt.Sprintf("%v", *s)
}

// Set new value
func (s *Strings) Set(value string) error {
	*s = append(*s, value)
	return nil
}

// Booleans The bool flag list, implemented flag.Value interface
type Booleans []bool

// String to string
func (s *Booleans) String() string {
	return fmt.Sprintf("%v", *s)
}

// Set new value
func (s *Booleans) Set(value string) error {
	boolVal, err := strconv.ParseBool(value)
	if err == nil {
		*s = append(*s, boolVal)
	}

	return err
}

// EnumString The string flag list, implemented flag.Value interface
type EnumString struct {
	val  string
	enum []string
}

// NewEnumString instance
func NewEnumString(enum ...string) EnumString {
	return EnumString{enum: enum}
}

// String to string
func (s *EnumString) String() string {
	return s.val
}

// SetEnum values
func (s *EnumString) SetEnum(enum []string) {
	s.enum = enum
}

// WithEnum values
func (s *EnumString) WithEnum(enum []string) *EnumString {
	s.enum = enum
	return s
}

// Set new value, will check value is right
func (s *EnumString) Set(value string) error {
	s.val = value

	if !arrutil.InStrings(value, s.enum) {
		return fmt.Errorf("value must one of the: %v", s.enum)
	}
	return nil
}

// String a special string
//
// Usage:
//
//	// case 1:
//	var names gcli.String
//	c.VarOpt(&names, "names", "", "multi name by comma split")
//
//	--names "tom,john,joy"
//	names.Split(",") -> []string{"tom","john","joy"}
//
//	// case 2:
//	var ids gcli.String
//	c.VarOpt(&ids, "ids", "", "multi id by comma split")
//
//	--names "23,34,56"
//	names.Ints(",") -> []int{23,34,56}
type String string

// Set value
func (s *String) Set(val string) error {
	*s = String(val)
	return nil
}

// String to string
func (s *String) String() string {
	return string(*s)
}

// Strings split value to []string by sep ','
func (s *String) Strings() []string {
	return strutil.Split(string(*s), ",")
}

// Split value to []string
func (s *String) Split(sep string) []string {
	return strutil.Split(string(*s), sep)
}

// Ints value to []int
func (s *String) Ints(sep string) []int {
	return strutil.Ints(string(*s), sep)
}

// KVString The kv-string flag, allow input multi.
// Implemented the flag.Value interface.
//
// Example:
//
//	--var name=inhere => string map {name:inhere}
//	--var name=inhere --var age=234 => string map {name:inhere, age:234}
type KVString struct {
	maputil.SMap
	Sep string
}

// NewKVString instance
func NewKVString() KVString {
	return KVString{
		Sep:  comdef.EqualStr,
		SMap: make(maputil.SMap),
	}
}

// String to string
func (s *KVString) String() string {
	return s.SMap.String()
}

// SetData value
func (s *KVString) SetData(mp map[string]string) {
	s.SMap = mp
}

// Data map get
func (s *KVString) Data() maputil.SMap {
	return s.SMap
}

// Set new value, will check value is right
func (s *KVString) Set(value string) error {
	if value != "" {
		key, val := strutil.SplitKV(value, s.Sep)
		if key != "" {
			s.SMap[key] = val
		}
	}
	return nil
}

// ConfString The config-string flag, INI format, like nginx-config.
// Implemented the flag.Value interface.
//
// Example:
//
//	--config 'k0=val0;k1=val1' => string map {k0:val0, k1:val1}
type ConfString struct {
	maputil.SMap
	val string
}

// String to string
func (s *ConfString) String() string {
	return s.val
}

// SetData value
func (s *ConfString) SetData(mp map[string]string) {
	s.SMap = mp
}

// Data map get
func (s *ConfString) Data() maputil.SMap {
	return s.SMap
}

// Set new value, will check value is right
func (s *ConfString) Set(value string) error {
	if value != "" {
		s.val = value
		mp, err := textutil.ParseInlineINI(value)

		if err != nil {
			return err
		}
		s.SMap = mp
	}
	return nil
}
