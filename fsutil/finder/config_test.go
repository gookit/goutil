package finder_test

import (
	"testing"

	"github.com/gookit/goutil/fsutil/finder"
	"github.com/gookit/goutil/testutil/assert"
)

func TestBasic_func(t *testing.T) {
	assert.Eq(t, finder.FlagBoth, finder.ToFlag("b"))
	assert.Eq(t, finder.FlagBoth, finder.ToFlag("both"))
	assert.Eq(t, finder.FlagFile, finder.ToFlag(""))
	assert.Eq(t, finder.FlagFile, finder.ToFlag("file"))
	assert.Eq(t, finder.FlagDir, finder.ToFlag("dir"))
}

func TestFinder_Config(t *testing.T) {
	ff := finder.NewEmpty().ConfigFn(func(c *finder.Config) {
		c.IncludeNames = []string{"name1"}
	})

	ff.AddScan("/some/path/dir").
		FileAndDir().
		WithExts([]string{".go"}).
		WithFileExt(".mod").
		WithoutFile(".keep").
		CacheResult(false)

	assert.NotEmpty(t, ff.Config())
}

func TestConfig_NewFinder(t *testing.T) {
	c := finder.NewConfig()
	c.IncludePaths = []string{"path/to"}
	c.ExcludePaths = []string{"exclude/path"}
	c.ExcludeDirs = []string{"dir1"}
	c.IncludeExts = []string{".go"}
	c.ExcludeFiles = []string{".keep"}

	ff := c.NewFinder()
	ff.UseAbsPath(true).
		CacheResult(true).
		WithoutDotFile(true)
	ff.IncludeExt(".txt")

	assert.NotEmpty(t, ff.Config())
}
