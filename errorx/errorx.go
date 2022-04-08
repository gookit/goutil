// Package errorx provide an enhanced error implements for go,
// allow with stacktraces and wrap another error.
package errorx

import (
	"fmt"
	"io"
	"strings"
)

// Causer interface for get first cause error
type Causer interface {
	// Cause returns the first cause error by call err.Cause().
	// Otherwise, will returns current error.
	Cause() error
}

// Unwrapper interface for get previous error
type Unwrapper interface {
	// Unwrap returns previous error by call err.Unwrap().
	// Otherwise, will returns nil.
	Unwrap() error
}

// ErrorX interface
type ErrorX interface {
	error
	Causer
	Unwrapper
}

/*************************************************************
 * implements ErrorX interface
 *************************************************************/

// errorX struct
type errorX struct {
	// trace stack
	*stack
	prev error
	msg  string
}

// Cause implements Causer.
func (e *errorX) Cause() error {
	if e.prev == nil {
		return e
	}

	if ex, ok := e.prev.(*errorX); ok {
		return ex.Cause()
	}
	return e.prev
}

// Unwrap implements Unwrapper.
func (e *errorX) Unwrap() error {
	return e.prev
}

// GoString to GO string
// printing an error with %#v will produce useful information.
func (e *errorX) GoString() string {
	var sb strings.Builder
	sb.WriteString(e.msg)

	if e.stack != nil {
		sb.WriteString("\nTRACE:\n")
		_, _ = e.stack.WriteTo(&sb)

		if e.prev != nil {
			sb.WriteString("\n----------------------------------\n")

			if ex, ok := e.prev.(*errorX); ok {
				_, _ = ex.WriteTo(&sb)
			} else {
				sb.WriteString(e.prev.Error())
			}
		}
	}

	return sb.String()
}

// String error to string, with stack trace
func (e *errorX) String() string {
	return e.GoString()
}

// Error to string
func (e *errorX) Error() string {
	msg := e.msg
	if e.stack != nil {
		msg = msg + e.stack.String()
	}

	return msg
}

// Format error
func (e *errorX) Format(s fmt.State, verb rune) {
	_, _ = io.WriteString(s, e.msg)

	if e.stack == nil {
		e.formatPrev(s, verb)
		return
	}

	e.stack.Format(s, verb)
	e.formatPrev(s, verb)
	// switch verb {
	// case 'v', 's':
	// 	e.stack.Format(s, verb)
	// 	e.formatPrev(s, verb)
	// }
}

// formatPrev error
func (e *errorX) formatPrev(s fmt.State, verb rune) {
	if e.prev == nil {
		return
	}

	_, _ = s.Write([]byte("\nPrevious: "))
	// _, _ = io.WriteString(s, "\nPrevious:")

	if ex, ok := e.prev.(*errorX); ok {
		ex.Format(s, verb)
	} else {
		_, _ = io.WriteString(s, e.prev.Error())
	}
}
