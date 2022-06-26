package stdutil_test

import (
	"reflect"
	"testing"

	"github.com/gookit/goutil/stdutil"
	"github.com/stretchr/testify/assert"
)

func TestValueIsEmpty(t *testing.T) {
	is := assert.New(t)

	is.True(stdutil.ValueIsEmpty(reflect.ValueOf(nil)))
	is.True(stdutil.ValueIsEmpty(reflect.ValueOf("")))

	type T struct{ v interface{} }
	rv := reflect.ValueOf(T{}).Field(0)
	is.True(stdutil.ValueIsEmpty(rv))
}

func TestValueLen(t *testing.T) {
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
		is.Equal(3, stdutil.ValueLen(reflect.ValueOf(sample)))
	}

	ptrArr := &[]string{"a", "b"}
	is.Equal(2, stdutil.ValueLen(reflect.ValueOf(ptrArr)))

	is.Equal(-1, stdutil.ValueLen(reflect.ValueOf(nil)))
}
