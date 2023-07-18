package goinfo_test

import (
	"testing"

	"github.com/gookit/goutil/goinfo"
	"github.com/gookit/goutil/testutil/assert"
)

func TestGoVersion(t *testing.T) {
	assert.NotEmpty(t, goinfo.GoVersion())
}
