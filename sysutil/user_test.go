package sysutil_test

import (
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/sysutil"
	"github.com/stretchr/testify/assert"
)

func TestUserDir(t *testing.T) {
	dir := sysutil.UserHomeDir()
	assert.NotEmpty(t, dir)
	dump.P(dir)

	dir = sysutil.UserDir("sub-path")
	assert.Contains(t, dir, "/sub-path")
	dump.P(dir)

	dir = sysutil.UserCacheDir("my-logs")
	assert.Contains(t, dir, ".cache/my-logs")
	dump.P(dir)

	dir = sysutil.UserConfigDir("my-conf")
	assert.Contains(t, dir, ".config/my-conf")
	dump.P(dir)
}

func TestWorkdir(t *testing.T) {
	assert.NotEmpty(t, sysutil.Workdir())
}

func TestLoginUser(t *testing.T) {
	cu := sysutil.LoginUser()
	assert.NotEmpty(t, cu)

	fu := sysutil.MustFindUser(cu.Username)
	assert.NotEmpty(t, fu)
	assert.Equal(t, cu.Uid, fu.Uid)
}
