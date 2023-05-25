package finder

import (
	"bytes"
	"io/fs"

	"github.com/gookit/goutil/fsutil"
)

// Filter for filter file path.
type Filter interface {
	// Apply check find elem. return False will filter this file.
	Apply(elem Elem) bool
}

// FilterFunc for filter file info, return False will filter this file
type FilterFunc func(elem Elem) bool

// Apply check file path. return False will filter this file.
func (fn FilterFunc) Apply(elem Elem) bool {
	return fn(elem)
}

// ------------------ raw filter wrapper ------------------

// RawFilter for filter file path.
type RawFilter interface {
	// Apply check file path. return False will filter this file.
	Apply(fPath string, ent fs.DirEntry) bool
}

// RawFilterFunc for filter file info, return False will filter this file
type RawFilterFunc func(fPath string, ent fs.DirEntry) bool

// Apply check file path. return False will filter this file.
func (fn RawFilterFunc) Apply(fPath string, ent fs.DirEntry) bool {
	return fn(fPath, ent)
}

// WrapRawFilter wrap a RawFilter to Filter
func WrapRawFilter(rf RawFilter) Filter {
	return FilterFunc(func(elem Elem) bool {
		return rf.Apply(elem.Path(), elem)
	})
}

// WrapRawFilters wrap RawFilter list to Filter list
func WrapRawFilters(rfs ...RawFilter) []Filter {
	fls := make([]Filter, len(rfs))
	for i, rf := range rfs {
		fls[i] = WrapRawFilter(rf)
	}
	return fls
}

// ------------------ Multi filter wrapper ------------------

// MultiFilter wrapper for multi filters
type MultiFilter struct {
	Before  Filter
	Filters []Filter
}

// AddFilter add filters
func (mf *MultiFilter) AddFilter(fls ...Filter) {
	mf.Filters = append(mf.Filters, fls...)
}

// Apply check file path. return False will filter this file.
func (mf *MultiFilter) Apply(el Elem) bool {
	if mf.Before != nil && !mf.Before.Apply(el) {
		return false
	}

	for _, fl := range mf.Filters {
		if !fl.Apply(el) {
			return false
		}
	}
	return true
}

// NewDirFilters create a new dir filters
func NewDirFilters(fls ...Filter) *MultiFilter {
	return &MultiFilter{
		Before:  OnlyDirFilter,
		Filters: fls,
	}
}

// NewFileFilters create a new dir filters
func NewFileFilters(fls ...Filter) *MultiFilter {
	return &MultiFilter{
		Before:  OnlyFileFilter,
		Filters: fls,
	}
}

// ------------------ Body Filter ------------------

// BodyFilter for filter file contents.
type BodyFilter interface {
	Apply(filePath string, buf *bytes.Buffer) bool
}

// BodyFilterFunc for filter file contents.
type BodyFilterFunc func(filePath string, buf *bytes.Buffer) bool

// Apply for filter file contents.
func (fn BodyFilterFunc) Apply(filePath string, buf *bytes.Buffer) bool {
	return fn(filePath, buf)
}

// BodyFilters multi body filters as Filter
type BodyFilters struct {
	Filters []BodyFilter
}

// NewBodyFilters create a new body filters
func NewBodyFilters(fls ...BodyFilter) *BodyFilters {
	return &BodyFilters{
		Filters: fls,
	}
}

// AddFilter add filters
func (mf *BodyFilters) AddFilter(fls ...BodyFilter) {
	mf.Filters = append(mf.Filters, fls...)
}

// Apply check file path. return False will filter this file.
func (mf *BodyFilters) Apply(fPath string, ent fs.DirEntry) bool {
	if ent.IsDir() {
		return false
	}

	// read file contents
	buf := bytes.NewBuffer(nil)
	file, err := fsutil.OpenReadFile(fPath)
	if err != nil {
		return false
	}

	_, err = buf.ReadFrom(file)
	if err != nil {
		file.Close()
		return false
	}
	file.Close()

	// apply filters
	for _, fl := range mf.Filters {
		if !fl.Apply(fPath, buf) {
			return false
		}
	}
	return true
}
