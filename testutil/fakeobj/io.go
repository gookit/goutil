package fakeobj

import (
	"errors"

	"github.com/gookit/goutil/byteutil"
)

// IOWriter only implements the io.Writer
type IOWriter struct {
	Buf []byte
	// ErrOnWrite return error on write, useful for testing
	ErrOnWrite bool
}

// NewIOWriter instance
func NewIOWriter() *IOWriter {
	return &IOWriter{
		Buf: make([]byte, 0, 1024),
	}
}

// Write implements
func (w *IOWriter) Write(p []byte) (n int, err error) {
	if w.ErrOnWrite {
		return 0, errors.New("fake write error")
	}

	w.Buf = append(w.Buf, p...)
	return len(p), nil
}

// Reset the buffer
func (w *IOWriter) Reset() {
	w.Buf = w.Buf[:0]
}

// Reader implements the io.Reader, io.Closer
type Reader struct {
	byteutil.Buffer
	// ErrOnRead return error on read, useful for testing
	ErrOnRead bool
}

// NewReader instance
func NewReader() *Reader {
	return &Reader{}
}

// NewStrReader instance
func NewStrReader(s string) *Reader {
	buf := byteutil.NewBuffer()
	buf.WriteStr1(s)

	return &Reader{
		Buffer: *buf,
	}
}

// SetErrOnRead mark
func (r *Reader) SetErrOnRead() {
	r.ErrOnRead = true
}

// Read implements the io.Reader
func (r *Reader) Read(p []byte) (n int, err error) {
	if r.ErrOnRead {
		return 0, errors.New("fake read error")
	}
	return r.Buffer.Read(p)
}

// Writer implements the io.Writer, stdio.Flusher, io.Closer.
type Writer struct {
	byteutil.Buffer
	// ErrOnWrite return error on write, useful for testing
	ErrOnWrite bool
	// ErrOnFlush return error on flush, useful for testing
	ErrOnFlush bool
	// ErrOnSync return error on flush, useful for testing
	ErrOnSync bool
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

// SetErrOnSync method
func (w *Writer) SetErrOnSync() *Writer {
	w.ErrOnSync = true
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

// Flush implements stdio.Flusher
func (w *Writer) Flush() error {
	if w.ErrOnFlush {
		return errors.New("fake flush error")
	}

	w.Reset()
	return nil
}

// Sync implements stdio.Syncer
func (w *Writer) Sync() error {
	if w.ErrOnSync {
		return errors.New("fake sync error")
	}

	w.Reset()
	return nil
}
