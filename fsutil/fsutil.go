package fsutil

import "github.com/mitchellh/go-homedir"

// ExpandPath will parse `~` as user home dir path.
func ExpandPath(path string) string {
	path, _ = homedir.Expand(path)
	return path
}
