package fmtutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/fmtutil"
	"github.com/stretchr/testify/assert"
)

func TestDataSize(t *testing.T) {
	tests := []struct {
		args uint64
		want string
	}{
		{346, "346B"},
		{3467, "3.39K"},
		{346778, "338.65K"},
		{1200346778, "1.12G"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, fmtutil.DataSize(tt.args))
	}
}

func TestPrettyJSON(t *testing.T) {
	tests := []interface{}{
		map[string]int{"a": 1},
		struct {
			A int `json:"a"`
		}{1},
	}
	want := `{
    "a": 1
}`
	for _, sample := range tests {
		got, err := fmtutil.PrettyJSON(sample)
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	}
}

func TestStringsToInts(t *testing.T) {
	is := assert.New(t)

	ints, err := fmtutil.StringsToInts([]string{"1", "2"})
	is.Nil(err)
	is.Equal("[]int{1, 2}", fmt.Sprintf("%#v", ints))

	_, err = fmtutil.StringsToInts([]string{"a", "b"})
	is.Error(err)
}
