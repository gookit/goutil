package mathutil_test

import (
	"testing"

	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/x/assert"
)

func TestHandlePtrUnsupportedValues(t *testing.T) {
	for _, value := range []any{struct{}{}, &struct{}{}} {
		t.Run("int", func(t *testing.T) {
			_, err := mathutil.ToIntWith(value, mathutil.WithHandlePtr[int])
			assert.Error(t, err)
		})
		t.Run("int64", func(t *testing.T) {
			_, err := mathutil.ToInt64With(value, mathutil.WithHandlePtr[int64])
			assert.Error(t, err)
		})
		t.Run("uint", func(t *testing.T) {
			_, err := mathutil.ToUintWith(value, mathutil.WithHandlePtr[uint])
			assert.Error(t, err)
		})
		t.Run("uint64", func(t *testing.T) {
			_, err := mathutil.ToUint64With(value, mathutil.WithHandlePtr[uint64])
			assert.Error(t, err)
		})
	}
}
