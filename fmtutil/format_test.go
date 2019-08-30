package fmtutil_test

import (
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
