package fs

import (
	"os"
	"path"
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

// IsAbsPath is abs path.
func IsAbsPath(filename string) bool {
	return path.IsAbs(filename)
}
