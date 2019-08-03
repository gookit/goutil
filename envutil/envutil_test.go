package envutil

import (
	"testing"

	"github.com/gookit/goutil/testutil"
	"github.com/stretchr/testify/assert"
)

func TestParseEnvValue(t *testing.T) {
	ris := assert.New(t)
	tests := []struct {
		eKey, eVal, rVal, nVal string
	}{
		{"EnvKey", "EnvKey val", "${EnvKey}", "EnvKey val"},
		{"EnvKey", "", "${EnvKey}", "${EnvKey}"},
		{"EnvKey0", "EnvKey0 val", "${ EnvKey0 }", "EnvKey0 val"},
		{"EnvKey1", "EnvKey1 val", "${EnvKey1|defValue}", "EnvKey1 val"},
		{"EnvKey1", "", "${EnvKey1|defValue}", "defValue"},
		{"EnvKey2", "", "${ EnvKey2 | defValue1 }", "defValue1"},
		{"EnvKey3", "EnvKey3 val", "${ EnvKey3 | app:run }", "EnvKey3 val"},
		{"EnvKey3", "", "${ EnvKey3 | app:run }", "app:run"},
		{"EnvKey6", "", "${ EnvKey6 | app=run }", "app=run"},
		{"EnvKey7", "", "${ EnvKey7 | app.run }", "app.run"},
		{"EnvKey8", "", "${ EnvKey7 | app/run }", "app/run"},
		{"TEST_SHELL", "/bin/zsh", "${TEST_SHELL|/bin/bash}", "/bin/zsh"},
		{"TEST_SHELL", "", "${TEST_SHELL|/bin/bash}", "/bin/bash"},
	}

	for _, tt := range tests {
		ris.Equal("", Getenv(tt.eKey))

		testutil.MockEnvValue(tt.eKey, tt.eVal, func(eVal string) {
			ris.Equal(tt.eVal, eVal)
			ris.Equal(tt.nVal, ParseEnvValue(tt.rVal))
		})
	}

	// test multi ENV key
	rVal := "${FirstEnv}/${ SecondEnv }"
	ris.Equal("", Getenv("FirstEnv"))
	ris.Equal("", Getenv("SecondEnv"))
	ris.Equal(rVal, ParseEnvValue(rVal))

	testutil.MockEnvValues(map[string]string{
		"FirstEnv":  "abc",
		"SecondEnv": "def",
	}, func() {
		ris.Equal("abc", Getenv("FirstEnv"))
		ris.Equal("def", Getenv("SecondEnv"))
		ris.Equal("abc/def", ParseEnvValue(rVal))
	})

	testutil.MockEnvValues(map[string]string{
		"FirstEnv": "abc",
	}, func() {
		ris.Equal("abc", Getenv("FirstEnv"))
		ris.Equal("", Getenv("SecondEnv"))
		ris.Equal("abc/${ SecondEnv }", ParseEnvValue(rVal))
	})
}
