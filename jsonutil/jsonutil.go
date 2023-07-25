// Package jsonutil provide some util functions for quick operate JSON data
package jsonutil

import (
	"bytes"
	"encoding/json"
	"os"
	"regexp"
	"strings"
	"text/scanner"
)

// WriteFile write data to JSON file
func WriteFile(filePath string, data any) error {
	jsonBytes, err := json.Marshal(data)
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

// MustPretty data to JSON string, will panic on error
func MustPretty(v any) string {
	out, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(out)
}

// Mapping src data(map,struct) to dst struct use json tags.
//
// On src, dst both is struct, equivalent to merging two structures (src should be a subset of dsc)
func Mapping(src, dst any) error {
	bts, err := json.Marshal(src)
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
