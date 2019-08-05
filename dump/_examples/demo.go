package main

import "github.com/gookit/goutil/dump"

// rum demo:
func main() {
	dump.P(
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

	otherFunc()
}

func otherFunc() {
	dump.P(234, int64(56))
	dump.P("abc", "def")
	dump.P([]string{"ab", "cd"})
	dump.P([]interface{}{"ab", 234, []int{1, 3}})
}
