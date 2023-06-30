package envutil

import (
	"fmt"
	"os"
	"testing"

	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
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
		assert.Eq(t, TestEnvValue, envValue, "env value not equals")
		envValue = Getenv(TestNoEnvName, defTestEnvValue)
		assert.Eq(t, defTestEnvValue, envValue, "env value not default")

		assert.Eq(t, 1, GetInt(TestEnvName), "int env value not equals")
		assert.Eq(t, 0, GetInt(TestNoEnvName))
		assert.Eq(t, 2, GetInt(TestNoEnvName, 2))

		assert.Len(t, GetMulti(TestEnvName, TestNoEnvName), 1)
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
	assert.NotEmpty(t, EnvPaths())
	assert.NotEmpty(t, EnvMap())

	testutil.MockOsEnv(map[string]string{
		TestEnvName: TestEnvValue,
	}, func() {
		envValue := Getenv("not_exist")
		assert.Eq(t, "", envValue)

		fmt.Println("os.Environ:", os.Environ())
		fmt.Println("new Environ:", Environ())
		assert.Contains(t, Environ(), TestEnvName)
	})
}

func TestSearchEnvKeys(t *testing.T) {
	testutil.MockOsEnv(map[string]string{
		TestEnvName: TestEnvValue,
		"APP_NAME":  "test",
	}, func() {
		keys := SearchEnvKeys(TestEnvName)
		assert.Contains(t, keys, TestEnvName)

		keys = SearchEnvKeys("APP")
		assert.Contains(t, keys, "APP_NAME")

		keys = SearchEnv("test", true)
		assert.Contains(t, keys, "APP_NAME")
	})
}
