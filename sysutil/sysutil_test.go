package sysutil_test

import (
	"os"
	"testing"

	"github.com/gookit/goutil/sysutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestProcessExists(t *testing.T) {
	pid := os.Getpid()

	assert.True(t, sysutil.ProcessExists(pid))
}
