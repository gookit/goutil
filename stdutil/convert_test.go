package stdutil_test

import (
	"testing"

	"github.com/gookit/goutil/stdutil"
	"github.com/gookit/goutil/strutil"
	"github.com/stretchr/testify/assert"
)

func TestMustString(t *testing.T) {
	assert.Equal(t, "23", stdutil.MustString(23))

	assert.PanicsWithError(t, "convert data type is failure", func() {
		stdutil.MustString([]string{"a", "b"})
	})
}

func TestToString(t *testing.T) {
	assert.Equal(t, "23", stdutil.ToString(23))
	assert.Equal(t, "[a b]", stdutil.ToString([]string{"a", "b"}))
}

func TestTryString(t *testing.T) {
	s, err := stdutil.TryString(23)
	assert.NoError(t, err)
	assert.Equal(t, "23", s)

	s, err = stdutil.TryString([]string{"a", "b"})
	assert.ErrorIs(t, err, strutil.ErrConvertFail)
	assert.Equal(t, "", s)
}
