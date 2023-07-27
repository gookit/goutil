package assert_test

import (
	"errors"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
)

func TestAssertions_Chain(t *testing.T) {
	err := errors.New("error message")

	as := assert.New(t).
		NotEmpty(err).
		NotNil(err).
		Err(err).
		Error(err).
		ErrIs(err, err).
		ErrMsg(err, "error message").
		ErrSubMsg(err, "message").
		Eq("error message", err.Error()).
		Neq("message", err.Error()).
		Equal("error message", err.Error()).
		Contains(err.Error(), "message").
		StrContains(err.Error(), "message").
		NotContains(err.Error(), "success").
		Gt(4, 3).
		Lt(2, 3)

	assert.True(t, as.IsOk())
	assert.False(t, as.IsFail())

	iv := 23
	ss := []string{"a", "b"}

	as = assert.New(t).WithMsg("prefix").
		IsType(1, iv).
		NotEq(22, iv).
		NotEqual(22, iv).
		Len(ss, 2).
		LenGt(ss, 1).
		Lte(iv, 23).
		Gte(iv, 23).
		Empty(0).
		True(true).
		False(false).
		NoErr(nil).
		NoError(nil).
		Nil(nil).
		ContainsKey(map[string]int{"a": 1}, "a")

	assert.True(t, as.IsOk())

	// Panics
	assert.New(t).
		Panics(func() { panic("panic") }).
		NotPanics(func() {}).
		PanicsMsg(func() { panic("panic") }, "panic").
		PanicsErrMsg(func() { panic(errors.New("panic")) }, "panic")
}

func TestAssertions_chain_fail(t *testing.T) {
	assert.HideFullPath()
	defer func() {
		assert.ShowFullPath = true
	}()

	tc := &tCustomTesting{T: t}
	ts := assert.New(tc)

	ts.Fail("fail message")
	str := tc.ResetGet()
	assert.StrContains(t, str, "fail message")
	assert.StrContains(t, str, "assertions_test.go")
	assert.NotContains(t, str, "testutil/assert/assertions_test.go")

	ts.FailNow("fail now message")
	assert.StrContains(t, tc.ResetGet(), "fail now message")
}
