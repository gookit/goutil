package mathutil_test

import (
	"testing"

	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/x/assert"
)

func TestNilNumericConversionPolicy(t *testing.T) {
	for _, input := range []any{nil, (*int)(nil), (*int64)(nil), (*uint)(nil), (*uint64)(nil), (*float64)(nil)} {
		t.Run("default", func(t *testing.T) {
			_, err := mathutil.ToIntWith(input, mathutil.WithStrictMode[int])
			assert.NoErr(t, err)
			_, err = mathutil.ToInt64With(input, mathutil.WithStrictMode[int64])
			assert.NoErr(t, err)
			_, err = mathutil.ToUintWith(input, mathutil.WithStrictMode[uint])
			assert.NoErr(t, err)
			_, err = mathutil.ToUint64With(input, mathutil.WithStrictMode[uint64])
			assert.NoErr(t, err)
			_, err = mathutil.ToFloatWith(input)
			assert.NoErr(t, err)
		})
		t.Run("fail", func(t *testing.T) {
			_, err := mathutil.ToIntWith(input, mathutil.WithNilAsFail[int])
			assert.Error(t, err)
			_, err = mathutil.ToInt64With(input, mathutil.WithNilAsFail[int64])
			assert.Error(t, err)
			_, err = mathutil.ToUintWith(input, mathutil.WithNilAsFail[uint])
			assert.Error(t, err)
			_, err = mathutil.ToUint64With(input, mathutil.WithNilAsFail[uint64])
			assert.Error(t, err)
			_, err = mathutil.ToFloatWith(input, mathutil.WithNilAsFail[float64])
			assert.Error(t, err)
		})
	}
}
