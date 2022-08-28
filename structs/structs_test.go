package structs_test

import (
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/structs"
	"github.com/gookit/goutil/testutil/assert"
)

func TestInitDefaults(t *testing.T) {
	type User struct {
		Name string `default:"inhere"`
		Age  int    `default:""`
		city string `default:""`
	}

	u := &User{}
	err := structs.InitDefaults(u)
	assert.NoErr(t, err)
	assert.Eq(t, "inhere", u.Name)
	assert.Eq(t, 0, u.Age)
	// dump.P(u)

	type User1 struct {
		Name string `default:"inhere"`
		Age  int32  `default:"30"`
		city string `default:"val0"`
	}

	u1 := &User1{}
	err = structs.InitDefaults(u1)
	assert.NoErr(t, err)
	assert.Eq(t, "inhere", u1.Name)
	assert.Eq(t, int32(30), u1.Age)
	assert.Eq(t, "", u1.city)
	// dump.P(u1)
	// fmt.Printf("%+v\n", u1)

	err = structs.InitDefaults([]string{"invalid"})
	assert.ErrMsg(t, err, "must be provider an pointer value")

	arr := []string{"invalid"}
	err = structs.InitDefaults(&arr)
	assert.ErrMsg(t, err, "must be provider an struct value")
}

func TestInitDefaults_convTypeError(t *testing.T) {
	type User struct {
		Name string `default:"inhere"`
		Age  int    `default:"abc"`
	}

	u := &User{}
	err := structs.InitDefaults(u)
	assert.ErrMsg(t, err, "convert value type is failure")
	assert.Eq(t, "inhere", u.Name)
	assert.Eq(t, 0, u.Age)
	// dump.P(u)
}

func TestInitDefaults_nestStruct(t *testing.T) {
	type Extra struct {
		City   string `default:"chengdu"`
		Github string `default:"https://github.com/inhere"`
	}
	type User struct {
		Name  string `default:"inhere"`
		Age   int    `default:"30"`
		Extra Extra
	}

	u := &User{}
	err := structs.InitDefaults(u)
	dump.P(u)
	assert.NoErr(t, err)
}

func TestInitDefaults_fieldPtr(t *testing.T) {
	type User struct {
		Name string `default:"inhere"`
		Age  *int   `default:"30"`
		city string `default:""`
	}

	u := &User{}
	err := structs.InitDefaults(u)
	dump.P(u)
	assert.NoErr(t, err)
	assert.Eq(t, "inhere", u.Name)
	assert.Eq(t, 30, *u.Age)
}
