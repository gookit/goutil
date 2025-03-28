// Package finder Provides a simple and convenient filedir lookup function,
// supports filtering, excluding, matching, ignoring, etc.
// and with some commonly built-in matchers.
package finder

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
)

type scanDir struct {
	path  string // dir path to scan
	depth int    // current depth
}

// FileFinder type alias.
type FileFinder = Finder

// Finder struct
type Finder struct {
	// config for finder
	c *Config
	// last error
	err error
	// num - founded fs elem number
	num uint32
	// ch - founded fs elem chan
	ch chan Elem
	// 等待组,跟踪任务完成
	wg sync.WaitGroup
	// dir queue channel, used for concurrency mode
	dirQueue chan scanDir
	// caches - cache found fs elem. if config.CacheResult is true
	caches []Elem
}

// New instance with source dir paths.
func New(dirs []string) *Finder {
	return NewWithConfig(NewConfig(dirs...))
}

// NewFinder new instance with source dir paths.
func NewFinder(dirPaths ...string) *Finder { return New(dirPaths) }

// NewWithConfig new instance with config.
func NewWithConfig(c *Config) *Finder { return &Finder{c: c} }

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
//	f := NewFinder("/path/to/dir")
//	for el := range f.Find() {
//		fmt.Println(el.Path())
//	}
func (f *Finder) Find() <-chan Elem { return f.find() }

// Elems find and return founded file Elem. alias of Find()
func (f *Finder) Elems() <-chan Elem { return f.find() }

// Results find and return founded file Elem. alias of Find()
func (f *Finder) Results() <-chan Elem { return f.find() }

// FindNames find and return founded file/dir names.
func (f *Finder) FindNames() []string {
	paths := make([]string, 0, 8*len(f.c.ScanDirs))
	for el := range f.find() {
		paths = append(paths, el.Name())
	}
	return paths
}

// FindPaths find and return founded file/dir paths.
func (f *Finder) FindPaths() []string {
	paths := make([]string, 0, 8*len(f.c.ScanDirs))
	for el := range f.find() {
		paths = append(paths, el.Path())
	}
	return paths
}

// Each founded file or dir Elem.
func (f *Finder) Each(fn func(el Elem)) { f.EachElem(fn) }

// EachElem founded file or dir Elem.
func (f *Finder) EachElem(fn func(el Elem)) {
	for el := range f.find() {
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
			f.setError(err)
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
			f.setError(err)
		}
	})
}

// prepare for find.
func (f *Finder) prepare() {
	f.err = nil
	if f.CacheNum() == 0 {
		f.num = 0
	}

	// ensure config
	if f.c == nil {
		f.c = NewConfig()
	} else {
		f.c.Init()
	}

	coNum := f.c.Concurrency
	f.debugf("PREPARE done. type-flag: %s, concurrency: %d", f.c.FindFlags, coNum)
	f.debugf("config: %+v", f.c)
	// 创建队列
	f.ch = make(chan Elem, coNum*8)
	f.dirQueue = make(chan scanDir, coNum*8)
}

// Do finding
//
// Usage:
//
//	for el := range f.find() {
//		fmt.Println(el.Path())
//	}
func (f *Finder) find() <-chan Elem {
	// has caches, return it
	if len(f.caches) > 0 {
		f.ch = make(chan Elem, 8)
		defer close(f.ch)
		for _, el := range f.caches {
			f.ch <- el
		}
		return f.ch
	}

	f.prepare()

	// 启动工作goroutine
	for i := 0; i < f.c.Concurrency; i++ {
		go f.worker(i)
	}

	// 添加初始任务
	f.addRootDirs()

	// 等待所有任务完成并关闭通道
	go func() {
		f.debugf("waiting all task complete ...")
		f.wg.Wait()

		close(f.ch)
		close(f.dirQueue)
		f.debugf("all find task DONE. total found: %d", f.num)

		// reset wg
		// f.wg = sync.WaitGroup{}
	}()

	f.debugf("find task STARTING ...")
	return f.ch
}

// worker 处理目录的工作goroutine
func (f *Finder) worker(index int) {
	for sd := range f.dirQueue {
		func() {
			defer f.wg.Done()
			f.debugf("worker#%d into dir: %s (depth: %d)", index, sd.path, sd.depth)
			f.findDir(sd.path, sd.depth)
		}()
	}
}

func (f *Finder) addRootDirs() {
	f.debugf("add scan root dirs: %v", f.c.ScanDirs)

	var err error
	for _, dirPath := range f.c.ScanDirs {
		if f.c.UseAbsPath {
			dirPath, err = filepath.Abs(dirPath)
			if err != nil {
				f.setError(err)
				continue
			}
		}

		// add task
		f.debugf("add root-dir: %s", dirPath)
		f.dirQueue <- scanDir{path: dirPath}
	}
}

// code refer filepath.glob()
func (f *Finder) findDir(dirPath string, depth int) {
	deList, err := os.ReadDir(dirPath)
	if err != nil {
		return // ignore I/O error
	}

	cfg := f.c
	depth++
	var ok bool

	for _, ent := range deList {
		name := ent.Name()
		isDir := ent.IsDir()
		if name[0] == '.' {
			if isDir {
				if cfg.ExcludeDotDir {
					continue
				}
			} else if cfg.ExcludeDotFile {
				continue
			}
		}

		fullPath := filepath.Join(dirPath, name)
		el := NewElem(fullPath, ent)

		// apply generic filters
		if !applyExMatchers(el, cfg.ExMatchers) {
			continue
		}

		// --- dir: apply dir filters
		if isDir {
			if !applyExMatchers(el, cfg.DirExMatchers) {
				continue
			}

			if len(cfg.Matchers) > 0 {
				ok = applyMatchers(el, cfg.Matchers)
				if !ok && len(cfg.DirMatchers) > 0 {
					ok = applyMatchers(el, cfg.DirMatchers)
				}
			} else {
				ok = applyMatchers(el, cfg.DirMatchers)
			}

			// match ok, send to consumer
			if ok && cfg.FindFlags&FlagDir > 0 {
				if cfg.CacheResult {
					f.caches = append(f.caches, el)
				}
				f.ch <- el
				atomic.AddUint32(&f.num, 1)

				// if cfg.FindFlags == FlagDir {
				// 	continue // only find sub-dir on ok=false
				// }
			}

			// find in sub dir. 添加子目录任务
			if cfg.MaxDepth == 0 || depth < cfg.MaxDepth {
				f.debugf("add sub-dir: %s (depth: %d)", fullPath, depth)
				f.dirQueue <- scanDir{path: fullPath, depth: depth}
			}
			continue
		}

		// --- type: file
		if cfg.FindFlags&FlagFile == 0 {
			continue
		}

		// apply file filters
		if !applyExMatchers(el, cfg.FileExMatchers) {
			continue
		}

		if len(cfg.Matchers) > 0 {
			ok = applyMatchers(el, cfg.Matchers)
			if !ok && len(cfg.FileMatchers) > 0 {
				ok = applyMatchers(el, cfg.FileMatchers)
			}
		} else {
			ok = applyMatchers(el, cfg.FileMatchers)
		}

		// write to consumer
		if ok && cfg.FindFlags&FlagFile > 0 {
			if cfg.CacheResult {
				f.caches = append(f.caches, el)
			}
			f.ch <- el
			atomic.AddUint32(&f.num, 1)
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
func (f *Finder) Num() uint { return uint(f.num) }

// Err get last error
func (f *Finder) Err() error { return f.err }

// Caches get cached results. only valid after finding.
func (f *Finder) Caches() []Elem { return f.caches }

// CacheNum get
func (f *Finder) CacheNum() int { return len(f.caches) }

// Config get, NOTE: it's a copy of config.
func (f *Finder) Config() Config { return *f.c }

// String all dir paths
func (f *Finder) String() string {
	return strings.Join(f.c.ScanDirs, ";")
}

func (f *Finder) debugf(tpl string, vs ...any) {
	if f.c.DebugMode {
		fmt.Printf("Finder: "+tpl+"\n", vs...)
	}
}

func (f *Finder) setError(err error) {
	if err != nil {
		f.err = err
		f.debugf("ERROR=%v", err)
	}
}
