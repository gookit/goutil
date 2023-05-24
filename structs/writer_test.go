package structs_test

import (
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/structs"
	"github.com/gookit/goutil/testutil/assert"
)

func TestSetValues(t *testing.T) {
	data := map[string]any{
		"Name": "inhere",
		"Age":  234,
		"Tags": []string{"php", "go"},
		"city": "chengdu",
	}

	type User struct {
		Name string
		Age  int
		Tags []string
		city string
	}

	u := &User{}
	err := structs.SetValues(u, data)
	assert.NoErr(t, err)
	assert.Eq(t, "inhere", u.Name)
	assert.Eq(t, 234, u.Age)
	assert.Eq(t, []string{"php", "go"}, u.Tags)
	assert.Eq(t, "", u.city)
	// dump.P(u)

	err = structs.SetValues(u, nil)
	assert.NoErr(t, err)
}

func TestSetValues_useFieldTag(t *testing.T) {
	data := map[string]any{
		"name": "inhere",
		"age":  234,
		"tags": []string{"php", "go"},
		"city": "chengdu",
	}

	type User struct {
		Name string   `json:"name"`
		Age  int      `json:"age"`
		Tags []string `json:"tags"`
		City string   `json:"city"`
	}

	u := &User{}
	err := structs.SetValues(u, data)
	dump.P(u)
	assert.NoErr(t, err)
	assert.Eq(t, "inhere", u.Name)
	assert.Eq(t, 234, u.Age)
	assert.Eq(t, []string{"php", "go"}, u.Tags)
	assert.Eq(t, "chengdu", u.City)

	// test for ptr field
	type User2 struct {
		Name *string  `json:"name"`
		Age  *int     `json:"age"`
		Tags []string `json:"tags"`
	}

	u2 := &User2{}
	err = structs.SetValues(u2, data)
	dump.P(u2)
	assert.NoErr(t, err)
	assert.Eq(t, "inhere", *u2.Name)
	assert.Eq(t, 234, *u2.Age)
	assert.Eq(t, []string{"php", "go"}, u2.Tags)
}

func TestSetValues_structField(t *testing.T) {
	type Address struct {
		City string `json:"city"`
	}

	data := map[string]any{
		"name": "inhere",
		"age":  234,
		"address": map[string]any{
			"city": "chengdu",
		},
	}

	// test for struct field
	t.Run("struct field", func(t *testing.T) {
		type User struct {
			Name    string  `json:"name"`
			Age     int     `json:"age"`
			Address Address `json:"address"`
		}

		u := &User{}
		err := structs.SetValues(u, data)
		dump.P(u)
		assert.NoErr(t, err)
		assert.Eq(t, "inhere", u.Name)
		assert.Eq(t, 234, u.Age)
		assert.Eq(t, "chengdu", u.Address.City)

		// test for error data
		assert.Err(t, structs.SetValues(u, map[string]any{
			"address": "string",
		}))
	})

	// test for struct ptr field
	t.Run("struct ptr field", func(t *testing.T) {
		type User2 struct {
			Name    string   `json:"name"`
			Age     int      `json:"age"`
			Address *Address `json:"address"`
		}

		u2 := &User2{}
		err := structs.SetValues(u2, data)
		dump.P(u2)
		assert.NoErr(t, err)
		assert.Eq(t, "inhere", u2.Name)
	})
}

func TestSetValues_useDefaultTag(t *testing.T) {
	data := map[string]any{
		"name": "inhere",
		// "age":  234,
		// "city": "chengdu",
	}

	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age" default:"345"`
		City string `json:"city" default:"shanghai"`
	}

	u := &User{}
	err := structs.SetValues(u, data, structs.WithParseDefault)
	assert.NoErr(t, err)
	assert.Eq(t, "inhere", u.Name)
	assert.Eq(t, 345, u.Age)
	assert.Eq(t, "shanghai", u.City)
}
