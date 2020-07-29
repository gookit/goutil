package dump

import (
	"bytes"
	"fmt"
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
	Config.ShowFlag = Ffile | Fline

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

	// print position
	Fprint(1, buf, 123)
	// "PRINT AT github.com/gookit/goutil/dump.TestPrint(dump_test.go:65)"
	str := buf.String()
	is.Contains(str, "PRINT AT github.com/gookit/goutil/dump.TestPrint(dump_test.go:")
	is.Contains(str, "int(123)")

	// disable position for test
	Config.ShowFlag = Fnopos

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
  ab: string("ab"),
  Cd: int(23),
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
	}{nil, true, 22, 34.45, "abc"}

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

type st0 struct {
	Sex int
}

type st1 struct {
	st0
	Age  int
	Name string
}

func TestStruct_WithNested(t *testing.T) {
	s1 := st1{st0{2}, 23, "inhere"}

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

	// Out:
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
	}{st1: s1, Github: "https://github.com/inhere"}
	Println(s3)

	// Out:
	// PRINT AT github.com/gookit/goutil/dump.TestStruct_WithNested(dump_test.go:278)
	// struct { dump.st1; Github string } {
	//  st1: dump.st1 {
	//    st0: dump.st0 {
	//      Sex: int(2),
	//    },
	//    Age: int(23),
	//    Name: string("inhere"),
	//  },
	//  Github: string("https://github.com/inhere"),
	// }
}

func TestStruct_InterfaceField(t *testing.T) {
	s1 := st1{st0{2}, 23, "inhere"}
	type st2 struct {
		st1
		Github string
		face interface{}
		faces map[string]interface{}
	}

	s2 := st2{
		st1: s1,
		Github: "https://github.com/inhere",
		face: s1,
		faces: map[string]interface{} {
			"key1": 12,
			"key2": "abc",
		},
	}

	Println(s2)
	fmt.Println(s2)
}

func TestMap_Simpled(t *testing.T) {
	m1 := map[int]int{
		23: 12,
		24: 13,
	}

	m2 := map[string]int{
		"key1": 12,
		"key2": 13,
	}

	m3 := map[string]string{
		"key1": "val1",
		"key2": "val2",
	}
	P(m1, m2, m3)
	/*
		Out:
		PRINT AT github.com/gookit/goutil/dump.TestMap_Simpled(dump_test.go:309)
		map[int]int {
		  24: int(13),
		  23: int(12),
		}
		map[string]int {
		  key1: int(12),
		  key2: int(13),
		}
		map[string]string {
		  key1: string("val1"),
		  key2: string("val2"),
		}

	*/

	m4 := map[string]interface{}{
		"key1": 12,
		"key2": "val1",
		"key3": 34,
		"key4": 3.14,
		"key5": -34,
		"key6": nil,
	}
	Print(m4)
	/*
		PRINT AT github.com/gookit/goutil/dump.TestMap_Simpled(dump_test.go:335)
		map[string]interface {} {
		  key4: float64(3.14),
		  key5: int(-34),
		  key6: <nil>,
		  key1: int(12),
		  key2: string("val1"),
		  key3: int(34),
		}
	*/
}

func TestMap_InterfaceNested(t *testing.T) {
	user := &struct {
		id   string
		Name string
		Age  int
	}{"ab1234", "inhere", 22}

	s1 := st1{st0{2}, 23, "inhere"}
	m1 := map[string]interface{}{
		"key1": 112,
		"key2": uint(112),
		"key3": int64(112),
		"key4": 112.23,
		"key5": nil,
		"key6": 'b', // rune
		"key7": byte('a'),
		"st1": s1,
		"user": user,
		"submap1": map[string]int{
			"key1": 12,
			"key2": 13,
		},
		"submap2": map[string]interface{}{
			"key1": 12,
			"key2": "abc",
			"submap21": map[string]string{
				"key1": "val1",
				"key2": "val2",
			},
		},
		"submap3": map[string]interface{}{
			"key1": 12,
			"key2": "abc",
			"submap31": map[string]interface{}{
				"key31": 12,
				"key32": 13,
				"user":  user,
				"submap311": map[string]int{
					"key1": 12,
					"key2": 13,
				},
			},
		},
	}

	Print(m1)
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
