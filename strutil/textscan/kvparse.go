package textscan

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gookit/goutil/comdef"
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/strutil"
)

// define special chars consts
const (
	MultiLineValMarkS = "'''"
	MultiLineValMarkD = `"""`
	MultiLineValMarkH = "<<<" // at start
	MultiLineValMarkQ = "\\"  // at end
	MultiLineCmtEnd   = "*/"
	// VarRefStartChars  = "${"
)

// KeyValueMatcher match key-value token.
type KeyValueMatcher struct {
	// Separator string for split key and value, default is "="
	Separator string
	// MergeComments collect previous comments token to value token.
	// If set as True, on each s.Scan() please notice skip TokComments
	MergeComments bool
	InlineComment bool
}

// Match text line.
func (m *KeyValueMatcher) Match(text string, prev Token) (Token, error) {
	str := strings.TrimSpace(text)
	ln := len(str)
	if ln == 0 {
		return nil, nil
	}

	if m.Separator == "" {
		m.Separator = comdef.EqualStr
	}

	nodes := strutil.SplitNTrimmed(str, m.Separator, 2)
	if len(nodes) != 2 {
		// err := errorx.Rawf("invalid contents %q(should be KEY=VALUE)", str)
		return nil, nil
	}

	key, val := nodes[0], nodes[1]
	if len(key) == 0 {
		return nil, errorx.Rawf("key cannot be empty: %q", str)
	}

	// handle value
	vln := len(val)
	tok := &ValueToken{
		m:   m,
		key: key,
	}

	tok.kind = TokValue

	// collect prev comments token
	if m.MergeComments && IsKindToken(TokComments, prev) {
		tok.comment = prev
	}

	// multi line value ended by \
	if vln > 0 && strings.HasSuffix(val, MultiLineValMarkQ) {
		val = val[:vln-1]
		tok.more = true
		tok.mark = MultiLineValMarkQ
		tok.value = val
		tok.values = []string{val}
		return tok, nil
	}

	if vln > 2 {
		// multi line value start
		hasPfx := strutil.HasOnePrefix(val, []string{MultiLineValMarkD, MultiLineValMarkS})
		if hasPfx {
			tok.more = true
			tok.mark = MultiLineValMarkS
			if val[0] == '"' {
				tok.mark = MultiLineValMarkD
			}

			val = val[3:] + "\n"
			tok.value = val
			tok.values = []string{val}
			return tok, nil
		}

		// clear quotes
		if val[0] == '"' || val[0] == '\'' {
			val = strutil.Unquote(val)
		} else if m.InlineComment {
			// split inline comments
			var comment string
			val, comment = strutil.SplitInlineComment(val)

			if len(comment) > 0 {
				cmt := NewCommentToken(comment)
				// merge comments token
				if tok.comment != nil {
					_ = tok.comment.MergeSame(cmt)
				} else {
					tok.comment = cmt
				}
			}
		}
	}

	tok.value = val
	return tok, nil
}

// DetectEnd for multi line value
func (m *KeyValueMatcher) DetectEnd(mark, text string) (ok bool, val string) {
	str := strings.TrimSpace(text)
	ln := len(str)

	// multi line value
	if mark == MultiLineValMarkS || mark == MultiLineValMarkD {
		if strings.HasSuffix(str, mark) { // end
			val = str[:ln-3]
			ok = true
		} else {
			val = str
		}
	} else if mark == MultiLineValMarkQ {
		if strings.HasSuffix(str, MultiLineValMarkQ) { // go on
			val += str[:ln-1]
		} else { // end
			val = str
			ok = true
		}
	}

	return
}

// ValueToken struct
type ValueToken struct {
	BaseToken
	m *KeyValueMatcher

	more bool
	mark string // end mark for multi line

	// key for token
	key string
	// for multi line value.
	values []string
	// comment for the item
	comment Token
}

// Key name
func (t *ValueToken) Key() string {
	return t.key
}

// Comment lines string
func (t *ValueToken) Comment() string {
	if t.comment != nil {
		return t.comment.Value()
	}
	return ""
}

// Value text string.
func (t *ValueToken) Value() string {
	if len(t.values) > 0 {
		return strings.Join(t.values, "\n")
	}
	return t.value
}

// HasMore is multi line values
func (t *ValueToken) HasMore() bool {
	return t.more
}

// MergeSame comments token
func (t *ValueToken) MergeSame(_ Token) error {
	return errors.New("merge value token not allowed")
}

// String of token
func (t *ValueToken) String() string {
	return fmt.Sprintf("key: %s\nvalue: %q\ncomments: %s", t.key, t.Value(), t.Comment())
}

// ErrMLineValueNotEnd error
var ErrMLineValueNotEnd = errors.New("not end of multi line value")

// ScanMore scan multi line values
func (t *ValueToken) ScanMore(ts *TextScanner) error {
	for {
		ok, line := ts.ScanNext()
		if !ok {
			return ErrMLineValueNotEnd
		}

		// detect value end line
		if ok, val := t.m.DetectEnd(t.mark, line); ok {
			t.values = append(t.values, val)
			return nil
		}

		t.values = append(t.values, line)
	}
}

// CommentsMatcher struct
type CommentsMatcher struct {
	// MatchFn for comments line
	// - mark 	useful on multi line comments
	MatchFn func(text string) (ok, more bool, err error)
	// DetectEnd for multi line comments
	DetectEnd func(text string) bool
}

// Match comments token
func (m *CommentsMatcher) Match(text string, prev Token) (Token, error) {
	if m.MatchFn == nil {
		m.MatchFn = CommentsDetect
	}

	ok, more, err := m.MatchFn(text)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	tok := &CommentToken{m: m}
	tok.more = more
	tok.kind = TokComments
	tok.value = text

	if more {
		tok.comments = []string{text}
	}

	if IsKindToken(TokComments, prev) {
		if err := tok.MergeSame(prev); err != nil {
			return nil, err
		}
	}

	return tok, nil
}

// CommentsDetect check.
func CommentsDetect(text string) (ok, more bool, err error) {
	str := strings.TrimSpace(text)
	ln := len(str)
	if ln == 0 {
		return
	}

	// a line comments
	if str[0] == '#' || str[0] == '!' {
		ok = true
		return
	}

	if str[0] == '/' {
		if ln < 2 {
			err = errorx.Rawf("invalid contents %q", str)
			return
		}

		if str[1] == '/' {
			ok = true
			return
		}

		// multi line comments start
		if str[1] == '*' {
			ok = true
			more = true

			// end at line
			if strings.HasSuffix(str, MultiLineCmtEnd) {
				more = false
			}
		}
	}
	return
}

// MatchEnd for multi line comments
func (m *CommentsMatcher) MatchEnd(text string) bool {
	if m.DetectEnd == nil {
		m.DetectEnd = CommentsDetectEnd
	}
	return m.DetectEnd(text)
}

// CommentToken struct
type CommentToken struct {
	BaseToken
	m *CommentsMatcher

	more bool
	mark string // end mark for multi line

	// for multi line comments.
	comments []string
}

// NewCommentToken instance.
func NewCommentToken(val string) *CommentToken {
	tok := &CommentToken{}
	tok.value = val
	tok.kind = TokComments
	return tok
}

func (t *CommentToken) Value() string {
	if len(t.comments) > 0 {
		return strings.Join(t.comments, "\n")
	}
	return t.value
}

func (t *CommentToken) String() string {
	return t.Value()
}

// MergeSame comments token
func (t *CommentToken) MergeSame(tok Token) error {
	if tok == nil {
		return nil
	}

	if tok.Kind() == t.Kind() {
		t.more = true
		t.comments = append(t.comments, tok.Value())
		return nil
	}
	return errorx.Rawf("cannot merge %s token to an Comments token", tok.Kind())
}

// HasMore is multi line values
func (t *CommentToken) HasMore() bool {
	return t.more
}

// ErrCommentsNotEnd error
var ErrCommentsNotEnd = errors.New("not end of multi-line comments")

// ScanMore scan multi line values
func (t *CommentToken) ScanMore(ts *TextScanner) error {
	if t.m.DetectEnd == nil {
		t.m.DetectEnd = CommentsDetectEnd
	}

	for {
		ok, line := ts.ScanNext()
		if !ok {
			return ErrCommentsNotEnd
		}

		t.comments = append(t.comments, line)

		// detect comments end line
		if t.m.DetectEnd(line) {
			return nil
		}
	}
}

func CommentsDetectEnd(line string) bool {
	return strings.HasSuffix(line, MultiLineCmtEnd)
}
