// Package jsonutil provide some util functions for quick operate JSON data
package jsonutil

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"regexp"
	"strings"
	"text/scanner"
)

// WriteFile write data to JSON file
func WriteFile(filePath string, data any) error {
	jsonBytes, err := Encode(data)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, jsonBytes, 0664)
}

// WritePretty write pretty data to JSON file
func WritePretty(filePath string, data any) error {
	bs, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, bs, 0664)
}

// ReadFile Read JSON file data
func ReadFile(filePath string, v any) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()
	return json.NewDecoder(file).Decode(v)
}

// Pretty JSON string and return
func Pretty(v any) (string, error) {
	out, err := json.MarshalIndent(v, "", "    ")
	return string(out), err
}

// Encode data to json bytes.
func Encode(v any) ([]byte, error) {
	return json.Marshal(v)
}

// EncodePretty encode pretty JSON data to json bytes.
func EncodePretty(v any) ([]byte, error) {
	return json.MarshalIndent(v, "", "    ")
}

// EncodeToWriter encode data to writer.
func EncodeToWriter(v any, w io.Writer) error {
	return json.NewEncoder(w).Encode(v)
}

// EncodeUnescapeHTML data to json bytes. will close escape HTML
func EncodeUnescapeHTML(v any) ([]byte, error) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)

	if err := enc.Encode(v); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Decode json bytes to data ptr.
func Decode(bts []byte, ptr any) error {
	return json.Unmarshal(bts, ptr)
}

// DecodeString json string to data ptr.
func DecodeString(str string, ptr any) error {
	return json.Unmarshal([]byte(str), ptr)
}

// DecodeReader decode JSON from io reader.
func DecodeReader(r io.Reader, ptr any) error {
	return json.NewDecoder(r).Decode(ptr)
}

// Mapping src data(map,struct) to dst struct use json tags.
//
// On src, dst both is struct, equivalent to merging two structures (src should be a subset of dsc)
func Mapping(src, dst any) error {
	bts, err := Encode(src)
	if err != nil {
		return err
	}
	return Decode(bts, dst)
}

// IsJSON check if the string is valid JSON. (Note: uses json.Unmarshal)
func IsJSON(s string) bool {
	if s == "" {
		return false
	}

	var js json.RawMessage
	return json.Unmarshal([]byte(s), &js) == nil
}

// IsJSONFast simple and fast check input string is valid JSON.
func IsJSONFast(s string) bool {
	ln := len(s)
	if ln < 2 {
		return false
	}
	if ln == 2 {
		return s == "{}" || s == "[]"
	}

	// object
	if s[0] == '{' {
		return s[ln-1] == '}' && s[1] == '"'
	}

	// array
	return s[0] == '[' && s[ln-1] == ']'
}

// `(?s:` enable match multi line
var jsonMLComments = regexp.MustCompile(`(?s:/\*.*?\*/\s*)`)

// StripComments strip comments for a JSON string
func StripComments(src string) string {
	// multi line comments
	if strings.Contains(src, "/*") {
		src = jsonMLComments.ReplaceAllString(src, "")
	}

	// single line comments
	if !strings.Contains(src, "//") {
		return strings.TrimSpace(src)
	}

	// strip inline comments
	var s scanner.Scanner

	s.Init(strings.NewReader(src))
	s.Filename = "comments"
	s.Mode ^= scanner.SkipComments // don't skip comments

	buf := new(bytes.Buffer)
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		txt := s.TokenText()
		if !strings.HasPrefix(txt, "//") && !strings.HasPrefix(txt, "/*") {
			buf.WriteString(txt)
			// } else {
			// fmt.Printf("%s: %s\n", s.Position, txt)
		}
	}

	return buf.String()
}
