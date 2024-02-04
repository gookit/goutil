package basefn_test

import (
	"errors"
	"testing"

	"github.com/gookit/goutil/basefn"
	"github.com/gookit/goutil/testutil/assert"
)

func TestPanicIf(t *testing.T) {
	basefn.PanicIf(false, "")
	assert.Panics(t, func() {
		basefn.PanicIf(true, "a error msg")
	})

	assert.Panics(t, func() {
		basefn.PanicIf(true)
	})
	assert.Panics(t, func() {
		basefn.PanicIf(true, 234)
	})
	assert.Panics(t, func() {
		basefn.PanicIf(true, 234, "abc")
	})

	assert.Panics(t, func() {
		basefn.PanicIf(true, "a error %s", "msg")
	})
}

func TestPanicErr(t *testing.T) {
	basefn.MustOK(nil)
	basefn.PanicErr(nil)
	assert.Panics(t, func() {
		basefn.PanicErr(errors.New("a error"))
	})

	// must ignore
	assert.NotPanics(t, func() {
		basefn.MustIgnore(nil, nil)
	})
	assert.Panics(t, func() {
		basefn.MustIgnore(nil, errors.New("a error"))
	})
}

func TestPanicf(t *testing.T) {
	basefn.MustOK(nil)
	assert.Panics(t, func() {
		basefn.Panicf("hi %s", "inhere")
	})

	assert.Eq(t, "hi", basefn.Must("hi", nil))
	assert.Panics(t, func() {
		basefn.Must("hi", errors.New("a error"))
	})
	assert.Panics(t, func() {
		basefn.MustOK(errors.New("a error"))
	})
}

func TestErrOnFail(t *testing.T) {
	err := errors.New("a error")
	assert.Err(t, basefn.ErrOnFail(false, err))
	assert.NoErr(t, basefn.ErrOnFail(true, err))
}

func TestOrValue(t *testing.T) {
	assert.Eq(t, "ab", basefn.OrValue(true, "ab", "dc"))
	assert.Eq(t, "dc", basefn.OrValue(false, "ab", "dc"))
	assert.Eq(t, 1, basefn.FirstOr([]int{1, 2}, 3))
	assert.Eq(t, 3, basefn.FirstOr(nil, 3))
}

func TestOrReturn(t *testing.T) {
	assert.Eq(t, "ab", basefn.OrReturn(true, func() string {
		return "ab"
	}, func() string {
		return "dc"
	}))
	assert.Eq(t, "dc", basefn.OrReturn(false, func() string {
		return "ab"
	}, func() string {
		return "dc"
	}))
}

func TestCallOn(t *testing.T) {
	assert.NoErr(t, basefn.CallOn(false, func() error {
		return errors.New("a error")
	}))
	assert.Err(t, basefn.CallOn(true, func() error {
		return errors.New("a error")
	}))

	assert.ErrMsg(t, basefn.CallOrElse(true, func() error {
		return errors.New("a error 001")
	}, func() error {
		return errors.New("a error 002")
	}), "a error 001")
	assert.ErrMsg(t, basefn.CallOrElse(false, func() error {
		return errors.New("a error 001")
	}, func() error {
		return errors.New("a error 002")
	}), "a error 002")
}
