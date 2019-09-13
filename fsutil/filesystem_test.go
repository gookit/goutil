package fsutil_test

import (
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/stretchr/testify/assert"
)

func TestPathExists(t *testing.T) {
	assert.False(t, fsutil.FileExists(""))
	assert.False(t, fsutil.PathExists("/not-exist"))
}

func TestIsFile(t *testing.T) {
	assert.False(t, fsutil.IsFile(""))
	assert.False(t, fsutil.IsFile("/not-exist"))
}

func TestIsDir(t *testing.T) {
	assert.False(t, fsutil.IsDir(""))
	assert.False(t, fsutil.IsDir("/not-exist"))
}

func TestIsAbsPath(t *testing.T) {
	assert.True(t, fsutil.IsAbsPath("/data/some.txt"))
}
