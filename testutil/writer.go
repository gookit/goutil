package testutil

import (
	"github.com/gookit/goutil/errorx"
)

// TestWriter struct, useful for testing
type TestWriter struct {
	Buffer
	// ErrOnWrite return error on write, useful for testing
	ErrOnWrite bool
	// ErrOnFlush return error on flush, useful for testing
	ErrOnFlush bool
	// ErrOnClose return error on close, useful for testing
	ErrOnClose bool
}

// NewTestWriter instance
func NewTestWriter() *TestWriter {
	return &TestWriter{}
}

// SetErrOnWrite method
func (w *TestWriter) SetErrOnWrite() *TestWriter {
	w.ErrOnWrite = true
	return w
}

// SetErrOnFlush method
func (w *TestWriter) SetErrOnFlush() *TestWriter {
	w.ErrOnFlush = true
	return w
}

// SetErrOnClose method
func (w *TestWriter) SetErrOnClose() *TestWriter {
	w.ErrOnClose = true
	return w
}

// Flush implements
func (w *TestWriter) Flush() error {
	if w.ErrOnFlush {
		return errorx.Raw("flush error")
	}

	w.Reset()
	return nil
}

// Close implements
func (w *TestWriter) Close() error {
	if w.ErrOnClose {
		return errorx.Raw("close error")
	}
	return nil
}

// Write implements
func (w *TestWriter) Write(p []byte) (n int, err error) {
	if w.ErrOnWrite {
		return 0, errorx.Raw("write error")
	}
	return w.Buffer.Write(p)
}
