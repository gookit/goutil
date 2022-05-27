package maputil

import (
	"bytes"
	"reflect"

	"github.com/gookit/goutil/strutil"
)

// MapFormatter struct
type MapFormatter struct {
	Buf *bytes.Buffer
	// Map source data
	Map interface{}
	// MaxDepth limit TODO
	MaxDepth int
	// Prefix string for each element
	Prefix string
	// Indent string for each element
	Indent string
	// ClosePrefix string for last "}"
	ClosePrefix string
	// AfterReset after reset on call Format().
	AfterReset bool
}

// NewFormatter instance
func NewFormatter(mp interface{}) *MapFormatter {
	return &MapFormatter{Map: mp}
}

// WithFn for config self
func (f *MapFormatter) WithFn(fn func(f *MapFormatter)) *MapFormatter {
	fn(f)
	return f
}

// WithIndent string
func (f *MapFormatter) WithIndent(indent string) *MapFormatter {
	f.Indent = indent
	return f
}

// Buffer get
func (f *MapFormatter) Buffer() *bytes.Buffer {
	if f.Buf == nil {
		f.Buf = new(bytes.Buffer)
	}
	return f.Buf
}

// Reset after format
func (f *MapFormatter) Reset() {
	f.Buf = nil
	f.Map = nil
}

// Format map data to string.
func (f *MapFormatter) Format() string {
	if f.Map == nil {
		return ""
	}

	rv, ok := f.Map.(reflect.Value)
	if !ok {
		rv = reflect.ValueOf(f.Map)
	}

	rv = reflect.Indirect(rv)
	if rv.Kind() != reflect.Map {
		return ""
	}

	ln := rv.Len()
	if ln == 0 {
		return "{}"
	}

	if f.AfterReset {
		defer f.Reset()
	}

	buf := f.Buffer()
	// buf.Grow(ln * 16)
	buf.WriteByte('{')

	indentLn := len(f.Indent)
	if indentLn > 0 {
		buf.WriteByte('\n')
	}

	for i, key := range rv.MapKeys() {
		kStr := strutil.QuietString(key.Interface())
		if indentLn > 0 {
			buf.WriteString(f.Indent)
		}

		buf.WriteString(kStr)
		buf.WriteByte(':')

		vStr := strutil.QuietString(rv.MapIndex(key).Interface())
		buf.WriteString(vStr)
		if i < ln-1 {
			buf.WriteByte(',')

			// no indent, with space
			if indentLn == 0 {
				buf.WriteByte(' ')
			}
		}

		// with newline
		if indentLn > 0 {
			buf.WriteByte('\n')
		}
	}

	if f.ClosePrefix != "" {
		buf.WriteString(f.ClosePrefix)
	}

	buf.WriteByte('}')
	return buf.String()
}

// FormatIndent map data to string.
func FormatIndent(mp interface{}, indent string) string {
	return NewFormatter(mp).WithIndent(indent).Format()
}
