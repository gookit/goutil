package reflects_test

import (
	"reflect"
	"testing"

	"github.com/gookit/goutil/reflects"
	"github.com/gookit/goutil/testutil/assert"
)

func TestIsFunc(t *testing.T) {
	assert.True(t, reflects.IsFunc(reflects.HasChild))
}

func TestIsEmpty(t *testing.T) {
	is := assert.New(t)

	is.True(reflects.IsEmpty(reflect.ValueOf(nil)))
	is.True(reflects.IsEmpty(reflect.ValueOf("")))

	type T struct{ v interface{} }
	rv := reflect.ValueOf(T{}).Field(0)
	is.True(reflects.IsEmpty(rv))
}
