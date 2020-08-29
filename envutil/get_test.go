package envutil

import (
	"testing"

	"github.com/gookit/goutil/testutil"
	"github.com/stretchr/testify/assert"
)

const (
	TestEnvName         = "TEST_GOUTIL_ENV"
	TestNoEnvName       = "TEST_GOUTIL_NO_ENV"
	TestEnvValue        = "1"
	DefaultTestEnvValue = "1"
)

func TestGetenv(t *testing.T) {
	testutil.MockEnvValues(map[string]string{
		TestEnvName: TestEnvValue,
	}, func() {
		envValue := Getenv(TestEnvName)
		assert.Equal(t, TestEnvValue, envValue, "env value not equals")
		envValue = Getenv(TestNoEnvName, DefaultTestEnvValue)
		assert.Equal(t, DefaultTestEnvValue, envValue, "env value not default")
	})
}
