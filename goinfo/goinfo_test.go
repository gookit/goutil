package goinfo_test

import (
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/goinfo"
	"github.com/gookit/goutil/testutil/assert"
)

func TestGoVersion(t *testing.T) {
	assert.NotEmpty(t, goinfo.GoVersion())
}

func TestOsGoInfo(t *testing.T) {
	assert.NotEmpty(t, goinfo.GoVersion())

	info, err := goinfo.ParseGoVersion("go version go1.19.2 darwin/amd64")
	assert.NoErr(t, err)
	assert.NotEmpty(t, info)
	assert.Eq(t, "1.19.2", info.Version)
	assert.Eq(t, "darwin", info.GoOS)
	assert.Eq(t, "amd64", info.Arch)

	_, err = goinfo.ParseGoVersion("invalid version")
	assert.Err(t, err)

	info, err = goinfo.OsGoInfo()
	assert.NoErr(t, err)
	assert.NotEmpty(t, info)
	dump.P(info)
}
