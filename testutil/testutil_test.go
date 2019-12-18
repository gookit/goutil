package testutil_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gookit/goutil/testutil"
	"github.com/stretchr/testify/assert"
)

func TestRewriteStdout(t *testing.T) {
	testutil.RewriteStdout()
	fmt.Println("Hello, playground")
	msg := testutil.RestoreStdout()

	assert.Equal(t, "Hello, playground\n", msg)
}

func TestMockEnvValue(t *testing.T) {
	ris := assert.New(t)
	ris.Equal("", os.Getenv("APP_COMMAND"))

	testutil.MockEnvValue("APP_COMMAND", "new val", func(nv string) {
		ris.Equal("new val", nv)
	})

	ris.Equal("", os.Getenv("APP_COMMAND"))
}

func TestMockEnvValues(t *testing.T) {
	ris := assert.New(t)
	ris.Equal("", os.Getenv("APP_COMMAND"))

	testutil.MockEnvValues(map[string]string{
		"APP_COMMAND": "new val",
	}, func() {
		ris.Equal("new val", os.Getenv("APP_COMMAND"))
	})

	ris.Equal("", os.Getenv("APP_COMMAND"))
}
