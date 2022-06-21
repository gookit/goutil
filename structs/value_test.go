package structs_test

import (
	"testing"

	"github.com/gookit/goutil/structs"
	"github.com/stretchr/testify/assert"
)

func TestNewValue(t *testing.T) {
	v := structs.NewValue(true)
	assert.True(t, v.Bool())

	v.Set("false")
	assert.False(t, v.Bool())
}

func TestValue_Val(t *testing.T) {
	v := structs.Value{V: 23}

	assert.Equal(t, 23, v.Val())
	assert.Equal(t, 23, v.Int())
	assert.Equal(t, int64(23), v.Int64())
	assert.Equal(t, float64(23), v.Float64())
	assert.Equal(t, "23", v.String())
	assert.False(t, v.IsEmpty())
	assert.False(t, v.Bool())

	v.V = []string{"a", "b"}
	assert.Equal(t, []string{"a", "b"}, v.Val())
	assert.Equal(t, []string{"a", "b"}, v.Strings())

	v.Reset()
	assert.Nil(t, v.V)
	assert.Nil(t, v.Val())
	assert.Nil(t, v.Strings())
	assert.True(t, v.IsEmpty())
	assert.False(t, v.Bool())
	assert.Equal(t, 0, v.Int())
	assert.Equal(t, int64(0), v.Int64())
	assert.Equal(t, float64(0), v.Float64())
	assert.Equal(t, "", v.String())
}
