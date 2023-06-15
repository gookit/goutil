package byteutil

import (
	"bytes"
	"fmt"
	"strings"
)

// Buffer wrap and extends the bytes.Buffer
type Buffer struct {
	bytes.Buffer
}

// NewBuffer instance
func NewBuffer() *Buffer {
	return &Buffer{}
}

// WriteAny type value to buffer
func (b *Buffer) WriteAny(vs ...any) {
	for _, v := range vs {
		_, _ = b.Buffer.WriteString(fmt.Sprint(v))
	}
}

// QuietWriteByte to buffer
func (b *Buffer) QuietWriteByte(c byte) {
	_ = b.WriteByte(c)
}

// QuietWritef write message to buffer
func (b *Buffer) QuietWritef(tpl string, vs ...any) {
	_, _ = b.WriteString(fmt.Sprintf(tpl, vs...))
}

// Writeln write message to buffer with newline
func (b *Buffer) Writeln(ss ...string) {
	b.QuietWriteln(ss...)
}

// QuietWriteln write message to buffer with newline
func (b *Buffer) QuietWriteln(ss ...string) {
	_, _ = b.WriteString(strings.Join(ss, ""))
	_ = b.WriteByte('\n')
}

// QuietWriteString to buffer
func (b *Buffer) QuietWriteString(ss ...string) {
	_, _ = b.WriteString(strings.Join(ss, ""))
}

// MustWriteString to buffer
func (b *Buffer) MustWriteString(ss ...string) {
	_, err := b.WriteString(strings.Join(ss, ""))
	if err != nil {
		panic(err)
	}
}

// ResetGet buffer string. alias of ResetAndGet()
func (b *Buffer) ResetGet() string {
	return b.ResetAndGet()
}

// ResetAndGet buffer string.
func (b *Buffer) ResetAndGet() string {
	s := b.String()
	b.Reset()
	return s
}

// Close buffer
func (b *Buffer) Close() error {
	return nil
}

// Flush buffer
func (b *Buffer) Flush() error {
	return nil
}
