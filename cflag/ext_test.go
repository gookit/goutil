package cflag_test

import (
	"errors"
	"flag"
	"testing"

	"github.com/gookit/goutil/cflag"
	"github.com/gookit/goutil/testutil/assert"
)

func TestIntVar_methods(t *testing.T) {
	iv := cflag.NewIntVar(cflag.LimitInt(1, 10))

	assert.Eq(t, 0, iv.Get())
	assert.NoErr(t, iv.Set("1"))
	assert.Eq(t, 1, iv.Get())
	assert.Eq(t, "1", iv.String())

	assert.Err(t, iv.Set("no-int"))
	assert.Err(t, iv.Set("11"))
}

func TestString_methods(t *testing.T) {
	var s cflag.String

	assert.Eq(t, "", s.Get())
	assert.NoErr(t, s.Set("val"))
	assert.Eq(t, "val", s.Get())
	assert.Eq(t, "val", s.String())

	assert.NoErr(t, s.Set("val1,val2"))
	assert.Eq(t, []string{"val1", "val2"}, s.Strings())

	assert.NoErr(t, s.Set("23,34"))
	assert.Eq(t, []int{23, 34}, s.Ints(","))
}

func TestStrVar_methods(t *testing.T) {
	sv := cflag.StrVar{}

	assert.Eq(t, "", sv.Get())
	assert.NoErr(t, sv.Set("val"))
	assert.Eq(t, "val", sv.Get())
	assert.Eq(t, "val", sv.String())

	sv = cflag.NewStrVar(func(val string) error {
		if val == "no" {
			return errors.New("invalid value")
		}
		return nil
	})
	assert.Err(t, sv.Set("no"))
}

func TestEnumString_methods(t *testing.T) {
	es := cflag.NewEnumString()
	es.WithEnum([]string{"php", "go"})

	assert.NotEmpty(t, es.Enum())
	assert.Err(t, es.Set("no-match"))

	assert.NoErr(t, es.Set("go"))
	assert.Eq(t, "go", es.String())
	assert.Eq(t, "go", es.Get())
	assert.Eq(t, "php,go", es.EnumString())
}

func TestConfString_methods(t *testing.T) {
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

func TestKVString_methods(t *testing.T) {
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

func TestInts_methods(t *testing.T) {
	its := cflag.Ints{}

	assert.NoErr(t, its.Set("23"))
	assert.NotEmpty(t, its.Get())
	assert.Eq(t, []int{23}, its.Ints())
	assert.Eq(t, "[23]", its.String())

	assert.True(t, its.IsRepeatable())
	assert.NoErr(t, its.Set("34"))
	assert.Eq(t, "[23,34]", its.String())
}

func TestIntsString_methods(t *testing.T) {
	its := cflag.IntsString{}

	assert.NoErr(t, its.Set("23"))
	assert.Eq(t, []int{23}, its.Ints())
	assert.Eq(t, []int{23}, its.Get())

	assert.NoErr(t, its.Set("34"))
	assert.Eq(t, "[23,34]", its.String())

	its.ValueFn = func(val int) error {
		if val < 10 {
			return errors.New("invalid value")
		}
		return nil
	}
	assert.Err(t, its.Set("3"))

	its.SizeFn = func(ln int) error {
		if ln > 2 {
			return errors.New("too many values")
		}
		return nil
	}
	assert.Err(t, its.Set("45"))
}

func TestStrings_methods(t *testing.T) {
	ss := cflag.Strings{}

	assert.NoErr(t, ss.Set("val"))
	assert.Eq(t, []string{"val"}, ss.Get())
	assert.Eq(t, []string{"val"}, ss.Strings())

	assert.NoErr(t, ss.Set("val2"))
	assert.Eq(t, "val,val2", ss.String())
	assert.True(t, ss.IsRepeatable())

	var v1 any

	v1 = &cflag.Strings{}
	val, ok := v1.(flag.Value)
	assert.True(t, ok)
	assert.NotNil(t, val)

	v1 = cflag.Strings{}
	val, ok = v1.(flag.Value)
	assert.False(t, ok)
	assert.Nil(t, val)
}

func TestBooleans_methods(t *testing.T) {
	bs := cflag.Booleans{}

	assert.NoErr(t, bs.Set("true"))
	assert.Eq(t, []bool{true}, bs.Bools())

	assert.NoErr(t, bs.Set("false"))
	assert.Eq(t, "[true,false]", bs.String())
	assert.True(t, bs.IsRepeatable())
}

func TestSafeFuncVar(t *testing.T) {
	var s string
	fv := cflag.SafeFuncVar(func(val string) {
		s = val
	})

	assert.NoErr(t, fv.Set("val"))
	assert.Eq(t, "val", s)
	assert.Eq(t, "", fv.String())
}
