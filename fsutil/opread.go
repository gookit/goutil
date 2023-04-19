package fsutil

import (
	"bufio"
	"errors"
	"io"
	"os"
	"text/scanner"
)

// NewIOReader instance by input file path or io.Reader
func NewIOReader(in any) (r io.Reader, err error) {
	switch typIn := in.(type) {
	case string: // as file path
		return OpenReadFile(typIn)
	case io.Reader:
		return typIn, nil
	}
	return nil, errors.New("invalid input type, allow: string, io.Reader")
}

// DiscardReader anything from the reader
func DiscardReader(src io.Reader) {
	_, _ = io.Copy(io.Discard, src)
}

// ReadFile read file contents, will panic on error
func ReadFile(filePath string) []byte {
	return MustReadFile(filePath)
}

// MustReadFile read file contents, will panic on error
func MustReadFile(filePath string) []byte {
	bs, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return bs
}

// ReadReader read contents from io.Reader, will panic on error
func ReadReader(r io.Reader) []byte { return MustReadReader(r) }

// MustReadReader read contents from io.Reader, will panic on error
func MustReadReader(r io.Reader) []byte {
	bs, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return bs
}

// ReadString read contents from path or io.Reader, will panic on in type error
func ReadString(in any) string {
	return string(GetContents(in))
}

// ReadStringOrErr read contents from path or io.Reader, will panic on in type error
func ReadStringOrErr(in any) (string, error) {
	r, err := NewIOReader(in)
	if err != nil {
		return "", err
	}

	bs, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

// ReadAll read contents from path or io.Reader, will panic on in type error
func ReadAll(in any) []byte { return GetContents(in) }

// GetContents read contents from path or io.Reader, will panic on in type error
func GetContents(in any) []byte {
	r, err := NewIOReader(in)
	if err != nil {
		panic(err)
	}
	return MustReadReader(r)
}

// ReadOrErr read contents from path or io.Reader, will panic on in type error
func ReadOrErr(in any) ([]byte, error) {
	r, err := NewIOReader(in)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(r)
}

// ReadExistFile read file contents if existed, will panic on error
func ReadExistFile(filePath string) []byte {
	if IsFile(filePath) {
		bs, err := os.ReadFile(filePath)
		if err != nil {
			panic(err)
		}
		return bs
	}
	return nil
}

// TextScanner from filepath or io.Reader, will panic on in type error
//
// Usage:
//
//	s := fsutil.TextScanner("/path/to/file")
//	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
//		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
//	}
func TextScanner(in any) *scanner.Scanner {
	var s scanner.Scanner
	r, err := NewIOReader(in)
	if err != nil {
		panic(err)
	}

	s.Init(r)
	s.Filename = "text-scanner"
	return &s
}

// LineScanner create from filepath or io.Reader
//
//	s := fsutil.LineScanner("/path/to/file")
//	for s.Scan() {
//		fmt.Println(s.Text())
//	}
func LineScanner(in any) *bufio.Scanner {
	r, err := NewIOReader(in)
	if err != nil {
		panic(err)
	}
	return bufio.NewScanner(r)
}
