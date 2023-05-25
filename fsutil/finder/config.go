package finder

// commonly dot file and dirs
var (
	CommonlyDotDir  = []string{".git", ".idea", ".vscode", ".svn", ".hg"}
	CommonlyDotFile = []string{".gitignore", ".dockerignore", ".npmignore", ".DS_Store"}
)

// FindFlag type for find result.
type FindFlag uint8

// flags for find result.
const (
	FlagFile FindFlag = iota + 1 // only find files(default)
	FlagDir
)

// Config for finder
type Config struct {
	curDepth int

	// DirPaths src paths for find file.
	DirPaths []string
	// FindFlags for find result. default is FlagFile
	FindFlags FindFlag
	// MaxDepth for find result. default is 0
	MaxDepth int
	// CacheResult cache result for find result. default is false
	CacheResult bool

	// IncludeDirs name list. eg: {"model"}
	IncludeDirs []string
	// IncludeExts name list. eg: {".go", ".md"}
	IncludeExts []string
	// IncludeFiles name list. eg: {"go.mod"}
	IncludeFiles []string
	// IncludePaths list. eg: {"path/to"}
	IncludePaths []string

	// ExcludeDirs name list. eg: {"test"}
	ExcludeDirs []string
	// ExcludeExts name list. eg: {".go", ".md"}
	ExcludeExts []string
	// ExcludeFiles name list. eg: {"go.mod"}
	ExcludeFiles []string
	// ExcludePaths list. eg: {"path/to"}
	ExcludePaths []string
	// ExcludeNames file/dir name list. eg: {"test", "some.go"}
	ExcludeNames []string

	ExcludeDotDir  bool
	ExcludeDotFile bool

	// Filters generic filters for filter file/dir paths
	Filters []Filter

	DirFilters  []Filter     // filters for filter dir paths
	FileFilters []Filter     // filters for filter file paths
	BodyFilters []BodyFilter // filters for filter file body
}

// NewConfig create a new Config
func NewConfig(dirs ...string) *Config {
	return &Config{
		DirPaths:  dirs,
		FindFlags: FlagFile,
		// with default setting.
		ExcludeDotDir: true,
	}
}

// NewFinder create a new FileFinder by config
func (c *Config) NewFinder() *FileFinder {
	return NewWithConfig(c)
}
