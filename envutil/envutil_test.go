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

func TestSplitLineToKv(t *testing.T) {
	tests := []struct {
		line string
		k, v string
	}{
		{"key=val", "key", "val"},
		{"key = val ", "key", "val"},
		{"key =val\n", "key", "val"},
		{"key= val\r\n", "key", "val"},
		{" key=val\r", "key", "val"},
		{"key=val\t ", "key", "val"},
		{" key=val\t\n", "key", "val"},
		{"key=val\t\r\n", "key", "val"},
		{"key = val\nue", "key", "val\nue"},
		{" key-one =val ", "key-one", "val"},
		{" key_one = val", "key_one", "val"},
		{" valid=", "valid", ""},
		// invalid input
		{"invalid", "", ""},
		{"=invalid", "", ""},
		{" = invalid", "", ""},
		{"  ", "", ""},
	}

	for _, tt := range tests {
		k, v := SplitLineToKv(tt.line)
		assert.Eq(t, tt.k, k)
		assert.Eq(t, tt.v, v)
	}
}

func TestSplitText2map(t *testing.T) {
	envMp := SplitText2map(`
# comment1
APP = kite
; comment2
DEBUG = true
RUN_OPT = 
# must split with =
INVALID
`)
	assert.NotEmpty(t, envMp)
	assert.Eq(t, "kite", envMp["APP"])
	assert.Eq(t, "true", envMp["DEBUG"])
	_, has := envMp["RUN_OPT"]
	assert.True(t, has)
	_, has = envMp["INVALID"]
	assert.False(t, has)
}
