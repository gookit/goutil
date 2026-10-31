//go:build unix

package fsutil_test

import (
	"os"
	"path/filepath"
	"syscall"
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/x/assert"
)

func TestCopyFileNamedPipeDestination(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "source")
	dst := filepath.Join(dir, "pipe")
	want := []byte{'p', 'i', 'p', 'e', 0, 255}
	assert.NoErr(t, os.WriteFile(src, want, 0600))
	if err := syscall.Mkfifo(dst, 0600); err != nil {
		t.Fatal(err)
	}
	type result struct {
		data []byte
		err  error
	}
	received := make(chan result, 1)
	go func() { data, err := os.ReadFile(dst); received <- result{data, err} }()
	err := fsutil.CopyFile(src, dst)
	got := <-received
	assert.NoErr(t, err)
	assert.NoErr(t, got.err)
	assert.Eq(t, want, got.data)
}
