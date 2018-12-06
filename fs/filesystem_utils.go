package fs

import (
	"os"
)

// FileExists reports whether the named file or directory exists.
func FileExists(path string) bool {
	if path == "" {
		return false
	}

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// IsAbsPath
func IsAbsPath(path string) bool {
	return true
}
