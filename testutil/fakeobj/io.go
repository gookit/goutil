package fakeobj

import (
	"bytes"
	"errors"
)

// Writer implements the io.Writer, io.Flusher, io.Closer.
type Writer struct {
	bytes.Buffer
	// ErrOnWrite return error on write, useful for testing
	ErrOnWrite bool
	// ErrOnFlush return error on flush, useful for testing
	ErrOnFlush bool
	// ErrOnClose return error on close, useful for testing
	ErrOnClose bool
}

// NewBuffer instance. alias of NewWriter()
func NewBuffer() *Writer {
	return NewWriter()
}

// NewWriter instance
func NewWriter() *Writer {
	return &Writer{}
}

// SetErrOnWrite method
func (w *Writer) SetErrOnWrite() *Writer {
	w.ErrOnWrite = true
	return w
}

// SetErrOnFlush method
func (w *Writer) SetErrOnFlush() *Writer {
	w.ErrOnFlush = true
	return w
}

// SetErrOnClose method
func (w *Writer) SetErrOnClose() *Writer {
	w.ErrOnClose = true
	return w
}

// ResetGet buffer string.
func (w *Writer) ResetGet() string {
	s := w.String()
	w.Reset()
	return s
}

// Write implements
func (w *Writer) Write(p []byte) (n int, err error) {
	if w.ErrOnWrite {
		return 0, errors.New("fake write error")
	}
	return w.Buffer.Write(p)
}

// Close implements
func (w *Writer) Close() error {
	if w.ErrOnClose {
		return errors.New("fake close error")
	}
	return nil
}

// Flush implements
func (w *Writer) Flush() error {
	if w.ErrOnFlush {
		return errors.New("fake flush error")
	}

	w.Reset()
	return nil
}
