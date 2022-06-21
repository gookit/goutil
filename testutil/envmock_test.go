package testutil_test

import (
	"os"
	"testing"

	"github.com/gookit/goutil/envutil"
	"github.com/gookit/goutil/testutil"
	"github.com/stretchr/testify/assert"
)

func TestMockEnvValue(t *testing.T) {
	is := assert.New(t)
	is.Equal("", os.Getenv("APP_COMMAND"))

	testutil.MockEnvValue("APP_COMMAND", "new val", func(nv string) {
		is.Equal("new val", nv)
	})

	is.Equal("", os.Getenv("APP_COMMAND"))
}

func TestMockEnvValues(t *testing.T) {
	is := assert.New(t)
	is.Equal("", os.Getenv("APP_COMMAND"))

	testutil.MockEnvValues(map[string]string{
		"APP_COMMAND": "new val",
	}, func() {
		is.Equal("new val", os.Getenv("APP_COMMAND"))
	})

	is.Equal("", os.Getenv("APP_COMMAND"))
}

func TestMockOsEnvByText(t *testing.T) {
	envStr := `
APP_COMMAND = login
APP_ENV = dev
APP_DEBUG = true
`

	testutil.MockOsEnvByText(envStr, func() {
		assert.Len(t, envutil.Environ(), 3)
		assert.True(t, envutil.GetBool("APP_DEBUG"))
		assert.Equal(t, "login", envutil.Getenv("APP_COMMAND"))
	})
}
