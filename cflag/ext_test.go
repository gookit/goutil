package cflag_test

import (
	"testing"

	"github.com/gookit/goutil/cflag"
	"github.com/gookit/goutil/testutil/assert"
)

func TestEnumString_Set(t *testing.T) {
	es := cflag.EnumString{}
	es.SetEnum([]string{"php", "go"})

	assert.Err(t, es.Set("no-match"))

	assert.NoErr(t, es.Set("go"))
	assert.Eq(t, "go", es.String())
}

func TestConfString_Set(t *testing.T) {
	cs := cflag.ConfString{}
	cs.SetData(map[string]string{"key": "val"})

	assert.NotEmpty(t, cs.Data())
	assert.Eq(t, "val", cs.Data().Str("key"))

	assert.NoErr(t, cs.Set(""))
	assert.Err(t, cs.Set("no-value"))

	cs = cflag.ConfString{}
	err := cs.Set("name=inhere;age=123")
	assert.NoErr(t, err)
	assert.NotEmpty(t, cs.Data())
	assert.Eq(t, "inhere", cs.Str("name"))
	assert.Eq(t, 123, cs.Int("age"))
}
