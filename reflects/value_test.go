package reflects_test

import (
	"reflect"
	"testing"

	"github.com/gookit/goutil/reflects"
	"github.com/stretchr/testify/assert"
)

func TestValueOf(t *testing.T) {
	rv := reflects.ValueOf(int64(23))

	assert.Equal(t, reflect.Int64, rv.Kind())
	assert.Equal(t, reflects.Int, rv.BaseKind())

	assert.Equal(t, uint64(23), rv.Uint())
	assert.Equal(t, int64(23), rv.Int())

	rv = reflects.ValueOf(uint64(23))
	assert.Equal(t, uint64(23), rv.Uint())
	assert.Equal(t, int64(23), rv.Int())
	assert.False(t, rv.HasChild())

	rv = reflects.ValueOf("abc")
	assert.Panics(t, func() {
		rv.Int()
	})
	assert.Panics(t, func() {
		rv.Uint()
	})
}

func TestValue_Indirect(t *testing.T) {
	type user struct {
		Name string
		Age  int
	}

	rv := reflects.ValueOf(&user{Age: 23})
	assert.Equal(t, reflects.BKind(reflect.Ptr), rv.BKind())
	assert.False(t, rv.HasChild())

	rv1 := reflects.Elem(rv.Value)
	assert.Equal(t, reflect.Struct, rv1.Kind())
	assert.True(t, reflects.HasChild(rv1))

	rv = rv.Indirect()
	assert.True(t, rv.HasChild())
	assert.Equal(t, reflects.BKind(reflect.Struct), rv.BKind())

}
