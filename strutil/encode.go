package strutil

import (
	"encoding/base32"
	"encoding/base64"
	"net/url"
	"strings"
	"text/template"
)

// EscapeJS escape javascript string
func EscapeJS(s string) string {
	return template.JSEscapeString(s)
}

// EscapeHTML escape html string
func EscapeHTML(s string) string {
	return template.HTMLEscapeString(s)
}

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
