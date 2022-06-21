package finder

// FileMeta struct
// type FileMeta struct {
// 	filePath string
// 	filename string
// }

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

// AddFilters Result get find paths
func (r *FindResults) AddFilters(filterFuncs ...FileFilter) *FindResults {
	// TODO
	return r
}

// Filter Result get find paths
func (r *FindResults) Filter() *FindResults {
	return r
}

// Each Result get find paths
func (r *FindResults) Each() *FindResults {
	return r
}

// Result get find paths
func (r *FindResults) Result() []string {
	return r.filePaths
}
