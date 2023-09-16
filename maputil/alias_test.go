package maputil_test

import (
	"testing"

	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestAliases_AddAlias(t *testing.T) {
	as := make(maputil.Aliases)
	as.AddAlias("a", "real")
	as.AddAliases("real", []string{"b"})
	as.AddAliasMap(map[string]string{"a1": "real1"})

	assert.True(t, as.HasAlias("a"))
	assert.True(t, as.HasAlias("b"))
	assert.True(t, as.HasAlias("a1"))
	assert.False(t, as.HasAlias("xyz"))

	assert.Eq(t, "real", as.ResolveAlias("a"))
	assert.Eq(t, "real", as.ResolveAlias("b"))
	assert.Eq(t, "real1", as.ResolveAlias("a1"))
	assert.Eq(t, "notExist", as.ResolveAlias("notExist"))

	assert.PanicsMsg(t, func() {
		as.AddAlias("a", "real3")
	}, "The alias 'a' is already used by 'real'")
}
