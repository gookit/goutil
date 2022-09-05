package stdutil_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/gookit/goutil/comdef"
	"github.com/gookit/goutil/stdutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestMustString(t *testing.T) {
	assert.Eq(t, "23", stdutil.MustString(23))

	assert.PanicsErrMsg(t, func() {
		stdutil.MustString([]string{"a", "b"})
	}, comdef.ErrConvType.Error())
}

func TestToString(t *testing.T) {
	assert.Eq(t, "23", stdutil.ToString(23))
	assert.Eq(t, "[a b]", stdutil.ToString([]string{"a", "b"}))
}

func TestTryString(t *testing.T) {
	s, err := stdutil.TryString(23)
	assert.NoErr(t, err)
	assert.Eq(t, "23", s)

	s, err = stdutil.TryString([]string{"a", "b"})
	assert.Eq(t, err, comdef.ErrConvType)
	assert.Eq(t, "", s)
}

func TestBaseTypeVal(t *testing.T) {
	tests := []interface{}{
		2,
		int8(2), int16(2), int32(2), int64(2),
		uint(2), uint8(2), uint16(2), uint32(2), uint64(2),
		time.Duration(2),
	}
	for _, el := range tests {
		val, err := stdutil.BaseTypeVal(el)
		assert.NoErr(t, err)
		assert.Eq(t, int64(2), val)
	}

	tests3 := []struct{ in, out interface{} }{
		{"adc", "adc"},
		{"2", "2"},
		{json.Number("2"), "2"},
	}
	for _, el := range tests3 {
		val, err := stdutil.BaseTypeVal(el.in)
		assert.NoErr(t, err)
		assert.Eq(t, el.out, val)
	}

	val, err := stdutil.BaseTypeVal(float32(2))
	assert.NoErr(t, err)
	assert.Eq(t, float64(2), val)

	_, err = stdutil.BaseTypeVal([]int{2})
	assert.Err(t, err)
}
