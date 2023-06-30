package testutil_test

import (
	"os"
	"testing"

	"github.com/gookit/goutil/internal/comfunc"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestMockEnvValue(t *testing.T) {
	is := assert.New(t)
	is.Eq("", os.Getenv("APP_COMMAND"))

	testutil.MockEnvValue("APP_COMMAND", "new val", func(nv string) {
		is.Eq("new val", nv)
	})

	shellVal := "custom-value"
	testutil.MockEnvValue("SHELL", shellVal, func(newVal string) {
		is.Eq(shellVal, newVal)
	})

	is.Eq("", os.Getenv("APP_COMMAND"))
	is.Panics(func() {
		testutil.MockEnvValue("invalid=", "value", nil)
	})
}

func TestMockEnvValues(t *testing.T) {
	is := assert.New(t)
	is.Eq("", os.Getenv("APP_COMMAND"))

	testutil.MockEnvValues(map[string]string{
		"APP_COMMAND": "new val",
	}, func() {
		is.Eq("new val", os.Getenv("APP_COMMAND"))
	})

	is.Eq("", os.Getenv("APP_COMMAND"))
}

func TestMockOsEnv(t *testing.T) {
	is := assert.New(t)
	is.Eq("", os.Getenv("APP_COMMAND"))

	testutil.MockOsEnv(map[string]string{
		"APP_COMMAND": "new val",
	}, func() {
		is.Eq("new val", os.Getenv("APP_COMMAND"))
	})

	is.Eq("", os.Getenv("APP_COMMAND"))
}

func TestClearOSEnv(t *testing.T) {
	testutil.ClearOSEnv()
	assert.Empty(t, os.Environ())

	testutil.RevertOSEnv()
	assert.NotEmpty(t, os.Environ())
}

func TestMockOsEnvByText(t *testing.T) {
	envStr := `
APP_COMMAND = login
APP_ENV = dev
APP_DEBUG = true
APP_PWD=
`

	testutil.MockOsEnvByText(envStr, func() {
		assert.Len(t, comfunc.Environ(), 4)
		assert.Eq(t, "true", os.Getenv("APP_DEBUG"))
		assert.Eq(t, "login", os.Getenv("APP_COMMAND"))
		assert.Eq(t, "", os.Getenv("APP_PWD"))
	})
}
