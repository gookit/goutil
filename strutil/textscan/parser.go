package textscan

import (
	"bytes"
	"io"
	"strings"
)

// HandleFn for token
type HandleFn func(t Token)

// Parser struct
type Parser struct {
	ts *TextScanner
	// Func for handle tokens
	Func HandleFn
}

func NewParser(fn HandleFn) *Parser {
	return &Parser{
		Func: fn,
		ts:   &TextScanner{},
	}
}

// AddMatchers register token matchers
func (p *Parser) AddMatchers(ms ...Matcher) {
	p.ts.AddMatchers(ms...)
}

func (p *Parser) Parse(bs []byte) error {
	return p.ParseFrom(bytes.NewReader(bs))
}

func (p *Parser) ParseText(text string) error {
	return p.ParseFrom(strings.NewReader(text))
}

func (p *Parser) ParseFrom(r io.Reader) error {
	ts := NewScanner(r)

	for ts.Scan() {
		p.Func(ts.Token())
	}

	return nil
}
