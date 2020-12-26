package goutil

import "strings"

// Strings type
type Strings []string

// String to string
func (ss Strings) String() string {
	return strings.Join(ss, ",")
}

// Has sub-string
func (ss Strings) Has(str string) bool {
	for _, s := range ss {
		if s == str {
			return true
		}
	}

	return false
}
