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

// MultiMatcher wrapper for multi matchers
type MultiMatcher struct {
	Before   Matcher
	Matchers []Matcher
}

// Add matchers
func (mf *MultiMatcher) Add(fls ...Matcher) {
	mf.Matchers = append(mf.Matchers, fls...)
}

// Apply check file path is match.
func (mf *MultiMatcher) Apply(el Elem) bool {
	if mf.Before != nil && !mf.Before.Apply(el) {
		return false
	}

	for _, fl := range mf.Matchers {
		if !fl.Apply(el) {
			return false
		}
	}
	return true
}

// NewDirMatchers create a new dir matchers
func NewDirMatchers(fls ...Matcher) *MultiMatcher {
	return &MultiMatcher{
		Before:   MatchDir,
		Matchers: fls,
	}
}

// NewFileMatchers create a new dir matchers
func NewFileMatchers(fls ...Matcher) *MultiMatcher {
	return &MultiMatcher{
		Before:   MatchFile,
		Matchers: fls,
	}
}

// ------------------ Body Matcher ------------------

// BodyMatcher for match file contents.
type BodyMatcher interface {
	Apply(filePath string, body *bytes.Buffer) bool
}

// BodyMatcherFunc for match file contents.
type BodyMatcherFunc func(filePath string, body *bytes.Buffer) bool

// Apply for match file contents.
func (fn BodyMatcherFunc) Apply(filePath string, body *bytes.Buffer) bool {
	return fn(filePath, body)
}

// BodyMatchers multi body matchers as Matcher
type BodyMatchers struct {
	Matchers []BodyMatcher
}

// NewBodyMatchers create a new body matchers
//
// Usage:
//
//		bf := finder.NewBodyMatchers(
//			finder.BodyMatcherFunc(func(filePath string, buf *bytes.Buffer) bool {
//				// match file contents
//				return true
//			}),
//		)
//
//	 es := finder.NewFinder('path/to/dir').Add(bf).Elems()
//	 for el := range es {
//			fmt.Println(el.Path())
//	 }
func NewBodyMatchers(fls ...BodyMatcher) *BodyMatchers {
	return &BodyMatchers{
		Matchers: fls,
	}
}

// AddMatcher add matchers
func (mf *BodyMatchers) AddMatcher(fls ...BodyMatcher) {
	mf.Matchers = append(mf.Matchers, fls...)
}

// Apply check file contents is match.
func (mf *BodyMatchers) Apply(el Elem) bool {
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
	for _, fl := range mf.Matchers {
		if !fl.Apply(el.Path(), buf) {
			return false
		}
	}
	return true
}
