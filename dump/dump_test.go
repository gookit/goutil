package dump

import (
	"bytes"
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

	// Output:
	// PRINT AT github.com/gookit/goutil/dump.ExamplePrint(LINE 13):
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
	str := buf.String()
	is.Contains(str, "\x1b[0;35mPRINT AT github.com/gookit/goutil/dump.TestPrint(LINE ")
	is.Contains(str, "\x1b[0m\nint(123)")

	// disable position for test
	Config.NoPosition = true

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

	ResetConfig()
}
