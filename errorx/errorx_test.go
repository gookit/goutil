package errorx_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/testutil/assert"
)

func returnErr(msg string) error {
	return errorx.Raw(msg)
}

func returnErrL2(msg string) error {
	return returnErr(msg)
}

func returnXErr(msg string) error {
	return errorx.New(msg)
}

func returnXErrL2(msg string) error {
	return returnXErr(msg)
}

func TestNew(t *testing.T) {
	err := returnXErrL2("the error message")
	assert.Err(t, err)

	fmt.Println(err)
	// fmt.Printf("%v\n", err)
	// fmt.Printf("%#v\n", err)
}

func TestRawGoErr(t *testing.T) {
	assert.Err(t, errorx.E("error message"))
	assert.Err(t, errorx.Err("error message"))
	assert.Err(t, errorx.Raw("error message"))
	assert.Err(t, errorx.Ef("error %s", "message"))
	assert.Err(t, errorx.Errf("error %s", "message"))
	assert.Err(t, errorx.Rawf("error %s", "message"))
}

func TestNewf(t *testing.T) {
	err := errorx.Newf("error %s", "message")
	assert.Err(t, err)
	fmt.Printf("%+v\n", err)

	err = errorx.Errorf("error %s", "message")
	assert.Err(t, err)
	fmt.Printf("%#v\n", err)
}

func TestWith_goerr(t *testing.T) {
	err1 := returnErr("first error message")
	assert.Err(t, err1)

	err2 := errorx.With(err1, "second error message")
	assert.Err(t, err2)
	assert.True(t, errorx.Has(err2, err1))
	assert.True(t, errorx.Is(err2, err1))

	fmt.Println(err2)
	// fmt.Printf("%v\n", err2)
}

func TestWith_errorx(t *testing.T) {
	err1 := returnXErr("first error message")
	assert.Err(t, err1)

	err2 := errorx.With(err1, "second error message")
	assert.Err(t, err2)
	assert.True(t, errorx.Has(err2, err1))
	assert.True(t, errorx.Is(err2, err1))

	fmt.Println(err2)
	// fmt.Printf("%v\n", err2)
}

func TestWithf_goerr(t *testing.T) {
	err1 := returnErr("first error message")
	assert.Err(t, err1)

	err2 := errorx.Withf(err1, "second error %s", "message")
	assert.Err(t, err2)
	assert.True(t, errorx.Has(err2, err1))
	assert.True(t, errorx.Is(err2, err1))

	// fmt.Println(err2)
	fmt.Printf("%v\n", err2)
}

func TestWithPrev_goerr(t *testing.T) {
	err1 := returnErr("first error message")
	assert.Err(t, err1)

	err2 := errorx.WithPrev(err1, "second error message")
	assert.Err(t, err2)
	assert.True(t, errorx.Has(err2, err1))
	assert.True(t, errorx.Is(err2, err1))

	fmt.Println(err2)
	// fmt.Printf("%v\n", err2)
}

func TestWithPrev_errorx(t *testing.T) {
	err1 := returnXErr("first error message")
	assert.Err(t, err1)

	err2 := errorx.WithPrev(err1, "second error message")
	assert.Err(t, err2)
	assert.True(t, errorx.Has(err2, err1))
	assert.True(t, errorx.Is(err2, err1))

	// fmt.Println(err2)
	fmt.Printf("%v\n", err2)
}

func TestWithPrev_errorx_l2(t *testing.T) {
	err1 := returnXErrL2("first error message")
	assert.Err(t, err1)

	err2 := errorx.WithPrevf(err1, "second error %s", "message")
	assert.Err(t, err2)
	assert.True(t, errorx.Has(err2, err1))
	assert.True(t, errorx.Is(err2, err1))
	// assert.True(t, errorx.Is(err2, &errorx.ErrorX{}))

	// fmt.Println(err2)
	fmt.Printf("%v\n", err2)

	fmt.Println("--- Use format flag: s")
	fmt.Printf("%s\n", err2)
}

func TestStacked_goerr(t *testing.T) {
	assert.Nil(t, errorx.Stacked(nil))
	assert.Nil(t, errorx.Traced(nil))

	err1 := errorx.E("first error message")
	assert.Err(t, err1)

	err2 := errorx.Stacked(err1)
	assert.Err(t, err2)
	fmt.Printf("%+v\n", err2)
}

func TestStacked_goerr_l2(t *testing.T) {
	err1 := returnErrL2("first error message")
	assert.Err(t, err1)

	err2 := errorx.Traced(err1)
	assert.Err(t, err2)
	fmt.Printf("%v\n", err2)

	fmt.Println("use format flag: s")
	fmt.Printf("%s\n", err2)
}

func TestStacked_errorx(t *testing.T) {
	err1 := returnXErr("first error message")
	assert.Err(t, err1)

	err2 := errorx.WithStack(err1)
	assert.Err(t, err2)
	fmt.Printf("%+v\n", err2)

	assert.Nil(t, errorx.WithStack(nil))
}

func TestTo_ErrorX(t *testing.T) {
	var ex *errorx.ErrorX
	err := returnXErrL2("an error message")
	assert.Err(t, err)

	assert.True(t, errorx.To(err, &ex))
	assert.Contains(t, ex.Location(), "github.com/gookit/goutil/errorx_test.returnXErr(), errorx_test.go")
	assert.Eq(t, "an error message", ex.Message())
	assert.Contains(t, ex.StackString(), "github.com/gookit/goutil/errorx_test.returnXErr()")

	fn := ex.CallerFunc()
	assert.NotNil(t, fn)
	assert.Eq(t, "github.com/gookit/goutil/errorx_test.returnXErr", fn.Name())
	assert.Contains(t, fn.String(), "errorx_test.returnXErr()")
}

func TestWrap(t *testing.T) {
	err := errors.New("first error message")
	assert.Nil(t, errorx.Unwrap(err))
	assert.Nil(t, errorx.Previous(err))
	assert.Eq(t, "first error message", errorx.Cause(err).Error())

	fmt.Println("----------------F------------------")
	fmt.Println(err)

	fmt.Println("----------------S------------------")
	err = errorx.Wrap(err, "second error message")
	assert.Err(t, err)
	_, ok := err.(*errorx.ErrorX)
	assert.True(t, ok)
	fmt.Println(err)

	var ex *errorx.ErrorX
	assert.True(t, errorx.To(err, &ex))
	assert.Nil(t, ex.CallerFunc())
	assert.Eq(t, "unknown", ex.Location())
	assert.Eq(t, "", ex.StackString())
	assert.Eq(t, "second error message", ex.Message())

	ex, ok = errorx.ToErrorX(err)
	assert.True(t, ok)
	assert.Eq(t, "second error message", ex.Message())

	fmt.Println("----------------T------------------")
	err = errorx.Wrap(err, "third error message")
	assert.Err(t, err)
	fmt.Println(err)
	fmt.Println("err.Error():")
	fmt.Println(err.Error())

	assert.Eq(t, "first error message", errorx.Cause(err).Error())
	assert.Contains(t, errorx.Unwrap(err).Error(), "second error message")

	err = errorx.Wrap(nil, "error message")
	assert.Err(t, err)
	_, ok = err.(*errorx.ErrorX)
	assert.False(t, ok)
	_, ok = errorx.ToErrorX(err)
	assert.False(t, ok)
}

func TestCause(t *testing.T) {
	assert.Nil(t, errorx.Cause(nil))
	assert.Nil(t, errorx.Unwrap(nil))
}

func TestWrapf(t *testing.T) {
	err := errorx.Rawf("first error %s", "message")
	assert.Panics(t, func() {
		_ = errorx.MustEX(err)
	})

	err = errorx.Wrapf(err, "second error %s", "message")
	assert.Err(t, err)
	assert.True(t, errorx.IsErrorX(err))

	ex, ok := errorx.ToErrorX(err)
	assert.True(t, ok)
	assert.Eq(t, "second error message", ex.Message())

	fmt.Println(err)
	fmt.Println("err.Error():")
	fmt.Println(err.Error())

	err = errorx.Wrapf(nil, "first error %s", "message")
	assert.Eq(t, "first error message", err.Error())
}

func TestWithOptions(t *testing.T) {
	err := errorx.WithOptions("err msg", errorx.SkipDepth(3), errorx.TraceDepth(10))
	// fmt.Println(err.Error())
	assert.Eq(t, "err msg", err.Error())
	assert.True(t, errorx.IsErrorX(err))

	s := errorx.MustEX(err).StackString()
	assert.StrContains(t, s, "TestWithOptions")
}

func TestErrorX_Format(t *testing.T) {
	err := errorx.New("first error message")
	err = errorx.Wrap(err, "second error message")
	err = errorx.Wrap(err, "third error message")
	assert.Err(t, err)
	fmt.Println(err)

	ex, ok := errorx.ToErrorX(err)
	assert.True(t, ok)

	s := ex.String()
	// fmt.Println(s)
	assert.StrContains(t, s, "third error message")
	assert.StrContains(t, s, "Previous: second error message")
	assert.StrContains(t, s, "Previous: first error message")

	err = ex.Cause()
	assert.Err(t, err)
	assert.Eq(t, "first error message", err.Error())
}

type MyError struct {
	Msg string
}

func (e MyError) Error() string {
	return e.Msg
}

func TestTo(t *testing.T) {
	err := returnMyErr()

	var mye *MyError
	assert.True(t, errorx.To(err, &mye))
	assert.Eq(t, err.Error(), mye.Msg)
	assert.Eq(t, "an error", mye.Msg)

	mye1 := new(MyError)
	// var mye1 *MyError
	assert.True(t, errorx.As(err, &mye1))
	assert.Eq(t, err.Error(), mye1.Msg)
	assert.Eq(t, "an error", mye1.Msg)
}

func returnMyErr() error {
	return &MyError{Msg: "an error"}
}
