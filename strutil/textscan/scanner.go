// Package textscan implements text scanner for quickly parse text contents.
// can use for parse like INI, Properties format contents
package textscan

import (
	"bufio"
	"bytes"
	"io"
	"strings"

	"github.com/gookit/goutil/errorx"
)

// Matcher interface
type Matcher interface {
	// Match text line by kind, if success returns a new Token
	Match(line string, prev Token) (tok Token, err error)
}

// TextScanner struct.
type TextScanner struct {
	in *bufio.Scanner

	// token matchers
	matchers []Matcher

	line int
	next string
	tok  Token
	err  error
}

// NewScanner instance
func NewScanner(in interface{}) *TextScanner {
	ts := &TextScanner{}
	if in != nil {
		ts.SetInput(in)
	}
	return ts
}

// SetInput for scan and parse
func (s *TextScanner) SetInput(in interface{}) {
	s.line = 1 // init

	switch typIn := in.(type) {
	case *bufio.Scanner:
		s.in = typIn
	case io.Reader:
		s.in = bufio.NewScanner(typIn)
	case []byte:
		s.in = bufio.NewScanner(bytes.NewReader(typIn))
	case string:
		s.in = bufio.NewScanner(strings.NewReader(typIn))
	default:
		panic("invalid input data for parse")
	}
}

// SetSplit set split func on scan
func (s *TextScanner) SetSplit(fn bufio.SplitFunc) {
	if s.in == nil {
		panic("must be set input before set split func")
	}
	s.in.Split(fn)
}

// AddMatchers register token matchers
func (s *TextScanner) AddMatchers(ms ...Matcher) {
	s.matchers = append(s.matchers, ms...)
}

// Each every token by given func
func (s *TextScanner) Each(fn func(t Token)) error {
	for s.Scan() {
		fn(s.Token())
	}
	return s.err
}

// Scan source input and parsing
//
// Usage:
//
//	for ts.Scan() {
//		ts.Token()
//	}
func (s *TextScanner) Scan() bool {
	if s.next != "" {
		return s.matchToken(s.next)
	}

	if ok, text := s.ScanNext(); ok {
		return s.matchToken(text)
	}

	s.tok = nil
	return false
}

func (s *TextScanner) matchToken(text string) (ok bool) {
	for _, m := range s.matchers {
		tok, err := m.Match(text, s.tok)
		if err != nil {
			s.err = errorx.Wrapf(err, "at line %d", s.line)
			return false
		}

		if tok != nil {
			if tok.HasMore() {
				if err := tok.ScanMore(s); err != nil {
					s.err = errorx.Wrapf(err, "at line %d", s.line)
					return false
				}
			}

			s.tok = tok
			return true
		}
	}

	// emtpy line, match next valid token
	if strings.TrimSpace(text) == "" {
		ok, text := s.ScanNext()
		if ok {
			return s.matchToken(text)
		}
		return false // end EOF
	}

	s.err = errorx.Rawf("match tokens fail. text: %s; at line %d", text, s.line)
	return false
}

// ScanNext advance and fetch next line text
func (s *TextScanner) ScanNext() (ok bool, text string) {
	if s.in.Scan() {
		s.line++
		return true, s.in.Text()
	}
	return
}

// SetNext text for scan and parse
func (s *TextScanner) SetNext(text string) {
	s.next = text
}

// Token on current
func (s *TextScanner) Token() Token {
	return s.tok
}

// Line on current
func (s *TextScanner) Line() Token {
	return s.tok
}

// Err get
func (s *TextScanner) Err() error {
	return s.err
}
