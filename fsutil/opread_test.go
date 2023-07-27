package fsutil_test

import (
	"strings"
	"testing"
	"text/scanner"

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

	fr.ErrOnRead = true
	assert.Panics(t, func() {
		fsutil.ReadReader(fr)
	})

	_, err := fsutil.ReadStringOrErr(fr)
	assert.Err(t, err)

	_, err = fsutil.ReadStringOrErr([]string{"invalid-type"})
	assert.Err(t, err)
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

func TestTextScanner(t *testing.T) {
	r := strings.NewReader("hello\ngolang")

	ts := fsutil.TextScanner(r)
	assert.Neq(t, scanner.EOF, ts.Scan())
	assert.Eq(t, "hello", ts.TokenText())

	assert.Panics(t, func() {
		fsutil.TextScanner([]string{"invalid-type"})
	})
}

func TestLineScanner(t *testing.T) {
	r := strings.NewReader("hello\ngolang")

	ls := fsutil.LineScanner(r)
	assert.True(t, ls.Scan())
	assert.Eq(t, "hello", ls.Text())

	assert.Panics(t, func() {
		fsutil.LineScanner([]string{"invalid-type"})
	})
}
