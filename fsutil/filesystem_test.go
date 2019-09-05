package fsutil_test

import (
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/stretchr/testify/assert"
)

func TestFileExists(t *testing.T) {
	assert.False(t, fsutil.FileExists(""))
	assert.False(t, fsutil.FileExists("/not-exist"))
}

func TestIsAbsPath(t *testing.T) {
	assert.True(t, fsutil.IsAbsPath("/data/some.txt"))
}
