package testutil_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestDiscardStdout(t *testing.T) {
	err := testutil.DiscardStdout()
	fmt.Println("Hello, playground")
	str := testutil.RestoreStdout()

	assert.NoErr(t, err)
	assert.Eq(t, "", str)
}

func TestRewriteStdout(t *testing.T) {
	testutil.RewriteStdout()

	assert.Eq(t, "", testutil.RestoreStdout())

	testutil.RewriteStdout()
	fmt.Println("Hello, playground")
	msg := testutil.RestoreStdout()

	assert.Eq(t, "Hello, playground\n", msg)
}

func TestRewriteStderr(t *testing.T) {
	testutil.RewriteStderr()
	assert.Eq(t, "", testutil.RestoreStderr())

	testutil.RewriteStderr()
	_, err := fmt.Fprint(os.Stderr, "Hello, playground")
	msg := testutil.RestoreStderr()

	assert.NoErr(t, err)
	assert.Eq(t, "Hello, playground", msg)
}

func TestNewDirEnt(t *testing.T) {
	de := testutil.NewDirEnt("testdata/some.txt")
	assert.NotEmpty(t, de)
	assert.False(t, de.IsDir())
}

// test testutil.SetTimeLocal
func TestSetTimeLocal(t *testing.T) {
	testutil.SetTimeLocalUTC()
	tt, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-01-01 12:12:12", time.Local)
	assert.NoError(t, err)
	assert.Eq(t, "2021-01-01 12:12:12", tt.Format("2006-01-02 15:04:05"))
	testutil.RestoreTimeLocal()
}
