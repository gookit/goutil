package fsutil_test

import (
	"strings"
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/testutil/fakeobj"
)

func TestDiscardReader(t *testing.T) {
	sr := strings.NewReader("hello")
	bs, err := fsutil.ReadOrErr(sr)
	assert.NoErr(t, err)
	assert.Eq(t, []byte("hello"), bs)

	sr = strings.NewReader("hello")
	assert.Eq(t, []byte("hello"), fsutil.GetContents(sr))

	sr = strings.NewReader("hello")
	fsutil.DiscardReader(sr)

	assert.Empty(t, fsutil.ReadReader(sr))
	assert.Empty(t, fsutil.ReadAll(sr))
}

func TestReadReader(t *testing.T) {
	fr := fakeobj.NewReader()
	assert.Empty(t, fsutil.ReadReader(fr))

	assert.Panics(t, func() {
		fr.ErrOnRead = true
		fsutil.ReadReader(fr)
	})
}

func TestGetContents(t *testing.T) {
	fpath := "./testdata/get-contents.txt"
	assert.NoErr(t, fsutil.RmFileIfExist(fpath))

	_, err := fsutil.PutContents(fpath, "hello")
	assert.NoErr(t, err)

	assert.Nil(t, fsutil.ReadExistFile("/path-not-exist"))
	assert.Eq(t, []byte("hello"), fsutil.ReadExistFile(fpath))

	assert.Panics(t, func() {
		fsutil.GetContents(45)
	})
	assert.Panics(t, func() {
		fsutil.ReadFile("/path-not-exist")
	})
}
