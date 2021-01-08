package fsutil

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

// FileFilter for filter file path.
type FileFilter interface {
	FilterFile(filePath, filename string) bool
}

// FileFilterFunc for filter file path.
type FileFilterFunc func(filePath, filename string) bool

// Filter for filter file path.
func (fn FileFilterFunc) FilterFile(filePath, filename string) bool {
	return fn(filePath, filename)
}

// DirFilter for filter dir path.
type DirFilter interface {
	FilterDir(dirPath, dirName string) bool
}

// DirFilterFunc for filter file path.
type DirFilterFunc func(dirPath, dirName string) bool

// Filter for filter file path.
func (fn DirFilterFunc) FilterDir(dirPath, dirName string) bool {
	return fn(dirPath, dirName)
}

// // BodyFilter for filter file contents.
// type BodyFilter interface {
// 	FilterBody(contents, filePath string) bool
// }
//
// // BodyFilterFunc for filter file contents.
// type BodyFilterFunc func(contents, filePath string) bool
//
// // Filter for filter file path.
// func (fn BodyFilterFunc) FilterBody(contents, filePath string) bool {
// 	return fn(contents, filePath)
// }

// FilterFunc for filter file path.
type FilterFunc func(filePath, filename string) bool

// Filter for filter file path.
func (fn FilterFunc) Filter(filePath, filename string) bool {
	return fn(filePath, filename)
}

// FileMeta struct
type FileMeta struct {
	filePath string
	filename string
}

// FindResults struct
type FindResults struct {
	f *FileFilter

	// founded file paths.
	filePaths []string

	// filters
	dirFilters  []DirFilter  // filters for filter dir paths
	fileFilters []FileFilter // filters for filter file paths
	// bodyFilters []BodyFilter // filters for filter file contents
}

func (r *FindResults) append(filePath ...string) {
	r.filePaths = append(r.filePaths, filePath...)
}

// Result get find paths
func (r *FindResults) AddFilters(filterFuncs ...FileFilter) *FindResults {
	return r
}

// Result get find paths
func (r *FindResults) Filter() *FindResults {
	return r
}

// Result get find paths
func (r *FindResults) Each() *FindResults {
	return r
}

// Result get find paths
func (r *FindResults) Result() []string {
	return r.filePaths
}

// TODO use excludeDotFlag 1 file 2 dir 1|2 both
type exDotFlag uint8

const (
	ExDotFile exDotFlag = 1
	ExDotDir  exDotFlag = 2
)

// FileFinder struct
type FileFinder struct {
	// r *FindResults

	// mark has been run find()
	founded bool
	// dir paths for find file.
	dirPaths []string
	// file paths for filter.
	srcFiles []string

	// builtin include filters
	includeDirs []string // include dir names. eg: {"model"}
	includeExts []string // include ext names. eg: {".go", ".md"}

	// builtin exclude filters
	excludeDirs  []string // exclude dir names. eg: {"test"}
	excludeExts  []string // exclude ext names. eg: {".go", ".md"}
	excludeNames []string // exclude file names. eg: {"go.mod"}

	// builtin dot filters.
	// TODO use excludeDotFlag 1 file 2 dir 1|2 both
	// excludeDotFlag exDotFlag
	excludeDotDir  bool
	excludeDotFile bool

	// fileFlags int

	dirFilters  []DirFilter  // filters for filter dir paths
	fileFilters []FileFilter // filters for filter file paths

	// founded file paths.
	filePaths []string

	// the founded file instances
	osFiles map[string]*os.File
	osInfos map[string]os.FileInfo
}

// EmptyFinder new empty FileFinder instance
func EmptyFinder() *FileFinder {
	return &FileFinder{
		osInfos: make(map[string]os.FileInfo),
	}
}

// NewFinder new instance with source dir paths.
func NewFinder(dirPaths []string, filePaths ...string) *FileFinder {
	return &FileFinder{
		dirPaths:  dirPaths,
		filePaths: filePaths,
		osInfos:   make(map[string]os.FileInfo),
	}
}

// AddDirPath add source dir for find
func (f *FileFinder) AddDirPath(dirPaths ...string) *FileFinder {
	f.dirPaths = append(f.dirPaths, dirPaths...)
	return f
}

// AddDir add source dir for find. alias of AddDirPath()
func (f *FileFinder) AddDir(dirPaths ...string) *FileFinder {
	f.dirPaths = append(f.dirPaths, dirPaths...)
	return f
}

// ExcludeDotDir exclude dot dir names. eg: ".idea"
func (f *FileFinder) ExcludeDotDir(exclude ...bool) *FileFinder {
	if len(exclude) > 0 {
		f.excludeDotDir = exclude[0]
	} else {
		f.excludeDotDir = true
	}
	return f
}

// NoDotDir exclude dot dir names. alias of ExcludeDotDir().
func (f *FileFinder) NoDotDir(exclude ...bool) *FileFinder {
	return f.ExcludeDotDir(exclude...)
}

// ExcludeDotFile exclude dot dir names. eg: ".gitignore"
func (f *FileFinder) ExcludeDotFile(exclude ...bool) *FileFinder {
	if len(exclude) > 0 {
		f.excludeDotFile = exclude[0]
	} else {
		f.excludeDotFile = true
	}
	return f
}

// NoDotFile exclude dot dir names. alias of ExcludeDotFile().
func (f *FileFinder) NoDotFile(exclude ...bool) *FileFinder {
	return f.ExcludeDotFile(exclude...)
}

// ExcludeDir exclude dir names.
func (f *FileFinder) ExcludeDir(dirs ...string) *FileFinder {
	f.excludeDirs = append(f.excludeDirs, dirs...)
	return f
}

// ExcludeName exclude file names.
func (f *FileFinder) ExcludeName(files ...string) *FileFinder {
	f.excludeNames = append(f.excludeNames, files...)
	return f
}

// AddFilter for filter filepath or dirpath
func (f *FileFinder) AddFilter(filterFuncs ...interface{}) *FileFinder {
	return f.WithFilter(filterFuncs...)
}

// WithFilter add filter func for filtering filepath or dirpath
func (f *FileFinder) WithFilter(filterFuncs ...interface{}) *FileFinder {
	for _, filterFunc := range filterFuncs {
		if fileFilter, ok := filterFunc.(FileFilter); ok {
			f.fileFilters = append(f.fileFilters, fileFilter)
		} else if dirFilter, ok := filterFunc.(DirFilter); ok {
			f.dirFilters = append(f.dirFilters, dirFilter)
		}
	}
	return f
}

// AddFileFilter for filter filepath
func (f *FileFinder) AddFileFilter(filterFuncs ...FileFilter) *FileFinder {
	f.fileFilters = append(f.fileFilters, filterFuncs...)
	return f
}

// WithFileFilter for filter func for filtering filepath
func (f *FileFinder) WithFileFilter(filterFuncs ...FileFilter) *FileFinder {
	f.fileFilters = append(f.fileFilters, filterFuncs...)
	return f
}

// AddDirFilter for filter file contents
func (f *FileFinder) AddDirFilter(filterFuncs ...DirFilter) *FileFinder {
	f.dirFilters = append(f.dirFilters, filterFuncs...)
	return f
}

// WithDirFilter for filter func for filtering file contents
func (f *FileFinder) WithDirFilter(filterFuncs ...DirFilter) *FileFinder {
	f.dirFilters = append(f.dirFilters, filterFuncs...)
	return f
}

// // AddBodyFilter for filter file contents
// func (f *FileFinder) AddBodyFilter(filterFuncs ...BodyFilter) *FileFinder {
// 	f.bodyFilters = append(f.bodyFilters, filterFuncs...)
// 	return f
// }
//
// // WithBodyFilter for filter func for filtering file contents
// func (f *FileFinder) WithBodyFilter(filterFuncs ...BodyFilter) *FileFinder {
// 	f.bodyFilters = append(f.bodyFilters, filterFuncs...)
// 	return f
// }

// AddFilePaths set founded files
func (f *FileFinder) AddFilePaths(filePaths []string) {
	f.filePaths = append(f.filePaths, filePaths...)
}

// AddFilePath add source file
func (f *FileFinder) AddFilePath(filePaths ...string) *FileFinder {
	f.filePaths = append(f.filePaths, filePaths...)
	return f
}

// AddFile add source file. alias of AddFilePath()
func (f *FileFinder) AddFile(filePaths ...string) *FileFinder {
	f.filePaths = append(f.filePaths, filePaths...)
	return f
}

// FindAll find and return founded file paths.
func (f *FileFinder) FindAll() []string {
	f.find()

	return f.filePaths
}

// Find find file paths.
func (f *FileFinder) Find() *FileFinder {
	f.find()
	return f
}

// do finding
func (f *FileFinder) find() {
	// mark found
	if f.founded {
		return
	}
	f.founded = true

	for _, filePath := range f.filePaths {
		fi, err := os.Stat(filePath)
		if err != nil {
			continue // ignore I/O error
		}
		if fi.IsDir() {
			continue // ignore I/O error
		}

		// cache file info
		f.osInfos[filePath] = fi
	}

	// do finding
	for _, dirPath := range f.dirPaths {
		f.findInDir(dirPath)
	}
}

// code refer filepath.glob()
func (f *FileFinder) findInDir(dirPath string) {
	dfi, err := os.Stat(dirPath)
	if err != nil {
		return // ignore I/O error
	}
	if !dfi.IsDir() {
		return // ignore I/O error
	}

	// opening
	d, err := os.Open(dirPath)
	if err != nil {
		return // ignore I/O error
	}

	names, _ := d.Readdirnames(-1)
	sort.Strings(names)

	hasDirFilter := len(f.dirFilters) > 0
	hasFileFilter := len(f.fileFilters) > 0
	for _, name := range names {
		fullPath := filepath.Join(dirPath, name)
		fi, err := os.Stat(fullPath)
		if err != nil {
			continue // ignore I/O error
		}

		// --- dir
		if fi.IsDir() {
			if f.excludeDotDir && name[0] == '.' {
				continue
			}

			var ok bool
			if hasDirFilter {
				for _, df := range f.dirFilters {
					ok = df.FilterDir(fullPath, name)
					if true == ok { // 有一个满足即可
						break
					}
				}

				// find in sub dir.
				if ok {
					f.findInDir(fullPath)
				}
			} else {
				// find in sub dir.
				f.findInDir(fullPath)
			}

			continue
		}

		// --- file
		if f.excludeDotFile && name[0] == '.' {
			continue
		}

		// use custom filter functions
		var ok bool
		if hasFileFilter {
			for _, ff := range f.fileFilters {
				ok = ff.FilterFile(fullPath, name)
				if true == ok { // 有一个满足即可
					break
				}
			}
		} else {
			ok = true
		}

		// append
		if ok {
			f.filePaths = append(f.filePaths, fullPath)
			// cache file info
			f.osInfos[fullPath] = fi
		}
	}

	d.Close()
}

// Each each file paths.
func (f *FileFinder) Each(fn func(filePath string)) {
	// ensure find is running
	f.find()

	for _, filePath := range f.filePaths {
		fn(filePath)
	}
}

// EachFile each file os.File
func (f *FileFinder) EachFile(fn func(file *os.File)) {
	// ensure find is running
	f.find()

	for _, filePath := range f.filePaths {
		file, err := os.Open(filePath)
		if err != nil {
			continue
		}

		fn(file)
	}
}

// EachStat each file os.FileInfo
func (f *FileFinder) EachStat(fn func(fi os.FileInfo, filePath string)) {
	// ensure find is running
	f.find()

	for filePath, fi := range f.osInfos {
		fn(fi, filePath)
	}
}

// EachBody each file contents
func (f *FileFinder) EachContents(fn func(contents, filePath string)) {
	// ensure find is running
	f.find()

	for _, filePath := range f.filePaths {
		bts, err := ioutil.ReadFile(filePath)
		if err != nil {
			continue
		}

		fn(string(bts), filePath)
	}
}

// Reset data setting.
func (f *FileFinder) Reset() {
	f.founded = false

	f.filePaths = make([]string, 0)

	f.excludeNames = make([]string, 0)
	f.excludeExts = make([]string, 0)
	f.excludeDirs = make([]string, 0)
}

// String all file paths
func (f *FileFinder) String() string {
	return strings.Join(f.filePaths, "\n")
}

//
// ------------------ built in file path filters ------------------
//

// ExtFilterFunc filter filepath by given file ext.
//
// Usage:
//	f := EmptyFiler()
//	f.AddFilter(ExtFilterFunc([]string{".go", ".md"}, true))
//	f.AddFilter(ExtFilterFunc([]string{".log", ".tmp"}, false))
func ExtFilterFunc(exts []string, include bool) FileFilterFunc {
	return func(filePath, _ string) bool {
		fExt := path.Ext(filePath)

		for _, ext := range exts {
			if fExt == ext {
				return include
			}
		}
		return !include
	}
}

// SuffixFilterFunc filter filepath by given file ext.
//
// Usage:
//	f := EmptyFiler()
//	f.AddFilter(SuffixFilterFunc([]string{"util.go", ".md"}, true))
//	f.AddFilter(SuffixFilterFunc([]string{"_test.go", ".log"}, false))
func SuffixFilterFunc(suffixes []string, include bool) FileFilterFunc {
	return func(filePath, _ string) bool {
		for _, sfx := range suffixes {
			if strings.HasSuffix(filePath, sfx) {
				return include
			}
		}
		return !include
	}
}

// PathNameFilterFunc filter filepath by given path names.
func PathNameFilterFunc(names []string, include bool) FileFilterFunc {
	return func(filePath, _ string) bool {
		for _, name := range names {
			if strings.Contains(filePath, name) {
				return include
			}
		}
		return !include
	}
}

// DotFileFilterFunc filter dot filename. eg: ".gitignore"
func DotFileFilterFunc(include bool) FileFilterFunc {
	return func(filePath, filename string) bool {
		// filename := path.Base(filePath)
		if filename[0] == '.' {
			return include
		}

		return !include
	}
}

// ModTimeFilterFunc filter file by modify time.
func ModTimeFilterFunc(limitSec int, op rune, include bool) FileFilterFunc {
	return func(filePath, filename string) bool {
		fi, err := os.Stat(filePath)
		if err != nil {
			return !include
		}

		now := time.Now().Second()
		if op == '>' {
			if now-fi.ModTime().Second() > limitSec {
				return include
			}

			return !include
		}

		// '<'
		if now-fi.ModTime().Second() < limitSec {
			return include
		}

		return !include
	}
}

// GlobFilterFunc filter filepath by given patterns.
//
// Usage:
//	f := EmptyFiler()
//	f.AddFilter(GlobFilterFunc([]string{"*_test.go"}, true))
func GlobFilterFunc(patterns []string, include bool) FileFilterFunc {
	return func(_, filename string) bool {
		for _, pattern := range patterns {
			if ok, _ := path.Match(pattern, filename); ok {
				return include
			}
		}
		return !include
	}
}

// RegexFilterFunc filter filepath by given regex pattern
//
// Usage:
//	f := EmptyFiler()
//	f.AddFilter(RegexFilterFunc(`[A-Z]\w+`, true))
func RegexFilterFunc(pattern string, include bool) FileFilterFunc {
	reg := regexp.MustCompile(pattern)

	return func(_, filename string) bool {
		return reg.MatchString(filename)
	}
}

//
// ----------------- built in dir path filters -----------------
//

// DotDirFilterFunc filter dot dirname. eg: ".idea"
func DotDirFilterFunc(include bool) DirFilterFunc {
	return func(_, dirname string) bool {
		if dirname[0] == '.' {
			return include
		}

		return !include
	}
}

// DirNameFilterFunc filter filepath by given dir names.
func DirNameFilterFunc(names []string, include bool) DirFilterFunc {
	return func(_, dirName string) bool {
		for _, name := range names {
			if dirName == name {
				return include
			}
		}
		return !include
	}
}
