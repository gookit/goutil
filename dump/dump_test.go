package dump

import (
	"testing"
)

func TestPrintln(t *testing.T) {
	// is := assert.New(t)

	// buf := new(bytes.Buffer)

	Print(
		23,
		[]string{"ab", "cd"},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		map[string]interface{}{
			"key": "val",
			"sub": map[string]string{"k": "v"},
		},
		struct {
			ab string
			Cd int
		}{
			"ab", 23,
		},
	)
}
