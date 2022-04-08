package errorx

import "github.com/gookit/goutil/mathutil"

// ErrorH useful for web service replay/response.
// code == 0 is successful. otherwise, is failed.
type ErrorH interface {
	error
	Code() int
	IsOk() bool
	IsErr() bool
}

type errorH struct {
	code int
	msg  string
}

// NewH error response
func NewH(code int, msg string) ErrorH {
	return &errorH{code: code, msg: msg}
}

// OkH success response
func OkH(msg string) ErrorH {
	return &errorH{code: 0, msg: msg}
}

// IsOk code value
func (e *errorH) IsOk() bool {
	return e.code == 0
}

// IsErr code value
func (e *errorH) IsErr() bool {
	return e.code != 0
}

// Code value
func (e *errorH) Code() int {
	return e.code
}

// Error string
func (e *errorH) Error() string {
	return e.msg
}

// String get
func (e *errorH) String() string {
	return e.msg + "(code: " + mathutil.String(e.code) + ")"
}
