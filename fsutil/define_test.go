package fsutil_test

import (
	"fmt"
	"sort"
	"testing"
	"time"

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

func TestFileInfos_Len(t *testing.T) {
	fi1 := fakeobj.NewFile("/path/to/file1.txt")
	fi1.Mt = time.Now()

	fi2 := fakeobj.NewFile("/path/to/file2.txt")
	fi2.Mt = time.Now().Add(-time.Minute)

	fis := fsutil.FileInfos{
		fsutil.NewFileInfo(fi1.Path, fi1),
		fsutil.NewFileInfo(fi2.Path, fi2),
	}

	assert.Equal(t, 2, fis.Len())
	assert.Eq(t, fi1.Path, fis[0].Path())

	sort.Sort(fis)
	assert.Eq(t, fi2.Path, fis[0].Path())
}
