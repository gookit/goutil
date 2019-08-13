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

// Config dump data
var Config = struct {
	NoPosition bool
	ShowMethod bool
	ShowFile   bool
	NoColor    bool
}{
	ShowMethod: true,
}

// ResetConfig reset config data
func ResetConfig() {
	Config.NoColor = false
	Config.ShowFile = false
	Config.ShowMethod = true
	Config.NoPosition = false
}

// P like fmt.Println, but the output is clearer and more beautiful
func P(vs ...interface{}) {
	Fprint(2, os.Stdout, vs...)
}

// V like fmt.Println, but the output is clearer and more beautiful
func V(vs ...interface{}) {
	Fprint(2, os.Stdout, vs...)
}

// Print like fmt.Println, but the output is clearer and more beautiful
func Print(vs ...interface{}) {
	Fprint(2, os.Stdout, vs...)
}

// Println like fmt.Println, but the output is clearer and more beautiful
func Println(vs ...interface{}) {
	Fprint(2, os.Stdout, vs...)
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
	rValue := reflect.ValueOf(v)
	rType := rValue.Type()

	switch rType.Kind() {
	case reflect.Slice, reflect.Array:
		eleNum := rValue.Len()

		if eleNum < 10 {
			mustFprintf(w, "%#v\n", v)
			return
		}

		mustFprint(w, rType.String(), " [\n")
		for i := 0; i < eleNum; i++ {
			mustFprintf(w, "  %v,\n", rValue.Index(i).Interface())
		}
		mustFprint(w, "]\n")
	case reflect.Struct:
		fldNum := rValue.NumField()
		if fldNum < 8 {
			mustFprintf(w, "%#v\n", v)
			return
		}

		mustFprint(w, rType.String(), " {\n")
		for i := 0; i < fldNum; i++ {
			tn := rType.Field(i).Name
			fv := rValue.Field(i)

			if fv.CanInterface() {
				mustFprintf(w, "  %s: %#v\n", tn, rValue.Field(i).Interface())
			} else {
				mustFprintf(w, "  %s: %#v\n", tn, rValue.Field(i).String())
			}
		}
		mustFprint(w, "}\n")
	case reflect.Map:
		mustFprint(w, rType.String(), " {\n")

		for _, key := range rValue.MapKeys() {
			mustFprintf(w, "  %v: %#v\n", key.Interface(), rValue.MapIndex(key).Interface())
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
