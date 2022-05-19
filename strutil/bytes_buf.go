package strutil

import (
	"bytes"
	"fmt"
)

// Buffer wrap and extends the bytes.Buffer
type Buffer struct {
	*bytes.Buffer
}

// NewEmptyBuffer instance
func NewEmptyBuffer() *Buffer {
	return &Buffer{
		Buffer: new(bytes.Buffer),
	}
}

// QuietWritef write message to buffer
func (b *Buffer) QuietWritef(tpl string, vs ...interface{}) {
	_, _ = b.WriteString(fmt.Sprintf(tpl, vs...))
}

// QuietWriteln write message to buffer with newline
func (b *Buffer) QuietWriteln(s string) {
	_, _ = b.WriteString(s)
	_ = b.WriteByte('\n')
}

// QuietWriteString to buffer
func (b *Buffer) QuietWriteString(ss ...string) {
	for _, s := range ss {
		_, _ = b.WriteString(s)
	}
}

// MustWriteString to buffer
func (b *Buffer) MustWriteString(ss ...string) {
	for _, s := range ss {
		_, err := b.WriteString(s)
		if err != nil {
			panic(err)
		}
	}
}
