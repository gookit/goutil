package main

import "github.com/gookit/goutil/dump"

// rum demo:
// 	go run ./dump/_examples/demo.go
func main() {
	// dump.Config.ShowFile = true
	// dump.Config.ShowMethod = true
	otherFunc()
}

func otherFunc() {
	dump.P(234, int64(56))
	dump.P("abc", "def")
	dump.P([]string{"ab", "cd"})
	dump.P(
		[]interface{}{"ab", 234, []int{1, 3}},
	)
}
