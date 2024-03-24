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

func TestPkgName(t *testing.T) {
	name := goutil.PkgName(goutil.FuncName(goutil.PanicIfErr))
	assert.Eq(t, "github.com/gookit/goutil", name)
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
