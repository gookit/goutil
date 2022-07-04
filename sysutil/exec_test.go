package sysutil_test

import (
	"testing"

	"github.com/gookit/goutil/sysutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestFindExecutable(t *testing.T) {
	path, err := sysutil.Executable("type")
	assert.NoErr(t, err)
	assert.NotEmpty(t, path)

	path, err = sysutil.FindExecutable("type")
	assert.NoErr(t, err)
	assert.NotEmpty(t, path)

	assert.True(t, sysutil.HasExecutable("type"))
}
