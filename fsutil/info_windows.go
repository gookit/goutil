package fsutil

import "path/filepath"

// Realpath returns the shortest path name equivalent to path by purely lexical processing.
func Realpath(pathStr string) string {
	return filepath.Clean(pathStr)
}
