package sysutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrentShell(t *testing.T) {
	path := CurrentShell(true)
	assert.NotEmpty(t, path)

	if path != "" {
		path = CurrentShell(false)
		assert.NotEmpty(t, path)
	}
}

func TestOS(t *testing.T) {
	if isw := IsWin(); isw {
		assert.True(t, isw)
		assert.False(t, IsMac())
		assert.False(t, IsLinux())
	}

	if ism := IsMac(); ism {
		assert.True(t, ism)
		assert.False(t, IsWin())
		assert.False(t, IsLinux())
	}

	if isl := IsLinux(); isl {
		assert.True(t, isl)
		assert.False(t, IsMac())
		assert.False(t, IsWin())
	}
}
