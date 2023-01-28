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

	assert.ErrMsg(t, goutil.CallOrElse(true, func() error {
		return errors.New("a error 001")
	}, func() error {
		return errors.New("a error 002")
	}), "a error 001")
	assert.ErrMsg(t, goutil.CallOrElse(false, func() error {
		return errors.New("a error 001")
	}, func() error {
		return errors.New("a error 002")
	}), "a error 002")
}
