package sysutil_test

import (
	"os"
	"testing"

	"github.com/gookit/goutil/dump"
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

func TestGoVersion(t *testing.T) {
	assert.NotEmpty(t, sysutil.GoVersion())

	info, err := sysutil.ParseGoVersion("go version go1.19.2 darwin/amd64")
	assert.NoErr(t, err)
	assert.NotEmpty(t, info)
	assert.Eq(t, "1.19.2", info.Version)
	assert.Eq(t, "darwin", info.GoOS)
	assert.Eq(t, "amd64", info.Arch)

	info, err = sysutil.OsGoInfo()
	assert.NoErr(t, err)
	assert.NotEmpty(t, info)
	dump.P(info)
}
