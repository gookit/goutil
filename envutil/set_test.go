package envutil

import (
	"testing"

	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

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

	// MustGet
	assert.Eq(t, "abc", MustGet("FirstEnv"))
	assert.Panics(t, func() {
		MustGet("NotExistEnvKey")
	})

	// UnsetEnvs
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

func TestLoadText(t *testing.T) {
	testutil.RunOnCleanEnv(func() {
		LoadText(`
# comment1
APP = kite
; comment2
DEBUG = true
RUN_OPT = 
# must split with =
INVALID
`)
		envMp := EnvMap()
		assert.NotEmpty(t, envMp)
		assert.Eq(t, "kite", envMp["APP"])
		assert.Eq(t, "true", envMp["DEBUG"])
		_, has := envMp["RUN_OPT"]
		assert.True(t, has)
		_, has = envMp["INVALID"]
		assert.False(t, has)
	})
}

func TestLoadString(t *testing.T) {
	tests := []struct {
		line string
		key, want string
		suc bool
	}{
		{
			"name= abc",
			"name",
			"abc",
			true,
		},
		{
			"name = abc\n",
			"name",
			"abc",
			true,
		},
		{
			"name=abc\nname2=def",
			"name",
			"abc\nname2=def",
			true,
		},
		{
			"invalid",
			"",
			"",
			false,
		},
	}

	testutil.ClearOSEnv()
	defer testutil.RevertOSEnv()

	for _, tt := range tests {
		suc := LoadString(tt.line)
		assert.Eq(t, tt.suc, suc)
		if suc {
			assert.Eq(t, tt.want, Getenv(tt.key))
		}
	}
}