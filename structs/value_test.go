package structs_test

import (
	"testing"

	"github.com/gookit/goutil/structs"
	"github.com/gookit/goutil/testutil/assert"
)

func TestNewValue(t *testing.T) {
	v := structs.NewValue(true)
	assert.True(t, v.Bool())

	v.Set("false")
	assert.False(t, v.Bool())

	// string
	v.Set("12, 34")
	assert.Eq(t, "12, 34", v.Val())
	assert.Eq(t, "12, 34", v.String())
	assert.Eq(t, []string{"12", "34"}, v.Strings())
	assert.Eq(t, []string{"12", "34"}, v.SplitToStrings())
	assert.Eq(t, []int{12, 34}, v.SplitToInts())

	// nil
	v.Set(nil)
	assert.Empty(t, v.Strings())
	assert.Empty(t, v.SplitToInts())
	assert.Empty(t, v.SplitToStrings())
}

func TestValue_Val(t *testing.T) {
	v := structs.Value{V: 23}

	assert.Eq(t, 23, v.Val())
	assert.Eq(t, 23, v.Int())
	assert.Eq(t, int64(23), v.Int64())
	assert.Eq(t, float64(23), v.Float64())
	assert.Eq(t, "23", v.String())
	assert.False(t, v.IsEmpty())
	assert.False(t, v.Bool())

	v.V = []string{"a", "b"}
	assert.Eq(t, []string{"a", "b"}, v.Val())
	assert.Eq(t, []string{"a", "b"}, v.Strings())

	v.Reset()
	assert.Nil(t, v.V)
	assert.Nil(t, v.Val())
	assert.Nil(t, v.Strings())
	assert.True(t, v.IsEmpty())
	assert.False(t, v.Bool())
	assert.Eq(t, 0, v.Int())
	assert.Eq(t, int64(0), v.Int64())
	assert.Eq(t, float64(0), v.Float64())
	assert.Eq(t, "", v.String())
}
