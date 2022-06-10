package stdio_test

import (
	"bytes"
	"testing"

	"github.com/gookit/goutil/stdio"
	"github.com/stretchr/testify/assert"
)

func TestNewWriteWrapper(t *testing.T) {
	buf := new(bytes.Buffer)

	w := stdio.NewWriteWrapper(buf)
	_, err := w.WriteString("inhere")
	assert.NoError(t, err)
	assert.Equal(t, "inhere", w.String())

	err = w.WriteByte(',')
	assert.NoError(t, err)

	_, err = w.Write([]byte("hi"))
	assert.NoError(t, err)
	assert.Equal(t, "inhere,hi", w.String())
}
