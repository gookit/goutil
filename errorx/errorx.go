// Package errorx provide an enhanced error implements,
// allow with call stack and wrap another error.
//
// refer there are packages: errgo.v2, pkg/errors, joomcode/errorx
package errorx

import (
	"fmt"
	"io"
	"strings"
)

// Causer interface
type Causer interface {
	Cause() error
}

// errorX struct
type errorX struct {
	prev error
	msg  string

	// trace stack
	*stack
}

// Cause implements Causer.
func (e *errorX) Cause() error {
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
		msg = msg + "\nTRACE:\n" + e.stack.TraceString()
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
