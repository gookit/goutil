package textscan

import (
	"github.com/gookit/goutil/errorx"
)

// Kind type
type Kind rune

// String name for kind
func (k Kind) String() string {
	return KindString(k)
}

// builtin defined kinds
const (
	TokInvalid Kind = iota
	TokKey
	TokValue
	TokComments
)

var allowKinds = map[Kind]string{
	TokInvalid:  "Invalid",
	TokKey:      "Key",
	TokValue:    "Value",
	TokComments: "Comments",
}

// AddKind to allowKinds
func AddKind(k Kind, name string) {
	if _, ok := allowKinds[k]; ok {
		panic("cannot repeat register exists kind: " + name)
	}
	allowKinds[k] = name
}

// KindString name
func KindString(k Kind) string {
	if name, ok := allowKinds[k]; ok {
		return name
	}
	return "Invalid"
}

// IsKindToken check
func IsKindToken(k Kind, tok Token) bool {
	if tok != nil {
		return tok.Kind() == k
	}
	return false
}

// LiteToken interface
type LiteToken interface {
	Kind() Kind
	Value() string
	IsValid() bool
}

// Token parser
type Token interface {
	LiteToken
	String() string
	// HasMore is multi line values
	HasMore() bool
	// ScanMore scan multi line values
	ScanMore(ts *TextScanner) error
	MergeSame(tok Token) error
}

// BaseToken struct
type BaseToken struct {
	kind  Kind
	value string

	// Offset int // byte offset, starting at 0
	// Line   int // line number, starting at 1
	// Column int // column number, starting at 1 (character count per line)
}

// Kind type
func (t *BaseToken) Kind() Kind {
	return t.kind
}

// IsValid token
func (t *BaseToken) IsValid() bool {
	return t.kind != TokInvalid
}

// Value of token
func (t *BaseToken) Value() string {
	return t.value
}

// String of token
func (t *BaseToken) String() string {
	if t.kind == TokInvalid {
		return "<Invalid>"
	}
	return t.value
}

// EmptyToken struct
type EmptyToken struct {
	BaseToken
}

// NewEmptyToken instance.
func NewEmptyToken() *EmptyToken {
	return &EmptyToken{}
}

// HasMore is multi line values
func (t *EmptyToken) HasMore() bool {
	return false
}

// ScanMore implements
func (t *EmptyToken) ScanMore(ts *TextScanner) error {
	return nil
}

// MergeSame implements
func (t *EmptyToken) MergeSame(tok Token) error {
	return errorx.Raw("cannot merge any token to Invalid token")
}
