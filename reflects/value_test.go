package reflects_test

import (
	"reflect"
	"testing"

	"github.com/gookit/goutil/reflects"
	"github.com/gookit/goutil/testutil/assert"
)

func TestValueOf(t *testing.T) {
	rv := reflects.ValueOf(int64(23))

	assert.Eq(t, reflect.Int64, rv.Kind())
	assert.Eq(t, reflects.Int, rv.BaseKind())

	assert.Eq(t, uint64(23), rv.Uint())
	assert.Eq(t, int64(23), rv.Int())

	rv = reflects.ValueOf(uint64(23))
	assert.Eq(t, uint64(23), rv.Uint())
	assert.Eq(t, int64(23), rv.Int())
	assert.False(t, rv.HasChild())
	assert.Eq(t, reflects.Uint, rv.Type().BaseKind())
	assert.Eq(t, reflects.Uint, rv.Indirect().BaseKind())

	rv = reflects.ValueOf(reflect.ValueOf("abc"))
	assert.Eq(t, "abc", rv.String())
	assert.Eq(t, "abc", rv.Elem().String())
	assert.Eq(t, reflect.String, rv.Type().BaseKind())

	// pointer
	s := new(string)
	*s = "abc"
	rv = reflects.ValueOf(s)
	assert.Eq(t, reflect.String, rv.Elem().Kind())

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
	assert.Eq(t, reflect.Ptr, rv.BKind())
	assert.False(t, rv.HasChild())

	rv1 := reflects.Elem(rv.Value)
	assert.Eq(t, reflect.Struct, rv1.Kind())
	assert.True(t, reflects.HasChild(rv1))

	rv1 = reflects.Indirect(rv.Value)
	assert.Eq(t, reflect.Struct, rv1.Kind())
	assert.True(t, reflects.HasChild(rv1))

	rv = rv.Indirect()
	assert.True(t, rv.HasChild())
	assert.Eq(t, reflect.Struct, rv.BKind())

	rv1 = reflects.Elem(reflect.ValueOf("abc"))
	assert.Eq(t, reflect.String, rv1.Kind())

	rv1 = reflects.Indirect(reflect.ValueOf("abc"))
	assert.Eq(t, reflect.String, rv1.Kind())
}
