//go:build !windows
// +build !windows

package sysutil_test

import (
	"testing"

	"github.com/gookit/goutil/sysutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestUser_func(t *testing.T) {
	assert.False(t, sysutil.IsAdmin())
	assert.NoErr(t, sysutil.ChangeUserUidGid(0, 1000))
}
