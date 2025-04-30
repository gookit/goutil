package main

import (
	"fmt"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/x/ccolor"
)

// rum demo:
//
//	go run ./struct.go
//	go run ./dump/_examples/struct.go
func main() {
	s1 := &struct {
		cannotExport map[string]any
	}{
		cannotExport: map[string]any{
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

	ccolor.Infoln("- Use fmt.Println:")
	fmt.Println(s1, s2)

	ccolor.Infoln("\n- Use dump.Println:")
	dump.P(s1, s2)
}
