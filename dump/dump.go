package dump

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"

	"github.com/gookit/color"
)

var Config = struct {
	ShowFile   bool
	ShowMethod bool
}{
	ShowMethod: true,
}

// P print input params for pretty
func P(vs ...interface{}) {
	Fprint(2, os.Stdout, vs...)
}

// Print print input params for pretty
func Print(vs ...interface{}) {
	Fprint(2, os.Stdout, vs...)
}

// Println print input params for pretty
func Println(vs ...interface{}) {
	Fprint(2, os.Stdout, vs...)
}

// Print print input params for pretty
func Fprint(skip int, w io.Writer, vs ...interface{}) {
	// get the print position
	pc, _, line, ok := runtime.Caller(skip)
	if ok {
		mName := runtime.FuncForPC(pc).Name()
		text := fmt.Sprint("<mga>PRINT AT ", mName, "(LINE ", line, ")</>:\n")
		color.Fprint(w, text)
		// mustFprint(w, )
		// new line
		// _, _ = w.Write([]byte("\n"))
	}

	for _, v := range vs {
		printOne(w, v)
	}

	// new line
	// _, _ = w.Write([]byte("\n"))
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
			mustFprintf(w, "  %#v\n", rValue.Index(i).Interface())
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
				mustFprintf(w, "  %v: %#v\n", tn, rValue.Field(i).Interface())
			} else {
				mustFprintf(w, "  %v: %#v\n", tn, rValue.Field(i).String())
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
		mustFprintf(w, "%v\n", v)
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
