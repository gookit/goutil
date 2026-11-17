package strutil

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

/*************************************************************
 * change string case
 *************************************************************/

// methods aliases
var (
	UpWords = UpperWord
	LoFirst = LowerFirst
	UpFirst = UpperFirst

	Snake = SnakeCase
)

// Title alias of the strings.ToTitle()
func Title(s string) string { return strings.ToTitle(s) }

// Lower alias of the strings.ToLower()
func Lower(s string) string { return strings.ToLower(s) }

// Lowercase alias of the strings.ToLower()
func Lowercase(s string) string { return strings.ToLower(s) }

// Upper alias of the strings.ToUpper()
func Upper(s string) string { return strings.ToUpper(s) }

// Uppercase alias of the strings.ToUpper()
func Uppercase(s string) string { return strings.ToUpper(s) }

// UpperWord Change the first character of each word to uppercase
func UpperWord(s string) string {
	if len(s) == 0 {
		return s
	}

	if len(s) == 1 {
		return strings.ToUpper(s)
	}

	inWord := true
	buf := make([]byte, 0, len(s))

	i := 0
	rs := []rune(s)
	if RuneIsLower(rs[i]) {
		buf = append(buf, []byte(string(unicode.ToUpper(rs[i])))...)
	} else {
		buf = append(buf, []byte(string(rs[i]))...)
	}

	for j := i + 1; j < len(rs); j++ {
		if !RuneIsWord(rs[i]) && RuneIsWord(rs[j]) {
			inWord = false
		}

		if RuneIsLower(rs[j]) && !inWord {
			buf = append(buf, []byte(string(unicode.ToUpper(rs[j])))...)
			inWord = true
		} else {
			buf = append(buf, []byte(string(rs[j]))...)
		}

		if RuneIsWord(rs[j]) {
			inWord = true
		}

		i++
	}

	return string(buf)
}

// LowerFirst lower first char
func LowerFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	rs := []rune(s)
	f := rs[0]

	if 'A' <= f && f <= 'Z' {
		return string(unicode.ToLower(f)) + string(rs[1:])
	}
	return s
}

// UpperFirst upper first char
func UpperFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	rs := []rune(s)
	f := rs[0]

	if 'a' <= f && f <= 'z' {
		return string(unicode.ToUpper(f)) + string(rs[1:])
	}
	return s
}

// snakeAcronyms are the initialisms that SnakeCase keeps in one word, so that
// "IDs" gives "ids" and "IPv4" gives "ipv4" instead of "i_ds" and "i_pv4".
//
// NOTE: "https" is not listed on purpose: the acronym run of "HTTPServer" is
// "HTTPS", which must still split to "http_server".
var snakeAcronyms = map[string]struct{}{
	"acl": {}, "api": {}, "ascii": {}, "cpu": {}, "css": {}, "dns": {}, "eof": {},
	"guid": {}, "html": {}, "http": {}, "id": {}, "ip": {}, "json": {}, "lhs": {},
	"qps": {}, "ram": {}, "rhs": {}, "rpc": {}, "sla": {}, "smtp": {}, "sql": {},
	"ssh": {}, "tcp": {}, "tls": {}, "ttl": {}, "udp": {}, "ui": {}, "uid": {},
	"uri": {}, "url": {}, "utf8": {}, "uuid": {}, "vm": {}, "xml": {}, "xmpp": {},
	"xsrf": {}, "xss": {},
}

// SnakeCase convert. eg "RangePrice" -> "range_price"
func SnakeCase(s string, sep ...string) string {
	sepChar := "_"
	if len(sep) > 0 {
		sepChar = sep[0]
	}

	// Split before an upper case letter that follows a lower case letter or digit, or that
	// starts a word after an acronym: "UserID" -> "user_id", "HTTPServer" -> "http_server".
	rs := []rune(s)
	var sb strings.Builder
	for i, r := range rs {
		if unicode.IsUpper(r) && i > 0 {
			prev := rs[i-1]
			nextLower := i+1 < len(rs) && unicode.IsLower(rs[i+1])
			if unicode.IsLower(prev) || unicode.IsDigit(prev) ||
				(nextLower && !isSnakeAcronymTail(rs, i) && string(prev) != sepChar) {
				sb.WriteString(sepChar)
			}
		}
		sb.WriteRune(unicode.ToLower(r))
	}

	return strings.TrimLeft(sb.String(), sepChar)
}

// isSnakeAcronymTail reports whether the upper case run that ends at rs[i] is a known
// acronym followed by a plural "s" ("IDs", "IDsList") or a version suffix ("IPv4").
// Such a run is one word, so no separator goes in front of its last letter.
func isSnakeAcronymTail(rs []rune, i int) bool {
	start := i
	for start > 0 && unicode.IsUpper(rs[start-1]) {
		start--
	}
	if _, ok := snakeAcronyms[strings.ToLower(string(rs[start:i+1]))]; !ok {
		return false
	}

	// rs[i+1] exists: the caller only asks when the next rune is lower case
	rest := rs[i+1:]
	if rest[0] == 's' && (len(rest) == 1 || unicode.IsUpper(rest[1])) {
		return true // plural of the acronym
	}
	return rest[0] == 'v' && len(rest) > 1 && unicode.IsDigit(rest[1])
}

// Camel alias of the CamelCase
func Camel(s string, sep ...string) string { return CamelCase(s, sep...) }

// CamelCase convert string to camel case.
//
// Support:
//
//	"range_price" -> "rangePrice"
//	"range price" -> "rangePrice"
//	"range-price" -> "rangePrice"
func CamelCase(s string, sep ...string) string {
	sepChar := "_"
	if len(sep) > 0 {
		sepChar = sep[0]
	}

	// Not contains sep char
	if !strings.Contains(s, sepChar) {
		return s
	}

	// Get regexp instance
	rgx, ok := toCamelRegs[sepChar]
	if !ok {
		rgx = regexp.MustCompile(regexp.QuoteMeta(sepChar) + "+[a-zA-Z]")
	}

	return rgx.ReplaceAllStringFunc(s, func(s string) string {
		s = strings.TrimLeft(s, sepChar)
		return UpperFirst(s)
	})
}

//
// Indent format multi line text
// from package: github.com/kr/text
//

// Indent inserts prefix at the beginning of each non-empty line of s. The
// end-of-line marker is NL.
func Indent(s, prefix string) string {
	return string(IndentBytes([]byte(s), []byte(prefix)))
}

// IndentBytes inserts prefix at the beginning of each non-empty line of b.
// The end-of-line marker is NL.
func IndentBytes(b, prefix []byte) []byte {
	if len(b) == 0 {
		return b
	}

	bol := true
	res := make([]byte, 0, len(b)+len(prefix)*4)

	for _, c := range b {
		if bol && c != '\n' {
			res = append(res, prefix...)
		}
		res = append(res, c)
		bol = c == '\n'
	}
	return res
}

// Replaces replace multi strings
//
//	pairs: {old1: new1, old2: new2, ...}
//
// Can also use:
//
//	strings.NewReplacer("old1", "new1", "old2", "new2").Replace(str)
func Replaces(str string, pairs map[string]string) string {
	return NewReplacer(pairs).Replace(str)
}

// ReplaceVars replaces simple variables in a string. format: {varName}
//
// Usage:
//
//	strutil.ReplaceVars("{name}, age is {age}", map[string]string{
//		"name": "Joe",
//		"age": "18"
//	})
func ReplaceVars(s string, vars map[string]string) string {
	if !ContainsByte(s, '{') {
		return s
	}

	// format var name to {name}
	pairs := make(map[string]string)
	for k, v := range vars {
		vName := "{" + k + "}"
		pairs[vName] = v
	}
	return NewReplacer(pairs).Replace(s)
}

// NewReplacer instance
func NewReplacer(pairs map[string]string) *strings.Replacer {
	ss := make([]string, 0, len(pairs)*2)
	for old, newVal := range pairs {
		ss = append(ss, old, newVal)
	}
	return strings.NewReplacer(ss...)
}

// WrapTag for given string.
func WrapTag(s, tag string) string {
	if s == "" {
		return s
	}
	return fmt.Sprintf("<%s>%s</%s>", tag, s, tag)
}
