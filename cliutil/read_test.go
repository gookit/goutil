package cliutil_test

import (
	"testing"

	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestReadFirst(t *testing.T) {
	// testutil.RewriteStdout()
	// _, err := os.Stdout.WriteString("haha")
	// ans, err1 := cliutil.ReadFirst("hi?")
	// testutil.RestoreStdout()
	// assert.NoError(t, err)
	// assert.NoError(t, err1)
	// assert.Equal(t, "haha", ans)
}

func TestInputIsYes(t *testing.T) {
	tests := []struct {
		in  string
		wnt bool
	}{
		{"y", true},
		{"yes", true},
		{"yES", true},
		{"Y", true},
		{"Yes", true},
		{"YES", true},
		{"h", false},
		{"n", false},
		{"no", false},
		{"NO", false},
	}
	for _, test := range tests {
		assert.Eq(t, test.wnt, cliutil.InputIsYes(test.in))
	}
}

func TestByteIsYes(t *testing.T) {
	tests := []struct {
		in  byte
		wnt bool
	}{
		{'y', true},
		{'Y', true},
		{'h', false},
		{'n', false},
		{'N', false},
	}
	for _, test := range tests {
		assert.Eq(t, test.wnt, cliutil.ByteIsYes(test.in))
	}
}
