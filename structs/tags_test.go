package structs_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gookit/goutil"
	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/structs"
	"github.com/gookit/goutil/testutil/assert"
)

func ExampleTagParser_Parse() {
	type User struct {
		Age   int    `json:"age" yaml:"age" default:"23"`
		Name  string `json:"name,omitempty" yaml:"name" default:"inhere"`
		inner string //lint:ignore U1000 for test
	}

	u := &User{}
	p := structs.NewTagParser("json", "yaml", "default")
	goutil.MustOK(p.Parse(u))

	tags := p.Tags()
	dump.P(tags)
	/*tags:
	map[string]maputil.SMap { #len=2
	  "Age": maputil.SMap { #len=3
	    "json": string("age"), #len=3
	    "yaml": string("age"), #len=3
	    "default": string("23"), #len=2
	  },
	  "Name": maputil.SMap { #len=3
	    "default": string("inhere"), #len=6
	    "json": string("name,omitempty"), #len=14
	    "yaml": string("name"), #len=4
	  },
	},
	*/

	dump.P(p.Info("name", "json"))
	/*info:
	maputil.SMap { #len=2
	  "name": string("name"), #len=4
	  "omitempty": string("true"), #len=4
	},
	*/

	fmt.Println(
		tags["Age"].Get("json"),
		tags["Age"].Get("default"),
	)

	// Output:
	// age 23
}

func ExampleTagParser_Parse_parseTagValueDefine() {
	// eg: "desc;required;default;shorts"
	type MyCmd struct {
		Name string `flag:"set your name;false;INHERE;n"`
	}

	c := &MyCmd{}
	p := structs.NewTagParser("flag")

	sepStr := ";"
	defines := []string{"desc", "required", "default", "shorts"}
	p.ValueFunc = structs.ParseTagValueDefine(sepStr, defines)

	goutil.MustOK(p.Parse(c))
	// dump.P(p.Tags())
	/*
		map[string]maputil.SMap { #len=1
		  "Name": maputil.SMap { #len=1
		    "flag": string("set your name;false;INHERE;n"), #len=28
		  },
		},
	*/
	fmt.Println("tags:", p.Tags())

	info, _ := p.Info("Name", "flag")
	dump.P(info)
	/*
		maputil.SMap { #len=4
		  "desc": string("set your name"), #len=13
		  "required": string("false"), #len=5
		  "default": string("INHERE"), #len=6
		  "shorts": string("n"), #len=1
		},
	*/

	// Output:
	// tags: map[Name:{flag:set your name;false;INHERE;n}]
}

func TestParseTags(t *testing.T) {
	type user struct {
		Age   int    `json:"age" default:"23"`
		Name  string `json:"name" default:"inhere"`
		inner string //lint:ignore U1000 unused
	}

	tags, err := structs.ParseTags(user{}, []string{"json", "default"})
	assert.NoErr(t, err)
	assert.NotEmpty(t, tags)
	assert.NotContains(t, tags, "inner")

	assert.Contains(t, tags, "Age")
	assert.Eq(t, "age", tags["Age"].Str("json"))
	assert.Eq(t, 23, tags["Age"].Int("default"))

	assert.Contains(t, tags, "Name")
	assert.Eq(t, "name", tags["Name"].Str("json"))
	assert.Eq(t, 0, tags["Name"].Int("default"))

	_, err = structs.ParseTags("invalid", []string{"json", "default"})
	assert.ErrMsg(t, err, "must input an struct value")
}

func TestParseReflectTags(t *testing.T) {
	type user struct {
		Age   int    `json:"age" default:"23"`
		Name  string `json:"name" default:"inhere"`
		inner string //lint:ignore U1000 unused
	}

	rt := reflect.TypeOf(user{})
	tags, err := structs.ParseReflectTags(rt, []string{"json", "default"})
	assert.NoErr(t, err)
	assert.NotEmpty(t, tags)
	assert.NotContains(t, tags, "inner")

	assert.Contains(t, tags, "Age")
	assert.Eq(t, "age", tags["Age"].Str("json"))
	assert.Eq(t, 23, tags["Age"].Int("default"))

	assert.Contains(t, tags, "Name")
	assert.Eq(t, "name", tags["Name"].Str("json"))
	assert.Eq(t, 0, tags["Name"].Int("default"))

	_, err = structs.ParseReflectTags(reflect.TypeOf("invalid"), []string{"json", "default"})
	assert.ErrMsg(t, err, "must input an struct value")
}

func TestTagParser_Parse(t *testing.T) {
	type user struct {
		Age  int    `json:"age" default:"23"`
		Name string `json:"name" default:"inhere"`
		City string
	}

	p := structs.NewTagParser("json", "default")
	err := p.Parse(user{})
	assert.NoErr(t, err)

	_, err = p.Info("invalid", "json")
	assert.ErrMsg(t, err, "field \"Invalid\" not found")

	info, err := p.Info("City", "json")
	assert.NoErr(t, err)
	assert.Empty(t, info)
	assert.True(t, info.IsEmpty())
}

func TestTagParser_Parse_err(t *testing.T) {
	p := structs.NewTagParser("json")
	err := p.Parse("invalid")

	assert.ErrMsg(t, err, "must input an struct value")
}

func TestParseTagValueDefault(t *testing.T) {
	mp, err := structs.ParseTagValueDefault("Age", "")
	assert.NoErr(t, err)
	assert.Eq(t, "Age", mp.Get("name"))

	mp, err = structs.ParseTagValueDefault("Age", ",")
	assert.NoErr(t, err)
	assert.Eq(t, "Age", mp.Get("name"))
}

func TestParseTagValueDefine(t *testing.T) {
	sepStr := ";"
	defines := []string{"desc", "required", "default", "shorts"}
	handler := structs.ParseTagValueDefine(sepStr, defines)

	info, err := handler("Name", "set your name;true;INHERE;n")
	assert.NoErr(t, err)
	assert.Eq(t, "set your name", info.Get("desc"))
	assert.Eq(t, "true", info.Get("required"))
	assert.Eq(t, "INHERE", info.Get("default"))
	assert.Eq(t, "n", info.Get("shorts"))
	assert.True(t, info.Bool("required"))

	info, err = handler("Name", "set your name;;;n")
	assert.NoErr(t, err)
	assert.Eq(t, "set your name", info.Get("desc"))
	assert.Eq(t, "", info.Get("required"))
	assert.Eq(t, "", info.Get("default"))
	assert.Eq(t, "n", info.Get("shorts"))
	assert.False(t, info.Bool("required"))
}

func TestParseTagValueNamed(t *testing.T) {
	mp, err := structs.ParseTagValueNamed("name", "")
	assert.NoErr(t, err)
	assert.Empty(t, mp)

	mp, err = structs.ParseTagValueNamed("name", "default=inhere")
	assert.NoErr(t, err)
	assert.NotEmpty(t, mp)
	assert.Eq(t, "inhere", mp.Str("default"))

	_, err = structs.ParseTagValueNamed("name", "no-value")
	assert.ErrSubMsg(t, err, "parse tag error on field 'name': must")

	_, err = structs.ParseTagValueNamed("name", "name=n;default=inhere", "name")
	assert.ErrSubMsg(t, err, "parse tag error on field 'name': invalid")
}

func TestParseTagValueQuick(t *testing.T) {
	fields := []string{"name", "default"}
	mp := structs.ParseTagValueQuick("", fields)
	assert.Empty(t, mp)

	mp = structs.ParseTagValueQuick("inhere", fields)
	assert.NotEmpty(t, mp)
	assert.Eq(t, "inhere", mp.Str("name"))

	mp = structs.ParseTagValueQuick(";tom", fields)
	assert.NotEmpty(t, mp)
	assert.Eq(t, "", mp.Str("name"))
	assert.Eq(t, "tom", mp.Str("default"))
}
