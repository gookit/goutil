package strutil

import (
	"bytes"
	"crypto/md5"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
	"text/template"
)

//
// -------------------- escape --------------------
//

// EscapeJS escape javascript string
func EscapeJS(s string) string {
	return template.JSEscapeString(s)
}

// EscapeHTML escape html string
func EscapeHTML(s string) string {
	return template.HTMLEscapeString(s)
}

// AddSlashes add slashes for the string.
func AddSlashes(s string) string {
	if ln := len(s); ln == 0 {
		return ""
	}

	var buf bytes.Buffer
	for _, char := range s {
		switch char {
		case '\'', '"', '\\':
			buf.WriteRune('\\')
		}
		buf.WriteRune(char)
	}

	return buf.String()
}

// StripSlashes strip slashes for the string.
func StripSlashes(s string) string {
	ln := len(s)
	if ln == 0 {
		return ""
	}

	var skip bool
	var buf bytes.Buffer

	for i, char := range s {
		if skip {
			skip = false
		} else if char == '\\' {
			if i+1 < ln && s[i+1] == '\\' {
				skip = true
			}
			continue
		}
		buf.WriteRune(char)
	}

	return buf.String()
}

//
// -------------------- encode --------------------
//

// Md5 Generate a 32-bit md5 string
func Md5(src interface{}) string { return GenMd5(src) }

// MD5 Generate a 32-bit md5 string
func MD5(src interface{}) string { return GenMd5(src) }

// GenMd5 Generate a 32-bit md5 string
func GenMd5(src interface{}) string {
	h := md5.New()
	if s, ok := src.(string); ok {
		h.Write([]byte(s))
	} else {
		h.Write([]byte(fmt.Sprint(src)))
	}

	return hex.EncodeToString(h.Sum(nil))
}

// URLEncode encode url string.
func URLEncode(s string) string {
	if pos := strings.IndexRune(s, '?'); pos > -1 { // escape query data
		return s[0:pos+1] + url.QueryEscape(s[pos+1:])
	}
	return s
}

// URLDecode decode url string.
func URLDecode(s string) string {
	if pos := strings.IndexRune(s, '?'); pos > -1 { // un-escape query data
		qy, err := url.QueryUnescape(s[pos+1:])
		if err == nil {
			return s[0:pos+1] + qy
		}
	}

	return s
}

//
// -------------------- base encode --------------------
//

// B32Encode base32 encode
func B32Encode(str string) string {
	return base32.StdEncoding.EncodeToString([]byte(str))
}

// B32Decode base32 decode
func B32Decode(str string) string {
	dec, _ := base32.StdEncoding.DecodeString(str)
	return string(dec)
}

// Base64 encode
func Base64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// B64Encode base64 encode
func B64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// B64Decode base64 decode
func B64Decode(str string) string {
	dec, _ := base64.StdEncoding.DecodeString(str)
	return string(dec)
}

// BaseEncoder struct
type BaseEncoder struct {
	// Base value
	Base int
}

// NewBaseEncoder instance
func NewBaseEncoder(base int) *BaseEncoder {
	return &BaseEncoder{Base: base}
}

// Encode handle
func (be *BaseEncoder) Encode(s string) string {
	return s
}

// Decode handle
func (be *BaseEncoder) Decode(s string) (string, error) {
	return s, nil
}
