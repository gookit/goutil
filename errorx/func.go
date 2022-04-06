package errorx

import (
	"fmt"
	"runtime"
)

// Cause returns the first cause error
func Cause(err error) error {
	if err == nil {
		return err
	}

	if err, ok := err.(Causer); ok {
		if cause := err.Cause(); cause != nil {
			return cause
		}
	}
	return err
}

// Unwrap error
func Unwrap(err error) error {
	return Cause(err)
}

var (
	skipDepth  = 1
	traceDepth = 20
)

var stdOpt = newErrOpt()

// ConfigOpt setting
func ConfigOpt(fns ...func(opt *ErrOpt)) {
	for _, fn := range fns {
		fn(stdOpt)
	}
}

// ErrOpt struct
type ErrOpt struct {
	SkipDepth  int
	TraceDepth int
}

func newErrOpt() *ErrOpt {
	return &ErrOpt{
		SkipDepth:  3,
		TraceDepth: 20,
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

// NewWith some option func
func NewWith(msg string, fns ...func(opt *ErrOpt)) error {
	opt := newErrOpt()
	for _, fn := range fns {
		fn(opt)
	}

	return &errorX{
		msg:   msg,
		stack: callersStack(opt.SkipDepth, opt.TraceDepth),
	}
}

// New error message
func New(msg string) error {
	return &errorX{
		msg:   msg,
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// Newf error message
func Newf(tpl string, vars ...interface{}) error {
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

// WithPrevf error
func WithPrevf(err error, tpl string, vars ...interface{}) error {
	return &errorX{
		msg:   fmt.Sprintf(tpl, vars...),
		prev:  err,
		stack: callersStack(stdOpt.SkipDepth, stdOpt.TraceDepth),
	}
}

// WithStack annotates err with a stack trace at the point WithStack was called.
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

// Wrap error and with message, but not with stack
func Wrap(err error, msg string) error {
	return &errorX{
		msg:  msg,
		prev: err,
	}
}

// Wrapf error and with message, but not with stack
func Wrapf(err error, tpl string, vars ...interface{}) error {
	return &errorX{
		msg:  fmt.Sprintf(tpl, vars...),
		prev: err,
	}
}

func callersStack(skip, depth int) *stack {
	// var pcs [traceDepth]uintptr
	pcs := make([]uintptr, depth)
	num := runtime.Callers(skip, pcs[:])

	var st stack = pcs[0:num]
	return &st
}
