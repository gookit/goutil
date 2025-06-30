package goutil_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gookit/goutil"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

var testSrvAddr string

func TestMain(m *testing.M) {
	s := testutil.NewEchoServer()
	defer s.Close()
	testSrvAddr = s.HTTPHost()
	fmt.Println("Test server listen on:", testSrvAddr)

	m.Run()
}

func TestPanicIfErr(t *testing.T) {
	goutil.PanicIf(false, "")
	assert.Panics(t, func() {
		goutil.PanicIf(true, "a error msg")
	})

	goutil.PanicIfErr(nil)
	assert.Panics(t, func() {
		goutil.PanicIfErr(errors.New("a error"))
	})

	goutil.PanicErr(nil)
	assert.Panics(t, func() {
		goutil.PanicErr(errors.New("a error"))
	})

	goutil.MustOK(nil)
	assert.Panics(t, func() {
		goutil.MustOK(errors.New("a error"))
	})

	assert.Eq(t, "hi", goutil.Must("hi", nil))
	assert.Panics(t, func() {
		goutil.Must("hi", errors.New("a error"))
	})

	assert.NotPanics(t, func() {
		goutil.MustIgnore(nil, nil)
	})
	assert.Panics(t, func() {
		goutil.MustIgnore(nil, errors.New("a error"))
	})
}

func TestPanicf(t *testing.T) {
	assert.Panics(t, func() {
		goutil.Panicf("hi %s", "inhere")
	})
}

func TestErrOnFail(t *testing.T) {
	err := errors.New("a error")
	assert.Err(t, goutil.ErrOnFail(false, err))
	assert.NoErr(t, goutil.ErrOnFail(true, err))
}

func TestOrValue(t *testing.T) {
	assert.Eq(t, "ab", goutil.OrValue(true, "ab", "dc"))
	assert.Eq(t, "dc", goutil.OrValue(false, "ab", "dc"))
}

func TestOrReturn(t *testing.T) {
	assert.Eq(t, "ab", goutil.OrReturn(true, func() string {
		return "ab"
	}, func() string {
		return "dc"
	}))
	assert.Eq(t, "dc", goutil.OrReturn(false, func() string {
		return "ab"
	}, func() string {
		return "dc"
	}))
}

func TestIsEmpty(t *testing.T) {
	is := assert.New(t)

	is.True(goutil.IsEmpty(nil))
	is.False(goutil.IsZero("abc"))
	is.False(goutil.IsEmpty("abc"))

	is.True(goutil.IsEmptyReal(nil))
	is.False(goutil.IsZeroReal("abc"))
	is.False(goutil.IsEmptyReal("abc"))
}

func TestIsFunc(t *testing.T) {
	is := assert.New(t)

	is.False(goutil.IsFunc(nil))
	is.True(goutil.IsFunc(goutil.IsEqual))
}

func TestIsEqual(t *testing.T) {
	is := assert.New(t)

	is.True(goutil.IsNil(nil))
	is.False(goutil.IsNil("abc"))

	is.True(goutil.IsEqual("a", "a"))
	is.True(goutil.IsEqual([]string{"a"}, []string{"a"}))
	is.True(goutil.IsEqual(23, 23))
	is.True(goutil.IsEqual(nil, nil))
	is.True(goutil.IsEqual([]byte("abc"), []byte("abc")))

	is.False(goutil.IsEqual([]byte("abc"), "abc"))
	is.False(goutil.IsEqual(nil, 23))
	is.False(goutil.IsEqual(goutil.IsEmpty, 23))
}

func TestIsContains(t *testing.T) {
	is := assert.New(t)

	is.True(goutil.Contains("abc", "a"))
	is.True(goutil.Contains([]string{"abc", "def"}, "abc"))
	is.True(goutil.Contains(map[int]string{2: "abc", 4: "def"}, 4))

	is.False(goutil.Contains("abc", "def"))
	is.False(goutil.IsContains("abc", "def"))
}

func TestPkgName(t *testing.T) {
	name := goutil.PkgName(goutil.FuncName(goutil.PanicIfErr))
	assert.Eq(t, "github.com/gookit/goutil", name)
}
