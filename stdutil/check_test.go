package stdutil_test

import (
	"reflect"
	"testing"

	"github.com/gookit/goutil/stdutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestIsEqual(t *testing.T) {
	is := assert.New(t)

	is.True(stdutil.IsEqual("a", "a"))
	is.True(stdutil.IsEqual([]string{"a"}, []string{"a"}))
	is.True(stdutil.IsEqual(23, 23))
	is.True(stdutil.IsEqual(nil, nil))
	is.True(stdutil.IsEqual([]byte("abc"), []byte("abc")))

	is.False(stdutil.IsEqual([]byte("abc"), "abc"))
	is.False(stdutil.IsEqual(nil, 23))
	is.False(stdutil.IsEqual(stdutil.IsEmpty, 23))
}

func TestContains(t *testing.T) {
	is := assert.New(t)

	is.True(stdutil.Contains("abc", "a"))
	is.True(stdutil.Contains([]string{"abc", "def"}, "abc"))
	is.True(stdutil.Contains(map[int]string{2: "abc", 4: "def"}, 4))

	is.False(stdutil.Contains("abc", "def"))
}

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
		is.Eq(3, stdutil.ValueLen(reflect.ValueOf(sample)))
	}

	ptrArr := &[]string{"a", "b"}
	is.Eq(2, stdutil.ValueLen(reflect.ValueOf(ptrArr)))

	is.Eq(-1, stdutil.ValueLen(reflect.ValueOf(nil)))
}
