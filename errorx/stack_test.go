package errorx

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrStackOpt(t *testing.T) {
	defer ResetStdOpt()

	assert.Equal(t, 3, stdOpt.SkipDepth)
	assert.Equal(t, 8, stdOpt.TraceDepth)

	Config(SkipDepth(5), TraceDepth(12))
	assert.Equal(t, 5, stdOpt.SkipDepth)
	assert.Equal(t, 12, stdOpt.TraceDepth)

}

func TestFuncForPC(t *testing.T) {
	fn := FuncForPC(uintptr(0))
	assert.Nil(t, fn)

	fn = FuncForPC(reflect.ValueOf(Config).Pointer())
	assert.Contains(t, fn.Location(), "gookit/goutil/errorx.Config()")
	assert.Contains(t, fn.String(), "goutil/errorx/stack.go")

	bs, err := fn.MarshalText()
	assert.NoError(t, err)
	str := string(bs)
	assert.Contains(t, str, "gookit/goutil/errorx.Config()")
	assert.Contains(t, str, "goutil/errorx/stack.go")
}

func TestStack_Format(t *testing.T) {
	st := new(stack)
	assert.Equal(t, 0, st.StackLen())

	buf := new(bytes.Buffer)
	_, err := st.WriteTo(buf)
	assert.NoError(t, err)
	assert.Equal(t, "", buf.String())
	assert.Equal(t, uintptr(0), st.CallerPC())

	st = callersStack(1, 5)
	assert.True(t, st.StackLen() > 0)
	assert.NotEmpty(t, st.StackFrames())
}
