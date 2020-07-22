// Package dump like fmt.Println but more clear and beautiful print data.
package dump

import (
	"fmt"
	"io"
	"os"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/gookit/color"
)

// These flags define which print information
const (
	Fnopos = 1 << iota // no position
	Ffunc
	Ffile
	Ffname
	Fline
)

type dumpConfig struct {
	NoColor  bool
	ShowFlag int
	// MoreLenNL array/slice elements length > MoreLenNL, will wrap new line
	MoreLenNL int
}

var flags = []int{Ffunc, Ffile, Ffname, Fline}

// Output set print content to the io.Writer
var Output io.Writer = os.Stdout

// Config dump data settings
var Config = newDefaultConfig()

// ResetConfig reset config data
func ResetConfig() {
	Config = newDefaultConfig()
}

func newDefaultConfig() dumpConfig {
	return dumpConfig{
		ShowFlag:  Ffunc | Ffname | Fline,
		MoreLenNL: 8,
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
	if Config.ShowFlag != Fnopos {
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
	// eg: github.com/gookit/goutil/dump.ExamplePrint
	fnName := runtime.FuncForPC(pc).Name()

	lineS := strconv.Itoa(line)
	nodes := []string{"PRINT AT "}

	// eg:
	// "PRINT AT github.com/gookit/goutil/dump.ExamplePrint(goutil/dump/dump_test.go:23)"
	// "PRINT AT github.com/gookit/goutil/dump.ExamplePrint(dump_test.go:23)"
	// "PRINT AT github.com/gookit/goutil/dump.ExamplePrint(:23)"
	for _, flag := range flags {
		// has flag
		if Config.ShowFlag&flag == 0 {
			continue
		}
		switch flag {
		case Ffunc: // full func name
			nodes = append(nodes, fnName, "(")
		case Ffile: // full file path
			nodes = append(nodes, file)
		case Ffname: // only file name
			fName := path.Base(file) // file name
			nodes = append(nodes, fName)
		case Fline:
			nodes = append(nodes, ":", lineS)
		}
	}

	// fallback. eg: "PRINT AT goutil/dump/dump_test.go:23"
	if len(nodes) == 1 {
		nodes = append(nodes, file, ":", lineS)
	} else if Config.ShowFlag & Ffunc != 0 { // has func, add ")"
		nodes = append(nodes, ")")
	}

	text := strings.Join(nodes, "")

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

	// if is an ptr, get real type and value
	if rType.Kind() == reflect.Ptr {
		rVal = rVal.Elem()
		rType = rType.Elem()
		// add "*" prefix
		mustFprintf(w, "*")
	}

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

			// TODO format print sub-struct
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
