package process_test

import (
	"os"
	"testing"

	"github.com/gookit/goutil/sysutil/process"
	"github.com/gookit/goutil/testutil/assert"
)

func TestProcessExists(t *testing.T) {
	pid := os.Getpid()

	assert.True(t, process.Exists(pid))
}

func TestPID(t *testing.T) {
	assert.True(t, process.PID() > 0)
}
