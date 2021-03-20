package sysutil_test

import (
	"os"
	"testing"

	"github.com/gookit/goutil/sysutil"
	"github.com/stretchr/testify/assert"
)

func TestProcessExists(t *testing.T) {
	pid := os.Getpid()

	assert.True(t, sysutil.ProcessExists(pid))
}
