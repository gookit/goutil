// Package finder provide a finder tool for find files
package finder

import (
	"os"
	"path/filepath"
	"strings"
)

// FileFinder struct
type FileFinder struct {
	// config for finder
	c *Config
	// last error
	err error
	// ch - founded file elem chan
	ch chan Elem
	// caches - cache found file elem. if config.CacheResult is true
	caches []Elem
}

// NewEmpty new empty FileFinder instance
func NewEmpty() *FileFinder { return New([]string{}) }

// EmptyFinder new empty FileFinder instance
func EmptyFinder() *FileFinder { return NewEmpty() }

// New instance with source dir paths.
func New(dirs []string, fls ...Filter) *FileFinder {
	c := NewConfig(dirs...)
	c.Filters = fls

	return NewWithConfig(c)
}

// NewFinder new instance with source dir paths.
func NewFinder(dirPaths ...string) *FileFinder {
	return New(dirPaths)
}

// NewWithConfig new instance with config.
func NewWithConfig(c *Config) *FileFinder {
	return &FileFinder{
		c: c,
	}
}

// WithConfig on the finder
func (f *FileFinder) WithConfig(c *Config) *FileFinder {
	f.c = c
	return f
}

// ConfigFn the finder
func (f *FileFinder) ConfigFn(fns ...func(c *Config)) *FileFinder {
	if f.c == nil {
		f.c = &Config{}
	}

	for _, fn := range fns {
		fn(f.c)
	}
	return f
}

// AddDirPath add source dir for find
func (f *FileFinder) AddDirPath(dirPaths ...string) *FileFinder {
	f.c.DirPaths = append(f.c.DirPaths, dirPaths...)
	return f
}

// AddDir add source dir for find. alias of AddDirPath()
func (f *FileFinder) AddDir(dirPaths ...string) *FileFinder {
	f.c.DirPaths = append(f.c.DirPaths, dirPaths...)
	return f
}

// CacheResult cache result for find result.
func (f *FileFinder) CacheResult(enable ...bool) *FileFinder {
	if len(enable) > 0 {
		f.c.CacheResult = enable[0]
	} else {
		f.c.CacheResult = true
	}
	return f
}

// ExcludeDotDir exclude dot dir names. eg: ".idea"
func (f *FileFinder) ExcludeDotDir(exclude ...bool) *FileFinder {
	if len(exclude) > 0 {
		f.c.ExcludeDotDir = exclude[0]
	} else {
		f.c.ExcludeDotDir = true
	}
	return f
}

// WithoutDotDir exclude dot dir names. alias of ExcludeDotDir().
func (f *FileFinder) WithoutDotDir(exclude ...bool) *FileFinder {
	return f.ExcludeDotDir(exclude...)
}

// NoDotDir exclude dot dir names. alias of ExcludeDotDir().
func (f *FileFinder) NoDotDir(exclude ...bool) *FileFinder {
	return f.ExcludeDotDir(exclude...)
}

// ExcludeDotFile exclude dot dir names. eg: ".gitignore"
func (f *FileFinder) ExcludeDotFile(exclude ...bool) *FileFinder {
	if len(exclude) > 0 {
		f.c.ExcludeDotFile = exclude[0]
	} else {
		f.c.ExcludeDotFile = true
	}
	return f
}

// WithoutDotFile exclude dot dir names. alias of ExcludeDotFile().
func (f *FileFinder) WithoutDotFile(exclude ...bool) *FileFinder {
	return f.ExcludeDotFile(exclude...)
}

// NoDotFile exclude dot dir names. alias of ExcludeDotFile().
func (f *FileFinder) NoDotFile(exclude ...bool) *FileFinder {
	return f.ExcludeDotFile(exclude...)
}

// ExcludeDir exclude dir names.
func (f *FileFinder) ExcludeDir(dirs ...string) *FileFinder {
	f.c.ExcludeDirs = append(f.c.ExcludeDirs, dirs...)
	return f
}

// ExcludeName exclude file names.
func (f *FileFinder) ExcludeName(files ...string) *FileFinder {
	f.c.ExcludeNames = append(f.c.ExcludeNames, files...)
	return f
}

// AddFilter for filter filepath or dir path
func (f *FileFinder) AddFilter(filters ...Filter) *FileFinder {
	return f.WithFilter(filters...)
}

// WithFilter add filter func for filtering filepath or dir path
func (f *FileFinder) WithFilter(filters ...Filter) *FileFinder {
	f.c.Filters = append(f.c.Filters, filters...)
	return f
}

// AddFileFilter for filter filepath
func (f *FileFinder) AddFileFilter(filters ...Filter) *FileFinder {
	return f.WithFileFilter(filters...)
}

// WithFileFilter for filter func for filtering filepath
func (f *FileFinder) WithFileFilter(filters ...Filter) *FileFinder {
	f.c.FileFilters = append(f.c.FileFilters, filters...)
	return f
}

// AddDirFilter for filter file contents
func (f *FileFinder) AddDirFilter(fls ...Filter) *FileFinder {
	return f.WithDirFilter(fls...)
}

// WithDirFilter for filter func for filtering file contents
func (f *FileFinder) WithDirFilter(filters ...Filter) *FileFinder {
	f.c.DirFilters = append(f.c.DirFilters, filters...)
	return f
}

// AddBodyFilter for filter file contents
func (f *FileFinder) AddBodyFilter(fls ...BodyFilter) *FileFinder {
	return f.WithBodyFilter(fls...)
}

// WithBodyFilter for filter func for filtering file contents
func (f *FileFinder) WithBodyFilter(fls ...BodyFilter) *FileFinder {
	f.c.BodyFilters = append(f.c.BodyFilters, fls...)
	return f
}

// Find files in given dir paths. will return a channel, you can use it to get the result.
//
// Usage:
//
//	f := NewFinder("/path/to/dir").Find()
//	for el := range f {
//		fmt.Println(el.Path())
//	}
func (f *FileFinder) Find() <-chan Elem {
	f.find()
	return f.ch
}

// Results find and return founded file Elem. alias of Find()
//
// Usage:
//
//	rs := NewFinder("/path/to/dir").Results()
//	for el := range rs {
//		fmt.Println(el.Path())
//	}
func (f *FileFinder) Results() <-chan Elem {
	f.find()
	return f.ch
}

// FindPaths find and return founded file paths.
func (f *FileFinder) FindPaths() []string {
	f.find()

	paths := make([]string, 0, 8*len(f.c.DirPaths))
	for el := range f.ch {
		paths = append(paths, el.Path())
	}
	return paths
}

// Each file or dir Elem.
func (f *FileFinder) Each(fn func(el Elem)) {
	f.EachElem(fn)
}

// EachElem file or dir Elem.
func (f *FileFinder) EachElem(fn func(el Elem)) {
	f.find()
	for el := range f.ch {
		fn(el)
	}
}

// EachPath file paths.
func (f *FileFinder) EachPath(fn func(filePath string)) {
	f.EachElem(func(el Elem) {
		fn(el.Path())
	})
}

// EachFile each file os.File
func (f *FileFinder) EachFile(fn func(file *os.File)) {
	f.EachElem(func(el Elem) {
		file, err := os.Open(el.Path())
		if err == nil {
			fn(file)
		} else {
			f.err = err
		}
	})
}

// EachStat each file os.FileInfo
func (f *FileFinder) EachStat(fn func(fi os.FileInfo, filePath string)) {
	f.EachElem(func(el Elem) {
		fi, err := el.Info()
		if err == nil {
			fn(fi, el.Path())
		} else {
			f.err = err
		}
	})
}

// EachContents handle each found file contents
func (f *FileFinder) EachContents(fn func(contents, filePath string)) {
	f.EachElem(func(el Elem) {
		bs, err := os.ReadFile(el.Path())
		if err == nil {
			fn(string(bs), el.Path())
		} else {
			f.err = err
		}
	})
}

// do finding
func (f *FileFinder) find() {
	f.err = nil
	f.ch = make(chan Elem, 8)

	if f.c == nil {
		f.c = NewConfig()
	}

	go func() {
		defer close(f.ch)

		// read from caches
		if f.c.CacheResult && len(f.caches) > 0 {
			for _, el := range f.caches {
				f.ch <- el
			}
			return
		}

		// do finding
		for _, dirPath := range f.c.DirPaths {
			f.findDir(dirPath, f.c)
		}
	}()
}

// code refer filepath.glob()
func (f *FileFinder) findDir(dirPath string, c *Config) {
	dfi, err := os.Stat(dirPath)
	if err != nil {
		return // ignore I/O error
	}
	if !dfi.IsDir() {
		return // ignore I/O error
	}

	des, err := os.ReadDir(dirPath)
	if err != nil {
		return // ignore I/O error
	}

	c.curDepth++
	for _, ent := range des {
		baseName := ent.Name()
		fullPath := filepath.Join(dirPath, baseName)

		ok := false
		el := NewElem(fullPath, ent)

		// apply generic filters
		for _, filter := range c.Filters {
			if filter.Apply(el) { // 有一个满足即可
				ok = true
				break
			}
		}
		if !ok {
			continue
		}

		// --- dir
		if ent.IsDir() {
			if c.ExcludeDotDir && baseName[0] == '.' {
				continue
			}

			// apply dir filters
			ok = false
			for _, df := range c.DirFilters {
				if df.Apply(el) { // 有一个满足即可
					ok = true
					break
				}
			}

			if ok {
				if c.FindFlags&FlagDir > 0 {
					if c.CacheResult {
						f.caches = append(f.caches, el)
					}
					f.ch <- el
				}

				// find in sub dir.
				if c.curDepth < c.MaxDepth {
					f.findDir(fullPath, c)
				}
			}
			continue
		}

		if c.FindFlags&FlagDir > 0 {
			continue
		}

		// --- type: file
		if c.ExcludeDotFile && baseName[0] == '.' {
			continue
		}

		// use custom filter functions
		ok = false
		for _, ff := range c.FileFilters {
			if ff.Apply(el) { // 有一个满足即可
				ok = true
				break
			}
		}

		// write to consumer
		if ok && c.FindFlags&FlagFile > 0 {
			if c.CacheResult {
				f.caches = append(f.caches, el)
			}
			f.ch <- el
		}
	}
}

// Reset filters config setting and results info.
func (f *FileFinder) Reset() {
	c := NewConfig(f.c.DirPaths...)
	c.ExcludeDotDir = f.c.ExcludeDotDir
	c.FindFlags = f.c.FindFlags
	c.MaxDepth = f.c.MaxDepth
	c.curDepth = 0

	f.c = c
	f.err = nil
	f.ch = make(chan Elem, 8)
	f.caches = []Elem{}
}

// Err get last error
func (f *FileFinder) Err() error {
	return f.err
}

// CacheNum get
func (f *FileFinder) CacheNum() int {
	return len(f.caches)
}

// Config get
func (f *FileFinder) Config() Config {
	return *f.c
}

// String all dir paths
func (f *FileFinder) String() string {
	return strings.Join(f.c.DirPaths, ",")
}
