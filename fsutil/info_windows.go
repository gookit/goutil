package fsutil

import (
	"github.com/gookit/goutil/internal/comfunc"
	"github.com/gookit/goutil/sysutil"
)

// Realpath returns the shortest path name equivalent to path by purely lexical processing.
func Realpath(pathStr string) string {
	pathStr = comfunc.ExpandHome(pathStr)

	if !IsAbsPath(pathStr) {
		pathStr = JoinSubPaths(sysutil.Workdir(), pathStr)
	}
	return filepath.Clean(pathStr)
}
