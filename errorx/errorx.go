// Package errorx provide an enhanced error implements for go,
// allow with stacktraces and wrap another error.
package errorx

import (
	"bytes"
	"errors"
	"fmt"
	"io"
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

// XErrorFace interface
type XErrorFace interface {
	error
	Causer
	Unwrapper
}

// Exception interface
// type Exception interface {
// 	XErrorFace
// 	Code() string
// 	Message() string
// 	StackString() string
// }

/*************************************************************
 * implements XErrorFace interface
 *************************************************************/

// ErrorX struct
//
// TIPS:
//  fmt pkg call order: Format > GoString > Error > String
type ErrorX struct {
	// trace stack
	*stack
	prev error
	msg  string
}

// Cause implements Causer.
func (e *ErrorX) Cause() error {
	if e.prev == nil {
		return e
	}

	if ex, ok := e.prev.(*ErrorX); ok {
		return ex.Cause()
	}
	return e.prev
}

// Unwrap implements Unwrapper.
func (e *ErrorX) Unwrap() error {
	return e.prev
}

// Format error
func (e *ErrorX) Format(s fmt.State, verb rune) {
	// format current error: only output on have msg
	if len(e.msg) > 0 {
		_, _ = io.WriteString(s, e.msg)
		if e.stack != nil {
			e.stack.Format(s, verb)
		}
	}

	// format prev error
	if e.prev == nil {
		return
	}

	_, _ = s.Write([]byte("\n---------\nPrevious: "))
	if ex, ok := e.prev.(*ErrorX); ok {
		ex.Format(s, verb)
	} else {
		_, _ = s.Write([]byte(e.prev.Error()))
	}
}

// GoString to GO string
// printing an error with %#v will produce useful information.
func (e *ErrorX) GoString() string {
	// var sb strings.Builder
	var buf bytes.Buffer
	_, _ = e.WriteTo(&buf)
	return buf.String()
}

// Error to string
func (e *ErrorX) Error() string {
	return e.GoString()
}

// String error to string, with stack trace
func (e *ErrorX) String() string {
	return e.GoString()
}

// WriteTo write the error to a writer
func (e *ErrorX) WriteTo(w io.Writer) (n int64, err error) {
	// current error: only output on have msg
	if len(e.msg) > 0 {
		_, _ = w.Write([]byte(e.msg))

		// with stack
		if e.stack != nil {
			_, _ = e.stack.WriteTo(w)
		}
	}

	// with prev error
	if e.prev != nil {
		_, _ = io.WriteString(w, "\n-------\n")

		if ex, ok := e.prev.(*ErrorX); ok {
			_, _ = ex.WriteTo(w)
		} else {
			_, _ = io.WriteString(w, e.prev.Error())
		}
	}
	return
}

// Message error message
func (e *ErrorX) Message() string {
	return e.msg
}

// StackString returns error stack string.
func (e *ErrorX) StackString() string {
	if e.stack != nil {
		return e.stack.String()
	}
	return ""
}

/*************************************************************
 * new error with call stacks
 *************************************************************/

// New error message and with caller stacks
func New(msg string) error {
	return &ErrorX{
		msg:   msg,
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// Newf error with format message, and with caller stacks.
// alias of Errorf()
func Newf(tpl string, vars ...interface{}) error {
	return &ErrorX{
		msg:   fmt.Sprintf(tpl, vars...),
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// Errorf error with format message, and with caller stacks
func Errorf(tpl string, vars ...interface{}) error {
	return &ErrorX{
		msg:   fmt.Sprintf(tpl, vars...),
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// With prev error and error message, and with caller stacks
func With(err error, msg string) error {
	return &ErrorX{
		msg:   msg,
		prev:  err,
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// Withf error and with format message, and with caller stacks
func Withf(err error, tpl string, vars ...interface{}) error {
	return &ErrorX{
		msg:   fmt.Sprintf(tpl, vars...),
		prev:  err,
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// WithPrev error and message, and with caller stacks. alias of With()
func WithPrev(err error, msg string) error {
	return &ErrorX{
		msg:   msg,
		prev:  err,
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// WithPrevf error and with format message, and with caller stacks. alias of Withf()
func WithPrevf(err error, tpl string, vars ...interface{}) error {
	return &ErrorX{
		msg:   fmt.Sprintf(tpl, vars...),
		prev:  err,
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// WithStack wrap err with a stacked trace. If err is nil, will returns nil.
func WithStack(err error) error {
	if err == nil {
		return nil
	}
	return &ErrorX{
		msg: err.Error(),
		// prev:  err,
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// Stacked warp a error and with caller stacks. alias of WithStack()
func Stacked(err error) error {
	if err == nil {
		return nil
	}
	return &ErrorX{
		msg: err.Error(),
		// prev:  err,
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// WithOptions new error with some option func
func WithOptions(msg string, fns ...func(opt *ErrStackOpt)) error {
	opt := newErrOpt()
	for _, fn := range fns {
		fn(opt)
	}

	return &ErrorX{
		msg:   msg,
		stack: callersStack(opt.SkipDepth, opt.TraceDepth),
	}
}

/*************************************************************
 * helper func for wrap error without stacks
 *************************************************************/

// Wrap error and with message, but not with stack
func Wrap(err error, msg string) error {
	if err == nil {
		return errors.New(msg)
	}

	return &ErrorX{
		msg:  msg,
		prev: err,
	}
}

// Wrapf error with format message, but not with stack
func Wrapf(err error, tpl string, vars ...interface{}) error {
	if err == nil {
		return errors.New(fmt.Sprintf(tpl, vars...))
	}

	return &ErrorX{
		msg:  fmt.Sprintf(tpl, vars...),
		prev: err,
	}
}
