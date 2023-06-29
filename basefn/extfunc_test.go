package basefn_test

import (
	"testing"

	"github.com/gookit/goutil/basefn"
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
		assert.Eq(t, tt.want, basefn.DataSize(tt.args))
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
		{1200346778, "13892 days"},
	}

	for _, tt := range tests {
		assert.Eq(t, tt.want, basefn.HowLongAgo(tt.args))
	}
}
