package testutil

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockEnvValue(t *testing.T) {
	ris := assert.New(t)
	ris.Equal("", os.Getenv("APP_COMMAND"))

	MockEnvValue("APP_COMMAND", "new val", func(nv string) {
		ris.Equal("new val", nv)
	})

	ris.Equal("", os.Getenv("APP_COMMAND"))
}

func TestMockEnvValues(t *testing.T) {
	ris := assert.New(t)
	ris.Equal("", os.Getenv("APP_COMMAND"))

	MockEnvValues(map[string]string{
		"APP_COMMAND": "new val",
	}, func() {
		ris.Equal("new val", os.Getenv("APP_COMMAND"))
	})

	ris.Equal("", os.Getenv("APP_COMMAND"))
}
