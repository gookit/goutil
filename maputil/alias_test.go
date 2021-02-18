package maputil_test

import (
	"testing"

	"github.com/gookit/goutil/maputil"
	"github.com/stretchr/testify/assert"
)

func TestAliases_AddAlias(t *testing.T) {
	as := make(maputil.Aliases)
	as.AddAlias("real", "a")
	as.AddAliases("real", []string{"b"})
	as.AddAliasMap(map[string]string{"a1": "real1"})

	assert.True(t, as.HasAlias("a"))
	assert.True(t, as.HasAlias("b"))
	assert.True(t, as.HasAlias("a1"))
	assert.False(t, as.HasAlias("xyz"))

	assert.Equal(t, "real", as.ResolveAlias("a"))
	assert.Equal(t, "real", as.ResolveAlias("b"))
	assert.Equal(t, "real1", as.ResolveAlias("a1"))
	assert.Equal(t, "notExist", as.ResolveAlias("notExist"))

	assert.PanicsWithValue(t, "The alias 'a' is already used by 'real'", func() {
		as.AddAlias("real3", "a")
	})
}
