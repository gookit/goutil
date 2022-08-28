package sysutil_test

import (
	"testing"

	"github.com/gookit/goutil/sysutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestCallersInfo(t *testing.T) {
	cs := sysutil.CallersInfos(0, 2)
	// dump.P(cs)
	assert.NotEmpty(t, cs)
	assert.Len(t, cs, 2)
	assert.StrContains(t, cs[0].String(), "gookit/goutil/sysutil/stack.go")
}
