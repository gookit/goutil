package finder

import "strings"

// commonly dot file and dirs
var (
	CommonlyDotDirs  = []string{".git", ".idea", ".vscode", ".svn", ".hg"}
	CommonlyDotFiles = []string{".gitignore", ".dockerignore", ".npmignore", ".DS_Store", ".env"}
)

// FindFlag type for find result.
type FindFlag uint8

// flags for find result.
const (
	FlagFile FindFlag = iota + 1 // only find files(default)
	FlagDir
)

// ToFlag convert string to FindFlag
func ToFlag(s string) FindFlag {
	switch strings.ToLower(s) {
	case "dir", "d":
		return FlagDir
	case "both", "b":
		return FlagFile | FlagDir
	default:
		return FlagFile
	}
}

// Config for finder
type Config struct {
	init  bool
	depth int

	// ScanDirs scan dir paths for find.
	ScanDirs []string
	// FindFlags for find result. default is FlagFile
	FindFlags FindFlag
	// MaxDepth for find result. default is 0 - not limit
	MaxDepth int
	// UseAbsPath use abs path for find result. default is false
	UseAbsPath bool
	// CacheResult cache result for find result. default is false
	CacheResult bool
	// ExcludeDotDir exclude dot dir. default is true
	ExcludeDotDir bool
	// ExcludeDotFile exclude dot dir. default is false
	ExcludeDotFile bool

	// Filters generic include filters for filter file/dir elems
	Filters []Filter
	// ExFilters generic exclude filters for filter file/dir elems
	ExFilters []Filter
	// DirFilters include filters for dir elems
	DirFilters []Filter
	// DirFilters exclude filters for dir elems
	DirExFilters []Filter
	// FileFilters include filters for file elems
	FileFilters []Filter
	// FileExFilters exclude filters for file elems
	FileExFilters []Filter

	// commonly settings for build filters

	// IncludeDirs include dir name list. eg: {"model"}
	IncludeDirs []string
	// IncludeExts include file ext name list. eg: {".go", ".md"}
	IncludeExts []string
	// IncludeFiles include file name list. eg: {"go.mod"}
	IncludeFiles []string
	// IncludePaths include file/dir path list. eg: {"path/to"}
	IncludePaths []string
	// IncludeNames include file/dir name list. eg: {"test", "some.go"}
	IncludeNames []string

	// ExcludeDirs exclude dir name list. eg: {"test"}
	ExcludeDirs []string
	// ExcludeExts exclude file ext name list. eg: {".go", ".md"}
	ExcludeExts []string
	// ExcludeFiles exclude file name list. eg: {"go.mod"}
	ExcludeFiles []string
	// ExcludePaths exclude file/dir path list. eg: {"path/to"}
	ExcludePaths []string
	// ExcludeNames exclude file/dir name list. eg: {"test", "some.go"}
	ExcludeNames []string
}

// NewConfig create a new Config
func NewConfig(dirs ...string) *Config {
	return &Config{
		ScanDirs:  dirs,
		FindFlags: FlagFile,
		// with default setting.
		ExcludeDotDir: true,
	}
}

// NewFinder create a new Finder by config
func (c *Config) NewFinder() *Finder {
	return NewWithConfig(c.Init())
}

// Init build filters by config and append to Filters.
func (c *Config) Init() *Config {
	if c.init {
		return c
	}

	// generic filters
	if len(c.IncludeNames) > 0 {
		c.Filters = append(c.Filters, WithNames(c.IncludeNames))
	}

	if len(c.IncludePaths) > 0 {
		c.Filters = append(c.Filters, WithPaths(c.IncludePaths))
	}

	if len(c.ExcludePaths) > 0 {
		c.ExFilters = append(c.ExFilters, ExcludePaths(c.ExcludePaths))
	}

	if len(c.ExcludeNames) > 0 {
		c.ExFilters = append(c.ExFilters, ExcludeNames(c.ExcludeNames))
	}

	// dir filters
	if len(c.IncludeDirs) > 0 {
		c.DirFilters = append(c.DirFilters, IncludeNames(c.IncludeDirs))
	}

	if len(c.ExcludeDirs) > 0 {
		c.DirExFilters = append(c.DirExFilters, ExcludeNames(c.ExcludeDirs))
	}

	// file filters
	if len(c.IncludeExts) > 0 {
		c.FileFilters = append(c.FileFilters, IncludeExts(c.IncludeExts))
	}

	if len(c.IncludeFiles) > 0 {
		c.FileFilters = append(c.FileFilters, IncludeNames(c.IncludeFiles))
	}

	if len(c.ExcludeExts) > 0 {
		c.FileExFilters = append(c.FileExFilters, ExcludeExts(c.ExcludeExts))
	}

	if len(c.ExcludeFiles) > 0 {
		c.FileExFilters = append(c.FileExFilters, ExcludeNames(c.ExcludeFiles))
	}

	return c
}

//
// --------- config for finder ---------
//

// WithConfig on the finder
func (f *Finder) WithConfig(c *Config) *Finder {
	f.c = c
	return f
}

// ConfigFn the finder. alias of WithConfigFn()
func (f *Finder) ConfigFn(fns ...func(c *Config)) *Finder {
	return f.WithConfigFn(fns...)
}

// WithConfigFn the finder
func (f *Finder) WithConfigFn(fns ...func(c *Config)) *Finder {
	if f.c == nil {
		f.c = &Config{}
	}

	for _, fn := range fns {
		fn(f.c)
	}
	return f
}

// AddScanDirs add source dir for find
func (f *Finder) AddScanDirs(dirPaths []string) *Finder {
	f.c.ScanDirs = append(f.c.ScanDirs, dirPaths...)
	return f
}

// AddScanDir add source dir for find. alias of AddScanDirs()
func (f *Finder) AddScanDir(dirPaths ...string) *Finder { return f.AddScanDirs(dirPaths) }

// ScanDir add source dir for find. alias of AddScanDirs()
func (f *Finder) ScanDir(dirPaths ...string) *Finder { return f.AddScanDirs(dirPaths) }

// CacheResult cache result for find result.
func (f *Finder) CacheResult(enable ...bool) *Finder {
	if len(enable) > 0 {
		f.c.CacheResult = enable[0]
	} else {
		f.c.CacheResult = true
	}
	return f
}

// WithCacheResult cache result for find result. alias of CacheResult()
func (f *Finder) WithCacheResult(enable ...bool) *Finder {
	return f.CacheResult(enable...)
}

// WithFlags set find flags.
func (f *Finder) WithFlags(flags FindFlag) *Finder {
	f.c.FindFlags = flags
	return f
}

// WithStrFlag set find flags by string.
func (f *Finder) WithStrFlag(s string) *Finder {
	f.c.FindFlags = ToFlag(s)
	return f
}

// OnlyFindDir only find dir.
func (f *Finder) OnlyFindDir() *Finder { return f.WithFlags(FlagDir) }

// FileAndDir both find file and dir.
func (f *Finder) FileAndDir() *Finder { return f.WithFlags(FlagDir | FlagFile) }

// UseAbsPath use absolute path for find result. alias of WithUseAbsPath()
func (f *Finder) UseAbsPath(enable ...bool) *Finder { return f.WithUseAbsPath(enable...) }

// WithUseAbsPath use absolute path for find result.
func (f *Finder) WithUseAbsPath(enable ...bool) *Finder {
	if len(enable) > 0 {
		f.c.UseAbsPath = enable[0]
	} else {
		f.c.UseAbsPath = true
	}
	return f
}

// WithMaxDepth set max depth for find.
func (f *Finder) WithMaxDepth(i int) *Finder {
	f.c.MaxDepth = i
	return f
}

// IncludeDir include dir names.
func (f *Finder) IncludeDir(dirs ...string) *Finder {
	f.c.IncludeDirs = append(f.c.IncludeDirs, dirs...)
	return f
}

// WithDirName include dir names. alias of IncludeDir()
func (f *Finder) WithDirName(dirs ...string) *Finder { return f.IncludeDir(dirs...) }

// IncludeFile include file names.
func (f *Finder) IncludeFile(files ...string) *Finder {
	f.c.IncludeFiles = append(f.c.IncludeFiles, files...)
	return f
}

// WithFileName include file names. alias of IncludeFile()
func (f *Finder) WithFileName(files ...string) *Finder { return f.IncludeFile(files...) }

// IncludeName include file or dir names.
func (f *Finder) IncludeName(names ...string) *Finder {
	f.c.IncludeNames = append(f.c.IncludeNames, names...)
	return f
}

// WithNames include file or dir names. alias of IncludeName()
func (f *Finder) WithNames(names []string) *Finder { return f.IncludeName(names...) }

// IncludeExt include file exts.
func (f *Finder) IncludeExt(exts ...string) *Finder {
	f.c.IncludeExts = append(f.c.IncludeExts, exts...)
	return f
}

// WithExts include file exts. alias of IncludeExt()
func (f *Finder) WithExts(exts []string) *Finder { return f.IncludeExt(exts...) }

// WithFileExt include file exts. alias of IncludeExt()
func (f *Finder) WithFileExt(exts ...string) *Finder { return f.IncludeExt(exts...) }

// IncludePath include file or dir paths.
func (f *Finder) IncludePath(paths ...string) *Finder {
	f.c.IncludePaths = append(f.c.IncludePaths, paths...)
	return f
}

// WithPaths include file or dir paths. alias of IncludePath()
func (f *Finder) WithPaths(paths []string) *Finder { return f.IncludePath(paths...) }

// WithSubPath include file or dir paths. alias of IncludePath()
func (f *Finder) WithSubPath(paths ...string) *Finder { return f.IncludePath(paths...) }

// ExcludeDir exclude dir names.
func (f *Finder) ExcludeDir(dirs ...string) *Finder {
	f.c.ExcludeDirs = append(f.c.ExcludeDirs, dirs...)
	return f
}

// WithoutDir exclude dir names. alias of ExcludeDir()
func (f *Finder) WithoutDir(dirs ...string) *Finder { return f.ExcludeDir(dirs...) }

// WithoutNames exclude file or dir names.
func (f *Finder) WithoutNames(names []string) *Finder {
	f.c.ExcludeNames = append(f.c.ExcludeNames, names...)
	return f
}

// ExcludeName exclude file names. alias of WithoutNames()
func (f *Finder) ExcludeName(names ...string) *Finder { return f.WithoutNames(names) }

// ExcludeFile exclude file names.
func (f *Finder) ExcludeFile(files ...string) *Finder {
	f.c.ExcludeFiles = append(f.c.ExcludeFiles, files...)
	return f
}

// WithoutFile exclude file names. alias of ExcludeFile()
func (f *Finder) WithoutFile(files ...string) *Finder { return f.ExcludeFile(files...) }

// ExcludeExt exclude file exts.
//
// eg: ExcludeExt(".go", ".java")
func (f *Finder) ExcludeExt(exts ...string) *Finder {
	f.c.ExcludeExts = append(f.c.ExcludeExts, exts...)
	return f
}

// WithoutExt exclude file exts. alias of ExcludeExt()
func (f *Finder) WithoutExt(exts ...string) *Finder { return f.ExcludeExt(exts...) }

// WithoutExts exclude file exts. alias of ExcludeExt()
func (f *Finder) WithoutExts(exts []string) *Finder { return f.ExcludeExt(exts...) }

// ExcludePath exclude file paths.
func (f *Finder) ExcludePath(paths ...string) *Finder {
	f.c.ExcludePaths = append(f.c.ExcludePaths, paths...)
	return f
}

// WithoutPath exclude file paths. alias of ExcludePath()
func (f *Finder) WithoutPath(paths ...string) *Finder { return f.ExcludePath(paths...) }

// WithoutPaths exclude file paths. alias of ExcludePath()
func (f *Finder) WithoutPaths(paths []string) *Finder { return f.ExcludePath(paths...) }

// ExcludeDotDir exclude dot dir names. eg: ".idea"
func (f *Finder) ExcludeDotDir(exclude ...bool) *Finder {
	if len(exclude) > 0 {
		f.c.ExcludeDotDir = exclude[0]
	} else {
		f.c.ExcludeDotDir = true
	}
	return f
}

// WithoutDotDir exclude dot dir names. alias of ExcludeDotDir().
func (f *Finder) WithoutDotDir(exclude ...bool) *Finder {
	return f.ExcludeDotDir(exclude...)
}

// NoDotDir exclude dot dir names. alias of ExcludeDotDir().
func (f *Finder) NoDotDir(exclude ...bool) *Finder {
	return f.ExcludeDotDir(exclude...)
}

// ExcludeDotFile exclude dot dir names. eg: ".gitignore"
func (f *Finder) ExcludeDotFile(exclude ...bool) *Finder {
	if len(exclude) > 0 {
		f.c.ExcludeDotFile = exclude[0]
	} else {
		f.c.ExcludeDotFile = true
	}
	return f
}

// WithoutDotFile exclude dot dir names. alias of ExcludeDotFile().
func (f *Finder) WithoutDotFile(exclude ...bool) *Finder {
	return f.ExcludeDotFile(exclude...)
}

// NoDotFile exclude dot dir names. alias of ExcludeDotFile().
func (f *Finder) NoDotFile(exclude ...bool) *Finder {
	return f.ExcludeDotFile(exclude...)
}

//
// --------- add filters to finder ---------
//

// Includes add include match filters
func (f *Finder) Includes(fls []Filter) *Finder {
	f.c.Filters = append(f.c.Filters, fls...)
	return f
}

// Include add include match filters. alias of Includes()
func (f *Finder) Include(fls ...Filter) *Finder { return f.Includes(fls) }

// With add include match filters. alias of Includes()
func (f *Finder) With(fls ...Filter) *Finder { return f.Includes(fls) }

// Adds include match filters. alias of Includes()
func (f *Finder) Adds(fls []Filter) *Finder { return f.Includes(fls) }

// Add include match filters. alias of Includes()
func (f *Finder) Add(fls ...Filter) *Finder { return f.Includes(fls) }

// Excludes add exclude match filters
func (f *Finder) Excludes(fls []Filter) *Finder {
	f.c.ExFilters = append(f.c.ExFilters, fls...)
	return f
}

// Exclude add exclude match filters. alias of Excludes()
func (f *Finder) Exclude(fls ...Filter) *Finder { return f.Excludes(fls) }

// Without add exclude match filters. alias of Excludes()
func (f *Finder) Without(fls ...Filter) *Finder { return f.Excludes(fls) }

// Nots add exclude match filters. alias of Excludes()
func (f *Finder) Nots(fls []Filter) *Finder { return f.Excludes(fls) }

// Not add exclude match filters. alias of Excludes()
func (f *Finder) Not(fls ...Filter) *Finder { return f.Excludes(fls) }

// WithFilters add include filters
func (f *Finder) WithFilters(fls []Filter) *Finder {
	f.c.Filters = append(f.c.Filters, fls...)
	return f
}

// WithFilter add include filters
func (f *Finder) WithFilter(fls ...Filter) *Finder { return f.WithFilters(fls) }

// MatchFiles add include file filters
func (f *Finder) MatchFiles(fls []Filter) *Finder {
	f.c.FileFilters = append(f.c.FileFilters, fls...)
	return f
}

// MatchFile add include file filters
func (f *Finder) MatchFile(fls ...Filter) *Finder { return f.MatchFiles(fls) }

// AddFiles add include file filters
func (f *Finder) AddFiles(fls []Filter) *Finder { return f.MatchFiles(fls) }

// AddFile add include file filters
func (f *Finder) AddFile(fls ...Filter) *Finder { return f.MatchFiles(fls) }

// NotFiles add exclude file filters
func (f *Finder) NotFiles(fls []Filter) *Finder {
	f.c.FileExFilters = append(f.c.FileExFilters, fls...)
	return f
}

// NotFile add exclude file filters
func (f *Finder) NotFile(fls ...Filter) *Finder { return f.NotFiles(fls) }

// MatchDirs add exclude dir filters
func (f *Finder) MatchDirs(fls []Filter) *Finder {
	f.c.DirFilters = append(f.c.DirFilters, fls...)
	return f
}

// MatchDir add exclude dir filters
func (f *Finder) MatchDir(fls ...Filter) *Finder { return f.MatchDirs(fls) }

// WithDirs add exclude dir filters
func (f *Finder) WithDirs(fls []Filter) *Finder { return f.MatchDirs(fls) }

// WithDir add exclude dir filters
func (f *Finder) WithDir(fls ...Filter) *Finder { return f.MatchDirs(fls) }

// NotDirs add exclude dir filters
func (f *Finder) NotDirs(fls []Filter) *Finder {
	f.c.DirExFilters = append(f.c.DirExFilters, fls...)
	return f
}

// NotDir add exclude dir filters
func (f *Finder) NotDir(fls ...Filter) *Finder {
	return f.NotDirs(fls)
}
