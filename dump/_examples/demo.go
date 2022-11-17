package main

import (
	"time"

	"github.com/gookit/goutil/dump"
)

// rum demo:
//
//	go run ./dump/_examples/demo.go
func main() {
	val := map[string]any{
		"bool":   true,
		"number": 1 + 1i,
		"bytes":  []byte{97, 98, 99},
		"lines":  "multiline string\nline two",
		"slice":  []any{1, 2},
		"time":   time.Now(),
		"struct": struct{ test int32 }{
			test: 13,
		},
	}
	val["slice"].([]any)[1] = val["slice"]
	dump.P(val)
	return
	// dump.Config.ShowFile = true
	// dump.Config.ShowMethod = true
	otherFunc()
}

func otherFunc() {
	dump.P(234, int64(56))
	dump.P("abc", "def")
	dump.P([]string{"ab", "cd"})
	dump.P(
		[]any{"ab", 234, []int{1, 3}},
	)
}
