package fsutil_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/x/assert"
)

func TestCopyFileRejectsSameFile(t *testing.T) {
	for _, mode := range []string{"same path", "relative alias", "hard link", "symbolic link"} {
		t.Run(mode, func(t *testing.T) {
			dir := t.TempDir()
			src := filepath.Join(dir, "source")
			dst := src
			want := []byte{'k', 'e', 'e', 'p', 0, 255}
			assert.NoErr(t, os.WriteFile(src, want, 0600))
			switch mode {
			case "relative alias":
				assert.NoErr(t, os.Mkdir(filepath.Join(dir, "sub"), 0700))
				dst = dir + string(os.PathSeparator) + "sub" + string(os.PathSeparator) + ".." + string(os.PathSeparator) + "source"
			case "hard link":
				dst = filepath.Join(dir, "alias")
				if err := os.Link(src, dst); err != nil {
					t.Skipf("hard links unavailable: %v", err)
				}
			case "symbolic link":
				dst = filepath.Join(dir, "alias")
				if err := os.Symlink(src, dst); err != nil {
					t.Skipf("symbolic links unavailable: %v", err)
				}
			}
			if err := fsutil.CopyFile(src, dst); err == nil {
				t.Error("CopyFile accepted source and destination resolving to same file")
			}
			got, err := os.ReadFile(src)
			assert.NoErr(t, err)
			assert.Eq(t, want, got)
		})
	}
}
func TestCopyFileDistinctDestination(t *testing.T) {
	for _, existing := range []bool{false, true} {
		t.Run(map[bool]string{false: "create parent", true: "overwrite shorter"}[existing], func(t *testing.T) {
			dir := t.TempDir()
			src := filepath.Join(dir, "source")
			dst := filepath.Join(dir, "nested", "destination")
			want := []byte{'n', 'e', 'w', 0, 255}
			assert.NoErr(t, os.WriteFile(src, want, 0600))
			var originalMode os.FileMode
			if existing {
				assert.NoErr(t, os.Mkdir(filepath.Dir(dst), 0700))
				assert.NoErr(t, os.WriteFile(dst, []byte("old and longer content"), 0600))
				info, err := os.Stat(dst)
				assert.NoErr(t, err)
				originalMode = info.Mode().Perm()
			}
			assert.NoErr(t, fsutil.CopyFile(src, dst))
			got, err := os.ReadFile(dst)
			assert.NoErr(t, err)
			assert.Eq(t, want, got)
			got, err = os.ReadFile(src)
			assert.NoErr(t, err)
			assert.Eq(t, want, got)
			if existing {
				info, err := os.Stat(dst)
				assert.NoErr(t, err)
				assert.Eq(t, originalMode, info.Mode().Perm())
			}
		})
	}
}
func TestCopyFileMissingSourcePreservesDestination(t *testing.T) {
	dir := t.TempDir()
	dst := filepath.Join(dir, "destination")
	assert.NoErr(t, os.WriteFile(dst, []byte("keep"), 0600))
	if err := fsutil.CopyFile(filepath.Join(dir, "absent"), dst); err == nil {
		t.Fatal("missing source accepted")
	}
	got, err := os.ReadFile(dst)
	assert.NoErr(t, err)
	assert.Eq(t, "keep", string(got))
}
func TestMustCopyFileRejectsSameFile(t *testing.T) {
	src := filepath.Join(t.TempDir(), "source")
	assert.NoErr(t, os.WriteFile(src, []byte("keep"), 0600))
	defer func() {
		if recover() == nil {
			t.Error("MustCopyFile did not panic")
		}
		got, err := os.ReadFile(src)
		assert.NoErr(t, err)
		assert.Eq(t, "keep", string(got))
	}()
	fsutil.MustCopyFile(src, src)
}
