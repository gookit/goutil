package structs_test

import (
	"testing"

	"github.com/gookit/goutil/structs"
	"github.com/gookit/goutil/testutil/assert"
)

func TestParseTagValueINI(t *testing.T) {
	mp, err := structs.ParseTagValueINI("name", "")
	assert.NoErr(t, err)
	assert.Empty(t, mp)

	mp, err = structs.ParseTagValueINI("name", "default=inhere")
	assert.NoErr(t, err)
	assert.NotEmpty(t, mp)
	assert.Eq(t, "inhere", mp.Str("default"))
}

func TestParseTags(t *testing.T) {
	type user struct {
		Age   int    `json:"age" default:"23"`
		Name  string `json:"name" default:"inhere"`
		inner string
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
}
