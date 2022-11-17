package reflects_test

import (
	"reflect"
	"testing"

	"github.com/gookit/goutil/reflects"
	"github.com/gookit/goutil/testutil/assert"
)

func TestIsNil(t *testing.T) {
	assert.False(t, reflects.IsNil(reflect.ValueOf(nil)))

	var v *reflects.Value
	assert.True(t, reflects.IsNil(reflect.ValueOf(v)))
}

func TestIsFunc(t *testing.T) {
	assert.False(t, reflects.IsFunc(nil))
	assert.True(t, reflects.IsFunc(reflects.HasChild))
}

func TestHasChild(t *testing.T) {
	is := assert.New(t)

	is.True(reflects.HasChild(reflect.ValueOf([]int{23})))
	is.False(reflects.HasChild(reflect.ValueOf("abc")))
}

func TestIsEqual(t *testing.T) {
	is := assert.New(t)

	is.False(reflects.IsEqual(nil, "abc"))
	is.False(reflects.IsEqual("abc", nil))

	is.False(reflects.IsEqual("abc", 123))
}

func TestIsEmpty(t *testing.T) {
	is := assert.New(t)

	is.True(reflects.IsEmpty(reflect.ValueOf(nil)))
	is.True(reflects.IsEmpty(reflect.ValueOf("")))
	is.True(reflects.IsEmpty(reflect.ValueOf([]string{})))
	is.True(reflects.IsEmpty(reflect.ValueOf(map[int]string{})))
	is.True(reflects.IsEmpty(reflect.ValueOf(0)))
	is.True(reflects.IsEmpty(reflect.ValueOf(uint(0))))
	is.True(reflects.IsEmpty(reflect.ValueOf(float32(0))))

	type T struct{ v any }
	rv := reflect.ValueOf(T{}).Field(0)
	is.True(reflects.IsEmpty(rv))
}

func TestIsEmptyValue(t *testing.T) {
	is := assert.New(t)

	is.True(reflects.IsEmptyValue(reflect.ValueOf(nil)))
	is.True(reflects.IsEmptyValue(reflect.ValueOf("")))
	is.True(reflects.IsEmptyValue(reflect.ValueOf([]string{})))
	is.True(reflects.IsEmptyValue(reflect.ValueOf(map[int]string{})))
	is.True(reflects.IsEmptyValue(reflect.ValueOf(0)))
	is.True(reflects.IsEmptyValue(reflect.ValueOf(uint(0))))
	is.True(reflects.IsEmptyValue(reflect.ValueOf(float32(0))))

	type T struct{ v any }
	rv := reflect.ValueOf(T{}).Field(0)
	is.True(reflects.IsEmptyValue(rv))
}
