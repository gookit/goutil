package dump

import (
	"bytes"
	"os"
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

func TestPrint(t *testing.T) {
	is := assert.New(t)
	buf := new(bytes.Buffer)

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
	buf := new(bytes.Buffer)

	// set output for test
	Output = buf
	// disable position for test
	Config.ShowFlag = Fnopos

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

	// reset
	resetDump()
}

func TestPrintPtr(t *testing.T) {
	user := &struct {
		id   string
		Name string
		Age  int
	}{}
	P(user)

	// Out:
	// PRINT AT github.com/gookit/goutil/dump.TestPrintPtr(dump_test.go:157)
	// *struct { id string; Name string; Age int } {
	//  id: "",
	//  Name: "",
	//  Age: 0,
	// }
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
	is.Contains(buf.String(), "goutil/dump/dump_test.go:1")
	buf.Reset()

	// reset
	resetDump()
}

func resetDump() {
	Output = os.Stdout
	ResetConfig()
}