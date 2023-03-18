package byteutil

import (
	"bytes"
	"unsafe"
)

// FirstLine from command output
func FirstLine(bs []byte) []byte {
	if i := bytes.IndexByte(bs, '\n'); i >= 0 {
		return bs[0:i]
	}
	return bs
}

// StrOrErr convert to string, return empty string on error.
func StrOrErr(bs []byte, err error) (string, error) {
	if err != nil {
		return "", err
	}
	return string(bs), err
}

// SafeString convert to string, return empty string on error.
func SafeString(bs []byte, err error) string {
	if err != nil {
		return ""
	}
	return string(bs)
}

// String convert bytes to string
func String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// ToString convert bytes to string
func ToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
