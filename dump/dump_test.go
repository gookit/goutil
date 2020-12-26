package dump

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gookit/color"
	"github.com/stretchr/testify/assert"
)

func ExamplePrint() {
	Config(func(d *Dumper) {
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

	Config(func(d *Dumper) {
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

	// print position
	Fprint(buf, 123)
	// "PRINT AT github.com/gookit/goutil/dump.TestPrint(dump_test.go:65)"
	str := buf.String()
	is.Contains(str, "PRINT AT github.com/gookit/goutil/dump.TestPrint(dump_test.go:")
	is.Contains(str, "int(123)")

	// disable caller position for test
	Config(func(d *Dumper) {
		d.ShowFlag = Fnopos
	})
	defer Reset()

	buf.Reset()
	Fprint(buf, "abc")
	is.Equal(`string("abc"),
`, buf.String())

	buf.Reset()
	Fprint(buf, []string{"ab", "cd"})
	is.Equal(`[]string{"ab", "cd"},
`, buf.String())

	buf.Reset()
	Fprint(buf, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
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
	Fprint(buf, struct {
		ab string
		Cd int
	}{
		"ab", 23,
	})
	is.Equal(`struct { ab string; Cd int } {
  ab: string("ab"),
  Cd: int(23),
},
`, buf.String())

	buf.Reset()
	Fprint(buf, map[string]interface{}{
		"key": "val",
		"sub": map[string]string{"k": "v"},
	})
	is.Contains(buf.String(), `
  "sub": map[string]string {
    "k": string("v"),
  },
`)

	buf.Reset()
}

func TestPrintNil(t *testing.T) {
	is := assert.New(t)

	buf := newBuffer()
	Config(func(d *Dumper) {
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
	assert.Contains(t, str, "opt0: <nil>,")
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
	myS2 := struct {
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

func newBuffer() *bytes.Buffer {
	buf := new(bytes.Buffer)

	Config(func(d *Dumper) {
		d.Output = buf
		d.NoColor = true
	})

	return buf
}
