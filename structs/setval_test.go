package structs_test

import (
	"testing"

	"github.com/gookit/goutil/comdef"
	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/structs"
	"github.com/gookit/goutil/testutil"
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

func TestInitDefaults_sliceField(t *testing.T) {
	type InitSliceFld struct {
		Name   string   `default:"inhere"`
		Age    int      `default:""`
		Tags   []string `default:"php,go"`
		TagIds []int64  `default:"34,456"`
	}

	u := &InitSliceFld{}
	err := structs.InitDefaults(u)
	dump.P(u)

	assert.NoErr(t, err)
	assert.Eq(t, "inhere", u.Name)
	assert.Eq(t, []string{"php", "go"}, u.Tags)
	assert.Eq(t, []int64{34, 456}, u.TagIds)
}

func TestInitDefaults_parseEnv(t *testing.T) {
	type App struct {
		Name  string `default:"${ APP_NAME | my-app }"`
		Env   string `default:"${ APP_ENV | dev}"`
		Debug bool   `default:"${ APP_DEBUG | false}"`
	}

	optFn := func(opt *structs.InitOptions) {
		opt.ParseEnv = true
	}

	// use default value
	obj := &App{}
	err := structs.InitDefaults(obj, optFn)
	assert.NoErr(t, err)
	assert.Eq(t, "my-app", obj.Name)
	assert.Eq(t, "dev", obj.Env)
	assert.False(t, obj.Debug)

	// load from env
	obj = &App{}
	testutil.MockEnvValues(map[string]string{
		"APP_NAME":  "goods",
		"APP_ENV":   "prod",
		"APP_DEBUG": "true",
	}, func() {
		err := structs.InitDefaults(obj, optFn)
		assert.NoErr(t, err)
	})

	assert.Eq(t, "goods", obj.Name)
	assert.Eq(t, "prod", obj.Env)
	assert.True(t, obj.Debug)
}

func TestInitDefaults_convTypeError(t *testing.T) {
	type User struct {
		Name string `default:"inhere"`
		Age  int    `default:"abc"`
	}

	u := &User{}
	err := structs.InitDefaults(u)
	assert.ErrMsg(t, err, comdef.ErrConvType.Error())
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
		City string `default:"cd"`
	}

	u := &User{City: "sh"}
	err := structs.InitDefaults(u)
	dump.P(u)
	assert.NoErr(t, err)
	assert.Eq(t, "inhere", u.Name)
	assert.Eq(t, 30, *u.Age)
	assert.Eq(t, "sh", u.City)
}

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
	err := structs.SetValues(u, data)
	assert.NoErr(t, err)
	assert.Eq(t, "inhere", u.Name)
	assert.Eq(t, 345, u.Age)
	assert.Eq(t, "shanghai", u.City)
}
