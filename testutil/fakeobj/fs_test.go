package fakeobj_test

import (
	"testing"
	"time"

	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/testutil/fakeobj"
)

func TestNewDirEntry(t *testing.T) {
	de := fakeobj.NewDirEntry("/some/path/to/dir", true)
	assert.True(t, de.IsDir())
	assert.Gt(t, int(de.Type()), 0)
	assert.Equal(t, "dir", de.Name())

	fi, err := de.Info()
	assert.NoError(t, err)
	assert.True(t, fi.IsDir())
	assert.Gt(t, int(fi.Mode()), 0)
	assert.Equal(t, "dir", fi.Name())
}

func TestNewFileInfo(t *testing.T) {
	fi := fakeobj.NewFile("/some/path/to/file")
	fi.WithBody("hello")
	fi.WithMtime(time.Now().Add(-time.Hour))

	assert.False(t, fi.IsDir())
	assert.Gt(t, int(fi.Mode()), 0)
	assert.Gt(t, fi.Size(), 0)
	assert.Equal(t, "file", fi.Name())
	assert.True(t, fi.ModTime().Before(time.Now()))
	assert.NoError(t, fi.Close())
	bs := make([]byte, 5)
	n, err := fi.Read(bs)
	assert.NoError(t, err)
	assert.Gt(t, n, 0)
	assert.Equal(t, "hello", string(bs))
	fi.Reset()
	st, err := fi.Stat()
	assert.NoError(t, err)
	assert.NotEmpty(t, st)

	fi = fakeobj.NewFileInfo("/some/path/to/dir", true)
	assert.True(t, fi.IsDir())
	assert.Gt(t, int(fi.Mode()), 0)
	assert.Equal(t, "dir", fi.Name())
	assert.Nil(t, fi.Sys())
}
