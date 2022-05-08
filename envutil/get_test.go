package envutil

import (
	"fmt"
	"os"
	"testing"

	"github.com/gookit/goutil/testutil"
	"github.com/stretchr/testify/assert"
)

const (
	TestEnvName     = "TEST_GOUTIL_ENV"
	TestNoEnvName   = "TEST_GOUTIL_NO_ENV"
	TestEnvValue    = "1"
	defTestEnvValue = "1"
)

func TestGetenv(t *testing.T) {
	testutil.MockEnvValues(map[string]string{
		TestEnvName: TestEnvValue,
	}, func() {
		envValue := Getenv(TestEnvName)
		assert.Equal(t, TestEnvValue, envValue, "env value not equals")
		envValue = Getenv(TestNoEnvName, defTestEnvValue)
		assert.Equal(t, defTestEnvValue, envValue, "env value not default")

		assert.Equal(t, 1, GetInt(TestEnvName), "int env value not equals")
		assert.Equal(t, 0, GetInt(TestNoEnvName))
		assert.Equal(t, 2, GetInt(TestNoEnvName, 2))
	})
}

func TestGetBool(t *testing.T) {
	testutil.MockEnvValues(map[string]string{
		TestEnvName: "on",
	}, func() {
		assert.True(t, GetBool(TestEnvName))
		assert.False(t, GetBool(TestNoEnvName))
		assert.True(t, GetBool(TestNoEnvName, true))
	})
}

func TestEnviron(t *testing.T) {
	testutil.MockOsEnv(map[string]string{
		TestEnvName: TestEnvValue,
	}, func() {
		envValue := Getenv("not_exist")
		assert.Equal(t, "", envValue)

		fmt.Println("os.Environ:", os.Environ())
		fmt.Println("new Environ:", Environ())
		assert.Contains(t, Environ(), TestEnvName)
	})
}
