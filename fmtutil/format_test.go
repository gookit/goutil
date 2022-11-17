package fmtutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/fmtutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestDataSize(t *testing.T) {
	tests := []struct {
		args uint64
		want string
	}{
		{346, "346B"},
		{3467, "3.39K"},
		{346778, "338.65K"},
		{12346778, "11.77M"},
		{1200346778, "1.12G"},
	}

	for _, tt := range tests {
		assert.Eq(t, tt.want, fmtutil.DataSize(tt.args))
	}
	assert.Eq(t, "1.12G", fmtutil.SizeToString(1200346778))
}

func TestParseByte(t *testing.T) {
	tests := []struct {
		bytes uint64
		sizeS string
	}{
		{346, "346B"},
		{3471, "3.39K"},
		{346777, "338.65Kb"},
		{12341739, "11.77M"},
		{1202590842, "1.12GB"},
	}

	for _, tt := range tests {
		assert.Eq(t, tt.bytes, fmtutil.StringToByte(tt.sizeS))
	}
}

func TestPrettyJSON(t *testing.T) {
	tests := []any{
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
		assert.NoErr(t, err)
		assert.Eq(t, want, got)
	}
}

func TestArgsWithSpaces(t *testing.T) {
	assert.Eq(t, "", fmtutil.ArgsWithSpaces(nil))
	assert.Eq(t, "", fmtutil.ArgsWithSpaces([]any{}))
	assert.Eq(t, "abc", fmtutil.ArgsWithSpaces([]any{"abc"}))
	assert.Eq(t, "23 abc", fmtutil.ArgsWithSpaces([]any{23, "abc"}))
}

func TestStringsToInts(t *testing.T) {
	is := assert.New(t)

	ints, err := fmtutil.StringsToInts([]string{"1", "2"})
	is.Nil(err)
	is.Eq("[]int{1, 2}", fmt.Sprintf("%#v", ints))

	_, err = fmtutil.StringsToInts([]string{"a", "b"})
	is.Err(err)
}
