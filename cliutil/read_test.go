package cliutil_test

import (
	"fmt"
	"os"
	"testing"
)

func TestReadFirst(t *testing.T) {
	os.Stdout.Write([]byte("hi"))

	b := make([]byte, 1)
	os.Stdout.Read(b)

	fmt.Println("read", string(b))
}
