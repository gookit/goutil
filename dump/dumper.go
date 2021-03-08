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
	"github.com/kortschak/utter"
)

// Options for dump vars
type Options struct {
	// Output output writer
	Output io.Writer
	// dont show type
	NoType bool
	// dont with color
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
	CallerSkip int
	// ColorTheme for print result.
	ColorTheme map[string]string
}

// printValue must keep track of already-printed pointer values to avoid
// infinite recursion. refer the pkg: github.com/kr/pretty
type visit struct {
	v   uintptr
	typ reflect.Type
}

// Dumper struct definition
type Dumper struct {
	*Options
	// visited struct records
	visited map[visit]int
	// is value of the map, struct
	msValue bool
	// context information
	prevDepth, curDepth, nextDepth int
	indentStr, indentPrev, lineEnd string
	//
	indentBytes []byte
}

// NewDumper create
func NewDumper(out io.Writer, skip int) *Dumper {
	return &Dumper{
		Options: NewDefaultOptions(out, skip),
		// init map
		visited: make(map[visit]int),
	}
}

func NewDefaultOptions(out io.Writer, skip int) *Options {
	if out == nil {
		out = os.Stdout
	}

	return &Options{
		Output: out,
		// ---
		MaxDepth:  5,
		ShowFlag:  Ffunc | Ffname | Fline,
		MoreLenNL: 8,
		// ---
		IndentLen:  2,
		IndentChar: ' ',
		CallerSkip: skip,
		ColorTheme: defaultTheme,
	}
}

// Config for dumper
func (d *Dumper) Config(fn func(d *Dumper)) *Dumper {
	fn(d)
	return d
}

// WithSkip for dumper
func (d *Dumper) WithSkip(skip int) *Dumper {
	d.CallerSkip = skip
	return d
}

// WithOptions for dumper
func (d *Dumper) WithOptions(fn func(opts *Options)) *Dumper {
	fn(d.Options)
	return d
}

// ResetOptions for dumper
func (d *Dumper) ResetOptions() {
	d.Options = NewDefaultOptions(os.Stdout, 2)
}

// Dump vars
func (d *Dumper) Dump(vs ...interface{}) {
	d.dump(vs...)
}

// Print vars. alias of Dump()
func (d *Dumper) Print(vs ...interface{}) {
	d.dump(vs...)
}

// Println vars. alias of Dump()
func (d *Dumper) Println(vs ...interface{}) {
	d.dump(vs...)
}

// Spew print vars.
func (d *Dumper) Spew(vs ...interface{}) {
	// show print position
	if d.ShowFlag != Fnopos {
		// get the print position
		pc, file, line, ok := runtime.Caller(d.CallerSkip - 1)
		if ok {
			d.printCaller(pc, file, line)
		}
	}

	// print var data
	for _, v := range vs {
		utter.Dump(v)
	}
}

// Fprint print vars to io.Writer
func (d *Dumper) Fprint(w io.Writer, vs ...interface{}) {
	out := d.Output // backup

	d.Output = w
	d.dump(vs...)
	d.Output = out // restore
}

// dump vars
func (d *Dumper) dump(vs ...interface{}) {
	// show print position
	if d.ShowFlag != Fnopos {
		// get the print position
		pc, file, line, ok := runtime.Caller(d.CallerSkip)
		if ok {
			d.printCaller(pc, file, line)
		}
	}

	// print var data
	for _, v := range vs {
		// d.advance(1)
		d.printOne(v)
		// d.advance(-1)
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
		d.indentPrint(text, "\n")
		return
	}

	color.Fprint(d.Output, "<mga>", text, "</>\n")
}

func (d *Dumper) advance(step int) {
	d.curDepth += step
	d.nextDepth = d.curDepth + step

	// strings.Repeat()
	d.indentStr = strings.Repeat(string(d.IndentChar), d.IndentLen*d.curDepth)
	d.indentBytes = strutil.RepeatBytes(d.IndentChar, d.IndentLen*d.curDepth)

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

func (d *Dumper) printOne(v interface{}) {
	if v == nil {
		d.indentPrint("<nil>,\n")
		return
	}

	rv := reflect.ValueOf(v)
	d.printRValue(rv.Type(), rv)
}

func (d *Dumper) printRValue(t reflect.Type, v reflect.Value) {
	var isPtr bool
	// if is an ptr, get real type and value
	if t.Kind() == reflect.Ptr {
		isPtr = true
		v = v.Elem()
		t = t.Elem()
		// add "*" prefix
		d.indentPrint("&")
	}

	switch t.Kind() {
	case reflect.Bool:
		d.printf("%s(%v),\n", t.String(), v.Bool())
	case reflect.Float32, reflect.Float64:
		d.printf("%s(%v),\n", t.String(), v.Float())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		d.printf("%s(%d),\n", t.String(), v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		d.printf("%s(%d),\n", t.String(), v.Uint())
	case reflect.String:
		d.printf("%s(\"%s\"),\n", t.String(), v.String())
	case reflect.Complex64, reflect.Complex128:
		d.printf("%#v\n", v.Complex())
	case reflect.Slice, reflect.Array:
		eleNum := v.Len()
		if eleNum < d.MoreLenNL {
			d.printf("%#v,\n", v.Interface())
			return
		}

		d.indentPrint(t.String(), " [#len=", eleNum, "\n")
		for i := 0; i < eleNum; i++ {
			sv := v.Index(i)
			d.advance(1)
			d.printRValue(sv.Type(), sv)
			// d.printf("%v,\n", v.Index(i).Interface())
			d.advance(-1)
		}
		d.indentPrint("]\n")
	case reflect.Struct:
		// refer https://stackoverflow.com/questions/42664837/how-to-access-unexported-struct-fields-in-golang
		// NOTICE: only re-reflect.New on curDepth=1
		if !isPtr && d.curDepth == 1 {
			// fmt.Println("re reflect.New")
			// oldRv := v
			// v = reflect.New(t).Elem()
			// v.Set(oldRv)

			// ele := reflect.NewAt(v.Field(0).Type(), unsafe.Pointer(v.Field(0).UnsafeAddr())).Elem()
			// fmt.Println("aaa", ele.CanInterface())
		}
		if d.curDepth > d.MaxDepth {
			if !v.CanInterface() {
				d.printf("%s,\n", v.String())
			} else {
				d.printf("%#v,\n", v.Interface())
			}
		} else {
			if v.CanAddr() {
				addr := v.UnsafeAddr()
				vis := visit{addr, t}
				if vd, ok := d.visited[vis]; ok && vd < d.MaxDepth {
					d.indentPrint(t.String()+"{(CYCLIC REFERENCE)}", false)
					break // don't print v again
				}
				d.visited[vis] = d.curDepth
			}

			d.indentPrint(t.String(), " {\n")

			fldNum := v.NumField()
			for i := 0; i < fldNum; i++ {
				fv := v.Field(i)
				d.advance(1)

				fieldName := t.Field(i).Name
				// mustFprintf(w, "%s<cyan>%s</>: ", indentStr, fieldName)
				d.indentPrint(fieldName, ": ")

				d.msValue = true
				d.printRValue(fv.Type(), fv)
				d.msValue = false

				d.advance(-1)
			} // end for

			d.indentPrint("},\n")
		}
	case reflect.Map:
		if d.curDepth > d.MaxDepth {
			if !v.CanInterface() {
				d.printf("%s,\n", v.String())
			} else {
				d.printf("%#v,\n", v.Interface())
			}
		} else {
			d.indentPrint(t.String(), " {\n")
			for _, key := range v.MapKeys() {
				mv := v.MapIndex(key)
				d.advance(1)

				// print key name
				if !key.CanInterface() {
					// d.printf("<cyan>%s</>: ", key.String())
					d.printf("%s: ", key.String())
				} else {
					d.printf("%#v: ", key.Interface())
				}

				// print field value
				d.msValue = true
				d.printRValue(mv.Type(), mv)
				d.msValue = false

				d.advance(-1)
			}

			d.indentPrint("},\n")
		}
	case reflect.Interface:
		switch e := v.Elem(); {
		case e.Kind() == reflect.Invalid:
			d.print("nil")
		case e.IsValid():
			// d.advance(1)
			d.printRValue(e.Type(), e)
		default:
			d.print(t.String(), "(nil)")
		}
	default:
		// intX, uintX, string
		if v.CanInterface() {
			d.printf("%s(%#v),\n", t.String(), v.Interface())
		} else {
			d.printf("%s(%v),\n", t.String(), v.String())
		}
	}
}

func (d *Dumper) print(v ...interface{}) {
	_, _ = fmt.Fprint(d.Output, v...)
	// color.Fprint(w, v...)
}

func (d *Dumper) printf(f string, v ...interface{}) {
	if !d.msValue {
		_, _ = d.Output.Write(d.indentBytes)
	}
	_, _ = fmt.Fprintf(d.Output, f, v...)
	// color.Fprintf(w, f, v...)
}

func (d *Dumper) indentPrint(v ...interface{}) {
	if !d.msValue {
		_, _ = d.Output.Write(d.indentBytes)
	}
	_, _ = fmt.Fprint(d.Output, v...)
	// color.Fprint(w, v...)
}
