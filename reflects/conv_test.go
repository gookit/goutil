package reflects_test

import (
	"reflect"
	"testing"

	"github.com/gookit/goutil/reflects"
	"github.com/gookit/goutil/testutil/assert"
)

func TestBaseTypeVal(t *testing.T) {
	tests := []struct {
		want, give interface{}
	}{
		{int64(23), 23},
		{int64(23), uint(23)},
		{23.4, 23.4},
		// {23.4, float32(23.4)},
		{"abc", "abc"},
	}
	for _, e := range tests {
		val, err := reflects.BaseTypeVal(reflect.ValueOf(e.give))
		assert.NoErr(t, err)
		assert.Eq(t, e.want, val)
	}

	// val = float64(23.399999618530273)
	val, err := reflects.BaseTypeVal(reflect.ValueOf(float32(23.4)))
	assert.NoErr(t, err)
	assert.NotEmpty(t, val)
}

func TestValueByKind(t *testing.T) {
	tests := []struct {
		want, give interface{}
		// want kind
		kind reflect.Kind
	}{
		{int8(23), 23, reflect.Int8},
		{int16(23), 23, reflect.Int16},
		{int32(23), 23, reflect.Int32},
		{int64(23), 23, reflect.Int64},
		{"23", 23, reflect.String},
		{23, uint(23), reflect.Int},
		{uint(23), 23, reflect.Uint},
		{uint8(23), 23, reflect.Uint8},
		{uint16(23), 23, reflect.Uint16},
		{uint32(23), 23, reflect.Uint32},
		{uint64(23), 23, reflect.Uint64},
		{float32(23), 23, reflect.Float32},
		{float64(23), 23, reflect.Float64},
	}
	for _, e := range tests {
		val, err := reflects.ValueByKind(e.give, e.kind)
		assert.NoErr(t, err)
		assert.Eq(t, e.want, val.Interface())
	}

	val, err := reflects.ValueByKind("abc", reflect.Int)
	assert.Err(t, err)
	assert.False(t, val.IsValid())
}
