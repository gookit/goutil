package structs_test

import (
	"testing"

	"github.com/gookit/goutil/structs"
	"github.com/stretchr/testify/assert"
)

func TestParseTagValueINI(t *testing.T) {
	mp, err := structs.ParseTagValueINI("name", "")
	assert.NoError(t, err)
	assert.Empty(t, mp)

	mp, err = structs.ParseTagValueINI("name", "default=inhere")
	assert.NoError(t, err)
	assert.NotEmpty(t, mp)
	assert.Equal(t, "inhere", mp.Str("default"))
}
