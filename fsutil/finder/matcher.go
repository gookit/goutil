package finder

import (
	"bytes"

	"github.com/gookit/goutil/fsutil"
)

// Matcher for match file path.
type Matcher interface {
	// Apply check find elem. return False will skip this file.
	Apply(elem Elem) bool
}

// MatcherFunc for match file info, return False will skip this file
type MatcherFunc func(elem Elem) bool

// Apply check file path. return False will skip this file.
func (fn MatcherFunc) Apply(elem Elem) bool {
	return fn(elem)
}

// ------------------ Multi matcher wrapper ------------------

// MultiFilter wrapper for multi matchers
type MultiFilter struct {
	Before  Matcher
	Filters []Matcher
}

// Add matchers
func (mf *MultiFilter) Add(fls ...Matcher) {
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

// NewDirFilters create a new dir matchers
func NewDirFilters(fls ...Matcher) *MultiFilter {
	return &MultiFilter{
		Before:  MatchDir,
		Filters: fls,
	}
}

// NewFileFilters create a new dir matchers
func NewFileFilters(fls ...Matcher) *MultiFilter {
	return &MultiFilter{
		Before:  MatchFile,
		Filters: fls,
	}
}

// ------------------ Body Matcher ------------------

// BodyFilter for filter file contents.
type BodyFilter interface {
	Apply(filePath string, buf *bytes.Buffer) bool
}

// BodyMatcherFunc for filter file contents.
type BodyMatcherFunc func(filePath string, buf *bytes.Buffer) bool

// Apply for filter file contents.
func (fn BodyMatcherFunc) Apply(filePath string, buf *bytes.Buffer) bool {
	return fn(filePath, buf)
}

// BodyFilters multi body matchers as Matcher
type BodyFilters struct {
	Filters []BodyFilter
}

// NewBodyFilters create a new body matchers
//
// Usage:
//
//		bf := finder.NewBodyFilters(
//			finder.BodyMatcherFunc(func(filePath string, buf *bytes.Buffer) bool {
//				// filter file contents
//				return true
//			}),
//		)
//
//	 es := finder.NewFinder('path/to/dir').Add(bf).Elems()
//	 for el := range es {
//			fmt.Println(el.Path())
//	 }
func NewBodyFilters(fls ...BodyFilter) *BodyFilters {
	return &BodyFilters{
		Filters: fls,
	}
}

// AddFilter add matchers
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

	// apply matchers
	for _, fl := range mf.Filters {
		if !fl.Apply(el.Path(), buf) {
			return false
		}
	}
	return true
}
