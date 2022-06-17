package strutil

import (
	"bytes"
	"fmt"
	"strings"
)

// Buffer wrap and extends the bytes.Buffer
type Buffer struct {
	bytes.Buffer
}

// NewEmptyBuffer instance
func NewEmptyBuffer() *Buffer {
	return &Buffer{}
}

// QuietWriteByte to buffer
func (b *Buffer) QuietWriteByte(c byte) {
	_ = b.WriteByte(c)
}

// QuietWritef write message to buffer
func (b *Buffer) QuietWritef(tpl string, vs ...interface{}) {
	_, _ = b.WriteString(fmt.Sprintf(tpl, vs...))
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
