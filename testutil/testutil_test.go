package testutil_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gookit/goutil/testutil"
	"github.com/stretchr/testify/assert"
)

func TestDiscardStdout(t *testing.T) {
	err := testutil.DiscardStdout()
	fmt.Println("Hello, playground")
	str := testutil.RestoreStdout()

	assert.NoError(t, err)
	assert.Equal(t, "", str)
}

func TestRewriteStdout(t *testing.T) {
	assert.Equal(t, "", testutil.RestoreStdout())

	testutil.RewriteStdout()
	fmt.Println("Hello, playground")
	msg := testutil.RestoreStdout()

	assert.Equal(t, "Hello, playground\n", msg)
}

func TestRewriteStderr(t *testing.T) {
	assert.Equal(t, "", testutil.RestoreStderr())

	testutil.RewriteStderr()
	_, err := fmt.Fprint(os.Stderr, "Hello, playground")
	msg := testutil.RestoreStderr()

	assert.NoError(t, err)
	assert.Equal(t, "Hello, playground", msg)
}
