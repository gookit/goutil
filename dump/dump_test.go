package dump

import (
	"bytes"
	"io"
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExamplePrint() {
	Config.NoColor = true

	Print(
		23,
		[]string{"ab", "cd"},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		map[string]string{"key": "val"},
		map[string]interface{}{
			"sub": map[string]string{"k": "v"},
		},
		struct {
			ab string
			Cd int
		}{
			"ab", 23,
		},
	)
	Config.NoColor = false

	// Output like:
	// PRINT AT github.com/gookit/goutil/dump.ExamplePrint(LINE 14):
	// int(23)
	// []string{"ab", "cd"}
	// []int [
	//   1,
	//   2,
	//   3,
	//   4,
	//   5,
	//   6,
	//   7,
	//   8,
	//   9,
	//   10,
	//   11,
	// ]
	// map[string]string {
	//   key: "val",
	// }
	// map[string]interface {} {
	//   sub: map[string]string{"k":"v"},
	// }
	// struct { ab string; Cd int } {
	//   ab: "ab",
	//   Cd: 23,
	// }
	//
}

func TestConfig(t *testing.T) {
	is := assert.New(t)
	buf := new(bytes.Buffer)

	Output = buf
	// no color on tests
	Config.NoColor = true

	// show file
	Config.ShowFlag = Ffile|Fline

	P("hi")
	// PRINT AT /Users/inhere/Workspace/godev/gookit/goutil/dump/dump_test.go:171
	is.Contains(buf.String(), "goutil/dump/dump_test.go:")
	buf.Reset()

	// reset
	resetDump()
}

func TestPrint(t *testing.T) {
	is := assert.New(t)
	buf := new(bytes.Buffer)
	// disable position for test
	Config.ShowFlag = Fnopos

	Fprint(1, buf, 123)
	// "PRINT AT github.com/gookit/goutil/dump.TestPrint(dump_test.go:65)"
	str := buf.String()
	is.Contains(str, "PRINT AT github.com/gookit/goutil/dump.TestPrint(dump_test.go:")
	is.Contains(str, "int(123)")

	buf.Reset()
	Fprint(1, buf, "abc")
	is.Equal("string(abc)\n", buf.String())

	buf.Reset()
	Fprint(1, buf, []string{"ab", "cd"})
	is.Equal(`[]string{"ab", "cd"}
`, buf.String())

	buf.Reset()
	Fprint(1, buf, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	is.Equal(`[]int [
  1,
  2,
  3,
  4,
  5,
  6,
  7,
  8,
  9,
  10,
  11,
]
`, buf.String())

	buf.Reset()
	Fprint(1, buf, struct {
		ab string
		Cd int
	}{
		"ab", 23,
	})
	is.Equal(`struct { ab string; Cd int } {
  ab: "ab",
  Cd: 23,
}
`, buf.String())

	buf.Reset()
	Fprint(1, buf, map[string]interface{}{
		"key": "val",
		"sub": map[string]string{"k": "v"},
	})
	is.Contains(buf.String(), `sub: map[string]string{"k":"v"},`)

	buf.Reset()
	ResetConfig()
}

func TestPrintNil(t *testing.T) {
	is := assert.New(t)

	buf := newBuffer()
	defer resetDump()

	Print(nil)
	is.Equal("<nil>\n", buf.String())
	buf.Reset()

	var i int
	Println(i)
	is.Equal("int(0)\n", buf.String())
	buf.Reset()

	var f interface{}
	V(f)
	is.Equal("<nil>\n", buf.String())
	buf.Reset()
}

func TestPrintPtr(t *testing.T) {
	user := &struct {
		id   string
		Name string
		Age  int
	}{"ab1234", "inhere", 22}
	P(user)

	// Out:
	// *struct { id string; Name string; Age int } {
	//  id: string("ab1234"),
	//  Name: string("inhere"),
	//  Age: int(22),
	// }

	buf := newBuffer()
	defer resetDump()

	Println(user)

	str := buf.String()
	assert.Contains(t, str, "*struct")
	assert.Contains(t, str, "Age: int(22),")
	assert.Contains(t, str, "id: string(\"ab1234\"),")
	assert.Contains(t, str, "Name: string(\"inhere\"),")
}

func TestStruct_CannotExportField(t *testing.T) {
	myOpts := struct {
		opt0 *int
		opt1 bool
		opt2 int
		opt3 float64
		opt4 string
	}{ nil,true, 22, 34.45, "abc"}

	Print(myOpts)

	// OUT:
	// PRINT AT github.com/gookit/goutil/dump.TestStruct_CannotExportField(dump_test.go:202)
	// struct { opt0 *int; opt1 bool; opt2 int; opt3 float64; opt4 string } {
	//  opt0: <nil>,
	//  opt1: true,
	//  opt2: int(22),
	//  opt3: float64(34.45),
	//  opt4: string("abc"),
	// }

	buf := newBuffer()
	defer resetDump()

	Println(myOpts)

	str := buf.String()
	assert.Contains(t, str, "struct {")
	assert.Contains(t, str, "opt0: <nil>,")
	assert.Contains(t, str, "opt2: int(22),")
	assert.Contains(t, str, "opt3: float64(34.45)")
	assert.Contains(t, str, "opt4: string(\"abc\"),")
}

func TestStruct_WithNested(t *testing.T)  {
	type st0 struct {
		Sex int
	}

	type st1 struct {
		st0
		Age int
		Name string
	}

	s1 := st1{ st0{2},23, "inhere"}

	Println(s1)
	// OUT:
	// PRINT AT github.com/gookit/goutil/dump.TestStruct_WithNested(dump_test.go:223)
	// struct { dump.st0; Age int; Name string } {
	//  st0: dump.st0 {
	//    Sex: 2,
	//  },
	//  Age: 23,
	//  Name: "inhere",
	// }

	type st2 struct {
		st1
		Github string
	}

	s2 := st2{st1: s1, Github: "https://github.com/inhere"}
	Println(s2)

	// Out
	// PRINT AT github.com/gookit/goutil/dump.TestStruct_WithNested(dump_test.go:257)
	// dump.st2 {
	//  st1: dump.st1 {
	//    st0: dump.st0 {
	//      Sex: int(2),
	//    },
	//    Age: int(23),
	//    Name: string("inhere"),
	//  },
	//  Github: string("https://github.com/inhere"),
	// }

	s3 := struct {
		st1
		Github string
	} {st1: s1, Github: "https://github.com/inhere"}
	Println(s3)
}

func newBuffer() *bytes.Buffer {
	buf := new(bytes.Buffer)

	// set output for test
	Output = buf
	// disable position for test
	Config.ShowFlag = Fnopos

	return buf
}

func resetDump() {
	Output = os.Stdout
	ResetConfig()
}

// Dumper struct
type Dumper struct {
	dumpConfig
	Skip int
	Out  io.Writer
}

// NewDumper create
func NewDumper(out io.Writer) *Dumper {
	return &Dumper{
		Out: out,
		Skip: 3,
	}
}

// Dump vars
func (d *Dumper) Dump(vars ...interface{}) {
	// show print position
	if d.ShowFlag != Fnopos {
		// get the print position
		pc, file, line, ok := runtime.Caller(d.Skip)
		if ok {
			printPosition(d.Out, pc, file, line)
		}
	}

	// print data
	for _, v := range vars {
		printOne(d.Out, v)
	}
}
