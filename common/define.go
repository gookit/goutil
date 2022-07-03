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

// StringWriteStringer interface
type StringWriteStringer interface {
	io.StringWriter
	fmt.Stringer
}

type (
	// MarshalFunc define
	MarshalFunc func(v interface{}) ([]byte, error)

	// UnmarshalFunc define
	UnmarshalFunc func(bts []byte, ptr interface{}) error
)
