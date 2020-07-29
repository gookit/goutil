package dump

import (
	"io"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/gookit/color"
	"github.com/gookit/goutil/strutil"
)

// Options for dump vars
type Options struct {
	NoColor bool
	// IndentLen width. default is 2
	IndentLen int
	// IndentChar default is one space
	IndentChar byte
	// MaxDepth for nested print
	MaxDepth int
	// ShowFlag for display caller position
	ShowFlag int
	// MoreLenNL array/slice elements length > MoreLenNL, will wrap new line
	MoreLenNL int
	// CallerSkip skip for call runtime.Caller()
	// CallerSkip int
}

// Dumper struct
type Dumper struct {
	Options
	Output io.Writer
	// Skip for call runtime.Caller()
	Skip int
	// context information
	prevDepth, curDepth, nextDepth int
	indentStr, indentPrev, lineEnd string
}

// NewDumper create
func NewDumper(out io.Writer, skip int) *Dumper {
	return &Dumper{
		Output:  out,
		Options: newDefaultOptions(),
		// other
		Skip: skip,
	}
}

func newDefaultOptions() Options {
	return Options{
		MaxDepth:  5,
		ShowFlag:  Ffunc | Ffname | Fline,
		MoreLenNL: 8,
		// indent
		IndentLen:  2,
		IndentChar: ' ',
	}
}

// Config for dumper
func (d *Dumper) Config(fn func(d *Dumper)) *Dumper {
	fn(d)
	return d
}

// WithSkip for dumper
func (d *Dumper) WithSkip(skip int) *Dumper {
	d.Skip = skip
	return d
}

// WithOptions for dumper
func (d *Dumper) WithOptions(opts Options) *Dumper {
	d.Options = opts
	return d
}

// ResetOptions for dumper
func (d *Dumper) ResetOptions() {
	d.Options = newDefaultOptions()
}

// Dump vars
func (d *Dumper) Dump(vars ...interface{}) {
	d.dump(vars...)
}

// Print vars. alias of Dump()
func (d *Dumper) Print(vars ...interface{}) {
	d.dump(vars...)
}

// Println vars. alias of Dump()
func (d *Dumper) Println(vars ...interface{}) {
	d.dump(vars...)
}

// Fprint print vars to io.Writer
func (d *Dumper) Fprint(w io.Writer, vars ...interface{}) {
	// backup
	out := d.Output

	d.Output = w
	d.dump(vars...)

	// restore
	d.Output = out
}

// dump vars
func (d *Dumper) dump(vars ...interface{}) {
	// show print position
	if d.ShowFlag != Fnopos {
		// get the print position
		pc, file, line, ok := runtime.Caller(d.Skip)
		if ok {
			d.printCaller(pc, file, line)
		}
	}

	// print var data
	for _, v := range vars {
		d.advance(1)
		d.printOne(v)
		d.advance(-1)
	}
}

func (d *Dumper) advance(step int) {
	d.curDepth += step
	d.nextDepth = d.curDepth + step

	// strings.Repeat()
	d.indentStr = strings.Repeat(string(d.IndentChar), d.IndentLen*d.curDepth)

	// current depth > 1
	if d.curDepth > 1 {
		d.prevDepth += step
		// d.lineEnd = ",\n"
		d.indentPrev = strutil.Repeat(string(d.IndentChar), d.IndentLen*d.prevDepth)
	} else {
		// d.lineEnd = "\n"
		d.prevDepth = 0
		d.indentPrev = ""
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
	for _, flag := range callerFlags {
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
		mustFprint(d.Output, text, "\n")
		return
	}

	color.Fprint(d.Output, "<mga>", text, "</>\n")
}

func (d *Dumper) printOne(v interface{}) {
	if v == nil {
		mustFprintf(d.Output, "<nil>,\n")
		return
	}

	switch val := v.(type) {
	case int, int8, int16, int32, int64:
		mustFprintf(d.Output, "%s(%v),\n", reflect.TypeOf(val).String(), val)
	case uint, uint8, uint16, uint32, uint64:
		mustFprintf(d.Output, "%s(%v),\n", reflect.TypeOf(val).String(), val)
	case float32, float64:
		mustFprintf(d.Output, "%s(%v),\n", reflect.TypeOf(val).String(), val)
	case string:
		mustFprintf(d.Output, "string(\"%s\"),\n", val)
	default: // must in switch, use asserted 'val' for continue handle
		rTyp := reflect.TypeOf(val)
		rVal := reflect.ValueOf(val)
		d.printReflectValue(rTyp, rVal, d.curDepth)
	}
}

func (d *Dumper) printReflectValue(rTyp reflect.Type, rVal reflect.Value, depth int) {
	// if is an ptr, get real type and value
	if rTyp.Kind() == reflect.Ptr {
		rVal = rVal.Elem()
		rTyp = rTyp.Elem()
		// add "*" prefix
		mustFprintf(d.Output, "*")
	}

	w := d.Output

	// prevDepth := depth - 1
	// nextDepth := depth + 1
	// strings.Repeat()
	// indentStr := strutil.Repeat(" ", d.IndentLen*depth)
	// indentPrev := strutil.Repeat(" ", d.IndentLen*prevDepth)
	switch rTyp.Kind() {
	case reflect.Slice, reflect.Array:
		eleNum := rVal.Len()
		if eleNum < d.MoreLenNL {
			mustFprintf(w, "%#v\n", rVal.Interface())
			return
		}

		mustFprint(w, rTyp.String(), " [\n")
		for i := 0; i < eleNum; i++ {
			mustFprintf(w, "%s%v,\n", d.indentStr, rVal.Index(i).Interface())
		}
		mustFprint(w, "]\n")
	case reflect.Struct:
		fldNum := rVal.NumField()

		mustFprint(w, rTyp.String(), " {\n")
		for i := 0; i < fldNum; i++ {
			fv := rVal.Field(i)
			fName := rTyp.Field(i).Name
			// print field name
			// mustFprintf(w, "%s<cyan>%s</>: ", indentStr, fName)
			mustFprintf(w, "%s%s: ", d.indentStr, fName)

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
				if d.curDepth > d.MaxDepth {
					if !fv.CanInterface() {
						mustFprintf(w, "%s,\n", fv.String())
					} else {
						mustFprintf(w, "%#v,\n", fv.Interface())
					}
				} else {
					d.advance(1)
					d.printReflectValue(fv.Type(), fv, d.curDepth)
					d.advance(-1)
				}
			case reflect.Interface:
				if !fv.CanInterface() {
					mustFprintf(w, "%#v,\n", fv.String())
					continue
				}

				// goon handle field value
				d.advance(1)
				d.printOne(fv.Interface())
				d.advance(-1)
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

		mustFprint(w, d.indentPrev, "},\n")
	case reflect.Map:
		mustFprint(w, rTyp.String(), " {\n")

		for _, key := range rVal.MapKeys() {
			// print key name
			if !key.CanInterface() {
				// mustFprintf(w, "%s<cyan>%s</>: ", indentStr, key.String())
				mustFprintf(w, "%s%s: ", d.indentStr, key.String())
			} else {
				// mustFprintf(w, "%s<cyan>%#v</>: ", indentStr, key.Interface())
				mustFprintf(w, "%s%#v: ", d.indentStr, key.Interface())
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
				if d.curDepth > d.MaxDepth {
					if !mv.CanInterface() {
						mustFprintf(w, "%s,\n", mv.String())
					} else {
						mustFprintf(w, "%#v,\n", mv.Interface())
					}
				} else {
					d.advance(1)
					d.printReflectValue(mv.Type(), mv, d.curDepth)
					d.advance(-1)
				}
			case reflect.Interface:
				if !mv.CanInterface() {
					mustFprintf(w, "%s,\n", mv.String())
					continue
				}

				// goon handle field value
				d.advance(1)
				d.printOne(mv.Interface())
				d.advance(-1)
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

		mustFprint(w, d.indentPrev, "},\n")
	default:
		// intX, uintX, string
		if rVal.CanInterface() {
			mustFprintf(w, "%s(%#v),\n", rTyp.String(), rVal.Interface())
		} else {
			mustFprintf(w, "%s(%v),\n", rTyp.String(), rVal.String())
		}
	}
}

func (d *Dumper) printMap(v interface{}) {

}

func (d *Dumper) printStruct(v interface{}) {

}
