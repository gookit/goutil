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

func TestLen(t *testing.T) {
	is := assert.New(t)
	tests := []interface{}{
		"abc",
		123,
		int8(123), int16(123), int32(123), int64(123),
		uint8(123), uint16(123), uint32(123), uint64(123),
		float32(123), float64(123),
		[]int{1, 2, 3}, []string{"a", "b", "c"},
		map[string]string{"k0": "v0", "k1": "v1", "k2": "v2"},
	}

	for _, sample := range tests {
		is.Eq(3, reflects.Len(reflect.ValueOf(sample)))
	}

	ptrArr := &[]string{"a", "b"}
	is.Eq(2, reflects.Len(reflect.ValueOf(ptrArr)))

	is.Eq(-1, reflects.Len(reflect.ValueOf(nil)))
}
