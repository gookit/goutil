package reflects_test

import (
	"reflect"
	"testing"

	"github.com/gookit/goutil/reflects"
	"github.com/stretchr/testify/assert"
)

func TestToBaseKind(t *testing.T) {
	assert.Equal(t, reflects.ToBaseKind(reflect.Int16), reflects.Int)
	assert.Equal(t, reflects.ToBaseKind(reflect.Uint16), reflects.Uint)
	assert.Equal(t, reflects.ToBaseKind(reflect.Float32), reflects.Float)
	assert.Equal(t, reflects.ToBaseKind(reflect.Slice), reflects.Array)
	assert.Equal(t, reflects.ToBaseKind(reflect.Complex64), reflects.Complex)
	assert.Equal(t, reflects.ToBaseKind(reflect.String), reflects.BKind(reflect.String))
}

func TestTypeOf(t *testing.T) {
	rt := reflects.TypeOf(int64(23))

	assert.Equal(t, reflect.Int64, rt.Kind())
	assert.Equal(t, reflects.Int, rt.BaseKind())
}
