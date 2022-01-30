package main

import "github.com/gookit/goutil/dump"

// rum demo:
// 	go run ./dump/_examples/basic_types.go
func main() {
	dump.P(
		nil, true,
		12, int8(12), int16(12), int32(12), int64(12),
		uint(22), uint8(22), uint16(22), uint32(22), uint64(22),
		float32(23.78), float64(56.45),
		'c', byte('d'),
		"string",
	)
}
