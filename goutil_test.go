package goutil_test

import (
	"errors"
	"testing"

	"github.com/gookit/goutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestPkgName(t *testing.T) {
	name := goutil.PkgName(goutil.FuncName(goutil.PanicIfErr))
	assert.Eq(t, "github.com/gookit/goutil", name)
}

func TestPanicIfErr(t *testing.T) {
	goutil.PanicIfErr(nil)
	goutil.PanicErr(nil)
	goutil.MustOK(nil)
}

func TestPanicf(t *testing.T) {
	assert.Panics(t, func() {
		goutil.Panicf("hi %s", "inhere")
	})
}

func TestErrOnFail(t *testing.T) {
	err := errors.New("a error")
	assert.Err(t, goutil.ErrOnFail(true, err))
	assert.NoErr(t, goutil.ErrOnFail(false, err))
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
