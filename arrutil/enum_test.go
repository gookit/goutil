package arrutil_test

import (
	"testing"

	"github.com/gookit/goutil/arrutil"
	"github.com/stretchr/testify/assert"
)

func TestInts_Has_String(t *testing.T) {
	tests := []struct {
		is    arrutil.Ints
		val   int
		want  bool
		want2 string
	}{
		{
			arrutil.Ints{12, 23},
			12,
			true,
			"12,23",
		},
	}

	for _, tt := range tests {
		assert.True(t, tt.want, tt.is.Has(tt.val))
		assert.Equal(t, tt.want2, tt.is.String())
	}
}

func TestStrings_Has_String(t *testing.T) {
	tests := []struct {
		ss    arrutil.Strings
		val   string
		want  bool
		want2 string
	}{
		{
			arrutil.Strings{"a", "b"},
			"a",
			true,
			"a,b",
		},
	}

	for _, tt := range tests {
		assert.True(t, tt.want, tt.ss.Has(tt.val))
		assert.Equal(t, tt.want2, tt.ss.String())
	}
}
