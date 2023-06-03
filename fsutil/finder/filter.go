package finder

import (
	"bytes"

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
//
// Usage:
//
//		bf := finder.NewBodyFilters(
//			finder.BodyFilterFunc(func(filePath string, buf *bytes.Buffer) bool {
//				// filter file contents
//				return true
//			}),
//		)
//
//	 es := finder.NewFinder('path/to/dir').WithFileFilter(bf).Elems()
//	 for el := range es {
//			fmt.Println(el.Path())
//	 }
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
func (mf *BodyFilters) Apply(el Elem) bool {
	if el.IsDir() {
		return false
	}

	// read file contents
	buf := bytes.NewBuffer(nil)
	file, err := fsutil.OpenReadFile(el.Path())
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
		if !fl.Apply(el.Path(), buf) {
			return false
		}
	}
	return true
}
