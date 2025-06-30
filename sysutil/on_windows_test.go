//go:build windows

package sysutil_test

import (
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/sysutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestWin_common(t *testing.T) {
	assert.False(t, sysutil.IsMac())
	assert.False(t, sysutil.IsDarwin())
	assert.False(t, sysutil.IsLinux())
	assert.Err(t, sysutil.Kill(0, 0))
}

func TestOsVersionByParse(t *testing.T) {
	// win10
	ov, err := sysutil.OsVersionByParse("Microsoft Windows [Version 10.0.21000]")
	assert.NoErr(t, err)
	assert.NotEmpty(t, ov)
	assert.True(t, ov.IsWindows10())
	assert.False(t, ov.IsWindows11())
	assert.False(t, ov.IsWindows8())
	assert.False(t, ov.IsWindows7())
	assert.False(t, ov.IsLtWindows7())

	// win8.1
	ov, err = sysutil.OsVersionByParse("Microsoft Windows [Version 6.3.21000]")
	assert.NoErr(t, err)
	assert.NotEmpty(t, ov)
	assert.True(t, ov.IsWindows8())
	// win8
	ov, err = sysutil.OsVersionByParse("Microsoft Windows [Version 6.2.21000]")
	assert.NoErr(t, err)
	assert.NotEmpty(t, ov)
	assert.True(t, ov.IsWindows8())
	// win7
	ov, err = sysutil.OsVersionByParse("Microsoft Windows [Version 6.1.21000]")
	assert.NoErr(t, err)
	assert.NotEmpty(t, ov)
	assert.True(t, ov.IsWindows7())

	// win xp
	ov, err = sysutil.OsVersionByParse("Microsoft Windows [Version 5.1.21000]")
	assert.NoErr(t, err)
	assert.NotEmpty(t, ov)
	assert.True(t, ov.IsLtWindows7())
}

func TestOsVersionByVerCmd(t *testing.T) {
	ov := sysutil.OsVersion()
	assert.NotEmpty(t, ov)

	ov, err := sysutil.OsVersionByVerCmd()
	assert.NoErr(t, err)
	assert.NotEmpty(t, ov)
	assert.NotEmpty(t, ov.String())

	if ov.IsWindows11() {
		assert.Eq(t, "win11", ov.Name())
		assert.True(t, ov.IsWindows11())
		assert.False(t, ov.IsWindows10())
		assert.False(t, ov.IsWindows8())
		assert.False(t, ov.IsWindows7())
		assert.False(t, ov.IsLtWindows7())
	}

	dump.P(ov.Name(), ov)
}

func TestUser_func(t *testing.T) {
	assert.NoErr(t, sysutil.ChangeUserByName("admin"))
	assert.NoErr(t, sysutil.ChangeUserUidGid(1, 1))

	if ok := sysutil.IsAdmin(); ok {
		assert.True(t, ok)
	} else {
		assert.False(t, ok)
	}
}
