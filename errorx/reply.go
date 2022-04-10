package errorx

import (
	"fmt"

	"github.com/gookit/goutil/mathutil"
)

// ErrorCoder interface
type ErrorCoder interface {
	error
	Code() int
}

// ErrorR useful for web service replay/response.
// code == 0 is successful. otherwise, is failed.
type ErrorR interface {
	ErrorCoder
	fmt.Stringer
	IsSuc() bool
	IsFail() bool
}

// error reply struct
type errorR struct {
	code int
	msg  string
}

// NewR code with error response
func NewR(code int, msg string) ErrorR {
	return &errorR{code: code, msg: msg}
}

// Fail code with error response
func Fail(code int, msg string) ErrorR {
	return &errorR{code: code, msg: msg}
}

// Suc success response reply
func Suc(msg string) ErrorR {
	return &errorR{code: 0, msg: msg}
}

// IsSuc code value check
func (e *errorR) IsSuc() bool {
	return e.code == 0
}

// IsFail code value check
func (e *errorR) IsFail() bool {
	return e.code != 0
}

// Code value
func (e *errorR) Code() int {
	return e.code
}

// Error string
func (e *errorR) Error() string {
	return e.msg
}

// String get
func (e *errorR) String() string {
	return e.msg + "(code: " + mathutil.String(e.code) + ")"
}

// GoString get.
func (e *errorR) GoString() string {
	return e.String()
}
