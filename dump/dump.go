// Package dump like fmt.Println but more clear and beautiful print data.
package dump

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"

	"github.com/gookit/color"
)

type dumpConfig struct {
	ShowMethod bool
	ShowFile   bool
	NoPosition bool
	NoColor    bool
	// MoreLenNL array/slice elements length > MoreLenNL, will wrap new line
	MoreLenNL int
}

// Output set print content to the io.Writer
var Output io.Writer = os.Stdout

// Config dump data settings
var Config = dumpConfig{
	ShowMethod: true,
	MoreLenNL:  8,
}

// ResetConfig reset config data
func ResetConfig() {
	Config = dumpConfig{
		ShowMethod: true,
		MoreLenNL:  8,
	}
}

// P like fmt.Println, but the output is clearer and more beautiful
func P(vs ...interface{}) {
	Fprint(2, Output, vs...)
}

// V like fmt.Println, but the output is clearer and more beautiful
func V(vs ...interface{}) {
	Fprint(2, Output, vs...)
}

// Print like fmt.Println, but the output is clearer and more beautiful
func Print(vs ...interface{}) {
	Fprint(2, Output, vs...)
}

// Println like fmt.Println, but the output is clearer and more beautiful
func Println(vs ...interface{}) {
	Fprint(2, Output, vs...)
}

// Fprint like fmt.Println, but the output is clearer and more beautiful
func Fprint(skip int, w io.Writer, vs ...interface{}) {
	// show print position
	if !Config.NoPosition {
		// get the print position
		pc, file, line, ok := runtime.Caller(skip)
		if ok {
			printPosition(w, pc, file, line)
		}
	}

	// print data
	for _, v := range vs {
		printOne(w, v)
	}
}

func printPosition(w io.Writer, pc uintptr, file string, line int) {
	var text string
	fnName := runtime.FuncForPC(pc).Name()

	if Config.ShowFile {
		text = fmt.Sprint("PRINT AT ", fnName, "(", file, " LINE ", line, "):")
	} else {
		text = fmt.Sprint("PRINT AT ", fnName, "(LINE ", line, "):")
	}

	if Config.NoColor {
		mustFprint(w, text, "\n")
		return
	}

	color.Fprint(w, "<mga>", text, "</>\n")
}

func printOne(w io.Writer, v interface{}) {
	if v == nil {
		mustFprintf(w, "<nil>\n")
		return
	}

	rVal := reflect.ValueOf(v)
	rType := rVal.Type()

	switch rType.Kind() {
	case reflect.Slice, reflect.Array:
		eleNum := rVal.Len()
		if eleNum < Config.MoreLenNL {
			mustFprintf(w, "%#v\n", v)
			return
		}

		mustFprint(w, rType.String(), " [\n")
		for i := 0; i < eleNum; i++ {
			mustFprintf(w, "  %v,\n", rVal.Index(i).Interface())
		}
		mustFprint(w, "]\n")
	case reflect.Struct:
		fldNum := rVal.NumField()

		mustFprint(w, rType.String(), " {\n")
		for i := 0; i < fldNum; i++ {
			tn := rType.Field(i).Name
			fv := rVal.Field(i)

			if fv.CanInterface() {
				mustFprintf(w, "  %s: %#v,\n", tn, rVal.Field(i).Interface())
			} else {
				mustFprintf(w, "  %s: %#v,\n", tn, rVal.Field(i).String())
			}
		}
		mustFprint(w, "}\n")
	case reflect.Map:
		mustFprint(w, rType.String(), " {\n")

		for _, key := range rVal.MapKeys() {
			mustFprintf(w, "  %v: %#v,\n", key.Interface(), rVal.MapIndex(key).Interface())
		}

		mustFprint(w, "}\n")
	default:
		mustFprintf(w, "%s(%v)\n", rType.String(), v)
	}
}

func mustFprint(w io.Writer, v ...interface{}) {
	_, err := fmt.Fprint(w, v...)
	if err != nil {
		panic(err)
	}
}

func mustFprintf(w io.Writer, f string, v ...interface{}) {
	_, err := fmt.Fprintf(w, f, v...)
	if err != nil {
		panic(err)
	}
}
