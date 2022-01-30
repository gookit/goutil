package main

import "github.com/gookit/goutil/dump"

// rum demo:
// 	go run ./map.go
// 	go run ./dump/_examples/map.go
func main() {
	dump.P(
		map[string]interface{}{
			"key0": 123,
			"key1": "value1",
			"key2": []int{1, 2, 3},
			"key3": map[string]string{
				"k0": "v0",
				"k1": "v1",
			},
		},
	)
}
