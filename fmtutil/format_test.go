package fmtutil_test

import (
	"testing"

	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/fmtutil"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/timex"
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
		assert.Eq(t, tt.bytes, fmtutil.ParseByte(tt.sizeS))
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

	tests := []struct {
		args []any
		want string
	}{
		{nil, ""},
		{[]any{"a", "b", "c"}, "a b c"},
		{[]any{"a", "b", "c", 1, 2, 3}, "a b c 1 2 3"},
		{[]any{"a", 1, nil}, "a 1 <nil>"},
		{[]any{12, int8(12), int16(12), int32(12), int64(12)}, "12 12 12 12 12"},
		{[]any{uint(12), uint8(12), uint16(12), uint32(12), uint64(12)}, "12 12 12 12 12"},
		{[]any{float32(12.12), 12.12}, "12.12 12.12"},
		{[]any{true, false}, "true false"},
		{[]any{[]byte("abc"), []byte("123")}, "abc 123"},
		{[]any{timex.OneHour}, "3600000000000"},
		{[]any{errorx.Raw("a error message")}, "a error message"},
		{[]any{[]int{1, 2, 3}}, "[1 2 3]"},
	}

	for _, tt := range tests {
		assert.Eq(t, tt.want, fmtutil.ArgsWithSpaces(tt.args))
	}

	assert.NotEmpty(t, fmtutil.ArgsWithSpaces([]any{timex.Now().T()}))
}
