package reflects_test

import (
	"reflect"
	"testing"

	"github.com/gookit/goutil/reflects"
	"github.com/gookit/goutil/testutil/assert"
)

func TestToBaseKind(t *testing.T) {
	assert.Eq(t, reflects.ToBaseKind(reflect.Int16), reflects.Int)
	assert.Eq(t, reflects.ToBaseKind(reflect.Uint16), reflects.Uint)
	assert.Eq(t, reflects.ToBaseKind(reflect.Float32), reflects.Float)
	assert.Eq(t, reflects.ToBaseKind(reflect.Slice), reflects.Array)
	assert.Eq(t, reflects.ToBaseKind(reflect.Complex64), reflects.Complex)
	assert.Eq(t, reflects.ToBaseKind(reflect.String), reflects.BKind(reflect.String))
}

func TestTypeOf(t *testing.T) {
	rt := reflects.TypeOf(int64(23))

	assert.Eq(t, reflect.Int64, rt.Kind())
	assert.Eq(t, reflects.Int, rt.BaseKind())

	assert.Eq(t, reflect.Int64, rt.RealType().Kind())
	assert.Eq(t, reflect.Int64, rt.SafeElem().Kind())

	s := new(string)
	*s = "abc"
	rt = reflects.TypeOf(s)
	assert.Eq(t, reflect.Pointer, rt.Kind())
	assert.Eq(t, reflect.String, rt.RealType().Kind())
	assert.Eq(t, reflect.Pointer, rt.SafeElem().Kind())

	ss := []string{"abc"}
	rt = reflects.TypeOf(ss)
	assert.Eq(t, reflect.Slice, rt.Kind())
	assert.Eq(t, reflect.String, rt.SafeElem().Kind())
}
