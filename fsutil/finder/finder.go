package finder

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// TODO use excludeDotFlag 1 file 2 dir 1|2 both
type exDotFlag uint8

const (
	ExDotFile exDotFlag = 1
	ExDotDir  exDotFlag = 2
)

// FileFinder struct
type FileFinder struct {
	// r *FindResults

	// dir paths for find file.
	dirPaths []string
	// file paths for handle.
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

	// mark has been run find()
	founded bool
	// founded file paths.
	filePaths []string

	// the founded file instances
	// osFiles map[string]*os.File
	osInfos map[string]os.FileInfo

	// handlers on found
	pathHandler func(filePath string)
	statHandler func(fi os.FileInfo, filePath string)
}

// EmptyFinder new empty FileFinder instance
func EmptyFinder() *FileFinder {
	return &FileFinder{}
}

// NewFinder new instance with source dir paths.
func NewFinder(dirPaths []string, filePaths ...string) *FileFinder {
	return &FileFinder{
		dirPaths:  dirPaths,
		filePaths: filePaths,
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

// Find files in given dir paths.
func (f *FileFinder) Find() *FileFinder {
	f.find()
	return f
}

// do finding
func (f *FileFinder) find() {
	if f.founded {
		return
	}

	// mark found
	f.founded = true
	for _, filePath := range f.filePaths {
		fi, err := os.Stat(filePath)
		if err != nil {
			continue // ignore I/O error
		}
		if fi.IsDir() {
			continue // ignore I/O error
		}

		// call handler
		if f.pathHandler != nil {
			f.pathHandler(filePath)
		}
		if f.statHandler != nil {
			f.statHandler(fi, filePath)
		}
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

	// sort.Strings(names)
	// names, _ := d.Readdirnames(-1)
	stats, _ := d.Readdir(-1)
	_ = d.Close() // close dir.

	for _, fi := range stats {
		baseName := fi.Name()
		fullPath := filepath.Join(dirPath, baseName)

		// --- dir
		if fi.IsDir() {
			if f.excludeDotDir && baseName[0] == '.' {
				continue
			}

			ok := true
			for _, df := range f.dirFilters {
				ok = df.FilterDir(fullPath, baseName)
				if true == ok { // 有一个满足即可
					break
				}
			}

			// find in sub dir.
			if ok {
				f.findInDir(fullPath)
			}
			continue
		}

		// --- file
		if f.excludeDotFile && baseName[0] == '.' {
			continue
		}

		// use custom filter functions
		ok := true
		for _, ff := range f.fileFilters {
			if ok = ff.FilterFile(fullPath, baseName); ok { // 有一个满足即可
				break
			}
		}

		// append
		if ok {
			f.filePaths = append(f.filePaths, fullPath)

			// call handler
			if f.pathHandler != nil {
				f.pathHandler(fullPath)
			}
			if f.statHandler != nil {
				f.statHandler(fi, fullPath)
			}
		}
	}
}

// EachFile each file os.File
func (f *FileFinder) EachFile(fn func(file *os.File)) {
	f.Each(func(filePath string) {
		file, err := os.Open(filePath)
		if err != nil {
			return
		}
		fn(file)
	})
}

// Each file paths.
func (f *FileFinder) Each(fn func(filePath string)) {
	f.pathHandler = fn
	if f.founded {
		for _, filePath := range f.filePaths {
			fn(filePath)
		}
		return
	}

	f.find()
}

// EachStat each file os.FileInfo
func (f *FileFinder) EachStat(fn func(fi os.FileInfo, filePath string)) {
	f.statHandler = fn
	f.find()
}

// EachContents handle each found file contents
func (f *FileFinder) EachContents(fn func(contents, filePath string)) {
	f.Each(func(filePath string) {
		bts, err := ioutil.ReadFile(filePath)
		if err != nil {
			return
		}

		fn(string(bts), filePath)
	})
}

// Reset data setting.
func (f *FileFinder) Reset() {
	f.founded = false
	f.filePaths = f.filePaths[:0]

	f.excludeNames = make([]string, 0)
	f.excludeExts = make([]string, 0)
	f.excludeDirs = make([]string, 0)
}

// FilePaths get
func (f *FileFinder) FilePaths() []string {
	return f.filePaths
}

// String all file paths
func (f *FileFinder) String() string {
	return strings.Join(f.filePaths, "\n")
}
