package strutil

import "strings"

// StrVal string
type StrVal string

// Set value
func (s *StrVal) Set(val string) error {
	*s = StrVal(val)
	return nil
}

// IsStartWith prefix
func (s StrVal) IsStartWith(sub string) bool {
	return strings.HasPrefix(string(s), sub)
}

// HasPrefix prefix
func (s StrVal) HasPrefix(sub string) bool {
	return strings.HasPrefix(string(s), sub)
}

// IsEndWith suffix
func (s StrVal) IsEndWith(sub string) bool {
	return strings.HasSuffix(string(s), sub)
}

// HasSuffix suffix
func (s StrVal) HasSuffix(sub string) bool {
	return strings.HasSuffix(string(s), sub)
}

// Bytes string to bytes
func (s StrVal) Bytes() []byte {
	return []byte(s)
}

// Val string
func (s StrVal) Val() string {
	return string(s)
}

// StrVal string
func (s StrVal) String() string {
	return string(s)
}

// TrimSpace string
func (s StrVal) TrimSpace() StrVal {
	return StrVal(strings.TrimSpace(string(s)))
}
