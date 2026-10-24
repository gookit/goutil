package mathutil_test

import (
	"testing"

	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/x/assert"
)

func TestStrictIntegerFloatInputs(t *testing.T) {
	for _, value := range []any{float32(1.5), float64(1.5)} {
		t.Run("int", func(t *testing.T) {
			_, err := mathutil.ToIntWith(value, mathutil.WithStrictMode[int])
			assert.Error(t, err)
		})
		t.Run("uint", func(t *testing.T) {
			_, err := mathutil.ToUintWith(value, mathutil.WithStrictMode[uint])
			assert.Error(t, err)
		})
		t.Run("uint64", func(t *testing.T) {
			_, err := mathutil.ToUint64With(value, mathutil.WithStrictMode[uint64])
			assert.Error(t, err)
		})
	}
}
