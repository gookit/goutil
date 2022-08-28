package goutil_test

import (
	"testing"

	"github.com/gookit/goutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestToBool(t *testing.T) {
	is := assert.New(t)

	blVal, err := goutil.ToBool("1")
	is.Nil(err)
	is.True(blVal)

	blVal = goutil.Bool("1")
	is.True(blVal)

	is.False(goutil.Bool(false))
	is.False(goutil.Bool(1))
}

func TestToString(t *testing.T) {
	is := assert.New(t)

	str, err := goutil.ToString(23)
	is.Nil(err)
	is.Eq("23", str)

	str = goutil.String(23)
	is.Eq("23", str)
}

func TestToInt(t *testing.T) {
	is := assert.New(t)

	// To int
	iVal, err := goutil.ToInt("2")
	is.Nil(err)
	is.Eq(int(2), iVal)

	iVal = goutil.Int("-2")
	is.Nil(err)
	is.Eq(int(-2), iVal)

	// To int64
	i64Val, err := goutil.ToInt64("2")
	is.Nil(err)
	is.Eq(int64(2), i64Val)

	i64Val = goutil.Int64("-2")
	is.Nil(err)
	is.Eq(int64(-2), i64Val)

	// To uint64
	u64Val, err := goutil.ToUint("2")
	is.Nil(err)
	is.Eq(uint64(2), u64Val)

	u64Val = goutil.Uint("2")
	is.Nil(err)
	is.Eq(uint64(2), u64Val)
}

func TestBaseTypeVal(t *testing.T) {
	is := assert.New(t)

	val, err := goutil.BaseTypeVal(uint64(23))
	is.NoErr(err)
	is.Eq(int64(23), val)

	val, err = goutil.BaseTypeVal(23)
	is.NoErr(err)
	is.Eq(int64(23), val)

	val, err = goutil.BaseTypeVal(nil)
	is.Err(err)
	is.Nil(val)
}
