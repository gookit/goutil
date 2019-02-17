package jsonutil

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"text/scanner"

	"github.com/json-iterator/go"
)

var parser = jsoniter.ConfigCompatibleWithStandardLibrary

// WriteFile write data to JSON file
func WriteFile(filePath string, data interface{}) error {
	jsonBytes, err := Encode(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, jsonBytes, 0664)
}

// ReadFile Read JSON file data
func ReadFile(filePath string, v interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	return Decode(content, v)
}

// Encode encode data to json bytes. use it instead of json.Marshal
func Encode(v interface{}) ([]byte, error) {
	return parser.Marshal(v)
}

// Decode decode json bytes to data. use it instead of json.Unmarshal
func Decode(json []byte, v interface{}) error {
	return parser.Unmarshal(json, v)
}

// Pretty get pretty JSON string
func Pretty(v interface{}) (string, error) {
	out, err := json.MarshalIndent(v, "", "    ")
	return string(out), err
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
