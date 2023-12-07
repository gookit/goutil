package envutil

import (
	"testing"

	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestParseEnvValue(t *testing.T) {
	is := assert.New(t)
	tests := []struct {
		eKey, eVal, rVal, nVal string
	}{
		{"EnvKey", "EnvKey val", "${EnvKey}", "EnvKey val"},
		{"EnvKey", "", "${EnvKey}", ""},
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
		// use JSON string as default value
		{"EnvKey10", "", `${ EnvKey10 | {"name": "inhere"} }`, `{"name": "inhere"}`},
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
	rVal := "${FirstEnv}/${ SecondEnv | def_val}"
	is.Eq("", Getenv("FirstEnv"))
	is.Eq("", Getenv("SecondEnv"))
	is.Eq("/def_val", ParseEnvValue(rVal))
	is.Eq("/def_val", VarParse(rVal))
	is.Eq("/", VarReplace(rVal)) // use os.ExpandEnv()

	testutil.MockEnvValues(map[string]string{
		"FirstEnv":  "abc",
		"SecondEnv": "def",
	}, func() {
		is.Eq("abc", Getenv("FirstEnv"))
		is.Eq("def", Getenv("SecondEnv"))
		is.Eq("abc/def", ParseValue(rVal))
		is.Eq("abc string", VarReplace("${FirstEnv} string"))
	})

	testutil.MockEnvValues(map[string]string{
		"FirstEnv": "abc",
	}, func() {
		is.Eq("abc", Getenv("FirstEnv"))
		is.Eq("", Getenv("SecondEnv"))
		is.Eq("abc/def_val", ParseEnvValue(rVal))
	})
}

func TestParseOrErr(t *testing.T) {
	val, err := ParseOrErr("${NotExist | ?error msg}")
	assert.ErrMsg(t, err, "error msg")
	assert.Eq(t, "", val)

	val, err = ParseOrErr("${NotExist | ?}")
	assert.ErrMsg(t, err, "value is required for var: NotExist")
	assert.Eq(t, "", val)

	testutil.MockEnvValue("NotExist", "val", func(eVal string) {
		val, err = ParseOrErr("${NotExist | ?}")
		assert.NoErr(t, err)
		assert.Eq(t, "val", val)
	})
}

func TestSetEnvs(t *testing.T) {
	envMp := map[string]string{
		"FirstEnv":  "abc",
		"SecondEnv": "def",
	}
	keys := maputil.Keys(envMp)
	for key := range envMp {
		assert.Empty(t, Getenv(key))
	}

	// SetEnvs
	SetEnvs(maputil.SMap(envMp).ToKVPairs()...)
	for key, val := range envMp {
		assert.Eq(t, val, Getenv(key))
	}
	assert.Panics(t, func() {
		SetEnvs("name", "one", "two")
	})

	UnsetEnvs(keys...)
	for key := range envMp {
		assert.Empty(t, Getenv(key))
	}

	// SetEnvMap
	SetEnvMap(envMp)
	for key, val := range envMp {
		assert.Eq(t, val, Getenv(key))
	}

	UnsetEnvs(keys...)
	for key := range envMp {
		assert.Empty(t, Getenv(key))
	}
}
