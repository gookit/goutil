//go:build !windows

package fsutil

import (
	"path"
)

// Realpath returns the shortest path name equivalent to path by purely lexical processing.
func Realpath(pathStr string) string {
	return path.Clean(pathStr)
}
