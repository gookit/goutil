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
	assert.Eq(t, "", u.city)
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

type ExtraDefault struct {
	City   string `default:"chengdu"`
	Github string `default:"https://github.com/inhere"`
}

func TestInitDefaults_nestStruct(t *testing.T) {
	type User struct {
		Name  string       `default:"inhere"`
		Age   int          `default:"30"`
		Extra ExtraDefault `default:""`
	}

	u := &User{}
	err := structs.InitDefaults(u)
	dump.P(u)
	assert.NoErr(t, err)
	assert.Eq(t, "inhere", u.Name)
	assert.Eq(t, 30, u.Age)
	assert.Eq(t, "chengdu", u.Extra.City)
	assert.Eq(t, "https://github.com/inhere", u.Extra.Github)

	u = &User{Extra: ExtraDefault{Github: "some url"}}
	err = structs.InitDefaults(u)
	// dump.P(u)
	assert.NoErr(t, err)
	assert.Eq(t, "chengdu", u.Extra.City)
	assert.Eq(t, "some url", u.Extra.Github)
}

func TestInitDefaults_ptrStructField(t *testing.T) {
	// test for pointer struct field
	type User struct {
		Name  string        `default:"inhere"`
		Age   int           `default:"30"`
		Extra *ExtraDefault `default:""`
	}

	u := &User{}
	err := structs.InitDefaults(u)
	dump.P(u)
	assert.NoErr(t, err)
	assert.Eq(t, "inhere", u.Name)
	assert.Eq(t, 30, u.Age)
	assert.Eq(t, "chengdu", u.Extra.City)
	assert.Eq(t, "https://github.com/inhere", u.Extra.Github)

	u = &User{Extra: &ExtraDefault{Github: "some url"}}
	err = structs.InitDefaults(u)
	// dump.P(u)
	assert.NoErr(t, err)
	assert.Eq(t, "chengdu", u.Extra.City)
	assert.Eq(t, "some url", u.Extra.Github)
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

func TestInitDefaults_initStructSlice(t *testing.T) {
	// test for slice struct field
	type User struct {
		Name  string         `default:"inhere"`
		Age   int            `default:"30"`
		Extra []ExtraDefault `default:""`
	}

	u := &User{}
	err := structs.Init(u)
	dump.P(u)
	assert.NoErr(t, err)
	assert.NotEmpty(t, u.Extra)
	assert.Eq(t, "chengdu", u.Extra[0].City)

	// test for slice struct field
	type User1 struct {
		Name  string          `default:"inhere"`
		Age   int             `default:"30"`
		Extra []*ExtraDefault `default:""`
	}

	u1 := &User1{}
	err = structs.Init(u1)
	dump.P(u1)
	assert.NoErr(t, err)
	assert.NotEmpty(t, u1.Extra)
	assert.Eq(t, "chengdu", u1.Extra[0].City)
}

func TestInitDefaults_ptrField(t *testing.T) {
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
