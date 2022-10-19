package errorx_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/testutil/assert"
)

func TestErrorR_usage(t *testing.T) {
	err := errorx.NewR(405, "param error")

	assert.Eq(t, 405, err.Code())
	assert.Eq(t, "param error", err.Error())
	assert.Eq(t, "param error(code: 405)", err.String())
	assert.False(t, err.IsSuc())
	assert.True(t, err.IsFail())

	fmt.Println(err)
	fmt.Printf("%v\n", err)
	fmt.Printf("%+v\n", err)
	fmt.Printf("%#v\n", err)

	err = errorx.Suc("ok")
	assert.Eq(t, 0, err.Code())
	assert.True(t, err.IsSuc())
	assert.False(t, err.IsFail())
}
