package fsutil_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/x/assert"
	"github.com/gookit/goutil/x/fakeobj"
)

func TestFindAllInParentDirs(t *testing.T) {
	// find all
	ss := fsutil.FindAllInParentDirs("testdata", "find.go", func(opt *fsutil.FindParentOption) {
		opt.MaxLevel = 3
	})
	assert.NotEmpty(t, ss)
	assert.Len(t, ss, 1)

	ss = fsutil.FindAllInParentDirs("testdata", "README.md", func(opt *fsutil.FindParentOption) {
		opt.MaxLevel = 5
	})
	assert.NotEmpty(t, ss)
	assert.Gt(t, len(ss), 1)
}

func TestFindOneInParentDirs(t *testing.T) {
	// find one
	s := fsutil.FindOneInParentDirs("testdata", "goutil.go")
	assert.NotEmpty(t, s)
	if runtime.GOOS == "windows" {
		assert.StrContains(t, s, "goutil\\goutil.go")
	} else {
		assert.StrContains(t, s, "goutil/goutil.go")
	}

	// find dir
	s = fsutil.FindOneInParentDirs("testdata", "errorx", func(opt *fsutil.FindParentOption) {
		opt.NeedDir = true
	})
	assert.NotEmpty(t, s)
}

func TestFilePathInDirs(t *testing.T) {
	result := fsutil.FilePathInDirs("not_existing_file.txt")
	assert.Empty(t, result)
	result = fsutil.FilePathInDirs("not_existing_file.txt", "testdata")
	assert.Empty(t, result)

	result = fsutil.FilePathInDirs("find.go")
	assert.NotEmpty(t, result)
}

func TestFilePathInDirsSearchOrder(t *testing.T) {
	root := t.TempDir()
	first := filepath.Join(root, "first")
	second := filepath.Join(root, "second")
	third := filepath.Join(root, "third")
	assert.NoErr(t, os.Mkdir(first, 0700))
	name := "lookup-target.conf"
	nested := filepath.Join("nested", name)
	for _, dir := range []string{second, third} {
		for _, file := range []string{name, nested} {
			fullPath := filepath.Join(dir, file)
			assert.NoErr(t, os.MkdirAll(filepath.Dir(fullPath), 0700))
			assert.NoErr(t, os.WriteFile(fullPath, []byte("present"), 0600))
		}
	}

	for _, tc := range []struct {
		name string
		file string
		dirs []string
		want string
	}{
		{"second directory", name, []string{first, second}, filepath.Join(second, name)},
		{"third directory", name, []string{first, filepath.Join(root, "missing"), third}, filepath.Join(third, name)},
		{"first match wins", name, []string{second, third}, filepath.Join(second, name)},
		{"reversed priority", name, []string{third, second}, filepath.Join(third, name)},
		{"direct path", filepath.Join(second, name), []string{first, third}, filepath.Join(second, name)},
		{"nested file", nested, []string{first, second}, filepath.Join(second, nested)},
		{"missing file", "absent.conf", []string{first, second, third}, ""},
		{"no directories", name, nil, ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Eq(t, tc.want, fsutil.FilePathInDirs(tc.file, tc.dirs...))
		})
	}
}

func TestMatchFirst(t *testing.T) {
	assert.Eq(t, "testdata", fsutil.MatchFirst([]string{"testdata"}, fsutil.IsDir, ""))

	assert.Eq(t, "testdata", fsutil.FirstExists("not-exists", "testdata"))
	assert.Eq(t, "testdata", fsutil.FirstExistsDir("not-exists", "testdata"))
	assert.Eq(t, "testdata/test.jpg", fsutil.FirstExistsFile("not-exists", "testdata/test.jpg"))

	ps := fsutil.MatchPaths([]string{"testdata", "testdata/test.jpg"}, fsutil.IsDir)
	assert.Eq(t, []string{"testdata"}, ps)

	assert.Eq(t, "default_dir", fsutil.MatchFirst([]string{"not_exist_dir"}, fsutil.IsDir, "default_dir"))
}

func TestSearchNameUp(t *testing.T) {
	p := fsutil.SearchNameUp("testdata", "dump")
	assert.NotEmpty(t, p)
	assert.True(t, strings.HasSuffix(p, "goutil"))

	p = fsutil.SearchNameUp("testdata", ".dotdir")
	assert.NotEmpty(t, p)
	assert.True(t, strings.HasSuffix(p, "testdata"))

	p = fsutil.SearchNameUp("testdata", "test.jpg")
	assert.NotEmpty(t, p)
	assert.True(t, strings.HasSuffix(p, "testdata"))

	p = fsutil.SearchNameUp("testdata", "not-exists")
	assert.Empty(t, p)
}

func TestGlobWithFunc(t *testing.T) {
	assert.NotEmpty(t, fsutil.Glob("testdata/*"))
	assert.NotEmpty(t, fsutil.Glob("testdata/*", func(s string) bool {
		return s[0] != '.'
	}))

	var paths []string
	err := fsutil.GlobWithFunc("testdata/*", func(fpath string) error {
		paths = append(paths, fpath)
		return nil
	})

	assert.NoErr(t, err)
	assert.NotEmpty(t, paths)
}

func TestApplyFilters(t *testing.T) {
	e1 := &fakeobj.DirEntry{Nam: "some-backup"}
	f1 := fsutil.ExcludeSuffix("-backup")

	assert.False(t, f1("", e1))
	assert.True(t, fsutil.ApplyFilters("", e1, []fsutil.FilterFunc{f1}))
	assert.True(t, fsutil.ApplyFilters("", e1, []fsutil.FilterFunc{fsutil.OnlyFindDir}))
	assert.False(t, fsutil.ApplyFilters("", e1, []fsutil.FilterFunc{fsutil.OnlyFindFile}))
	assert.False(t, fsutil.ApplyFilters("", e1, []fsutil.FilterFunc{fsutil.ExcludeDotFile}))
	assert.False(t, fsutil.ApplyFilters("", e1, []fsutil.FilterFunc{fsutil.IncludeSuffix("-backup")}))
	assert.True(t, fsutil.ApplyFilters("", e1, []fsutil.FilterFunc{fsutil.ExcludeNames("some-backup")}))
}

func TestFindInDir(t *testing.T) {
	err := fsutil.FindInDir("path-not-exist", nil)
	assert.NoErr(t, err)

	err = fsutil.FindInDir("testdata/test.jpg", nil)
	assert.NoErr(t, err)

	files := make([]string, 0, 8)
	err = fsutil.FindInDir("testdata", func(fPath string, de fs.DirEntry) error {
		files = append(files, fPath)
		return nil
	})

	dump.P(files)
	assert.NoErr(t, err)
	assert.True(t, len(files) > 0)

	files = files[:0]
	err = fsutil.FindInDir("testdata", func(fPath string, de fs.DirEntry) error {
		files = append(files, fPath)
		return nil
	}, func(fPath string, de fs.DirEntry) bool {
		return !strings.HasPrefix(de.Name(), ".")
	})
	assert.NoErr(t, err)
	assert.True(t, len(files) > 0)

	err = fsutil.FindInDir("testdata", func(fPath string, de fs.DirEntry) error {
		return errorx.Raw("handle error")
	})
	assert.Err(t, err)
}
