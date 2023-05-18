package structs_test

import (
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/structs"
	"github.com/gookit/goutil/testutil/assert"
)

func TestTryToMap(t *testing.T) {
	mp, err := structs.TryToMap(nil)
	assert.Empty(t, mp)
	assert.NoErr(t, err)

	type User struct {
		Name string
		Age  int
		city string
	}

	u := User{
		Name: "inhere",
		Age:  34,
		city: "somewhere",
	}

	mp, err = structs.TryToMap(u)
	assert.NoErr(t, err)
	dump.P(mp)
	assert.Contains(t, mp, "Name")
	assert.Contains(t, mp, "Age")
	assert.NotContains(t, mp, "city")

	mp, err = structs.TryToMap(&u)
	assert.NoErr(t, err)
	assert.NotEmpty(t, mp)
	// dump.P(mp)

	mp = structs.MustToMap(&u)
	assert.NotEmpty(t, mp)
	// dump.P(mp)

	// test to string map
	smp := structs.MustToSMap(&u)
	assert.NotEmpty(t, smp)
	assert.ContainsKeys(t, smp, []string{"Name", "Age"})

	smp = structs.ToSMap(&u)
	assert.NotEmpty(t, smp)
	assert.ContainsKeys(t, smp, []string{"Name", "Age"})

	_, err = structs.TryToSMap("invalid")
	assert.Err(t, err)

	smp, err = structs.TryToSMap(&u)
	assert.NoErr(t, err)
	assert.NotEmpty(t, smp)
	assert.ContainsKeys(t, smp, []string{"Name", "Age"})

	assert.NotEmpty(t, structs.ToString(&u))

	// test error
	assert.Panics(t, func() {
		structs.MustToMap("abc")
	})
	assert.Panics(t, func() {
		structs.MustToSMap("abc")
	})
}

func TestToMap_useTag(t *testing.T) {
	type User1 struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		city string
	}

	u1 := &User1{
		Name: "inhere",
		Age:  34,
		city: "somewhere",
	}

	mp := structs.ToMap(u1)
	dump.P(mp)
	assert.ContainsKeys(t, mp, []string{"name", "age"})
	assert.NotContains(t, mp, "city")

	// export unexported field
	mp = structs.MustToMap(u1, structs.ExportPrivate)
	dump.P(mp)
	assert.ContainsKeys(t, mp, []string{"name", "age", "city"})
}

type Extra struct {
	City   string `json:"city"`
	Github string `json:"github"`
}

type Extra1 struct {
	ExtraSub
	City   string `json:"city"`
	Github string `json:"github"`
}

type ExtraSub struct {
	ESubKey string `json:"e_sub_key"`
}

type User struct {
	Extra `json:"extra"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type User1 struct {
	Extra
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type User2 struct {
	Extra1
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestToMap_nestStruct(t *testing.T) {
	e := Extra{
		City:   "chengdu",
		Github: "https://github.com/inhere",
	}

	u := &User{
		Name:  "inhere",
		Age:   30,
		Extra: e,
	}

	mp := structs.MustToMap(u)
	dump.P(mp)
	assert.ContainsKeys(t, mp, []string{"name", "age", "extra"})
	assert.ContainsKeys(t, mp["extra"], []string{"city", "github"})

	// use pointer
	type UserPtrSub struct {
		*Extra `json:"extra"`
		Name   string `json:"name"`
		Age    int    `json:"age"`
	}

	u2 := &UserPtrSub{
		Name:  "inhere",
		Age:   30,
		Extra: &e,
	}

	mp = structs.MustToMap(u2)
	dump.P(mp)
	assert.ContainsKeys(t, mp, []string{"name", "age", "extra"})
	assert.ContainsKeys(t, mp["extra"], []string{"city", "github"})
}

func TestToMap_anonymousStruct(t *testing.T) {
	u := &User1{
		Name: "inhere",
		Age:  30,
		Extra: Extra{
			City:   "chengdu",
			Github: "https://github.com/inhere",
		},
	}

	mp := structs.MustToMap(u, structs.MergeAnonymous)
	dump.P(mp)

	assert.ContainsKeys(t, mp, []string{"name", "age", "city", "github"})
	assert.NotContainsKey(t, mp, "extra")
	assert.NotContainsKeys(t, mp, []string{"extra"})

	u2 := &User2{
		Name: "inhere",
		Age:  30,
		Extra1: Extra1{
			ExtraSub: ExtraSub{
				ESubKey: "sub key",
			},
			City:   "chengdu",
			Github: "https://github.com/inhere",
		},
	}

	mp = structs.MustToMap(u2, structs.MergeAnonymous)
	dump.P(mp)

	assert.ContainsKeys(t, mp, []string{"name", "age", "city", "github", "e_sub_key"})
	assert.NotContainsKey(t, mp, "extra")
}

func TestTryToMap_customTag(t *testing.T) {
	type User struct {
		Name     string `export:"name"`
		Age      int    `export:"age"`
		FullName string `export:"full_name"`
	}

	u1 := User{
		Name:     "inhere",
		Age:      34,
		FullName: "inhere xyz",
	}

	mp, err := structs.TryToMap(u1, structs.WithMapTagName("export"))
	assert.NoErr(t, err)
	assert.NotEmpty(t, mp)

	assert.ContainsKeys(t, mp, []string{"name", "age", "full_name"})
}
