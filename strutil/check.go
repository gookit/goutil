package strutil

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Equal check
var Equal = strings.EqualFold

// IsNumeric returns true if the given character is a numeric, otherwise false.
func IsNumeric(c byte) bool {
	return c >= '0' && c <= '9'
}

// IsAlphabet char
func IsAlphabet(char uint8) bool {
	// A 65 -> Z 90
	if char >= 'A' && char <= 'Z' {
		return true
	}

	// a 97 -> z 122
	if char >= 'a' && char <= 'z' {
		return true
	}

	return false
}

// IsAlphaNum reports whether the byte is an ASCII letter, number, or underscore
func IsAlphaNum(c uint8) bool {
	return c == '_' || '0' <= c && c <= '9' || 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
}

// StrPos alias of the strings.Index
func StrPos(s, sub string) int {
	return strings.Index(s, sub)
}

// BytePos alias of the strings.IndexByte
func BytePos(s string, bt byte) int {
	return strings.IndexByte(s, bt)
}

// RunePos alias of the strings.IndexRune
func RunePos(s string, ru rune) int {
	return strings.IndexRune(s, ru)
}

// HasOneSub substr in the given string.
func HasOneSub(s string, subs []string) bool {
	for _, sub := range subs {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}

// HasAllSubs all substr in the given string.
func HasAllSubs(s string, subs []string) bool {
	for _, sub := range subs {
		if !strings.Contains(s, sub) {
			return false
		}
	}
	return true
}

// IsStartsOf alias of the HasOnePrefix
func IsStartsOf(s string, prefixes []string) bool {
	return HasOnePrefix(s, prefixes)
}

// HasOnePrefix the string start withs one of the subs
func HasOnePrefix(s string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

// HasPrefix substr in the given string.
func HasPrefix(s string, prefix string) bool { return strings.HasPrefix(s, prefix) }

// IsStartOf alias of the strings.HasPrefix
func IsStartOf(s, prefix string) bool { return strings.HasPrefix(s, prefix) }

// HasSuffix substr in the given string.
func HasSuffix(s string, suffix string) bool { return strings.HasSuffix(s, suffix) }

// IsEndOf alias of the strings.HasSuffix
func IsEndOf(s, suffix string) bool { return strings.HasSuffix(s, suffix) }

// Len of the string
func Len(s string) int { return len(s) }

// RuneLen of the string
func RuneLen(s string) int { return len([]rune(s)) }

// Utf8Len of the string
func Utf8Len(s string) int { return utf8.RuneCountInString(s) }

// Utf8len of the string
func Utf8len(s string) int { return utf8.RuneCountInString(s) }

// IsValidUtf8 valid utf8 string check
func IsValidUtf8(s string) bool { return utf8.ValidString(s) }

// ----- refer from github.com/yuin/goldmark/util

// refer from github.com/yuin/goldmark/util
var spaceTable = [256]int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// IsSpace returns true if the given character is a space, otherwise false.
func IsSpace(c byte) bool {
	return spaceTable[c] == 1
}

// IsSpaceRune returns true if the given rune is a space, otherwise false.
func IsSpaceRune(r rune) bool {
	return r <= 256 && IsSpace(byte(r)) || unicode.IsSpace(r)
}

// IsEmpty returns true if the given string is empty.
func IsEmpty(s string) bool { return len(s) == 0 }

// IsBlank returns true if the given string is all space characters.
func IsBlank(s string) bool {
	return IsBlankBytes([]byte(s))
}

// IsNotBlank returns true if the given string is not blank.
func IsNotBlank(s string) bool {
	return !IsBlankBytes([]byte(s))
}

// IsBlankBytes returns true if the given []byte is all space characters.
func IsBlankBytes(bs []byte) bool {
	for _, b := range bs {
		if !IsSpace(b) {
			return false
		}
	}
	return true
}

// IsSymbol reports whether the rune is a symbolic character.
func IsSymbol(r rune) bool {
	return unicode.IsSymbol(r)
}

// VersionCompare for two version string.
func VersionCompare(v1, v2, op string) bool {
	switch op {
	case ">":
		return v1 > v2
	case "<":
		return v1 < v2
	case ">=":
		return v1 >= v2
	case "<=":
		return v1 <= v2
	default:
		return v1 == v2
	}
}
