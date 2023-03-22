package cflag

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/comdef"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/strutil/textutil"
)

// RepeatableFlag interface.
type RepeatableFlag interface {
	flag.Value
	// IsRepeatable mark option flag can be set multi times
	IsRepeatable() bool
}

/*************************************************************************
 * options: some special flag vars
 * - implemented flag.Value interface
 *************************************************************************/

// IntValue int, allow limit min and max value TODO
type IntValue struct {
	val string
	// validate
	Min, Max int
}

// IntsString The ints-string flag. eg: --get 1,2,3
//
// Implemented the flag.Value interface
type IntsString struct {
	val  string // input
	ints []int
	// value and size validate
	ValueFn func(val int) error
	SizeFn  func(ln int) error
}

// String input value to string
func (o *IntsString) String() string {
	return o.val
}

// Get value
func (o *IntsString) Get() any {
	return o.ints
}

// Ints value
func (o *IntsString) Ints() []int {
	return o.ints
}

// Set new value
func (o *IntsString) Set(value string) error {
	intVal, err := strconv.Atoi(value)
	if err != nil {
		return err
	}

	if o.ValueFn != nil {
		if err = o.ValueFn(intVal); err != nil {
			return err
		}
	}
	if o.SizeFn != nil {
		if err = o.SizeFn(len(o.ints) + 1); err != nil {
			return err
		}
	}

	o.ints = append(o.ints, intVal)
	return nil
}

// Ints The int flag list, repeatable
//
// implemented flag.Value interface
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

// IsRepeatable on input
func (s *Ints) IsRepeatable() bool {
	return true
}

// Strings The string flag list, repeatable
//
// implemented flag.Value interface
type Strings []string

// String input value to string
func (s *Strings) String() string {
	return strings.Join(*s, ",")
}

// Set new value
func (s *Strings) Set(value string) error {
	*s = append(*s, value)
	return nil
}

// IsRepeatable on input
func (s *Strings) IsRepeatable() bool {
	return true
}

// Booleans The bool flag list, repeatable
// implemented flag.Value interface
type Booleans []bool

// String input value to string
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

// IsRepeatable on input
func (s *Booleans) IsRepeatable() bool {
	return true
}

// EnumString The string flag list.
// implemented flag.Value interface
type EnumString struct {
	val  string
	enum []string
}

// NewEnumString instance
func NewEnumString(enum ...string) EnumString {
	return EnumString{enum: enum}
}

// Get value
func (s *EnumString) Get() any {
	return s.val
}

// String input value to string
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

// EnumString to string
func (s *EnumString) EnumString() string {
	return strings.Join(s.enum, ",")
}

// Set new value, will check value is right
func (s *EnumString) Set(value string) error {
	if !arrutil.InStrings(value, s.enum) {
		return fmt.Errorf("value must one of the: %v", s.enum)
	}

	s.val = value
	return nil
}

// Enum to string
func (s *EnumString) Enum() []string {
	return s.enum
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

// Get value
func (s *String) Get() any {
	return s
}

// Set value
func (s *String) Set(val string) error {
	*s = String(val)
	return nil
}

// String input value to string
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
//
// Implemented the flag.Value interface.
//
// Usage:
//
//		type myOpts struct {
//			vars cflag.KVString
//		}
//	 var mo &myOpts{ vars: cflag.NewKVString() }
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
	return *(&KVString{}).Init()
}

// Init settings
func (s *KVString) Init() *KVString {
	if s.Sep == "" {
		s.Sep = comdef.EqualStr
	}
	if s.SMap == nil {
		s.SMap = make(maputil.SMap)
	}
	return s
}

// Get value
func (s *KVString) Get() any {
	return s.SMap
}

// Data map get
func (s *KVString) Data() maputil.SMap {
	return s.SMap
}

// Set new value, will check value is right
func (s *KVString) Set(value string) error {
	if value != "" {
		s.Init()

		key, val := strutil.SplitKV(value, s.Sep)
		if key != "" {
			s.SMap[key] = val
		}
	}
	return nil
}

// IsRepeatable on input
func (s *KVString) IsRepeatable() bool {
	return true
}

// ConfString The config-string flag, INI format, like nginx-config.
//
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

// Get value
func (s *ConfString) Get() any {
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
