package structs_test

import (
	"regexp"
	"testing"

	"github.com/gookit/goutil/structs"
	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, "real", as.ResolveAlias("a"))
	assert.Equal(t, "real", as.ResolveAlias("b"))
	assert.Equal(t, "real1", as.ResolveAlias("a1"))
	assert.Equal(t, "notExist", as.ResolveAlias("notExist"))

	assert.PanicsWithValue(t, "The alias 'a' is already used by 'real'", func() {
		as.AddAlias("real3", "a")
	})
}

func TestAliases_AddAlias(t *testing.T) {
	as := structs.NewAliases(func(alias string) {
		rg := regexp.MustCompile(`^[a-zA-Z][\w-]*$`)
		if !rg.MatchString(alias) {
			panic("alias must match: ^[a-zA-Z][\\w-]*$")
		}
	})

	assert.PanicsWithValue(t, "alias must match: ^[a-zA-Z][\\w-]*$", func() {
		as.AddAlias("real3", "a:b")
	})
}
