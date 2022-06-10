package stdio_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/stdio"
	"github.com/stretchr/testify/assert"
)

func TestQuietWriteString(t *testing.T) {
	buf := new(bytes.Buffer)
	stdio.QuietWriteString(buf, "inhere")

	assert.Equal(t, "inhere", buf.String())
}

func TestDiscardReader(t *testing.T) {
	sr := strings.NewReader("hello")
	fsutil.DiscardReader(sr)

	assert.Empty(t, stdio.MustReadReader(sr))
}
