package testutil

import (
	"bytes"
	"sync"

	"github.com/gookit/goutil/byteutil"
)

// Buffer wrap and extends the bytes.Buffer
type Buffer = byteutil.Buffer

// NewBuffer instance
func NewBuffer() *byteutil.Buffer {
	return byteutil.NewBuffer()
}

// SafeBuffer Thread-safe buffer for testing
type SafeBuffer struct {
	bytes.Buffer
	mu sync.Mutex
}

// NewSafeBuffer instance
func NewSafeBuffer() *SafeBuffer {
	return &SafeBuffer{Buffer: bytes.Buffer{}}
}

// Write implements io.Writer
func (sb *SafeBuffer) Write(p []byte) (n int, err error) {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	return sb.Buffer.Write(p)
}

// WriteString implements io.StringWriter
func (sb *SafeBuffer) WriteString(s string) (n int, err error) {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	return sb.Buffer.WriteString(s)
}

// ResetGet get buffer content and reset
func (sb *SafeBuffer) ResetGet() string {
	s := sb.String()
	sb.Reset()
	return s
}
