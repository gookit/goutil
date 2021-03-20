package fsutil

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/mitchellh/go-homedir"
)

const (
	// sniff Length, use for detect file mime type
	MimeSniffLen = 512
)

// OSTempFile create an temp file on os.TempDir()
// Usage:
// 	fsutil.OSTempFile("example.*.txt")
func OSTempFile(pattern string) (*os.File, error) {
	return ioutil.TempFile(os.TempDir(), pattern)
}

// TempFile is alias of ioutil.TempFile
// Usage:
// 	fsutil.TempFile("", "example.*.txt")
func TempFile(dir, pattern string) (*os.File, error) {
	return ioutil.TempFile(dir, pattern)
}

// OSTempDir creates a new temp dir on os.TempDir and return the temp dir path
// Usage:
// 	fsutil.OSTempDir("example.*.txt")
func OSTempDir(pattern string) (string, error) {
	return ioutil.TempDir(os.TempDir(), pattern)
}

// TempDir creates a new temp dir and return the temp dir path
// Usage:
// 	fsutil.TempDir("", "example.*.txt")
func TempDir(dir, pattern string) (string, error) {
	return ioutil.TempDir(dir, pattern)
}

// ExpandPath will parse `~` as user home dir path.
func ExpandPath(path string) string {
	path, _ = homedir.Expand(path)
	return path
}

// MimeType get File Mime Type name. eg "image/png"
func MimeType(path string) (mime string) {
	if path == "" {
		return
	}

	file, err := os.Open(path)
	if err != nil {
		return
	}

	return ReaderMimeType(file)
}

// ReaderMimeType get the io.Reader mimeType
// Usage:
// 	file, err := os.Open(filepath)
// 	if err != nil {
// 		return
// 	}
//	mime := ReaderMimeType(file)
func ReaderMimeType(r io.Reader) (mime string) {
	var buf [MimeSniffLen]byte
	n, _ := io.ReadFull(r, buf[:])
	if n == 0 {
		return ""
	}

	return http.DetectContentType(buf[:n])
}
