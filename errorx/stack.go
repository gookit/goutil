package errorx

import (
	"bytes"
	"fmt"
	"io"
	"runtime"
	"strconv"
)

// stack represents a stack of program counters.
type stack []uintptr

// Format stack trace
func (s *stack) Format(fs fmt.State, verb rune) {
	switch verb {
	case 'v', 's':
		_, _ = s.WriteTo(fs)
	}
}

// StackLen for error
func (s *stack) StackLen() int {
	return len(*s)
}

// WriteTo for error
func (s *stack) WriteTo(w io.Writer) (int64, error) {
	if len(*s) == 0 {
		return 0, nil
	}

	nn, _ := w.Write([]byte("\nSTACK:\n"))
	for _, pc := range *s {
		// For historical reasons if pc is interpreted as a uintptr
		// its value represents the program counter + 1.
		fc := runtime.FuncForPC(pc - 1)
		if fc == nil {
			continue
		}

		// file eg: workspace/godev/gookit/goutil/errorx/errorx_test.go
		file, line := fc.FileLine(pc - 1)
		// f.Name() eg: github.com/gookit/goutil/errorx_test.TestWithPrev()
		location := fc.Name() + "()\n  " + file + ":" + strconv.Itoa(line) + "\n"

		n, _ := w.Write([]byte(location))
		nn += n
	}

	return int64(nn), nil
}

// String format to string
func (s *stack) String() string {
	var buf bytes.Buffer
	_, _ = s.WriteTo(&buf)
	return buf.String()
}

// StackFrames stack frame list
func (s *stack) StackFrames() *runtime.Frames {
	return runtime.CallersFrames(*s)
}

// Location for error report
func (s *stack) Location() (file string, line int) {
	if len(*s) > 0 {
		// For historical reasons if pc is interpreted as a uintptr
		// its value represents the program counter + 1.
		pc := (*s)[0] - 1
		fc := runtime.FuncForPC(pc)
		if fc != nil {
			return fc.FileLine(pc)
		}
	}
	return
}

// CallerName for error report.
//
// Returns eg:
// 	github.com/gookit/goutil/errorx_test.TestWithPrev()
func (s *stack) CallerName() (fcName string) {
	if len(*s) > 0 {
		// For historical reasons if pc is interpreted as a uintptr
		// its value represents the program counter + 1.
		pc := (*s)[0] - 1
		fc := runtime.FuncForPC(pc)
		if fc != nil {
			return fc.Name()
		}
	}
	return
}

/*************************************************************
 * helper func for callers stacks
 *************************************************************/

// ErrStackOpt struct
type ErrStackOpt struct {
	SkipDepth  int
	TraceDepth int
}

// default option
var stdOpt = newErrOpt()

func newErrOpt() *ErrStackOpt {
	return &ErrStackOpt{
		SkipDepth:  3,
		TraceDepth: 15,
	}
}

// Config the stdOpt setting
func Config(fns ...func(opt *ErrStackOpt)) {
	for _, fn := range fns {
		fn(stdOpt)
	}
}

// SkipDepth setting
func SkipDepth(skipDepth int) func(opt *ErrStackOpt) {
	return func(opt *ErrStackOpt) {
		opt.SkipDepth = skipDepth
	}
}

// TraceDepth setting
func TraceDepth(traceDepth int) func(opt *ErrStackOpt) {
	return func(opt *ErrStackOpt) {
		opt.TraceDepth = traceDepth
	}
}

func callersStack(skip, depth int) *stack {
	pcs := make([]uintptr, depth)
	num := runtime.Callers(skip, pcs[:])

	var st stack = pcs[0:num]
	return &st
}
