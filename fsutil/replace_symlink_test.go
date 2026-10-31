package fsutil_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/x/assert"
)

func TestCreateSymlinkReplacesExistingPath(t *testing.T) {
	for _, kind := range []string{"dangling link", "directory link", "file link", "regular file", "missing path", "real directory"} {
		t.Run(kind, func(t *testing.T) {
			dir := t.TempDir()
			oldTarget := filepath.Join(dir, "old")
			newTarget := filepath.Join(dir, "new")
			linkPath := filepath.Join(dir, "link")
			assert.NoErr(t, os.WriteFile(newTarget, []byte("new contents"), 0600))
			probe := filepath.Join(dir, "probe")
			if err := os.Symlink(newTarget, probe); err != nil {
				t.Skipf("symbolic links unavailable: %v", err)
			}
			assert.NoErr(t, os.Remove(probe))
			switch kind {
			case "directory link":
				assert.NoErr(t, os.Mkdir(oldTarget, 0700))
				assert.NoErr(t, os.WriteFile(filepath.Join(oldTarget, "marker"), []byte("keep directory contents"), 0600))
			case "file link":
				assert.NoErr(t, os.WriteFile(oldTarget, []byte("keep file contents"), 0600))
			case "regular file":
				assert.NoErr(t, os.WriteFile(linkPath, []byte("replace this file"), 0600))
			case "real directory":
				assert.NoErr(t, os.Mkdir(linkPath, 0700))
				assert.NoErr(t, os.WriteFile(filepath.Join(linkPath, "marker"), []byte("keep real directory"), 0600))
			}
			if kind == "dangling link" || kind == "directory link" || kind == "file link" {
				assert.NoErr(t, os.Symlink(oldTarget, linkPath))
			}
			err := fsutil.CreateSymlink(newTarget, linkPath)
			if kind == "real directory" {
				if err == nil {
					t.Fatal("CreateSymlink replaced a real directory")
				}
				got, err := os.ReadFile(filepath.Join(linkPath, "marker"))
				assert.NoErr(t, err)
				assert.Eq(t, "keep real directory", string(got))
				info, err := os.Lstat(linkPath)
				assert.NoErr(t, err)
				assert.True(t, info.IsDir())
				return
			}
			assert.NoErr(t, err)
			gotTarget, err := os.Readlink(linkPath)
			assert.NoErr(t, err)
			assert.Eq(t, newTarget, gotTarget)
			got, err := os.ReadFile(linkPath)
			assert.NoErr(t, err)
			assert.Eq(t, "new contents", string(got))
			switch kind {
			case "directory link":
				got, err = os.ReadFile(filepath.Join(oldTarget, "marker"))
				assert.NoErr(t, err)
				assert.Eq(t, "keep directory contents", string(got))
			case "file link":
				got, err = os.ReadFile(oldTarget)
				assert.NoErr(t, err)
				assert.Eq(t, "keep file contents", string(got))
			case "dangling link":
				_, err := os.Lstat(oldTarget)
				assert.True(t, os.IsNotExist(err))
			}
		})
	}
}
