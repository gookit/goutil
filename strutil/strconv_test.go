package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/stretchr/testify/assert"
)

func TestStringToBool(t *testing.T) {
	is := assert.New(t)

	tests1 := map[string]bool{
		"1":     true,
		"on":    true,
		"yes":   true,
		"true":  true,
		"false": false,
		"off":   false,
		"no":    false,
		"0":     false,
	}

	for str, want := range tests1 {
		is.Equal(want, strutil.MustBool(str))
	}

	blVal, err := strutil.ToBool("1")
	is.Nil(err)
	is.True(blVal)

	blVal, err = strutil.Bool("10")
	is.Error(err)
	is.False(blVal)
}

func TestValToString(t *testing.T) {
	is := assert.New(t)

	tests := []interface{}{
		2,
		int8(2), int16(2), int32(2), int64(2),
		uint(2), uint8(2), uint16(2), uint32(2), uint64(2),
		"2",
	}
	for _, in := range tests {
		is.Equal("2", strutil.MustString(in))
	}

	tests1 := []interface{}{
		float32(2.3), 2.3,
	}
	for _, in := range tests1 {
		is.Equal("2.3", strutil.MustString(in))
	}

	str, err := strutil.String(2.3)
	is.NoError(err)
	is.Equal("2.3", str)

	str, err = strutil.String(nil)
	is.NoError(err)
	is.Equal("", str)

	_, err = strutil.String([]string{"a"})
	is.Error(err)
}
