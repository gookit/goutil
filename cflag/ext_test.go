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
	assert.NotEmpty(t, cs.Get())
	assert.NotEmpty(t, cs.String())
	assert.Eq(t, "inhere", cs.Str("name"))
	assert.Eq(t, 123, cs.Int("age"))
}

func TestKVString_Set(t *testing.T) {
	kv := cflag.NewKVString()
	assert.Empty(t, kv.Data())

	assert.NoErr(t, kv.Set("age=234"))
	assert.NotEmpty(t, kv.Data())
	assert.NotEmpty(t, kv.Get())
	assert.Eq(t, 234, kv.Int("age"))
	assert.Eq(t, "{age:234}", kv.String())
	assert.False(t, kv.IsEmpty())
	assert.True(t, kv.IsRepeatable())

	assert.NoErr(t, kv.Set("name=inhere"))
	assert.Eq(t, "inhere", kv.Str("name"))
}
