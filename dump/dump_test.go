package dump

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/gookit/color"
	"github.com/stretchr/testify/assert"
)

func ExamplePrint() {
	Config(func(d *Options) {
		d.NoColor = true
	})
	defer Reset()

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

func TestStd(t *testing.T) {
	assert.Equal(t, Std().NoColor, false)
	assert.Equal(t, Std().IndentLen, 2)
}

func TestConfig(t *testing.T) {
	is := assert.New(t)
	buf := new(bytes.Buffer)

	Config(func(d *Options) {
		d.Output = buf
		// no color on tests
		d.NoColor = true
		// show file
		d.ShowFlag = Ffile | Fline
	})
	defer Reset()

	P("hi")
	// PRINT AT /Users/inhere/Workspace/godev/gookit/goutil/dump/dump_test.go:171
	is.Contains(buf.String(), "goutil/dump/dump_test.go:")
	buf.Reset()
}

func TestPrint(t *testing.T) {
	is := assert.New(t)
	buf := new(bytes.Buffer)

	// disable caller position for test
	Config(func(d *Options) {
		d.NoColor = true
	})
	defer Reset()

	// print position
	Fprint(buf, 123)
	// "PRINT AT github.com/gookit/goutil/dump.TestPrint(dump_test.go:65)"
	str := buf.String()
	is.Contains(str, "PRINT AT github.com/gookit/goutil/dump.TestPrint(dump_test.go:")
	is.Contains(str, "int(123)")

	// dont print position
	Std().ShowFlag = Fnopos

	buf.Reset()
	Fprint(buf, "abc")
	is.Equal(`string("abc"), #len=3
`, buf.String())

	buf.Reset()
	Fprint(buf, []string{"ab", "cd"})
	is.Contains(buf.String(), `[]string [ #len=2`)

	buf.Reset()
	Fprint(buf, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	is.Equal(`[]int [ #len=11
  int(1),
  int(2),
  int(3),
  int(4),
  int(5),
  int(6),
  int(7),
  int(8),
  int(9),
  int(10),
  int(11),
],
`, buf.String())

	buf.Reset()
	Fprint(buf, struct {
		ab string
		Cd int
	}{
		"ab", 23,
	})
	is.Equal(`struct { ab string; Cd int } {
  ab: string("ab"), #len=2
  Cd: int(23),
},
`, buf.String())

	buf.Reset()
	Fprint(buf, map[string]interface{}{
		"key": "val",
		"sub": map[string]string{"k": "v"},
	})
	is.Contains(buf.String(), `
  "sub": map[string]string { #len=1
    "k": string("v"), #len=1
  },
`)

	buf.Reset()
}

func TestPrintNil(t *testing.T) {
	is := assert.New(t)

	buf := newBuffer()
	Config(func(d *Options) {
		d.ShowFlag = Fnopos
	})
	defer Reset()

	Print(nil)
	is.Equal("<nil>,\n", buf.String())
	buf.Reset()

	var i int
	Println(i)
	is.Equal("int(0),\n", buf.String())
	buf.Reset()

	var f interface{}
	V(f)
	is.Equal("<nil>,\n", buf.String())
	buf.Reset()
}

func TestStruct_CannotExportField(t *testing.T) {
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
	defer Reset()

	Println(myOpts)

	str := buf.String()
	assert.Contains(t, str, "struct {")
	assert.Contains(t, str, "opt0: *int<nil>,")
	assert.Contains(t, str, "opt2: int(22),")
	assert.Contains(t, str, "opt3: float64(34.45)")
	assert.Contains(t, str, "opt4: string(\"abc\"),")
}

func TestStruct_InterfaceField(t *testing.T) {
	myS1 := struct {
		// cannotExport interface{} // ok
		cannotExport st1 // ok
		// CanExport interface{} ok
		CanExport st1 // ok
	}{
		cannotExport: s1,
		CanExport:    s1,
	}

	Println(myS1)
	color.Infoln("\nUse fmt.Println:")
	fmt.Println(myS1)
}

func TestStruct_MapInterfacedValue(t *testing.T) {
	myS2 := &struct {
		cannotExport map[string]interface{}
	}{
		cannotExport: map[string]interface{}{
			"key1": 12,
			"key2": "abcd123",
		},
	}

	Println(myS2)
	color.Infoln("\nUse fmt.Println:")
	fmt.Println(myS2)
	fmt.Println("---------------------------------------------------------------")

	type st2 struct {
		st1
		Github string
		Face1  interface{}
		face2  interface{}
		faces  map[string]interface{}
	}

	s2 := st2{
		st1:    s1,
		Github: "https://github.com/inhere",
		Face1:  s1,
		face2:  s1,
		faces: map[string]interface{}{
			"key1": 12,
			"key2": "abc2344",
		},
	}

	Println(s2)
	color.Infoln("\nUse fmt.Println:")
	fmt.Println(s2)
}

func TestStruct_ptrField(t *testing.T) {
	type userOpts struct {
		Int *int
		// use ptr
		Str *string
	}

	aint := 2
	astr := "xyz"
	opt := &userOpts{
		Int: &aint,
		Str: &astr,
	}

	Println(opt)
	color.Infoln("\nUse fmt.Println:")
	fmt.Println(opt)
	fmt.Println("---------------------------------------------------------------")
}

func TestFormat(t *testing.T) {
	s := Format(23, "abc", map[string]interface{}{
		"key1": 12,
		"key2": "abc2344",
	})

	assert.NotEmpty(t, s)
	fmt.Println(s)
}

func TestPrint_over_max_depth(t *testing.T) {
	a := map[string]interface{}{}
	a["circular"] = map[string]interface{}{
		"a": a,
	}

	// TIP: will stack overflow
	// fmt.Println(a)

	P(a)
	s := Format(a)
	assert.NotEmpty(t, s)
	assert.Contains(t, s, "!OVER MAX DEPTH!")
}

func TestPrint_cyclic_slice(t *testing.T) {
	a := map[string]interface{}{
		"bool":   true,
		"number": 1 + 1i,
		"bytes":  []byte{97, 98, 99},
		"lines":  "first line\nsecond line",
		"slice":  []interface{}{1, 2},
		"time":   time.Now(),
		"struct": struct{ test int32 }{
			test: 13,
		},
	}
	a["slice"].([]interface{})[1] = a["slice"]

	// TIP: will stack overflow
	// fmt.Println(a)

	P(a)
	s := Format(a)
	assert.NotEmpty(t, s)
	assert.Contains(t, s, "!CYCLIC REFERENCE!")
}

func newBuffer() *bytes.Buffer {
	buf := new(bytes.Buffer)

	Config(func(d *Options) {
		d.Output = buf
		d.NoColor = true
	})

	return buf
}
