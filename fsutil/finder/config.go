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
	FlagBoth = FlagFile | FlagDir
)

// ToFlag convert string to FindFlag
func ToFlag(s string) FindFlag {
	switch strings.ToLower(s) {
	case "dirs", "dir", "d":
		return FlagDir
	case "both", "b":
		return FlagBoth
	default:
		return FlagFile
	}
}

// Config for finder
type Config struct {
	init  bool
	depth int

	// ScanDirs scan dir paths for find.
	ScanDirs []string `json:"scan_dirs"`
	// FindFlags for find result. default is FlagFile
	FindFlags FindFlag `json:"find_flags"`
	// MaxDepth for find result. default is 0 - not limit
	MaxDepth int `json:"max_depth"`
	// UseAbsPath use abs path for find result. default is false
	UseAbsPath bool `json:"use_abs_path"`
	// CacheResult cache result for find result. default is false
	CacheResult bool `json:"cache_result"`
	// ExcludeDotDir exclude dot dir. default is true
	ExcludeDotDir bool `json:"exclude_dot_dir"`
	// ExcludeDotFile exclude dot dir. default is false
	ExcludeDotFile bool `json:"exclude_dot_file"`

	// Matchers generic include matchers for file/dir elems
	Matchers []Matcher
	// ExMatchers generic exclude matchers for file/dir elems
	ExMatchers []Matcher
	// DirMatchers include matchers for dir elems
	DirMatchers []Matcher
	// DirExMatchers exclude matchers for dir elems
	DirExMatchers []Matcher
	// FileMatchers include matchers for file elems
	FileMatchers []Matcher
	// FileExMatchers exclude matchers for file elems
	FileExMatchers []Matcher

	// commonly settings for build matchers

	// IncludeDirs include dir name list. eg: {"model"}
	IncludeDirs []string `json:"include_dirs"`
	// IncludeExts include file ext name list. eg: {".go", ".md"}
	IncludeExts []string `json:"include_exts"`
	// IncludeFiles include file name list. eg: {"go.mod"}
	IncludeFiles []string `json:"include_files"`
	// IncludePaths include file/dir path list. eg: {"path/to"}
	IncludePaths []string `json:"include_paths"`
	// IncludeNames include file/dir name list. eg: {"test", "some.go"}
	IncludeNames []string `json:"include_names"`

	// ExcludeDirs exclude dir name list. eg: {"test"}
	ExcludeDirs []string `json:"exclude_dirs"`
	// ExcludeExts exclude file ext name list. eg: {".go", ".md"}
	ExcludeExts []string `json:"exclude_exts"`
	// ExcludeFiles exclude file name list. eg: {"go.mod"}
	ExcludeFiles []string `json:"exclude_files"`
	// ExcludePaths exclude file/dir path list. eg: {"path/to"}
	ExcludePaths []string `json:"exclude_paths"`
	// ExcludeNames exclude file/dir name list. eg: {"test", "some.go"}
	ExcludeNames []string `json:"exclude_names"`
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

// NewEmptyConfig create a new Config
func NewEmptyConfig() *Config {
	return &Config{FindFlags: FlagFile}
}

// NewFinder create a new Finder by config
func (c *Config) NewFinder() *Finder {
	return NewWithConfig(c.Init())
}

// Init build matchers by config and append to Matchers.
func (c *Config) Init() *Config {
	if c.init {
		return c
	}

	// generic matchers
	if len(c.IncludeNames) > 0 {
		c.Matchers = append(c.Matchers, MatchNames(c.IncludeNames))
	}

	if len(c.IncludePaths) > 0 {
		c.Matchers = append(c.Matchers, MatchPaths(c.IncludePaths))
	}

	if len(c.ExcludePaths) > 0 {
		c.ExMatchers = append(c.ExMatchers, MatchPaths(c.ExcludePaths))
	}

	if len(c.ExcludeNames) > 0 {
		c.ExMatchers = append(c.ExMatchers, MatchNames(c.ExcludeNames))
	}

	// dir matchers
	if len(c.IncludeDirs) > 0 {
		c.DirMatchers = append(c.DirMatchers, MatchNames(c.IncludeDirs))
	}

	if len(c.ExcludeDirs) > 0 {
		c.DirExMatchers = append(c.DirExMatchers, MatchNames(c.ExcludeDirs))
	}

	// file matchers
	if len(c.IncludeExts) > 0 {
		c.FileMatchers = append(c.FileMatchers, MatchExts(c.IncludeExts))
	}

	if len(c.IncludeFiles) > 0 {
		c.FileMatchers = append(c.FileMatchers, MatchNames(c.IncludeFiles))
	}

	if len(c.ExcludeExts) > 0 {
		c.FileExMatchers = append(c.FileExMatchers, MatchExts(c.ExcludeExts))
	}

	if len(c.ExcludeFiles) > 0 {
		c.FileExMatchers = append(c.FileExMatchers, MatchNames(c.ExcludeFiles))
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
func (f *Finder) ConfigFn(fns ...func(c *Config)) *Finder { return f.WithConfigFn(fns...) }

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

// AddScan add source dir for find. alias of AddScanDirs()
func (f *Finder) AddScan(dirPaths ...string) *Finder { return f.AddScanDirs(dirPaths) }

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
// --------- add matchers to finder ---------
//

// Includes add include match matchers
func (f *Finder) Includes(fls []Matcher) *Finder {
	f.c.Matchers = append(f.c.Matchers, fls...)
	return f
}

// Collect add include match matchers. alias of Includes()
func (f *Finder) Collect(fls ...Matcher) *Finder { return f.Includes(fls) }

// Include add include match matchers. alias of Includes()
func (f *Finder) Include(fls ...Matcher) *Finder { return f.Includes(fls) }

// With add include match matchers. alias of Includes()
func (f *Finder) With(fls ...Matcher) *Finder { return f.Includes(fls) }

// Adds include match matchers. alias of Includes()
func (f *Finder) Adds(fls []Matcher) *Finder { return f.Includes(fls) }

// Add include match matchers. alias of Includes()
func (f *Finder) Add(fls ...Matcher) *Finder { return f.Includes(fls) }

// Excludes add exclude match matchers
func (f *Finder) Excludes(fls []Matcher) *Finder {
	f.c.ExMatchers = append(f.c.ExMatchers, fls...)
	return f
}

// Exclude add exclude match matchers. alias of Excludes()
func (f *Finder) Exclude(fls ...Matcher) *Finder { return f.Excludes(fls) }

// Without add exclude match matchers. alias of Excludes()
func (f *Finder) Without(fls ...Matcher) *Finder { return f.Excludes(fls) }

// Nots add exclude match matchers. alias of Excludes()
func (f *Finder) Nots(fls []Matcher) *Finder { return f.Excludes(fls) }

// Not add exclude match matchers. alias of Excludes()
func (f *Finder) Not(fls ...Matcher) *Finder { return f.Excludes(fls) }

// WithMatchers add include matchers
func (f *Finder) WithMatchers(fls []Matcher) *Finder {
	f.c.Matchers = append(f.c.Matchers, fls...)
	return f
}

// WithFilter add include matchers
func (f *Finder) WithFilter(fls ...Matcher) *Finder { return f.WithMatchers(fls) }

// MatchFiles add include file matchers
func (f *Finder) MatchFiles(fls []Matcher) *Finder {
	f.c.FileMatchers = append(f.c.FileMatchers, fls...)
	return f
}

// MatchFile add include file matchers
func (f *Finder) MatchFile(fls ...Matcher) *Finder { return f.MatchFiles(fls) }

// AddFiles add include file matchers
func (f *Finder) AddFiles(fls []Matcher) *Finder { return f.MatchFiles(fls) }

// AddFile add include file matchers
func (f *Finder) AddFile(fls ...Matcher) *Finder { return f.MatchFiles(fls) }

// NotFiles add exclude file matchers
func (f *Finder) NotFiles(fls []Matcher) *Finder {
	f.c.FileExMatchers = append(f.c.FileExMatchers, fls...)
	return f
}

// NotFile add exclude file matchers
func (f *Finder) NotFile(fls ...Matcher) *Finder { return f.NotFiles(fls) }

// MatchDirs add exclude dir matchers
func (f *Finder) MatchDirs(fls []Matcher) *Finder {
	f.c.DirMatchers = append(f.c.DirMatchers, fls...)
	return f
}

// MatchDir add exclude dir matchers
func (f *Finder) MatchDir(fls ...Matcher) *Finder { return f.MatchDirs(fls) }

// WithDirs add exclude dir matchers
func (f *Finder) WithDirs(fls []Matcher) *Finder { return f.MatchDirs(fls) }

// WithDir add exclude dir matchers
func (f *Finder) WithDir(fls ...Matcher) *Finder { return f.MatchDirs(fls) }

// NotDirs add exclude dir matchers
func (f *Finder) NotDirs(fls []Matcher) *Finder {
	f.c.DirExMatchers = append(f.c.DirExMatchers, fls...)
	return f
}

// NotDir add exclude dir matchers
func (f *Finder) NotDir(fls ...Matcher) *Finder { return f.NotDirs(fls) }
