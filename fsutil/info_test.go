package fsutil_test

import (
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestExpandPath(t *testing.T) {
	path := "~/.kite"

	assert.NotEq(t, path, fsutil.Expand(path))
	assert.NotEq(t, path, fsutil.ExpandPath(path))
	assert.NotEq(t, path, fsutil.ResolvePath(path))

	assert.Eq(t, "", fsutil.Expand(""))
	assert.Eq(t, "/path/to", fsutil.Expand("/path/to"))
}

func TestPathNoExt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"path/to/file.txt", "path/to/file"},
		{"path/to/file", "path/to/file"},
		{"path/to/.hiddenfile", "path/to/"},
		{"path/to/file.tar.gz", "path/to/file.tar"},
		{"", ""},
	}

	for _, test := range tests {
		result := fsutil.PathNoExt(test.input)
		assert.Eq(t, test.expected, result, "input: %s", test.input)
	}
}

func TestNameNoExt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"path/to/file.txt", "file"},
		{"path/to/file", "file"},
		{"path/to/.hiddenfile", ".hiddenfile"},
		{"path/to/file.tar.gz", "file.tar"},
		{"", ""},
	}

	for _, test := range tests {
		result := fsutil.NameNoExt(test.input)
		assert.Eq(t, test.expected, result, "input: %s", test.input)
	}
}
