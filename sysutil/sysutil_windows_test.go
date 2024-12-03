//go:build windows
// +build windows

package sysutil_test

import (
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/sysutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestFetchOsVersion(t *testing.T) {
	ov, err := sysutil.FetchOsVersion()
	assert.NoErr(t, err)
	assert.NotEmpty(t, ov)
	assert.NotEmpty(t, ov.String())

	dump.P(ov.Name(), ov)
	assert.NotEmpty(t, sysutil.OsVersion())
}
