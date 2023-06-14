// Package finder Provides a simple and convenient filedir lookup function,
// supports filtering, excluding, matching, ignoring, etc.
// and with some commonly built-in matchers.
package finder

import (
	"os"
	"path/filepath"
	"strings"
)

// FileFinder type alias.
type FileFinder = Finder

// Finder struct
type Finder struct {
	// config for finder
	c *Config
	// last error
	err error
	// num - founded fs elem number
	num int
	// ch - founded fs elem chan
	ch chan Elem
	// caches - cache found fs elem. if config.CacheResult is true
	caches []Elem
}

// New instance with source dir paths.
func New(dirs []string) *Finder {
	c := NewConfig(dirs...)
	return NewWithConfig(c)
}

// NewFinder new instance with source dir paths.
func NewFinder(dirPaths ...string) *Finder { return New(dirPaths) }

// NewWithConfig new instance with config.
func NewWithConfig(c *Config) *Finder {
	return &Finder{c: c}
}

// NewEmpty new empty Finder instance
func NewEmpty() *Finder {
	return &Finder{c: NewEmptyConfig()}
}

// EmptyFinder new empty Finder instance. alias of NewEmpty()
func EmptyFinder() *Finder { return NewEmpty() }

//
// --------- do finding ---------
//

// Find files in given dir paths. will return a channel, you can use it to get the result.
//
// Usage:
//
//	f := NewFinder("/path/to/dir").Find()
//	for el := range f {
//		fmt.Println(el.Path())
//	}
func (f *Finder) Find() <-chan Elem {
	f.find()
	return f.ch
}

// Elems find and return founded file Elem. alias of Find()
func (f *Finder) Elems() <-chan Elem { return f.Find() }

// Results find and return founded file Elem. alias of Find()
func (f *Finder) Results() <-chan Elem { return f.Find() }

// FindNames find and return founded file/dir names.
func (f *Finder) FindNames() []string {
	paths := make([]string, 0, 8*len(f.c.ScanDirs))
	for el := range f.Find() {
		paths = append(paths, el.Name())
	}
	return paths
}

// FindPaths find and return founded file/dir paths.
func (f *Finder) FindPaths() []string {
	paths := make([]string, 0, 8*len(f.c.ScanDirs))
	for el := range f.Find() {
		paths = append(paths, el.Path())
	}
	return paths
}

// Each founded file or dir Elem.
func (f *Finder) Each(fn func(el Elem)) { f.EachElem(fn) }

// EachElem founded file or dir Elem.
func (f *Finder) EachElem(fn func(el Elem)) {
	f.find()
	for el := range f.ch {
		fn(el)
	}
}

// EachPath founded file paths.
func (f *Finder) EachPath(fn func(filePath string)) {
	f.EachElem(func(el Elem) {
		fn(el.Path())
	})
}

// EachFile each file os.File
func (f *Finder) EachFile(fn func(file *os.File)) {
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
func (f *Finder) EachStat(fn func(fi os.FileInfo, filePath string)) {
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
func (f *Finder) EachContents(fn func(contents, filePath string)) {
	f.EachElem(func(el Elem) {
		bs, err := os.ReadFile(el.Path())
		if err == nil {
			fn(string(bs), el.Path())
		} else {
			f.err = err
		}
	})
}

// prepare for find.
func (f *Finder) prepare() {
	f.err = nil
	f.ch = make(chan Elem, 8)

	if f.CacheNum() == 0 {
		f.num = 0
	}

	if f.c == nil {
		f.c = NewConfig()
	} else {
		f.c.Init()
	}
}

// do finding
func (f *Finder) find() {
	f.prepare()

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
		var err error
		for _, dirPath := range f.c.ScanDirs {
			if f.c.UseAbsPath {
				dirPath, err = filepath.Abs(dirPath)
				if err != nil {
					f.err = err
					continue
				}
			}

			f.c.depth = 0
			f.findDir(dirPath, f.c)
		}
	}()
}

// code refer filepath.glob()
func (f *Finder) findDir(dirPath string, c *Config) {
	des, err := os.ReadDir(dirPath)
	if err != nil {
		return // ignore I/O error
	}

	var ok bool
	c.depth++
	for _, ent := range des {
		name := ent.Name()
		isDir := ent.IsDir()
		if name[0] == '.' {
			if isDir {
				if c.ExcludeDotDir {
					continue
				}
			} else if c.ExcludeDotFile {
				continue
			}
		}

		fullPath := filepath.Join(dirPath, name)
		el := NewElem(fullPath, ent)

		// apply generic filters
		if !applyExMatchers(el, c.ExMatchers) {
			continue
		}

		// --- dir: apply dir filters
		if isDir {
			if !applyExMatchers(el, c.DirExMatchers) {
				continue
			}

			if len(c.Matchers) > 0 {
				ok = applyMatchers(el, c.Matchers)
				if !ok && len(c.DirMatchers) > 0 {
					ok = applyMatchers(el, c.DirMatchers)
				}
			} else {
				ok = applyMatchers(el, c.DirMatchers)
			}

			if ok && c.FindFlags&FlagDir > 0 {
				if c.CacheResult {
					f.caches = append(f.caches, el)
				}
				f.num++
				f.ch <- el

				if c.FindFlags == FlagDir {
					continue // only find subdir on ok=false
				}
			}

			// find in sub dir.
			if c.MaxDepth == 0 || c.depth < c.MaxDepth {
				f.findDir(fullPath, c)
				c.depth-- // restore depth
			}
			continue
		}

		// --- type: file
		if c.FindFlags&FlagFile == 0 {
			continue
		}

		// apply file filters
		if !applyExMatchers(el, c.FileExMatchers) {
			continue
		}

		if len(c.Matchers) > 0 {
			ok = applyMatchers(el, c.Matchers)
			if !ok && len(c.FileMatchers) > 0 {
				ok = applyMatchers(el, c.FileMatchers)
			}
		} else {
			ok = applyMatchers(el, c.FileMatchers)
		}

		// write to consumer
		if ok && c.FindFlags&FlagFile > 0 {
			if c.CacheResult {
				f.caches = append(f.caches, el)
			}
			f.num++
			f.ch <- el
		}
	}
}

func applyMatchers(el Elem, fls []Matcher) bool {
	for _, f := range fls {
		if f.Apply(el) {
			return true
		}
	}
	return len(fls) == 0
}

func applyExMatchers(el Elem, fls []Matcher) bool {
	for _, f := range fls {
		if f.Apply(el) {
			return false
		}
	}
	return true
}

// Reset filters config setting and results info.
func (f *Finder) Reset() {
	c := NewConfig(f.c.ScanDirs...)
	c.ExcludeDotDir = f.c.ExcludeDotDir
	c.FindFlags = f.c.FindFlags
	c.MaxDepth = f.c.MaxDepth

	f.c = c
	f.ResetResult()
}

// ResetResult reset result info.
func (f *Finder) ResetResult() {
	f.num = 0
	f.err = nil
	f.ch = make(chan Elem, 8)
	f.caches = []Elem{}
}

// Num get found elem num. only valid after finding.
func (f *Finder) Num() int {
	return f.num
}

// Err get last error
func (f *Finder) Err() error {
	return f.err
}

// Caches get cached results. only valid after finding.
func (f *Finder) Caches() []Elem {
	return f.caches
}

// CacheNum get
func (f *Finder) CacheNum() int {
	return len(f.caches)
}

// Config get
func (f *Finder) Config() Config {
	return *f.c
}

// String all dir paths
func (f *Finder) String() string {
	return strings.Join(f.c.ScanDirs, ";")
}
