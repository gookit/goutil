package errorx

import (
	"fmt"
	"strconv"
	"strings"
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
	return e.msg + "(code: " + strconv.FormatInt(int64(e.code), 10) + ")"
}

// GoString get.
func (e *errorR) GoString() string {
	return e.String()
}

// ErrMap type
type ErrMap map[string]error

// Error string
func (e ErrMap) Error() string {
	var sb strings.Builder
	for name, err := range e {
		sb.WriteString(name)
		sb.WriteByte(':')
		sb.WriteString(err.Error())
		sb.WriteByte('\n')
	}
	return sb.String()
}

// IsEmpty error
func (e ErrMap) IsEmpty() bool {
	return len(e) == 0
}

// One error
func (e ErrMap) One() error {
	for _, err := range e {
		return err
	}
	return nil
}

// ErrList type
type ErrList []error

// Error string
func (el ErrList) Error() string {
	var sb strings.Builder
	for _, err := range el {
		sb.WriteString(err.Error())
		sb.WriteByte('\n')
	}
	return sb.String()
}

// IsEmpty error
func (el ErrList) IsEmpty() bool {
	return len(el) == 0
}

// First error
func (el ErrList) First() error {
	if len(el) > 0 {
		return el[0]
	}
	return nil
}
