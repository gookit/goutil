package goutil_test

import (
	"errors"
	"testing"

	"github.com/gookit/goutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestFuncName(t *testing.T) {
	name := goutil.FuncName(goutil.PkgName)
	assert.Eq(t, "github.com/gookit/goutil.PkgName", name)

	name = goutil.FuncName(goutil.PanicIfErr)
	assert.Eq(t, "github.com/gookit/goutil.PanicIfErr", name)

	err := goutil.Go(func() error {
		return nil
	})
	assert.NoErr(t, err)
}

func TestCallOn(t *testing.T) {
	assert.NoErr(t, goutil.CallOn(false, func() error {
		return errors.New("a error")
	}))
	assert.Err(t, goutil.CallOn(true, func() error {
		return errors.New("a error")
	}))

	err := goutil.CallOrElse(true, func() error {
		return errors.New("a error 001")
	}, func() error {
		return errors.New("a error 002")
	})
	assert.ErrMsg(t, err, "a error 001")

	err = goutil.CallOrElse(false, func() error {
		return errors.New("a error 001")
	}, func() error {
		return errors.New("a error 002")
	})
	assert.ErrMsg(t, err, "a error 002")
}

func TestSafeRun(t *testing.T) {
	t.Run("NoPanic_ReturnsNil", func(t *testing.T) {
		err := goutil.SafeRun(func() {})
		assert.Nil(t, err)
	})

	t.Run("PanicWithError_ReturnsError", func(t *testing.T) {
		expectedErr := errors.New("test error")
		err := goutil.SafeRun(func() {
			panic(expectedErr)
		})
		assert.Equal(t, expectedErr, err)
	})

	t.Run("PanicWithNonError_ReturnsFormattedError", func(t *testing.T) {
		expectedMsg := "test message"
		err := goutil.SafeRun(func() {
			panic(expectedMsg)
		})
		assert.ErrMsg(t, err, expectedMsg)
	})
}

func TestSafeRunWithError(t *testing.T) {
	t.Run("NormalExecution_ReturnsNil", func(t *testing.T) {
		err := goutil.SafeRunWithError(func() error {
			return nil
		})
		assert.NoErr(t, err)
	})

	t.Run("NormalExecution_ReturnsError", func(t *testing.T) {
		expectedErr := errors.New("test error")
		err := goutil.SafeRunWithError(func() error {
			return expectedErr
		})
		assert.Err(t, expectedErr, err)
	})

	t.Run("PanicWithErrorMessage_ReturnsError", func(t *testing.T) {
		expectedErr := errors.New("panic error")
		err := goutil.SafeRunWithError(func() error {
			panic(expectedErr)
		})
		assert.Err(t, expectedErr, err)
	})

	t.Run("PanicWithNonErrorValue_ReturnsError", func(t *testing.T) {
		expectedMsg := "non-error panic"
		err := goutil.SafeRunWithError(func() error {
			panic(expectedMsg)
		})
		assert.ErrMsg(t, err, expectedMsg)
	})
}
