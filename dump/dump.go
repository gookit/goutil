// Package dump like fmt.Println but more pretty and beautiful print Go values.
package dump

import (
	"io"
	"os"
)

// These flags define which print caller information
const (
	Fnopos = 1 << iota // no position
	Ffunc
	Ffile
	Ffname
	Fline
)

var (
	// valid flag for print caller info
	callerFlags = []int{Ffunc, Ffile, Ffname, Fline}

	// std dumper
	std = NewDumper(os.Stdout, 3)

	defaultTheme = map[string]string{
		"caller": "magenta",
		"key":    "green", // key name color of the map, struct.
		"value":  "normal",
	}
)

// Std dumper
func Std() *Dumper {
	return std
}

// Reset std dumper
func Reset() {
	std = NewDumper(os.Stdout, 3)
}

// Config std dumper
func Config(fn func(*Dumper)) {
	fn(std)
}

// Spew print
func Spew(vs ...interface{}) {
	std.Spew(vs...)
}

// V like fmt.Println, but the output is clearer and more beautiful
func V(vs ...interface{}) {
	std.Dump(vs...)
}

// P like fmt.Println, but the output is clearer and more beautiful
func P(vs ...interface{}) {
	std.Print(vs...)
}

// Print like fmt.Println, but the output is clearer and more beautiful
func Print(vs ...interface{}) {
	std.Print(vs...)
}

// Println like fmt.Println, but the output is clearer and more beautiful
func Println(vs ...interface{}) {
	std.Println(vs...)
}

// Fprint like fmt.Println, but the output is clearer and more beautiful
func Fprint(w io.Writer, vs ...interface{}) {
	std.Fprint(w, vs...)
}
