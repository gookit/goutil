package fsutil

import (
	"io"
	"strings"
)

// QuietWriteString to writer
func QuietWriteString(w io.Writer, ss ...string) {
	_, _ = io.WriteString(w, strings.Join(ss, ""))
}

// WriteWrapper struct
type WriteWrapper struct {
	Out io.Writer
}

// NewWriteWrapper instance
func NewWriteWrapper(w io.Writer) *WriteWrapper {
	return &WriteWrapper{Out: w}
}

// Write bytes data
func (w *WriteWrapper) Write(p []byte) (n int, err error) {
	return w.Out.Write(p)
}

// WriteByte data
func (w *WriteWrapper) WriteByte(c byte) error {
	_, err := w.Out.Write([]byte{c})
	return err
}

// WriteString data
func (w *WriteWrapper) WriteString(s string) (n int, err error) {
	return w.Out.Write([]byte(s))
}
