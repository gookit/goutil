package sysutil_test

import (
	"os"
	"testing"

	"github.com/gookit/goutil/sysutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestBasic_usage(t *testing.T) {
	assert.NotEmpty(t, sysutil.BinDir())
	assert.NotEmpty(t, sysutil.BinFile())
}

func TestProcessExists(t *testing.T) {
	pid := os.Getpid()

	assert.True(t, sysutil.ProcessExists(pid))
}
