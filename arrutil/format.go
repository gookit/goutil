package arrutil

import (
	"bytes"
	"reflect"

	"github.com/gookit/goutil/strutil"
)

// ArrFormatter struct
type ArrFormatter struct {
	Buf *bytes.Buffer
	// Arr src array data
	Arr interface{}
	// MaxDepth limit TODO
	MaxDepth int
	// Prefix string for each element
	Prefix string
	// Indent string for format each element
	Indent string
	// ClosePrefix string for last "]"
	ClosePrefix string
	// AfterReset after reset on call Format().
	// AfterReset bool
}

// NewFormatter instance
func NewFormatter(arr interface{}) *ArrFormatter {
	return &ArrFormatter{
		Arr: arr,
	}
}

// Buffer get
func (f *ArrFormatter) Buffer() *bytes.Buffer {
	if f.Buf == nil {
		f.Buf = new(bytes.Buffer)
	}
	return f.Buf
}

// WithFn for config self
func (f *ArrFormatter) WithFn(fn func(f *ArrFormatter)) *ArrFormatter {
	fn(f)
	return f
}

// WithIndent string
func (f *ArrFormatter) WithIndent(indent string) *ArrFormatter {
	f.Indent = indent
	return f
}

// Reset after format
func (f *ArrFormatter) Reset() {
	f.Buf = nil
	f.Arr = nil
}

// Format to string
func (f *ArrFormatter) String() string {
	f.Format()
	return f.Buf.String()
}

// FormatTo to custom buffer
func (f *ArrFormatter) FormatTo(buf *bytes.Buffer) {
	f.Buf = buf
	f.Format()
}

// Format to string
func (f *ArrFormatter) Format() string {
	if f.Arr == nil {
		return ""
	}

	rv, ok := f.Arr.(reflect.Value)
	if !ok {
		rv = reflect.ValueOf(f.Arr)
	}

	rv = reflect.Indirect(rv)
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return ""
	}

	arrLn := rv.Len()
	if arrLn == 0 {
		return "[]"
	}

	// if f.AfterReset {
	// 	defer f.Reset()
	// }

	buf := f.Buffer()
	// sb.Grow(arrLn * 4)
	buf.WriteByte('[')

	indentLn := len(f.Indent)
	if indentLn > 0 {
		buf.WriteByte('\n')
	}

	for i := 0; i < arrLn; i++ {
		if indentLn > 0 {
			buf.WriteString(f.Indent)
		}
		buf.WriteString(strutil.QuietString(rv.Index(i).Interface()))

		if i < arrLn-1 {
			buf.WriteByte(',')

			// no indent, with space
			if indentLn == 0 {
				buf.WriteByte(' ')
			}
		}
		if indentLn > 0 {
			buf.WriteByte('\n')
		}
	}

	if f.ClosePrefix != "" {
		buf.WriteString(f.ClosePrefix)
	}

	buf.WriteByte(']')
	return buf.String()
}

// FormatIndent array data to string.
func FormatIndent(arr interface{}, indent string) string {
	return NewFormatter(arr).WithIndent(indent).String()
}
