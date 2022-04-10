// Package errorx provide an enhanced error implements for go,
// allow with stacktraces and wrap another error.
package errorx

import (
	"errors"
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
//
// TIPS:
//  fmt pkg call order: Format > GoString > Error > String
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

// Each previous errors
// func (e *errorX) Each(fn func(err error)) {
// 	return e.prev
// }

// WriteTo write the error to a writer
func (e *errorX) WriteTo(w io.Writer) (n int64, err error) {
	_, _ = w.Write([]byte(e.msg))

	// with stack
	if e.stack != nil {
		_, _ = e.stack.WriteTo(w)
	}

	// with prev error
	if e.prev != nil {
		_, _ = io.WriteString(w, "\n----------------------------------\n")

		if ex, ok := e.prev.(*errorX); ok {
			_, _ = ex.WriteTo(w)
		} else {
			_, _ = io.WriteString(w, e.prev.Error())
		}
	}
	return
}

// GoString to GO string
// printing an error with %#v will produce useful information.
func (e *errorX) GoString() string {
	var sb strings.Builder

	_, _ = e.WriteTo(&sb)
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

	if e.stack != nil {
		e.stack.Format(s, verb)
	}

	e.formatPrev(s, verb)
}

// formatPrev error
func (e *errorX) formatPrev(s fmt.State, verb rune) {
	if e.prev == nil {
		return
	}

	_, _ = s.Write([]byte("\nPrevious: "))
	if ex, ok := e.prev.(*errorX); ok {
		ex.Format(s, verb)
	} else {
		_, _ = s.Write([]byte(e.prev.Error()))
	}
}

/*************************************************************
 * new error with call stacks
 *************************************************************/

// New error message and with caller stacks
func New(msg string) error {
	return &errorX{
		msg:   msg,
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// Newf error with format message, and with caller stacks.
// alias of Errorf()
func Newf(tpl string, vars ...interface{}) error {
	return &errorX{
		msg:   fmt.Sprintf(tpl, vars...),
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// Errorf error with format message, and with caller stacks
func Errorf(tpl string, vars ...interface{}) error {
	return &errorX{
		msg:   fmt.Sprintf(tpl, vars...),
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// With prev error and error message, and with caller stacks
func With(err error, msg string) error {
	if err == nil {
		return nil
	}

	return &errorX{
		msg:   msg,
		prev:  err,
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// Withf error and with format message, and with caller stacks
func Withf(err error, tpl string, vars ...interface{}) error {
	if err == nil {
		return nil
	}

	return &errorX{
		msg:   fmt.Sprintf(tpl, vars...),
		prev:  err,
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// WithPrev error and message, and with caller stacks
func WithPrev(err error, msg string) error {
	if err == nil {
		return nil
	}

	return &errorX{
		msg:   msg,
		prev:  err,
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// WithPrevf error and with format message, and with caller stacks
func WithPrevf(err error, tpl string, vars ...interface{}) error {
	if err == nil {
		return nil
	}

	return &errorX{
		msg:   fmt.Sprintf(tpl, vars...),
		prev:  err,
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// WithStack annotates err with a stacked trace at the point WithStack was called.
// If err is nil, WithStack returns nil.
func WithStack(err error) error {
	if err == nil {
		return nil
	}

	return &errorX{
		prev:  err,
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// Stacked warp a error and with caller stacks.
// alias of WithStack
func Stacked(err error) error {
	if err == nil {
		return nil
	}

	return &errorX{
		prev:  err,
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// WithOptions new error with some option func
func WithOptions(msg string, fns ...func(opt *ErrStackOpt)) error {
	opt := newErrOpt()
	for _, fn := range fns {
		fn(opt)
	}

	return &errorX{
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

	return &errorX{
		msg:  msg,
		prev: err,
	}
}

// Wrapf error with format message, but not with stack
func Wrapf(err error, tpl string, vars ...interface{}) error {
	if err == nil {
		return errors.New(fmt.Sprintf(tpl, vars...))
	}

	return &errorX{
		msg:  fmt.Sprintf(tpl, vars...),
		prev: err,
	}
}
