package testutil

import (
	"bytes"
	"fmt"
)

// Buffer wrap and extends the bytes.Buffer
type Buffer struct {
	bytes.Buffer
}

// NewBuffer instance
func NewBuffer() *Buffer {
	return &Buffer{}
}

// WriteString rewrite
func (b *Buffer) WriteString(ss ...string) {
	for _, s := range ss {
		_, _ = b.Buffer.WriteString(s)
	}
}

// WriteAny method
func (b *Buffer) WriteAny(vs ...interface{}) {
	for _, v := range vs {
		_, _ = b.Buffer.WriteString(fmt.Sprint(v))
	}
}

// Writeln method
func (b *Buffer) Writeln(s string) {
	_, _ = b.Buffer.WriteString(s)
	_ = b.Buffer.WriteByte('\n')
}

// ResetAndGet buffer string.
func (b *Buffer) ResetAndGet() string {
	s := b.String()
	b.Reset()
	return s
}
