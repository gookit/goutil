package goutil_test

import (
	"errors"
	"runtime"
	"strings"
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

	err = goutil.IfElseFn(false, func() error {
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
func TestSafeGo(t *testing.T) {
	t.Run("Normal execution", func(t *testing.T) {
		var called bool
		goutil.SafeGo(func() {
			called = true
		}, func(err error) {
			t.Fatal("Expected no error")
		})
		waitForGoroutine()
		if !called {
			t.Fail()
		}
	})

	t.Run("Panic captured", func(t *testing.T) {
		expected := "panic occurred"
		goutil.SafeGo(func() {
			panic(expected)
		}, func(err error) {
			if err == nil || !strings.Contains(err.Error(), expected) {
				t.Fatalf("Expected error containing %q, got %v", expected, err)
			}
		})
		waitForGoroutine()
	})
}

func TestSafeGoWithError(t *testing.T) {
	t.Run("Function returns error", func(t *testing.T) {
		expected := errors.New("test error")
		goutil.SafeGoWithError(func() error {
			return expected
		}, func(err error) {
			if err != expected {
				t.Fail()
			}
		})
		waitForGoroutine()
	})

	t.Run("Panic captured in SafeGoWithError", func(t *testing.T) {
		expected := "panic inside SafeGoWithError"
		goutil.SafeGoWithError(func() error {
			panic(expected)
		}, func(err error) {
			if err == nil || !strings.Contains(err.Error(), expected) {
				t.Fatalf("Expected error containing %q, got %v", expected, err)
			}
		})
		waitForGoroutine()
	})
}

// 等待 goroutine 执行完成
func waitForGoroutine() {
	// 简单等待 goroutine 完成（适用于简单测试）
	for i := 0; i < 10; i++ {
		if active := runtime.NumGoroutine(); active <= 1 {
			break
		}
		runtime.Gosched()
	}
}
