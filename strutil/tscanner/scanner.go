// Package tscanner text-scanner
package tscanner

import (
	"io"
	"text/scanner"
)

const (
	TokInvalid = iota
	TokComments
	TokILComments
	TokMLComments
	TokValue
	TokMLValue
)

const (
	Value = iota
	MValue
)

type tokenItem struct {
	// see TokValueLine
	kind rune
	// key string. eg: top.sub.some-key
	key string

	// token value
	value string
	// for multi line value.
	values []string
	// for multi line comments.
	comments []string
}

func newTokenItem(key, value string, kind rune) *tokenItem {
	tk := &tokenItem{
		key:   key,
		kind:  kind,
		value: value,
	}

	return tk
}

// Valid of the token data.
func (ti *tokenItem) addValue(val string) {
	ti.values = append(ti.values, val)
}

const bufLen = 1024 // at least utf8.UTFMax

// TextScanner struct.
// refer text/scanner.Scanner
type TextScanner struct {
	scanner.Position
	// Input
	src io.Reader

	// k-v split char, default is "="
	kvSep string
}

func (s TextScanner) Next() *tokenItem {

}
