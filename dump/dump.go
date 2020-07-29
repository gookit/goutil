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
	"github.com/gookit/goutil/strutil"
)

// These flags define which print information
const (
	Fnopos = 1 << iota // no position
	Ffunc
	Ffile
	Ffname
	Fline
)

// Options for dump vars
type Options struct {
	NoColor bool
	// Indent space. default is 2
	Indent int
	// MaxDepth for nested print
	MaxDepth int
	// ShowFlag for display caller position
	ShowFlag int
	// MoreLenNL array/slice elements length > MoreLenNL, will wrap new line
	MoreLenNL int
}

// Dumper struct
type Dumper struct {
	Options
	Skip int
	Out  io.Writer
}

// NewDumper create
func NewDumper(out io.Writer) *Dumper {
	return &Dumper{
		Out:  out,
		Skip: 3,
	}
}

// Config options for dumper
func (d *Dumper) Config(fn func(opts *Options)) *Dumper {
	fn(&d.Options)
	return d
}

// WithOptions for dumper
func (d *Dumper) WithOptions(opts Options) *Dumper {
	d.Options = opts
	return d
}

// Dump vars
func (d *Dumper) Dump(vars ...interface{}) {
	// show print position
	if d.ShowFlag != Fnopos {
		// get the print position
		pc, file, line, ok := runtime.Caller(d.Skip)
		if ok {
			d.printCaller(pc, file, line)
		}
	}

	// print data
	for _, v := range vars {
		printOne(d.Out, v)
	}
}

func (d *Dumper) printCaller(pc uintptr, file string, line int) {
	// eg: github.com/gookit/goutil/dump.ExamplePrint
	fnName := runtime.FuncForPC(pc).Name()

	lineS := strconv.Itoa(line)
	nodes := []string{"PRINT AT "}

	// eg:
	// "PRINT AT github.com/gookit/goutil/dump.ExamplePrint(goutil/dump/dump_test.go:23)"
	// "PRINT AT github.com/gookit/goutil/dump.ExamplePrint(dump_test.go:23)"
	// "PRINT AT github.com/gookit/goutil/dump.ExamplePrint(:23)"
	for _, flag := range posFlags {
		// has flag
		if d.ShowFlag&flag == 0 {
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
	} else if d.ShowFlag&Ffunc != 0 { // has func, add ")"
		nodes = append(nodes, ")")
	}

	text := strings.Join(nodes, "")

	if d.NoColor {
		mustFprint(d.Out, text, "\n")
		return
	}

	color.Fprint(d.Out, "<mga>", text, "</>\n")
}

func (d *Dumper) printOne(v interface{}) {
	if v == nil {
		mustFprintf(w, "<nil>\n")
		return
	}


}

func (d *Dumper) printMap(v interface{}) {

}

func (d *Dumper) printStruct(v interface{}) {

}

var (
	// valid flag for position
	posFlags = []int{Ffunc, Ffile, Ffname, Fline}

	// Output set print content to the io.Writer
	Output io.Writer = os.Stdout

	// Config dump data settings
	Config = newDefaultOptions()
)

// ResetConfig reset config data
func ResetConfig() {
	Config = newDefaultOptions()
}

func newDefaultOptions() Options {
	return Options{
		Indent:    2,
		MaxDepth:  5,
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
	for _, flag := range posFlags {
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
	} else if Config.ShowFlag&Ffunc != 0 { // has func, add ")"
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

	printReflectedValue(w, rType, rVal, 1)
}

func printReflectedValue(w io.Writer, rType reflect.Type, rVal reflect.Value, depth int) {
	// if is an ptr, get real type and value
	if rType.Kind() == reflect.Ptr {
		rVal = rVal.Elem()
		rType = rType.Elem()
		// add "*" prefix
		mustFprintf(w, "*")
	}

	prevDepth := depth - 1
	nextDepth := depth + 1
	indentStr := strutil.Repeat(" ", Config.Indent*depth)
	indentPrev := strutil.Repeat(" ", Config.Indent*prevDepth)
	switch rType.Kind() {
	case reflect.Slice, reflect.Array:
		eleNum := rVal.Len()
		if eleNum < Config.MoreLenNL {
			mustFprintf(w, "%#v\n", rVal.Interface())
			return
		}

		mustFprint(w, rType.String(), " [\n")
		for i := 0; i < eleNum; i++ {
			mustFprintf(w, "%s%v,\n", indentStr, rVal.Index(i).Interface())
		}
		mustFprint(w, "]\n")
	case reflect.Struct:
		fldNum := rVal.NumField()

		mustFprint(w, rType.String(), " {\n")
		for i := 0; i < fldNum; i++ {
			fv := rVal.Field(i)
			fName := rType.Field(i).Name
			// print field name
			mustFprintf(w, "%s%s: ", indentStr, fName)

			fTypeName := fv.Type().String()

			// print field value
			switch fv.Kind() {
			case reflect.Bool:
				mustFprintf(w, "%v,\n", fv.Bool())
			case reflect.String:
				mustFprintf(w, "%s(\"%s\"),\n", fTypeName, fv.String())
			case reflect.Float32, reflect.Float64:
				mustFprintf(w, "%s(%v),\n", fTypeName, fv.Float())
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				mustFprintf(w, "%s(%d),\n", fTypeName, fv.Int())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				mustFprintf(w, "%s(%d),\n", fTypeName, fv.Uint())
			case reflect.Map, reflect.Struct:
				if depth > Config.MaxDepth {
					mustFprintf(w, "%s,\n", fv.String())
				} else {
					printReflectedValue(w, fv.Type(), fv, nextDepth)
				}
			case reflect.Interface:
				if !fv.CanInterface() {
					mustFprintf(w, "%#v,\n", fv.String())
					continue
				}

				switch typData := fv.Interface().(type) {
				case int, int8, int16, int32, int64:
					mustFprintf(w, "%s(%v),\n", reflect.TypeOf(typData).String(), typData)
				case uint, uint8, uint16, uint32, uint64:
					mustFprintf(w, "%s(%v),\n", reflect.TypeOf(typData).String(), typData)
				case float32, float64:
					mustFprintf(w, "%s(%v),\n", reflect.TypeOf(typData).String(), typData)
				case string:
					mustFprintf(w, "%s(\"%v\"),\n", "string", fv.Interface())
				case map[int]int, map[string]int, map[string]string, map[int]interface{}, map[string]interface{}:
					if depth > Config.MaxDepth {
						mustFprintf(w, "%v,\n", typData)
					} else {
						typRVal := reflect.ValueOf(typData)
						printReflectedValue(w, typRVal.Type(), typRVal, nextDepth)
					}
				default:
					if fv.IsNil() {
						mustFprint(w, "<nil>,\n")
					} else if fv.CanInterface() {
						// mustFprintf(w, "%s(%v),\n", vTypeName, fv.Interface())
						mustFprintf(w, "%#v,\n", fv.Interface())
					} else {
						mustFprintf(w, "%#v,\n", fv.String())
					}
				}
			default:
				if fv.IsNil() {
					mustFprint(w, "<nil>,\n")
				} else if fv.CanInterface() {
					mustFprintf(w, "%#v,\n", fv.Interface())
				} else {
					mustFprintf(w, "%#v,\n", fv.String())
				}
			} // end switch
		} // end for

		if prevDepth > 0 {
			mustFprint(w, indentPrev, "},\n")
		} else {
			mustFprint(w, "}\n")
		}
	case reflect.Map:
		mustFprint(w, rType.String(), " {\n")

		for _, key := range rVal.MapKeys() {
			// old: direct print value
			// mustFprintf(w, "%s%v: %#v,\n", indentStr, key.Interface(), rVal.MapIndex(key).Interface())

			// print key name
			if !key.CanInterface() {
				mustFprintf(w, "%s%v: ", indentStr, key.String())
			} else {
				mustFprintf(w, "%s%v: ", indentStr, key.Interface())
			}

			mv := rVal.MapIndex(key)
			vTypeName := mv.Type().String()

			// print field value
			switch mv.Kind() {
			case reflect.Bool:
				mustFprintf(w, "%v,\n", mv.Bool())
			case reflect.String:
				mustFprintf(w, "%s(\"%s\"),\n", vTypeName, mv.String())
			case reflect.Float32, reflect.Float64:
				mustFprintf(w, "%s(%v),\n", vTypeName, mv.Float())
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				mustFprintf(w, "%s(%d),\n", vTypeName, mv.Int())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				mustFprintf(w, "%s(%d),\n", vTypeName, mv.Uint())
			case reflect.Map, reflect.Struct:
				if depth > Config.MaxDepth {
					mustFprintf(w, "%s,\n", mv.String())
				} else {
					printReflectedValue(w, mv.Type(), mv, nextDepth)
				}
			case reflect.Interface:
				if !mv.CanInterface() {
					mustFprintf(w, "%#v,\n", mv.String())
					continue
				}
				switch typData := mv.Interface().(type) {
				case int, int8, int16, int32, int64:
					mustFprintf(w, "%s(%v),\n", reflect.TypeOf(typData).String(), typData)
				case uint, uint8, uint16, uint32, uint64:
					mustFprintf(w, "%s(%v),\n", reflect.TypeOf(typData).String(), typData)
				case float32, float64:
					mustFprintf(w, "%s(%v),\n", reflect.TypeOf(typData).String(), typData)
				case string:
					mustFprintf(w, "%s(\"%v\"),\n", "string", mv.Interface())
				case map[int]int, map[string]int, map[string]string, map[int]interface{}, map[string]interface{}:
					if depth > Config.MaxDepth {
						mustFprintf(w, "%v,\n", typData)
					} else {
						typRVal := reflect.ValueOf(typData)
						printReflectedValue(w, typRVal.Type(), typRVal, nextDepth)
					}
				default:
					if mv.IsNil() {
						mustFprint(w, "<nil>,\n")
					} else if mv.CanInterface() {
						// mustFprintf(w, "%s(%v),\n", vTypeName, mv.Interface())
						mustFprintf(w, "%#v,\n", mv.Interface())
					} else {
						mustFprintf(w, "%#v,\n", mv.String())
					}
				}
			default:
				if mv.IsNil() {
					mustFprint(w, "<nil>,\n")
				} else if mv.CanInterface() {
					mustFprintf(w, "%s(%#v),\n", vTypeName, mv.Interface())
				} else {
					mustFprintf(w, "%#v,\n", mv.String())
				}
			}
		}

		if prevDepth > 0 {
			mustFprint(w, indentPrev, "},\n")
		} else {
			mustFprint(w, "}\n")
		}
	default:
		if rVal.CanInterface() {
			mustFprintf(w, "%s(%v)\n", rType.String(), rVal.Interface())
		} else {
			mustFprintf(w, "%s(%v)\n", rType.String(), rVal.String())
		}
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
