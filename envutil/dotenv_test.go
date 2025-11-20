package envutil_test

import (
	"testing"

	"github.com/gookit/goutil/envutil"
	"github.com/gookit/goutil/testutil/assert"
)

func Test_dotenv(t *testing.T) {
	std := envutil.StdDotenv()
	assert.NotNil(t, std)
	assert.False(t, std.UnloadEnv())

	err := envutil.LoadEnvFiles("testdata")
	assert.NoErr(t, err)
	// dump.P(std.LoadedData())

	// check env value
	assert.Eq(t, "dev", envutil.Getenv("T_APP_ENV"))
	assert.True(t, envutil.GetBool("T_APP_DEBUG"))
	assert.Eq(t, "", envutil.Getenv("T_APP_NAME"))
	assert.Eq(t, "value has space", envutil.Getenv("T_APP_VAL1"))
	assert.Eq(t, "value has = char", envutil.Getenv("T_APP_VAL2"))
	assert.False(t, envutil.HasEnv("T_INVALID"))
	assert.Empty(t, envutil.Getenv("T_INVALID"))

	// check loaded
	assert.NotEmpty(t, std.LoadedFiles())
	assert.NotEmpty(t, std.LoadedData())

	// LoadFiles
	err = std.LoadFiles(".env.prod")
	assert.NoErr(t, err)
	assert.Len(t, std.LoadedFiles(), 2)

	// check env value
	assert.Eq(t, "prod", envutil.Getenv("T_APP_ENV"))
	assert.False(t, envutil.GetBool("T_APP_DEBUG"))

	// load contents
	err = std.LoadText(`
# comments
T_APP_KEY009 = value009
`)
	assert.NoErr(t, err)
	assert.Eq(t, "value009", envutil.Getenv("T_APP_KEY009"))

	// unload env
	assert.True(t, std.UnloadEnv())
	assert.False(t, envutil.HasEnv("T_APP_ENV"))
	assert.False(t, envutil.HasEnv("T_APP_DEBUG"))

	// reset
	std.Reset()
	assert.Empty(t, std.LoadedFiles())
}

func Test_Dotenv_loadGlob(t *testing.T) {
	std := envutil.StdDotenv()
	defer std.Reset()

	err := envutil.LoadEnvFiles("testdata", ".env.*")
	assert.NoErr(t, err)

	assert.NotEmpty(t, std.LoadedFiles())
}

func Test_Dotenv_notExistFile(t *testing.T) {
	err := envutil.LoadEnvFiles("testdata", "not-exist-file.env")
	assert.Err(t, err)

	err = envutil.DotenvLoad(func(cfg *envutil.Dotenv) {
		cfg.IgnoreNotExist = true
		cfg.Files = []string{"not-exist-file.env"}
	})
	assert.NoErr(t, err)
}