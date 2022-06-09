package fsutil_test

import (
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/stretchr/testify/assert"
)

func TestExpandPath(t *testing.T) {
	path := "~/.kite"

	assert.NotEqual(t, path, fsutil.Expand(path))
	assert.NotEqual(t, path, fsutil.ExpandPath(path))

	assert.Equal(t, "", fsutil.Expand(""))
	assert.Equal(t, "/path/to", fsutil.Expand("/path/to"))
}
