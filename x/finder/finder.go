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
	// caches - cache found fs elem. if config.CacheResult is true
	caches []Elem
	// mu guard the c, err and caches fields, they are changed by the worker goroutines
	mu sync.Mutex
}

// findRun holds the channels and the wait group of one find run. they can not be fields
// of the Finder: a new find can start while the goroutines of the previous run are still
// finishing their work(eg: the closer goroutine after it closed the result channel).
type findRun struct {
	wg       sync.WaitGroup
	ch       chan Elem
	dirQueue chan scanDir
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
			f.setError(err)
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
func (f *Finder) prepare() *findRun {
	f.mu.Lock()
	f.err = nil
	c := f.c
	f.mu.Unlock()

	if f.CacheNum() == 0 {
		atomic.StoreUint32(&f.num, 0)
	}

	// ensure config
	if c == nil {
		c = NewConfig()
		f.setConfig(c)
	} else {
		c.Init()
	}

	coNum := c.Concurrency
	f.debugf("PREPARE done. type-flag: %s, concurrency: %d", c.FindFlags, coNum)
	f.debugf("config: %+v", c)
	// 创建队列
	return &findRun{
		ch:       make(chan Elem, coNum*8*3),
		dirQueue: make(chan scanDir, coNum*8*2),
	}
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
	if caches := f.Caches(); len(caches) > 0 {
		ch := make(chan Elem, 8)
		for _, el := range caches {
			ch <- el
		}
		close(ch)
		return ch
	}

	run := f.prepare()

	// 添加初始任务
	f.addRootDirs(run)

	// 启动工作goroutine
	coNum := f.config().Concurrency
	for i := 0; i < coNum; i++ {
		go f.worker(run, i)
	}

	// 等待所有任务完成并关闭通道
	go func() {
		f.debugf("waiting all task complete ...")
		run.wg.Wait()

		// tip: close run.ch at the last. the consumer can start a new find on the same
		// Finder as soon as run.ch is closed, and a new run will not touch this one.
		close(run.dirQueue)
		f.debugf("all find task DONE. total found: %d", f.Num())

		close(run.ch)
	}()

	f.debugf("find task STARTING ...")
	return run.ch
}

// worker 处理目录的工作goroutine
func (f *Finder) worker(run *findRun, index int) {
	f.debugf("worker#%d STARTING ...", index)
	for sd := range run.dirQueue {
		f.safeFindDir(run, index, sd.path, sd.depth)
	}
	f.debugf("worker#%d DONE.", index)
}

func (f *Finder) safeFindDir(run *findRun, index int, dirPath string, depth int) {
	f.debugf("worker#%d into dir: %s (depth: %d)", index, dirPath, depth)

	// recover error and always call wg.Done()
	defer func() {
		if err := recover(); err != nil {
			f.debugf("worker#%d panic in dir: %s, ERROR: %v", index, dirPath, err)
			f.setError(fmt.Errorf("worker#%d findDir panic, dir: %s, ERROR: %v", index, dirPath, err))
		}
		run.wg.Done()
	}()

	f.findDir(run, dirPath, depth)
}

func (f *Finder) addRootDirs(run *findRun) {
	cfg := f.config()
	f.debugf("add scan root dirs: %v", cfg.ScanDirs)

	var err error
	for _, dirPath := range cfg.ScanDirs {
		if cfg.UseAbsPath {
			dirPath, err = filepath.Abs(dirPath)
			if err != nil {
				f.setError(err)
				continue
			}
		}

		// add task
		f.debugf("add root-dir: %s", dirPath)
		run.wg.Add(1)
		run.dirQueue <- scanDir{path: dirPath}
	}
}

// code refer filepath.glob()
func (f *Finder) findDir(run *findRun, dirPath string, depth int) {
	deList, err := os.ReadDir(dirPath)
	if err != nil {
		return // ignore I/O error
	}

	cfg := f.config()
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
					f.cacheElem(el)
				}
				run.ch <- el
				atomic.AddUint32(&f.num, 1)

				// if cfg.FindFlags == FlagDir {
				// 	continue // only find sub-dir on ok=false
				// }
			}

			// find in sub dir. 添加子目录任务
			if cfg.MaxDepth == 0 || depth < cfg.MaxDepth {
				f.debugf("add sub-dir: %s (depth: %d)", fullPath, depth)
				run.wg.Add(1)
				// fix: 创建一个 goroutine 添加子目录任务，不然会造成阻塞
				go func(p string, d int) {
					run.dirQueue <- scanDir{path: p, depth: d}
				}(fullPath, depth)
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
				f.cacheElem(el)
			}
			run.ch <- el
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
	old := f.config()
	c := NewConfig(old.ScanDirs...)
	c.ExcludeDotDir = old.ExcludeDotDir
	c.FindFlags = old.FindFlags
	c.MaxDepth = old.MaxDepth

	f.setConfig(c)
	f.ResetResult()
}

// ResetResult reset result info.
func (f *Finder) ResetResult() {
	atomic.StoreUint32(&f.num, 0)

	f.mu.Lock()
	f.err = nil
	f.caches = []Elem{}
	f.mu.Unlock()
}

// Num get found elem num. only valid after finding.
func (f *Finder) Num() uint { return uint(atomic.LoadUint32(&f.num)) }

// Err get last error
func (f *Finder) Err() error {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.err
}

// Caches get cached results. only valid after finding.
func (f *Finder) Caches() []Elem {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.caches
}

// CacheNum get the cached result number.
func (f *Finder) CacheNum() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return len(f.caches)
}

// cacheElem add a found elem to the caches. it is called by the worker goroutines.
func (f *Finder) cacheElem(el Elem) {
	f.mu.Lock()
	f.caches = append(f.caches, el)
	f.mu.Unlock()
}

// Config get, NOTE: it's a copy of config.
func (f *Finder) Config() Config { return *f.config() }

// String all dir paths
func (f *Finder) String() string {
	return strings.Join(f.config().ScanDirs, ";")
}

// config get the config pointer. it's guarded by f.mu.
func (f *Finder) config() *Config {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.c
}

// setConfig set the config pointer. it's guarded by f.mu.
func (f *Finder) setConfig(c *Config) {
	f.mu.Lock()
	f.c = c
	f.mu.Unlock()
}

func (f *Finder) debugf(tpl string, vs ...any) {
	if c := f.config(); c != nil && c.DebugMode {
		fmt.Printf("Finder: "+tpl+"\n", vs...)
	}
}

func (f *Finder) setError(err error) {
	if err != nil {
		f.mu.Lock()
		f.err = err
		f.mu.Unlock()
		f.debugf("ERROR=%v", err)
	}
}
