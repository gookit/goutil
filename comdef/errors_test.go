package comdef_test

import (
	"errors"
	"testing"

	"github.com/gookit/goutil/comdef"
	"github.com/gookit/goutil/testutil/assert"
)

func TestErrConvType(t *testing.T) {
	expected := "convert value type error"
	assert.Eq(t, expected, comdef.ErrConvType.Error())
}

func TestErrors_Error(t *testing.T) {
	errs := comdef.Errors{
		errors.New("error 1"),
		errors.New("error 2"),
	}

	expected := "error 1\nerror 2\n"
	assert.Eq(t, expected, errs.Error())
}

func TestErrors_ErrOrNil(t *testing.T) {
	var empty comdef.Errors
	assert.Nil(t, empty.ErrOrNil())

	errs := comdef.Errors{errors.New("error 1")}
	assert.NotNil(t, errs.ErrOrNil())
}

func TestErrors_First(t *testing.T) {
	var empty comdef.Errors
	assert.Nil(t, empty.First())

	errs := comdef.Errors{errors.New("error 1"), errors.New("error 2")}
	assert.Eq(t, "error 1", errs.First().Error())
}
