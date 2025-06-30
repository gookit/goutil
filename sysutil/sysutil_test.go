package sysutil_test

import (
	"os"
	"runtime"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/sysutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestSysUtil_basic(t *testing.T) {
	assert.NotEmpty(t, sysutil.BinDir())
	assert.NotEmpty(t, sysutil.BinName())
	assert.NotEmpty(t, sysutil.BinFile())
	// echo $PSVersionTable.PSVersion.ToString()

	// 	return windows.ShellExecute(0, nil, windows.StringToUTF16Ptr(url), nil, nil, windows.SW_SHOWNORMAL)
	dump.P(sysutil.ExecCmd("echo", []string{"$PSVersionTable.PSVersion.ToString()"}))
}

func TestOpenFile(t *testing.T) {
	assert.Err(t, sysutil.Open("open_a_invalid_path"))
	assert.Err(t, sysutil.OpenFile("open_a_invalid_path"))
	assert.Err(t, sysutil.OpenBrowser("open_a_invalid_path"))
}

func TestProcessExists(t *testing.T) {
	if runtime.GOOS != "windows" {
		pid := os.Getpid()
		assert.True(t, sysutil.ProcessExists(pid))
	} else {
		assert.Panics(t, func() { sysutil.ProcessExists(0) })
	}
}
