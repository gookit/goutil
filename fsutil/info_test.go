package fsutil_test

import (
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestExpandPath(t *testing.T) {
	path := "~/.kite"

	assert.NotEq(t, path, fsutil.Expand(path))
	assert.NotEq(t, path, fsutil.ExpandPath(path))
	assert.NotEq(t, path, fsutil.ResolvePath(path))

	assert.Eq(t, "", fsutil.Expand(""))
	assert.Eq(t, "/path/to", fsutil.Expand("/path/to"))
}
