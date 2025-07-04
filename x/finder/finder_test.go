package finder_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/x/finder"
)

func TestMain(m *testing.M) {
	_, _ = fsutil.PutContents("./testdata/test.txt", "hello, in test.txt")
	m.Run()
}

func elemChanToStrings(ch <-chan finder.Elem) (list []string) {
	for elem := range ch {
		list = append(list, elem.Path())
	}
	return
}

func TestFinder_find_file(t *testing.T) {
	f := finder.EmptyFinder().
		WithDebug().
		ScanDir(".").
		// ScanDir("C:\\Users\\inhere\\workspace\\godev\\gookit\\goutil").
		NoDotFile().
		NoDotDir().
		ExcludeName("*_test.go").
		WithoutExt(".jpg")

	// find paths
	assert.NotEmpty(t, f.FindNames())
	assert.Eq(t, 0, f.CacheNum())

	assert.NotEmpty(t, elemChanToStrings(f.Elems()))
	assert.NotEmpty(t, elemChanToStrings(f.Results()))
}

func TestFinder_find_file_withCache(t *testing.T) {
	f := finder.EmptyFinder().
		TypeFile().
		WithDebug().
		ScanDir("./testdata").
		NoDotFile().
		NoDotDir().
		WithoutExt(".jpg").
		CacheResult()

	assert.Nil(t, f.Err())
	assert.NotEmpty(t, f.String())
	assert.NotEmpty(t, f.Config())
	assert.Eq(t, 0, f.CacheNum())

	// find paths
	assert.NotEmpty(t, f.FindPaths())
	assert.Gt(t, f.CacheNum(), 0)
	assert.NotEmpty(t, f.Caches())

	f.Each(func(elem finder.Elem) {
		fmt.Println(elem)
	})

	assert.NotEmpty(t, elemChanToStrings(f.Find()))
	assert.NotEmpty(t, elemChanToStrings(f.Elems()))
	assert.NotEmpty(t, elemChanToStrings(f.Results()))

	t.Run("each elem", func(t *testing.T) {
		f.EachElem(func(elem finder.Elem) {
			fmt.Println(elem)
		})
	})

	t.Run("each file", func(t *testing.T) {
		f.EachFile(func(file *os.File) {
			fmt.Println(file.Name())
		})
	})

	t.Run("each path", func(t *testing.T) {
		f.EachPath(func(filePath string) {
			fmt.Println(filePath)
		})
	})

	t.Run("each stat", func(t *testing.T) {
		f.EachStat(func(fi os.FileInfo, filePath string) {
			fmt.Println(filePath, "=>", fi.ModTime())
		})
	})

	t.Run("reset", func(t *testing.T) {
		f.Reset()
		assert.Empty(t, f.Caches())
		assert.NotEmpty(t, f.FindPaths())

		f.EachElem(func(elem finder.Elem) {
			fmt.Println(elem)
		})
	})
}

func TestFinder_OnlyFindDir(t *testing.T) {
	// ff := finder.NewFinder("./../").
	ff := finder.NewFinder("./../../").
		// ff := finder.NewFinder("C:\\Users\\inhere\\workspace\\godev\\gookit").
		// ff := finder.EmptyFinder().ScanDir("./../").
		WithConcurrency(2).
		OnlyFindDir().
		UseAbsPath().
		WithoutDotDir().
		WithDirName("testdata").
		WithDebug()

	fmt.Println(ff.FindPaths())
	// return

	t.Run("EachPath", func(t *testing.T) {
		ff.EachPath(func(filePath string) {
			fmt.Println(filePath)
		})
		assert.Gt(t, ff.Num(), 0)
		assert.Eq(t, 0, ff.CacheNum())
	})

	t.Run("Each", func(t *testing.T) {
		ff.Each(func(el finder.Elem) {
			fmt.Println(el)
		})
	})

	ff.ResetResult()
	assert.Eq(t, uint(0), ff.Num())
	assert.Eq(t, 0, ff.CacheNum())

	t.Run("max depth", func(t *testing.T) {
		ff.WithMaxDepth(2)
		ff.EachPath(func(filePath string) {
			fmt.Println(filePath)
		})
		assert.Gt(t, ff.Num(), 0)
	})
}

func TestFileFinder_NoDotFile(t *testing.T) {
	f := finder.NewEmpty().
		CacheResult().
		ScanDir("./testdata")
	assert.NotEmpty(t, f.String())

	fileName := ".env"
	assert.NotEmpty(t, f.FindPaths())
	assert.Contains(t, f.FindNames(), fileName)

	f = finder.EmptyFinder().
		ScanDir("./testdata").
		NoDotFile()
	assert.NotContains(t, f.FindNames(), fileName)

	t.Run("Not MatchDotFile", func(t *testing.T) {
		f = finder.EmptyFinder().
			ScanDir("./testdata").
			Not(finder.MatchDotFile())

		assert.NotContains(t, f.FindNames(), fileName)
	})
}

func TestFileFinder_IncludeName(t *testing.T) {
	f := finder.NewFinder(".").
		IncludeName("elem.go").
		WithNames([]string{"not-exist.file"})

	names := f.FindNames()
	assert.Len(t, names, 1)
	assert.Contains(t, names, "elem.go")
	assert.NotContains(t, names, "not-exist.file")

	f.Reset()
	t.Run("name in subdir", func(t *testing.T) {
		f.WithFileName("test.jpg")
		names = f.FindNames()
		assert.Len(t, names, 1)
		assert.Contains(t, names, "test.jpg")
	})
}

func TestFileFinder_ExcludeName(t *testing.T) {
	f := finder.NewEmpty().
		AddScanDir(".").
		WithMaxDepth(1).
		ExcludeName("elem.go").
		WithoutNames([]string{"config.go"})
	f.Exclude(finder.MatchSuffix("_test.go"), finder.MatchExt(".md"))

	names := f.FindNames()
	fmt.Println(names)
	assert.Contains(t, names, "matcher.go")
	assert.NotContains(t, names, "elem.go")
}

func TestFileFinder_WithConcurrency(t *testing.T) {
	// test with concurrency
	f := finder.NewFinder("..").
		WithDebug().
		WithConcurrency(2).
		ExcludeExt(".md").
		ExcludeName("*_test.go")

	for i, s := range f.FindPaths() {
		fmt.Println(i+1, s)
	}
}
