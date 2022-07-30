package envutil

import (
	"testing"

	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestParseEnvValue(t *testing.T) {
	is := assert.New(t)
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
		{"EnvKey9", "", "test_value", "test_value"},
		{"TEST_SHELL", "/bin/zsh", "${TEST_SHELL|/bin/bash}", "/bin/zsh"},
		{"TEST_SHELL", "", "${TEST_SHELL|/bin/bash}", "/bin/bash"},
	}

	for _, tt := range tests {
		is.Eq("", Getenv(tt.eKey))

		testutil.MockEnvValue(tt.eKey, tt.eVal, func(eVal string) {
			is.Eq(tt.eVal, eVal)
			is.Eq(tt.nVal, ParseEnvValue(tt.rVal))
		})
	}

	// test multi ENV key
	rVal := "${FirstEnv}/${ SecondEnv }"
	is.Eq("", Getenv("FirstEnv"))
	is.Eq("", Getenv("SecondEnv"))
	is.Eq(rVal, ParseEnvValue(rVal))
	is.Eq(rVal, VarParse(rVal))

	testutil.MockEnvValues(map[string]string{
		"FirstEnv":  "abc",
		"SecondEnv": "def",
	}, func() {
		is.Eq("abc", Getenv("FirstEnv"))
		is.Eq("def", Getenv("SecondEnv"))
		is.Eq("abc/def", ParseEnvValue(rVal))
		is.Eq("abc string", VarReplace("${FirstEnv} string"))
	})

	testutil.MockEnvValues(map[string]string{
		"FirstEnv": "abc",
	}, func() {
		is.Eq("abc", Getenv("FirstEnv"))
		is.Eq("", Getenv("SecondEnv"))
		is.Eq("abc/${ SecondEnv }", ParseEnvValue(rVal))
	})
}
