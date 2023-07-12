package fsutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/testutil/fakeobj"
)

func TestNewEntry(t *testing.T) {
	fPath := "/path/to/some.txt"
	ent := fakeobj.NewDirEntry(fPath)

	fe := fsutil.NewEntry(fPath, ent)
	assert.False(t, fe.IsDir())
	assert.Eq(t, fPath, fe.Path())
	assert.Eq(t, fsutil.Name(fPath), fe.Name())
	assert.Equal(t, "file: /path/to/some.txt", fmt.Sprint(fe))

	fi, err := fe.Info()
	assert.NoError(t, err)
	// dump.P(fi)
	assert.Equal(t, "some.txt", fi.Name())

	nfi := fsutil.NewFileInfo(fPath, fi)
	assert.Equal(t, fPath, nfi.Path())
}
