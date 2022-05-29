package common

import (
	"fmt"
	"io"
)

// ByteStringWriter interface
type ByteStringWriter interface {
	io.Writer
	io.ByteWriter
	io.StringWriter
	fmt.Stringer
}
