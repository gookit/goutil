package fsutil_test

import (
	"bytes"
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestExtByMimeType(t *testing.T) {
	tests := []struct {
		mime string
		ok   bool
		exts []string
		ext  string
	}{
		{"application/json", true, []string{".json"}, ".json"},
		{"application/xml", true, []string{".xml"}, ".xml"},
		{"application/x-yaml", true, []string{".yaml", ".yml"}, ".yaml"},
		{"text/plain; charset=utf-8", true, []string{".asc", ".txt"}, ".txt"},
	}

	for _, tt := range tests {
		es, err := fsutil.ExtsByMimeType(tt.mime)
		if tt.ok {
			assert.NoErr(t, err)
			assert.Eq(t, tt.exts, es)
		} else {
			assert.Err(t, err)
			assert.Empty(t, es)
		}

		ext, err := fsutil.ExtByMimeType(tt.mime, "")
		if tt.ok {
			assert.NoErr(t, err)
			assert.Eq(t, tt.ext, ext)
		} else {
			assert.Err(t, err)
			assert.Empty(t, ext)
		}
	}

	ext, err := fsutil.ExtByMimeType("application/not-exists", "default")
	assert.NoErr(t, err)
	assert.Eq(t, "default", ext)
}

func TestMimeType(t *testing.T) {
	assert.Eq(t, "", fsutil.DetectMime(""))
	assert.Eq(t, "", fsutil.MimeType("not-exist"))
	assert.Eq(t, "image/jpeg", fsutil.MimeType("testdata/test.jpg"))

	buf := new(bytes.Buffer)
	buf.Write([]byte("\xFF\xD8\xFF"))
	assert.Eq(t, "image/jpeg", fsutil.ReaderMimeType(buf))
	buf.Reset()

	buf.Write([]byte("text"))
	assert.Eq(t, "text/plain; charset=utf-8", fsutil.ReaderMimeType(buf))
	buf.Reset()

	buf.Write([]byte(""))
	assert.Eq(t, "", fsutil.ReaderMimeType(buf))
	buf.Reset()

	assert.True(t, fsutil.IsImageFile("testdata/test.jpg"))
	assert.False(t, fsutil.IsImageFile("testdata/not-exists"))
}
