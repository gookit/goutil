//go:build windows

package fsutil_test

import (
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestJoinPaths_win(t *testing.T) {
	assert.Eq(t, "path\\to\\dir", fsutil.JoinPaths("path", "to", "dir"))
	assert.Eq(t, "path\\to\\dir", fsutil.JoinPaths3("path", "to", "dir"))
	assert.Eq(t, "path\\to\\dir", fsutil.JoinSubPaths("path", "to", "dir"))
}

func TestRealpath_win(t *testing.T) {
	inPath := "/path/to/some/../dir"
	assert.Eq(t, "\\path\\to\\dir", fsutil.Realpath(inPath))
}
