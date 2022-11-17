package main

import (
	"github.com/kortschak/utter"
	"github.com/kr/pretty"
)

// rum demo:
//
//	go run ./refer_kr_pretty.go
//	go run ./dump/_examples/refer_kr_pretty.go
func main() {
	vs := []any{
		23,
		[]string{"ab", "cd"},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, // len > 10
		map[string]any{
			"key": "val", "sub": map[string]string{"k": "v"},
		},
		struct {
			ab string
			Cd int
		}{
			"ab", 23,
		},
	}

	// print var data
	_, err := pretty.Println(vs...)
	if err != nil {
		panic(err)
	}

	// print var data
	for _, v := range vs {
		utter.Dump(v)
	}
}
