package structs_test

import (
	"regexp"
	"testing"

	"github.com/gookit/goutil/structs"
	"github.com/gookit/goutil/testutil/assert"
)

func TestAliases_AddAliases(t *testing.T) {
	as := structs.Aliases{}

	as.AddAlias("real", "a")
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

	assert.NotEmpty(t, as.Mapping())

	assert.PanicsMsg(t, func() {
		as.AddAlias("real3", "a")
	}, "The alias 'a' is already used by 'real'")
}

func TestAliases_AddAlias(t *testing.T) {
	as := structs.NewAliases(func(alias string) {
		rg := regexp.MustCompile(`^[a-zA-Z][\w-]*$`)
		if !rg.MatchString(alias) {
			panic("alias must match: ^[a-zA-Z][\\w-]*$")
		}
	})

	assert.PanicsMsg(t, func() {
		as.AddAlias("real3", "a:b")
	}, "alias must match: ^[a-zA-Z][\\w-]*$")
}
