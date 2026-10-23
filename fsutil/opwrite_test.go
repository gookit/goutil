package fsutil_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/x/assert"
)

func TestMustCopyFile(t *testing.T) {
	srcPath := "./testdata/cp-file-src.txt"
	dstPath := "./testdata/cp-file-dst.txt"

	assert.NoErr(t, fsutil.RmIfExist(srcPath))
	assert.NoErr(t, fsutil.RmFileIfExist(dstPath))

	fsutil.Must2(fsutil.PutContents(srcPath, "hello"))
	fsutil.MustCopyFile(srcPath, dstPath)
	assert.Eq(t, []byte("hello"), fsutil.GetContents(dstPath))
	assert.Eq(t, "hello", fsutil.ReadString(dstPath))

	assert.Panics(t, func() {
		fsutil.MustCopyFile("testdata/not-exists-file", "")
	})

	str, err := fsutil.ReadStringOrErr(dstPath)
	assert.NoErr(t, err)
	assert.Eq(t, "hello", str)

	assert.NoErr(t, fsutil.RmFileIfExist(srcPath))
	assert.NoErr(t, fsutil.RmIfExist(srcPath)) // repeat delete
}

func TestWriteFile(t *testing.T) {
	tempFile := "./testdata/write-file.txt"

	err := fsutil.WriteFile(tempFile, []byte("hello\ngolang"), 0644)
	assert.NoErr(t, err)
	assert.Eq(t, []byte("hello\ngolang"), fsutil.MustReadFile(tempFile))

	// write invalid data
	assert.Panics(t, func() {
		_ = fsutil.WriteFile(tempFile, []string{"hello"}, 0644)
	})
}

func TestMustSave(t *testing.T) {
	opt := fsutil.OpenOptOrNew(nil)
	assert.NotNil(t, opt)
	opt = fsutil.OpenOptOrNew(fsutil.NewOpenOption())
	assert.NotNil(t, opt)

	testFile := "./testdata/must-save.txt"

	fsutil.MustSave(testFile, []byte("hello"),
		fsutil.WithFlag(fsutil.FsCWTFlags),
		fsutil.WithPerm(0644),
	)
	assert.Eq(t, "hello", fsutil.ReadString(testFile))

	// write invalid data
	assert.Panics(t, func() {
		fsutil.MustSave(testFile, []string{"hello"})
	})
}

func TestUpdateContents(t *testing.T) {
	// the file not exists
	err := fsutil.UpdateContents("testdata/not-exists-file", func(bs []byte) []byte {
		return bs
	})
	assert.Err(t, err)

	of, err := fsutil.TempFile("testdata", "test-update-contents-*.txt")
	assert.NoErr(t, err)

	dump.P(of.Name())
	_, err = of.WriteString("hello")
	assert.NoErr(t, err)
	assert.NoErr(t, of.Close())

	err = fsutil.UpdateContents(of.Name(), func(bs []byte) []byte {
		return []byte("hello, world")
	})
	assert.NoErr(t, err)
	assert.Eq(t, "hello, world", fsutil.ReadString(of.Name()))
}

func TestUpdateContentsTransformsExistingContent(t *testing.T) {
	for _, tt := range []struct {
		name string
		old  string
		ret  []byte
		want string
	}{
		{"append", "hello", []byte("hello, world"), "hello, world"},
		{"longer multiline", "a\nb", []byte("a\nb\nc\nd"), "a\nb\nc\nd"},
		{"shorten", "hello", []byte("hi"), "hi"},
		{"unchanged", "hello", []byte("hello"), "hello"},
		{"empty result", "hello", []byte(""), ""},
		{"nil result", "hello", nil, ""},
		{"from empty file", "", []byte("filled"), "filled"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "contents.txt")
			assert.NoErr(t, os.WriteFile(path, []byte(tt.old), 0600))

			err := fsutil.UpdateContents(path, func(bs []byte) []byte {
				assert.Eq(t, tt.old, string(bs))
				return tt.ret
			})
			assert.NoErr(t, err)
			got, err := os.ReadFile(path)
			assert.NoErr(t, err)
			assert.Eq(t, tt.want, string(got))
		})
	}
}

func TestUpdateContentsNilHandler(t *testing.T) {
	path := filepath.Join(t.TempDir(), "contents.txt")
	assert.NoErr(t, os.WriteFile(path, []byte("hello"), 0600))

	err := fsutil.UpdateContents(path, nil)
	assert.ErrMsgContains(t, err, "handleFn is nil")

	// the file must be left as it was
	assert.Eq(t, "hello", fsutil.ReadString(path))
}

func TestUpdateContentsPanicInHandler(t *testing.T) {
	path := filepath.Join(t.TempDir(), "contents.txt")
	assert.NoErr(t, os.WriteFile(path, []byte("hello"), 0600))

	assert.Panics(t, func() {
		_ = fsutil.UpdateContents(path, func(bs []byte) []byte {
			panic("boom")
		})
	})

	// the file is not emptied before the handleFn is called
	assert.Eq(t, "hello", fsutil.ReadString(path))
}

func TestOSTempFile(t *testing.T) {
	of, err := fsutil.OSTempFile("test-os-tmpfile-*.txt")
	assert.NoErr(t, err)
	defer of.Close()

	dump.P(of.Name())
	assert.StrContains(t, of.Name(), "test-os-tmpfile-")
}

func TestTempDir(t *testing.T) {
	dir, err := fsutil.TempDir("testdata", "temp-dir-*")
	assert.NoErr(t, err)
	assert.True(t, fsutil.IsDir(dir))
	assert.NoErr(t, fsutil.Remove(dir))

	dir, err = fsutil.OSTempDir("os-temp-dir-*")
	assert.NoErr(t, err)
	assert.True(t, fsutil.IsDir(dir))
	assert.True(t, fsutil.IsEmptyDir(dir))
}

func TestCreateSymlink(t *testing.T) {
	srcFile := "./testdata/symlink-src.txt"
	assert.NoErr(t, fsutil.SaveFile(srcFile, "hello"))
	assert.Eq(t, "hello", fsutil.ReadString(srcFile))

	linkFile := "./testdata/symlink-dst.txt"
	assert.NoErr(t, fsutil.CreateSymlink(srcFile, linkFile))
	assert.False(t, fsutil.IsFile(linkFile))
	assert.True(t, fsutil.IsSymlink(linkFile))
}
