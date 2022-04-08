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

// Options for dump vars
type Options struct {
	// Output the output writer
	Output io.Writer
	// NoType dont show data type TODO
	NoType bool
	// NoColor don't with color
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
	// MoreLenNL int
	// CallerSkip skip for call runtime.Caller()
	CallerSkip int
	// ColorTheme for print result.
	ColorTheme Theme
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
	// is value in the slice, map, struct. will not apply indent.
	msValue bool
	// current depth
	curDepth int
	// current indent string bytes
	indentBytes []byte
	// prevDepth, nextDepth int
	// indentStr, indentPrev, lineEnd string
}

// NewDumper create
func NewDumper(out io.Writer, skip int) *Dumper {
	return &Dumper{
		Options: NewDefaultOptions(out, skip),
		// init map
		visited: make(map[visit]int),
	}
}

// NewWithOptions create
func NewWithOptions(fn func(opts *Options)) *Dumper {
	d := NewDumper(os.Stdout, 3)
	fn(d.Options)
	return d
}

// NewDefaultOptions create.
func NewDefaultOptions(out io.Writer, skip int) *Options {
	if out == nil {
		out = os.Stdout
	}

	return &Options{
		Output: out,
		// ---
		MaxDepth: 5,
		ShowFlag: Ffunc | Ffname | Fline,
		// MoreLenNL: 8,
		// ---
		IndentLen:  2,
		IndentChar: ' ',
		CallerSkip: skip,
		ColorTheme: defaultTheme,
	}
}

// WithSkip for dumper
func (d *Dumper) WithSkip(skip int) *Dumper {
	d.CallerSkip = skip
	return d
}

// WithoutColor for dumper
func (d *Dumper) WithoutColor() *Dumper {
	d.NoColor = true
	return d
}

// WithOptions for dumper
func (d *Dumper) WithOptions(fn func(opts *Options)) *Dumper {
	fn(d.Options)
	return d
}

// ResetOptions for dumper
func (d *Dumper) ResetOptions() {
	d.curDepth = 0
	d.visited = make(map[visit]int)
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

// Fprint print vars to io.Writer
func (d *Dumper) Fprint(w io.Writer, vs ...interface{}) {
	backup := d.Output // backup

	d.Output = w
	d.dump(vs...)
	d.Output = backup // restore
}

// dump go vars
func (d *Dumper) dump(vs ...interface{}) {
	// reset some settings.
	d.curDepth = 0
	d.visited = make(map[visit]int)

	// clear all theme settings.
	if d.NoColor {
		d.ColorTheme = make(Theme)
	}

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
	d.print(d.ColorTheme.caller(text), "\n")
}

func (d *Dumper) advance(step int) {
	d.curDepth += step
	// d.nextDepth = d.curDepth + step
	d.indentBytes = strutil.RepeatBytes(d.IndentChar, d.IndentLen*d.curDepth)
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
	// var isPtr bool
	// if is a ptr, get real type and value
	if t.Kind() == reflect.Ptr {
		if v.IsNil() {
			d.printf("%s<nil>,\n", t.String())
			return
		}

		v = v.Elem()
		t = t.Elem()
		// add "*" prefix
		d.indentPrint("&")
	}

	if !v.IsValid() {
		d.indentPrint(t.String(), "<nil>, #invalid\n")
	}

	// if v.CanAddr() && !d.checkCyclicRef(t, v) {
	// 	return // don't print v again
	// }

	if d.curDepth > d.MaxDepth {
		// if !v.CanInterface() {
		// 	d.printf("%s,\n", v.String())
		// } else {
		// 	// v.Interface() will stack overflow on cyclic refer
		// 	d.printf("%#v,\n", v.Interface())
		// }
		d.printf("%s(!OVER MAX DEPTH!),\n", v.String())
		return
	}

	switch t.Kind() {
	case reflect.Bool:
		d.printf("%s(%v),\n", t.String(), v.Bool())
	case reflect.Float32, reflect.Float64:
		d.printf("%s(%v),\n", t.String(), v.Float())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intStr := strconv.FormatInt(v.Int(), 10)
		intStr = d.ColorTheme.integer(intStr)
		d.printf("%s(%s),\n", t.String(), intStr)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		intStr := strconv.FormatUint(v.Uint(), 10)
		intStr = d.ColorTheme.integer(intStr)
		d.printf("%s(%s),\n", t.String(), intStr)
	case reflect.String:
		strVal := d.ColorTheme.string(v.String())
		lenTip := d.ColorTheme.lenTip("#len=" + strconv.Itoa(v.Len()))
		d.printf("%s(\"%s\"), %s\n", t.String(), strVal, lenTip)
	case reflect.Complex64, reflect.Complex128:
		d.printf("%#v\n", v.Complex())
	case reflect.Slice, reflect.Array:
		if v.CanAddr() && !d.checkCyclicRef(t, v) {
			break // don't print v again
		}

		eleNum := v.Len()
		lenTip := d.ColorTheme.lenTip("#len=" + strconv.Itoa(eleNum))

		d.indentPrint(t.String(), " [ ", lenTip, "\n")
		d.msValue = false
		for i := 0; i < eleNum; i++ {
			sv := v.Index(i)
			d.advance(1)

			// d.msValue = true
			d.printRValue(sv.Type(), sv)
			// d.msValue = false

			// d.printf("%v,\n", v.Index(i).Interface())
			d.advance(-1)
		}

		d.indentPrint("],\n")
	case reflect.Struct:
		if v.CanAddr() && !d.checkCyclicRef(t, v) {
			break // don't print v again
		}

		d.indentPrint(d.ColorTheme.msType(t.String()), " {\n")
		d.msValue = false

		fldNum := v.NumField()
		for i := 0; i < fldNum; i++ {
			fv := v.Field(i)
			d.advance(1)

			fName := t.Field(i).Name
			d.indentPrint(d.ColorTheme.field(fName), ": ")

			d.msValue = true
			d.printRValue(fv.Type(), fv)
			d.msValue = false

			d.advance(-1)
		}

		d.indentPrint("},\n")
	case reflect.Map:
		lenTip := d.ColorTheme.lenTip("#len=" + strconv.Itoa(v.Len()))
		d.indentPrint(d.ColorTheme.msType(t.String()), " { ", lenTip, "\n")
		d.msValue = false

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

			if mv.CanAddr() && !d.checkCyclicRef(mv.Type(), mv) {
				d.advance(-1)
				continue // don't print mv again
			}

			// print field value
			d.msValue = true
			d.printRValue(mv.Type(), mv)
			d.msValue = false

			d.advance(-1)
		}

		d.indentPrint("},\n")
	case reflect.Interface:
		if v.CanAddr() && !d.checkCyclicRef(t, v) {
			break // don't print v again
		}

		switch e := v.Elem(); {
		case e.Kind() == reflect.Invalid:
			d.indentPrint("nil,\n")
		case e.IsValid():
			// d.advance(1)
			d.printRValue(e.Type(), e)
		default:
			d.indentPrint(t.String(), "(nil),\n")
		}
	// case reflect.Ptr:
	case reflect.Chan:
		d.printf("(%s)(%#v),\n", t.String(), v.Pointer())
	case reflect.Func:
		d.printf("(%s) {...},\n", t.String())
	case reflect.UnsafePointer:
		d.printf("(%#v),\n", v.Pointer())
	case reflect.Invalid:
		d.indentPrint(t.String(), "(nil),\n")
	default:
		if v.CanAddr() && !d.checkCyclicRef(t, v) {
			break // don't print v again
		}

		if v.CanInterface() {
			d.printf("%s(%#v),\n", t.String(), v.Interface())
		} else {
			d.printf("%s(%v),\n", t.String(), v.String())
		}
	}
}

func (d *Dumper) checkCyclicRef(t reflect.Type, v reflect.Value) (goon bool) {
	addr := v.UnsafeAddr()
	vis := visit{addr, t}

	if vd, ok := d.visited[vis]; ok && vd < d.MaxDepth {
		d.indentPrint(t.String(), "{(!CYCLIC REFERENCE!)}\n")
		return false // don't print v again
	}

	// record
	d.visited[vis] = d.curDepth
	return true
}

func (d *Dumper) print(v ...interface{}) {
	if d.NoColor {
		_, _ = fmt.Fprint(d.Output, v...)
	} else {
		color.Fprint(d.Output, v...)
	}
}

func (d *Dumper) printf(f string, v ...interface{}) {
	if !d.msValue {
		_, _ = d.Output.Write(d.indentBytes)
	}

	if d.NoColor {
		_, _ = fmt.Fprintf(d.Output, f, v...)
	} else {
		color.Fprintf(d.Output, f, v...)
	}
}

func (d *Dumper) indentPrint(v ...interface{}) {
	if !d.msValue {
		_, _ = d.Output.Write(d.indentBytes)
	}

	if d.NoColor {
		_, _ = fmt.Fprint(d.Output, v...)
	} else {
		color.Fprint(d.Output, v...)
	}
}
