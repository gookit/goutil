package reflects_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/gookit/goutil/reflects"
	"github.com/gookit/goutil/testutil/assert"
)

var testFunc1 = func(a, b int) int {
	return a + b
}

var testFunc2 = func(a, b int) (int, error) {
	return 0, errors.New("test error")
}

func TestNewFunc(t *testing.T) {
	fx := reflects.NewFunc(reflect.ValueOf(testFunc1))
	assert.Eq(t, "func(int, int) int", fx.String())
	assert.Eq(t, 2, fx.NumIn())
	assert.Eq(t, 1, fx.NumOut())

	assert.Panics(t, func() {
		reflects.NewFunc(nil)
	})
	assert.Panics(t, func() {
		reflects.NewFunc("invalid")
	})
}

func TestFuncX_Call(t *testing.T) {
	fx := reflects.NewFunc(testFunc1)

	ret, err := fx.Call(1, 2)
	assert.NoErr(t, err)
	assert.Equal(t, 3, ret[0])

	// test return error
	fx = reflects.NewFunc(testFunc2)
	ret, err = fx.Call(1, 2)
	assert.NoErr(t, err)
	assert.Equal(t, 0, ret[0])
	err = ret[1].(error)
	assert.Equal(t, "test error", err.Error())
}

func TestFuncX_Call_err(t *testing.T) {
	fx := reflects.NewFunc(testFunc1)

	// arg number error
	_, err := fx.Call(1, 2, 3)
	assert.ErrMsg(t, err, "wrong number of args: got 3 want 2")

	// arg type error
	_, err = fx.Call(1, "2")
	assert.ErrMsg(t, err, "arg 1: value has type string; should be int")
}

func TestFuncX_Call2(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		fx := reflects.NewFunc(testFunc1)
		val, err := fx.Call2(1, 2)
		assert.NoErr(t, err)
		assert.Equal(t, 3, val)

		fx = reflects.NewFunc(func() error { return nil })
		_, err = fx.Call2()
		assert.NoErr(t, err)
	})

	t.Run("arg type error", func(t *testing.T) {
		fx := reflects.NewFunc(testFunc1)
		_, err := fx.Call2(1, "2")
		assert.Err(t, err)
	})

	t.Run("conv arg type", func(t *testing.T) {
		fx := reflects.NewFunc(testFunc1).WithEnhanceConv()

		val, err := fx.Call2(1, "2")
		assert.NoErr(t, err)
		assert.Equal(t, 3, val)
	})

	t.Run("out number error", func(t *testing.T) {
		fx := reflects.NewFunc(func() {})
		_, err := fx.Call2()
		assert.Err(t, err)
	})
}

func TestCall2(t *testing.T) {
	emptyFn := func() {}
	_, err := reflects.Call2(reflect.ValueOf(emptyFn), nil)
	assert.Err(t, err)
}
