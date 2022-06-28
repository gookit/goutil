package stdio_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/stdio"
	"github.com/gookit/goutil/testutil"
	"github.com/stretchr/testify/assert"
)

func TestQuietFprint(t *testing.T) {
	buf := testutil.NewBuffer()

	stdio.QuietFprint(buf, "hi, inhere")
	assert.Equal(t, "hi, inhere", buf.ResetAndGet())

	stdio.QuietFprintf(buf, "hi, %s", "inhere")
	assert.Equal(t, "hi, inhere", buf.ResetAndGet())

	stdio.QuietFprintln(buf, "hi, inhere")
	assert.Equal(t, "hi, inhere\n", buf.ResetAndGet())
}

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
