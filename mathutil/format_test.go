package mathutil_test

import (
	"testing"

	"github.com/gookit/goutil/mathutil"
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
		assert.Eq(t, tt.want, mathutil.DataSize(tt.args))
	}
}

func TestFormatBytes(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{0, "0 B"},
		{1024, "1.00 KB"},
		{1536, "1.50 KB"},
		{1048576, "1.00 MB"},
		{1073741824, "1.00 GB"},
		{1200346778, "1.12 GB"},
	}

	for _, tt := range tests {
		result := mathutil.FormatBytes(tt.input)
		assert.Eq(t, tt.expected, result)
	}
}

func TestHowLongAgo(t *testing.T) {
	tests := []struct {
		args int64
		want string
	}{
		{-36, "unknown"},
		{36, "36 secs"},
		{60, "1 min"},
		{346, "5 mins"},
		{3467, "57 mins"},
		{346778, "4 days"},
		{2592000*7 + 2, "7 months"},
	}

	for _, tt := range tests {
		assert.Eq(t, tt.want, mathutil.HowLongAgo(tt.args))
	}
}
