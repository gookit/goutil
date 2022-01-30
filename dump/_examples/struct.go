package main

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/gookit/goutil/dump"
)

// rum demo:
// 	go run ./struct.go
// 	go run ./dump/_examples/struct.go
func main() {
	s1 := &struct {
		cannotExport map[string]interface{}
	}{
		cannotExport: map[string]interface{}{
			"key1": 12,
			"key2": "abcd123",
		},
	}

	s2 := struct {
		ab string
		Cd int
	}{
		"ab", 23,
	}

	color.Infoln("- Use fmt.Println:")
	fmt.Println(s1, s2)

	color.Infoln("\n- Use dump.Println:")
	dump.P(
		s1,
		s2,
	)

}
