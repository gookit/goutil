package testutil_test

import (
	"testing"

	"github.com/gookit/goutil/testutil"
	"github.com/stretchr/testify/assert"
)

func TestNewTestWriter(t *testing.T) {
	tw := testutil.NewTestWriter()
	_, err := tw.Write([]byte("hello"))
	assert.NoError(t, err)
	assert.Equal(t, "hello", tw.String())
	assert.NoError(t, tw.Flush())
	assert.Equal(t, "", tw.String())
	assert.NoError(t, tw.Close())

	tw.ErrOnWrite = true
	_, err = tw.Write([]byte("hello"))
	assert.Error(t, err)
	assert.Equal(t, "", tw.String())

	tw.ErrOnFlush = true
	assert.Error(t, tw.Flush())

	tw.ErrOnClose = true
	assert.Error(t, tw.Close())
}
