// Package dump like fmt.Println but more clear and beautiful print data.
package dump

import (
	"fmt"
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

func mustFprint(w io.Writer, v ...interface{}) {
	_, _ = fmt.Fprint(w, v...)
	// color.Fprint(w, v...)
}

func mustFprintf(w io.Writer, f string, v ...interface{}) {
	_, _ = fmt.Fprintf(w, f, v...)
	// color.Fprintf(w, f, v...)
}
