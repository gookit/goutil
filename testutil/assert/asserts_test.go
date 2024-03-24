package assert_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
)

type tCustomTesting struct {
	*testing.T
	// logs []string
	errs []string
}

// FailNow marks the function as having failed and stops its execution.
func (tc *tCustomTesting) FailNow() {}

// Error logs args to error
func (tc *tCustomTesting) Error(args ...interface{}) {
	tc.errs = append(tc.errs, fmt.Sprint(args...))
}

// First get the first error message
func (tc *tCustomTesting) First() string {
	if len(tc.errs) > 0 {
		return tc.errs[0]
	}
	return ""
}

// Reset error messages
func (tc *tCustomTesting) Reset() {
	tc.errs = []string{}
}

// ResetGet reset error messages and return first message
func (tc *tCustomTesting) ResetGet() string {
	if len(tc.errs) > 0 {
		first := tc.errs[0]
		tc.errs = []string{}
		return first
	}
	return ""
}

func TestCommon_success(t *testing.T) {
	assert.Nil(t, nil)
	assert.False(t, false)
	assert.True(t, true)

	assert.Empty(t, "")
	assert.Empty(t, nil)
	assert.NotEmpty(t, "abc")

	// eq
	assert.Eq(t, 1, 1)
	assert.Eq(t, nil, nil)
	assert.Equal(t, 1, 1)

	// neq
	assert.Neq(t, 1, 2)
	assert.NotEq(t, 1, 2)
	assert.NotEqual(t, 1, 2)

	// kind
	assert.IsType(t, 1, 1)
	assert.IsKind(t, reflect.Int, 1)

	// same
	val := new(int)
	*val = 1

	assert.Same(t, val, val)
	assert.NotSame(t, 2, 2)
	assert.NotSame(t, 1, 2)

	// panics
	assert.Panics(t, func() {
		panic("hh")
	})

	assert.NotPanics(t, func() {})

	assert.PanicsMsg(t, func() {
		panic("hh")
	}, "hh")

	assert.PanicsErrMsg(t, func() {
		panic(errors.New("hh"))
	}, "hh")
}

func TestCommon_fail(t *testing.T) {
	assert.DisableColor()
	defer func() {
		assert.EnableColor = true
	}()

	tc := &tCustomTesting{T: t}

	// nil
	assert.Nil(tc, 1)
	str := tc.First()
	assert.StrContains(t, str, "TestCommon_fail")
	assert.StrContains(t, str, "goutil/testutil/assert/asserts_test.go:")
	assert.StrContains(t, str, "Expected nil, but got:")
	assert.StrNotContains(t, str, "NOT EXIST")
	tc.Reset()

	assert.NotNil(tc, nil)
	assert.StrContains(t, tc.ResetGet(), "Should not nil value")

	// false
	assert.False(tc, true)
	assert.StrContains(t, tc.ResetGet(), "Result should be False")

	// true
	assert.True(tc, false)
	assert.StrContains(t, tc.ResetGet(), "Result should be True")
	assert.True(tc, false, "user custom message")
	assert.StrContains(t, tc.ResetGet(), "user custom message")

	// empty
	assert.Empty(tc, "abc")
	assert.StrContains(t, tc.ResetGet(), "Should be empty, but was:")

	assert.NotEmpty(tc, "")
	assert.StrContains(t, tc.ResetGet(), "Should not be empty, but was:")

	// error
	assert.Error(tc, nil)
	assert.StrContains(t, tc.ResetGet(), "An error is expected but got nil.")

	err := errors.New("want error msg")
	err1 := errors.New("another error msg")
	assert.ErrIs(tc, nil, err)
	assert.StrContains(t, tc.ResetGet(), "An error is expected but got nil.")

	assert.ErrIs(tc, err, err1)
	assert.StrContains(t, tc.ResetGet(), "Expect given err is equals")

	assert.ErrMsg(tc, nil, "want error msg")
	assert.StrContains(t, tc.ResetGet(), "An error is expected but got nil.")

	assert.ErrMsg(tc, err1, "want error msg")
	assert.StrContains(t, tc.ResetGet(), "Error message not equal:")

	assert.ErrSubMsg(tc, nil, "want error msg")
	assert.StrContains(t, tc.ResetGet(), "An error is expected but got nil.")

	assert.ErrSubMsg(tc, err1, "want error msg")
	assert.StrContains(t, tc.ResetGet(), "Error message check fail:")

	// eq
	assert.Eq(tc, 1, 2)
	assert.StrContains(t, tc.ResetGet(), "Not equal:")

	assert.Eq(tc, tc.ResetGet, 2)
	assert.StrContains(t, tc.ResetGet(), "cannot take func type as argument")

	assert.Equal(tc, 1, 2)
	assert.StrContains(t, tc.ResetGet(), "Not equal:")

	// neq
	assert.Neq(tc, 1, 1)
	assert.StrContains(t, tc.ResetGet(), "Given should not be: 1")

	assert.NotEq(tc, 1, 1)
	assert.StrContains(t, tc.ResetGet(), "Given should not be: 1")

	assert.NotEq(tc, tc.ResetGet, 2)
	assert.StrContains(t, tc.ResetGet(), "cannot take func type as argument")

	assert.NotEqual(tc, 1, 1)
	assert.StrContains(t, tc.ResetGet(), "Given should not be: 1")

	// kind
	assert.IsType(tc, 1, "1")
	assert.StrContains(t, tc.ResetGet(), "Expected to be of type int, but was string")

	assert.IsKind(tc, reflect.Int, "1")
	assert.StrContains(t, tc.ResetGet(), "Expected to be of kind int, but was string")

	// same
	val := new(int)
	*val = 1

	assert.Same(tc, val, 2)
	assert.StrContains(t, tc.ResetGet(), "Not same: ")

	assert.NotSame(tc, val, val)
	assert.StrContains(t, tc.ResetGet(), "Expect and actual is same object:")

	// compare
	assert.Lt(tc, 2, 1)
	assert.StrContains(t, tc.ResetGet(), "Given 2 should less than 1")

	assert.Lte(tc, 2, 1)
	assert.StrContains(t, tc.ResetGet(), "Given 2 should less than or equal 1")

	assert.Gt(tc, 1, 2)
	assert.StrContains(t, tc.ResetGet(), "Given 1 should greater than 2")

	assert.Gte(tc, 1, 2)
	assert.StrContains(t, tc.ResetGet(), "Given 1 should greater than or equal 2")

	// contains
	assert.Contains(tc, "abc", "d")
	assert.StrContains(t, tc.ResetGet(), "Should contain:")
	assert.Contains(tc, nil, "d")
	assert.StrContains(t, tc.ResetGet(), "could not be applied builtin len()")

	assert.NotContains(tc, "abc", "a")
	assert.StrContains(t, tc.ResetGet(), "Should not contain:")

	assert.StrContains(tc, "abc", "d")
	assert.StrContains(t, tc.ResetGet(), "String check fail:")

	// contains key
	assert.ContainsKey(tc, map[string]int{"a": 1}, "b")
	assert.StrContains(t, tc.ResetGet(), "Map should contains the key:")

	assert.NotContainsKey(tc, map[string]int{"a": 1}, "a")
	assert.StrContains(t, tc.ResetGet(), "Map should not contains the key:")

	// contains keys
	assert.ContainsKeys(tc, map[string]int{"a": 1}, []string{"a", "b"})
	assert.StrContains(t, tc.ResetGet(), "Map should contains the key:")

	assert.ContainsKeys(tc, map[string]int{"a": 1}, "invalid-type")
	assert.StrContains(t, tc.ResetGet(), "input param type is invalid")

	assert.NotContainsKeys(tc, map[string]int{"a": 1}, []string{"a"})
	assert.StrContains(t, tc.ResetGet(), "Map should not contains the key:")

	assert.NotContainsKeys(tc, map[string]int{"a": 1}, "invalid-type")
	assert.StrContains(t, tc.ResetGet(), "input param type is invalid")

	// len
	assert.Len(tc, "abc", 4)
	assert.StrContains(t, tc.ResetGet(), "should have 4 item(s), but has 3")

	assert.Len(tc, tc.ResetGet, 4)
	assert.StrContains(t, tc.ResetGet(), "type 'func() string' could not be calc length")

	assert.LenGt(tc, "abc", 3)
	assert.StrContains(t, tc.ResetGet(), "should have more than 3 item(s), but has 3")

	assert.LenGt(tc, tc.ResetGet, 4)
	assert.StrContains(t, tc.ResetGet(), "type 'func() string' could not be calc length")

	// panics
	assert.Panics(tc, func() {})
	assert.StrContains(t, tc.ResetGet(), "should panic")

	assert.PanicsMsg(tc, func() {}, "custom message")
	assert.StrContains(t, tc.ResetGet(), "should panic")

	assert.PanicsMsg(tc, func() {
		panic("user custom message")
	}, "custom message")
	assert.StrContains(t, tc.ResetGet(), "custom message")

	assert.PanicsErrMsg(tc, func() {}, "custom message")
	assert.StrContains(t, tc.ResetGet(), "should panic")

	assert.PanicsErrMsg(tc, func() {
		panic("user custom message")
	}, "user custom message")
	assert.StrContains(t, tc.ResetGet(), "should panic and is error type")

	assert.PanicsErrMsg(tc, func() {
		panic(errors.New("user custom message"))
	}, "not custom message")
	assert.StrContains(t, tc.ResetGet(), "not custom message")

	assert.NotPanics(tc, func() {
		panic("user custom message")
	})
	assert.StrContains(t, tc.ResetGet(), "should not panic")

	// fail
	assert.Fail(tc, "custom message1")
	assert.StrContains(t, tc.ResetGet(), "custom message1")

	assert.FailNow(tc, "custom message2")
	assert.StrContains(t, tc.ResetGet(), "custom message2")
}

func TestErr(t *testing.T) {
	assert.NoErr(t, nil)
	assert.NoError(t, nil)

	err := errors.New("this is a error")
	// assert2.EqualError(t, err, "user custom message")
	assert.Err(t, err, "user custom message")
	assert.Error(t, err)
	assert.ErrMsg(t, err, "this is a error")
}

func TestContains(t *testing.T) {
	str := "abc+123"
	assert.StrContains(t, str, "123")
	assert.Contains(t, str, "123")
	assert.NotContains(t, str, "456")
	assert.StrCount(t, str, "123", 1)

	mp := map[string]any{
		"age":  456,
		"name": "inhere",
	}
	assert.ContainsKey(t, mp, "name")
	assert.ContainsKeys(t, mp, []string{"name", "age"})
	assert.NotContainsKey(t, mp, "addr")
	assert.NotContainsKeys(t, mp, []string{"addr"})

	assert.ContainsElems(t, []string{"def"}, []string{"def"})
	assert.ContainsElems(t, []string{"def", "abc"}, []string{"def"})
}
