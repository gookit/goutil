package assert_test

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
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
		ErrMsgContains(err, "message").
		ErrHasMsg(err, "message").
		Eq("error message", err.Error()).
		Neq("message", err.Error()).
		Equal("error message", err.Error()).
		Contains(err.Error(), "message").
		StrContains(err.Error(), "message").
		StrNotContains(err.Error(), "success").
		StrContainsAll(err.Error(), []string{"error", "message"}).
		StrCount(err.Error(), "e", 3).
		NotContains(err.Error(), "success").
		Gt(4, 3).
		Lt(2, 3)

	assert.True(t, as.IsOk())
	assert.False(t, as.IsFail())

	iv := 23
	ss := []string{"a", "b"}
	mp := map[string]int{"a": 1, "b": 2}

	as = assert.New(t).WithMsg("prefix").
		IsType(1, iv).
		IsKind(reflect.Int, iv).
		NotEq(22, iv).
		NotEqual(22, iv).
		Len(ss, 2).
		LenGt(ss, 1).
		Lte(iv, 23).
		Gte(iv, 23).
		EqInt(23, iv).
		InDelta(1.0, 1.001, 0.01).
		EqFloat(1.0, 1.001, 0.01).
		Empty(0).
		Zero(0).
		NotZero(iv).
		True(true).
		False(false).
		NoErr(nil).
		NoError(nil).
		Nil(nil).
		ContainsKey(mp, "a").
		NotContainsKey(mp, "c").
		ContainsKeys(mp, []string{"a", "b"}).
		NotContainsKeys(mp, []string{"c", "d"}).
		ContainsElems([]int{1, 2, 3}, []int{1, 2})

	assert.True(t, as.IsOk())

	// Test Same/NotSame
	p1 := &iv
	p3 := new(int)
	*p3 = iv
	as = assert.New(t).
		Same(p1, p1).
		NotSame(p1, p3)

	assert.True(t, as.IsOk())

	// Panics
	assert.New(t).
		Panics(func() { panic("panic") }).
		NotPanics(func() {}).
		PanicsMsg(func() { panic("panic") }, "panic").
		PanicsErrMsg(func() { panic(errors.New("panic")) }, "panic")
}

func TestAssertions_Filesystem(t *testing.T) {
	// Create temp file and dir for testing
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")

	// Write temp file
	err := os.WriteFile(tmpFile, []byte("test"), 0644)
	assert.NoErr(t, err)

	as := assert.New(t).
		FileExists(tmpFile).
		FileNotExists(filepath.Join(tmpDir, "notexist.txt")).
		DirExists(tmpDir).
		DirNotExists(filepath.Join(tmpDir, "notexist"))

	assert.True(t, as.IsOk())
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
