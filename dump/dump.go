// Package dump like fmt.Println but more pretty and beautiful print Go values.
package dump

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/gookit/color"
)

// These flags define which print caller information
const (
	Fnopos = 1 << iota // no position
	Ffunc
	Ffile
	Ffname
	Fline
)

const defaultSkip = 3

var (
	// valid flag for print caller info
	callerFlags = []int{Ffunc, Ffile, Ffname, Fline}
	// default theme
	defaultTheme = Theme{
		"caller": "magenta",
		"field":  "green", // field name color of the map, struct.
		"value":  "normal",
		// special type
		"msType":  "green", // for keywords map, struct type
		"valTip":  "gray",  // tips comments for string, slice, map len
		"string":  "green",
		"integer": "lightBlue",
	}

	// std dumper
	std = NewDumper(os.Stdout, defaultSkip)
	// no location dumper.
	std2 = NewWithOptions(func(opts *Options) {
		opts.Output = os.Stdout
		opts.ShowFlag = Fnopos
	})

	// some type init
	stringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
)

// Theme color code/tag map for dump
type Theme map[string]string

func (ct Theme) caller(s string) string  { return ct.wrap("caller", s) }
func (ct Theme) field(s string) string   { return ct.wrap("field", s) }
func (ct Theme) value(s string) string   { return ct.wrap("value", s) }
func (ct Theme) msType(s string) string  { return ct.wrap("msType", s) }
func (ct Theme) valTip(s string) string  { return ct.wrap("valTip", s) }
func (ct Theme) string(s string) string  { return ct.wrap("string", s) }
func (ct Theme) integer(s string) string { return ct.wrap("integer", s) }

// wrap color tag.
func (ct Theme) wrap(key string, s string) string {
	if tag := ct[key]; tag != "" {
		return color.WrapTag(s, tag)
	}
	return s
}

// Std dumper
func Std() *Dumper { return std }

// Reset std dumper
func Reset() { std = NewDumper(os.Stdout, 3) }

// Config std dumper
func Config(fns ...OptionFunc) { std.WithOptions(fns...) }

// V like fmt.Println, but the output is clearer and more beautiful
func V(vs ...any) {
	std.Dump(vs...)
}

// P like fmt.Println, but the output is clearer and more beautiful
func P(vs ...any) {
	std.Print(vs...)
}

// Print like fmt.Println, but the output is clearer and more beautiful
func Print(vs ...any) {
	std.Print(vs...)
}

// Println like fmt.Println, but the output is clearer and more beautiful
func Println(vs ...any) {
	std.Println(vs...)
}

// Fprint like fmt.Println, but the output is clearer and more beautiful
func Fprint(w io.Writer, vs ...any) {
	std.Fprint(w, vs...)
}

// Std2 dumper
func Std2() *Dumper { return std2 }

// Reset2 reset std2 dumper
func Reset2() {
	std2 = NewWithOptions(func(opts *Options) {
		opts.Output = os.Stdout
		opts.ShowFlag = Fnopos
	})
}

// Format like fmt.Println, but the output is clearer and more beautiful
func Format(vs ...any) string {
	w := &bytes.Buffer{}

	std2.Fprint(w, vs...)
	return w.String()
}

// NoLoc dump vars data, without location.
func NoLoc(vs ...any) {
	std2.Println(vs...)
}

// Clear dump clear data, without location.
func Clear(vs ...any) {
	std2.Println(vs...)
}

// is unexported field name on struct
func isUnexported(fieldName string) bool {
	return fieldName[0] < 'A' || fieldName[0] > 'Z'
}

func isNilOrInvalid(v reflect.Value) bool {
	if !v.IsValid() {
		return true
	}

	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}
