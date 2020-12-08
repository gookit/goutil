package strutil

import (
	"fmt"
	"testing"
)

func TestMicroTimeID(t *testing.T) {
	for i := 0; i < 10; i++ {
		id := MicroTimeID()
		fmt.Println(id)
	}
}

func TestMicroTimeHexID(t *testing.T) {
	for i := 0; i < 10; i++ {
		id := MicroTimeHexID()
		fmt.Println(id, "len:", len(id))
	}
}
