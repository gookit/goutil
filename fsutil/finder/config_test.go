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

// TestFinder_IncludeRule tests the IncludeRule method of Finder
func TestFinder_IncludeRule(t *testing.T) {
	f := finder.NewFinder()
	f.IncludeRule("ext:.go", "name:*_test.go")

	assert.Len(t, f.Config().IncludeExts, 1)
	assert.Len(t, f.Config().IncludeNames, 1)
	assert.Eq(t, ".go", f.Config().IncludeExts[0])
	assert.Eq(t, "*_test.go", f.Config().IncludeNames[0])
}

func TestFinder_IncludeRules(t *testing.T) {
	f := finder.NewFinder()
	f.IncludeRules([]string{"ext:.go", "name:*_test.go"})

	assert.Len(t, f.Config().IncludeExts, 1)
	assert.Len(t, f.Config().IncludeNames, 1)
	assert.Eq(t, ".go", f.Config().IncludeExts[0])
	assert.Eq(t, "*_test.go", f.Config().IncludeNames[0])
}

func TestFinder_ExcludeRule(t *testing.T) {
	f := finder.NewFinder()
	f.ExcludeRule("ext:.go", "name:*_test.go")

	assert.Len(t, f.Config().ExcludeExts, 1)
	assert.Len(t, f.Config().ExcludeNames, 1)
	assert.Eq(t, ".go", f.Config().ExcludeExts[0])
	assert.Eq(t, "*_test.go", f.Config().ExcludeNames[0])
}

func TestFinder_ExcludeRules(t *testing.T) {
	f := finder.NewFinder()
	f.ExcludeRules([]string{"ext:.go", "name:*_test.go"})

	assert.Len(t, f.Config().ExcludeExts, 1)
	assert.Len(t, f.Config().ExcludeNames, 1)
	assert.Eq(t, ".go", f.Config().ExcludeExts[0])
	assert.Eq(t, "*_test.go", f.Config().ExcludeNames[0])
}

// TestConfig_LoadRules_Add tests the LoadRules method of Config with addOrNot set to true
func TestConfig_LoadRules(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		c := &finder.Config{}
		err := c.LoadRules(true, []string{"ext:.go", "name:*_test.go"})

		assert.Nil(t, err)
		assert.Len(t, c.IncludeExts, 1)
		assert.Len(t, c.IncludeNames, 1)
		assert.Eq(t, ".go", c.IncludeExts[0])
		assert.Eq(t, "*_test.go", c.IncludeNames[0])
	})

	// tests the LoadRules method of Config with addOrNot set to false
	t.Run("Exclude", func(t *testing.T) {
		c := &finder.Config{}
		err := c.LoadRules(false, []string{"ext:.go", "name:*_test.go"})

		assert.Nil(t, err)
		assert.Len(t, c.ExcludeExts, 1)
		assert.Len(t, c.ExcludeNames, 1)
		assert.Eq(t, ".go", c.ExcludeExts[0])
		assert.Eq(t, "*_test.go", c.ExcludeNames[0])
	})

	// tests the LoadRules method of Config with an invalid rule
	t.Run("InvalidRule", func(t *testing.T) {
		c := &finder.Config{}
		err := c.LoadRules(true, []string{"invalidrule"})
		assert.Err(t, err)

		err = c.LoadRules(false, []string{""})
		assert.Err(t, err)
	})

	// tests the LoadRules method of Config with multiple rules
	t.Run("MultipleRules", func(t *testing.T) {
		c := &finder.Config{}
		err := c.LoadRules(true, []string{"ext:.go,.yaml", "name:*_test.go,go.mod"})

		assert.Nil(t, err)
		assert.Len(t, c.IncludeExts, 2)
		assert.Len(t, c.IncludeNames, 2)
		assert.Eq(t, ".go", c.IncludeExts[0])
		assert.Eq(t, ".yaml", c.IncludeExts[1])
		assert.Eq(t, "*_test.go", c.IncludeNames[0])
		assert.Eq(t, "go.mod", c.IncludeNames[1])
	})

	// test file rule
	t.Run("File", func(t *testing.T) {
		c := &finder.Config{}
		err := c.LoadRules(true, []string{"file:*.go"})

		assert.Nil(t, err)
		assert.Len(t, c.FileMatchers, 1)

		// exclude
		err = c.LoadRules(false, []string{"file:*.log"})
		assert.Nil(t, err)
		assert.Len(t, c.FileExMatchers, 1)
	})

	// tests the LoadRules method of Config with size rules
	t.Run("Size", func(t *testing.T) {
		c := &finder.Config{}
		err := c.LoadRules(true, []string{"size:>=1M,<=10M"})

		assert.Nil(t, err)
		assert.Len(t, c.FileMatchers, 2)
	})

	// tests the LoadRules method of Config with time rules
	t.Run("Time", func(t *testing.T) {
		c := &finder.Config{}
		err := c.LoadRules(true, []string{"time:>=1d,<=10d"})
		assert.Nil(t, err)
		assert.Len(t, c.FileMatchers, 2)

		// exclude
		err = c.LoadRules(false, []string{"time:>=1d"})
		assert.Nil(t, err)
		assert.Len(t, c.FileExMatchers, 1)
	})

	// tests the LoadRules method of Config with empty rules
	t.Run("EmptyRules", func(t *testing.T) {
		c := &finder.Config{}
		err := c.LoadRules(true, []string{})

		assert.Nil(t, err)
		assert.Len(t, c.FileMatchers, 0)
		assert.Len(t, c.FileExMatchers, 0)
	})
}
