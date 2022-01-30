package main

import "github.com/gookit/goutil/dump"

// rum demo:
// 	go run ./dump/_examples/slice.go
func main() {
	dump.P(
		[]byte("abc"),
		[]int{1, 2, 3},
		[]string{"ab", "cd"},
		[]interface{}{
			"ab",
			234,
			[]int{1, 3},
			[]string{"ab", "cd"},
		},
	)
}
