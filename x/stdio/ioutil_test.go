package stdio_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/x/stdio"
)

func TestFprintTo(t *testing.T) {
	buf := testutil.NewBuffer()

	stdio.Fprint(buf, "hi, inhere")
	assert.Eq(t, "hi, inhere", buf.ResetAndGet())

	stdio.Fprintf(buf, "hi, %s", "inhere")
	assert.Eq(t, "hi, inhere", buf.ResetAndGet())

	stdio.Fprintln(buf, "hi, inhere")
	assert.Eq(t, "hi, inhere\n", buf.ResetAndGet())
}

func TestWriteStringTo(t *testing.T) {
	buf := new(bytes.Buffer)
	stdio.WriteStringTo(buf, "inhere")

	assert.Eq(t, "inhere", buf.String())
}

func TestDiscardReader(t *testing.T) {
	sr := strings.NewReader("hello")
	fsutil.DiscardReader(sr)

	assert.Empty(t, stdio.MustReadReader(sr))
}
