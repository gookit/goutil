package textutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/strutil/textutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestStrVarRenderer(t *testing.T) {
	sv := textutil.NewStrVarRenderer()
	sv.SetFunc("demo_func", func(...any) any {
		return "2025-01-01"
	})
	sv.SetFuncMap(map[string]func(...any) any{
		"func_with_arg": func(as ...any) any {
			return "value, params: " + strutil.SafeString(as)
		},
	})
	sv.SetVars(map[string]any{
		"g1": "global value1",
	})
	sv.SetVar("g2", "global value2")

	vars := map[string]any{
		"name": "inhere",
		"int_var": 345,
		"some_var": "value1",
		"mapdata": map[string]string{
			"key1": "value1",
		},
	}

	text := "name: $name, int: ${int_var}. mapdata.key1: ${mapdata.key1}. $ENV_VAR"

	// render var
	t.Run("render var", func(t *testing.T) {
		text = sv.Render(text, vars)
		fmt.Println(text)
		assert.Eq(t, "name: inhere, int: 345. mapdata.key1: value1. $ENV_VAR", text)
	})

	// render global vars
	t.Run("global vars", func(t *testing.T) {
		text = sv.Render(`hi $g1, $g2`, vars)
		fmt.Println(text)
		assert.Eq(t, "hi global value1, global value2", text)
	})

	// render var like $1, $2, $3
	t.Run("render var01", func(t *testing.T) {
		vars["1"] = "name"
		vars["2"] = "number2"
		text = sv.Render(`hi $1, $2`, vars)
		fmt.Println(text)
		assert.Eq(t, "hi name, number2", text)
	})

	// render var like $*, $@
	t.Run("render var02", func(t *testing.T) {
		vars["*"] = "all var string"
		vars["@"] = "all vars"
		text = sv.Render(`*: $*, @: $@`, vars)
		fmt.Println(text)
		assert.Eq(t, "*: all var string, @: all vars", text)
	})

	// test getter func
	t.Run("getter func", func(t *testing.T) {
		sv.SetGetter(func(name string) (string, bool) {
			if name == "ENV_VAR" {
				return "ENV_VALUE", true
			}
			return "", false
		})

		text = "name: $name, int: ${int_var}. mapdata.key1: ${mapdata.key1}. $ENV_VAR"
		text = sv.Render(text, vars)
		fmt.Println(text)
		assert.Eq(t, "name: inhere, int: 345. mapdata.key1: value1. ENV_VALUE", text)
	})

	// call func
	t.Run("call func", func(t *testing.T) {
		text = sv.Render(`call func: ${demo_func()}, has params: ${func_with_arg("v001", $name, "arg002", 345)}`, vars)
		fmt.Println(text)
		assert.Eq(t, `call func: 2025-01-01, has params: value, params: [v001 inhere arg002 345]`, text)
	})
}
