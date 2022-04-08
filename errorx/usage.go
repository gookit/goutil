package errorx

import (
	"errors"
	"fmt"
	"runtime"
)

/*************************************************************
 * new error with call stacks
 *************************************************************/

// New error message
func New(msg string) error {
	return &errorX{
		msg:   msg,
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// Newf error message. alias of Errorf()
func Newf(tpl string, vars ...interface{}) error {
	return &errorX{
		msg:   fmt.Sprintf(tpl, vars...),
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// Errorf error with format message
func Errorf(tpl string, vars ...interface{}) error {
	return &errorX{
		msg:   fmt.Sprintf(tpl, vars...),
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// With prev error
func With(err error, msg string) error {
	return &errorX{
		msg:   msg,
		prev:  err,
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// WithPrev error
func WithPrev(err error, msg string) error {
	return &errorX{
		msg:   msg,
		prev:  err,
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// WithPrevf error and with format message
func WithPrevf(err error, tpl string, vars ...interface{}) error {
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

// Stacked alias of WithStack
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
func WithOptions(msg string, fns ...func(opt *ErrOpt)) error {
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
	return &errorX{
		msg:  msg,
		prev: err,
	}
}

// Wrapf error with format message, but not with stack
func Wrapf(err error, tpl string, vars ...interface{}) error {
	return &errorX{
		msg:  fmt.Sprintf(tpl, vars...),
		prev: err,
	}
}

/*************************************************************
 * helper func for error
 *************************************************************/

// Cause returns the first cause error by call err.Cause().
// Otherwise, will returns current error.
func Cause(err error) error {
	if err == nil {
		return nil
	}

	if err, ok := err.(Causer); ok {
		return err.Cause()
	}
	return err
}

// Unwrap returns previous error by call err.Unwrap().
// Otherwise, will returns nil.
func Unwrap(err error) error {
	if err == nil {
		return nil
	}

	if err, ok := err.(Unwrapper); ok {
		return err.Unwrap()
	}
	return nil
}

// Previous alias of Unwrap()
func Previous(err error) error { return Unwrap(err) }

// Has check err has contains target, or err is eq target.
// alias of errors.Is()
func Has(err, target error) bool {
	return errors.Is(err, target)
}

// Is alias of errors.Is()
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// To try convert err to target, returns is result.
//
// NOTICE: target must be ptr and not nil
//
// alias of errors.As()
func To(err error, target interface{}) bool {
	return errors.As(err, target)
}

// As alias of errors.As()
//
// NOTICE: target must be ptr and not nil
//
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

/*************************************************************
 * helper func for callers stacks
 *************************************************************/

// ErrOpt struct
type ErrOpt struct {
	SkipDepth  int
	TraceDepth int
}

// default option
var stdOpt = newErrOpt()

func newErrOpt() *ErrOpt {
	return &ErrOpt{
		SkipDepth:  3,
		TraceDepth: 20,
	}
}

// Config the stdOpt setting
func Config(fns ...func(opt *ErrOpt)) {
	for _, fn := range fns {
		fn(stdOpt)
	}
}

// SkipDepth setting
func SkipDepth(skipDepth int) func(opt *ErrOpt) {
	return func(opt *ErrOpt) {
		opt.SkipDepth = skipDepth
	}
}

// TraceDepth setting
func TraceDepth(traceDepth int) func(opt *ErrOpt) {
	return func(opt *ErrOpt) {
		opt.TraceDepth = traceDepth
	}
}

func callersStack(skip, depth int) *stack {
	pcs := make([]uintptr, depth)
	num := runtime.Callers(skip, pcs[:])

	var st stack = pcs[0:num]
	return &st
}
