package ccolor_test

import (
	"bytes"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/x/ccolor"
)

func TestStyleString(t *testing.T) {
	assert.Eq(t, "32", ccolor.Info.String())
}

func TestStyle_Fprint(t *testing.T) {
	ccolor.ForceEnableColor()
	defer ccolor.RevertColorSupport()

	buf := &bytes.Buffer{}

	ccolor.Info.Fprint(buf, "text")
	assert.Equal(t, "\x1b[32mtext\x1b[0m", buf.String())
	buf.Reset()

	ccolor.Warn.Fprint(buf, "text", "more")
	assert.Equal(t, "\x1b[33mtextmore\x1b[0m", buf.String())
	buf.Reset()

	ccolor.Debug.Fprint(buf)
	assert.Equal(t, "", buf.String())
	buf.Reset()
}

func TestStyle_Fprintf(t *testing.T) {
	ccolor.ForceEnableColor()
	defer ccolor.RevertColorSupport()

	buf := &bytes.Buffer{}

	ccolor.Info.Fprintf(buf, "string %s", "arg0")
	assert.Equal(t, "\x1b[32mstring arg0\x1b[0m", buf.String())
	buf.Reset()

	ccolor.Warn.Fprintf(buf, "number %d", 42)
	assert.Equal(t, "\x1b[33mnumber 42\x1b[0m", buf.String())
	buf.Reset()
}

func TestStyle_Fprintln(t *testing.T) {
	ccolor.ForceEnableColor()
	defer ccolor.RevertColorSupport()

	buf := &bytes.Buffer{}

	ccolor.Info.Fprintln(buf, "text")
	assert.Equal(t, "\x1b[32mtext\x1b[0m\n", buf.String())
	buf.Reset()

	ccolor.Warn.Fprintln(buf, "text", "more")
	assert.Equal(t, "\x1b[33mtext more\x1b[0m\n", buf.String())
	buf.Reset()

	ccolor.Debug.Fprintln(buf)
	assert.Equal(t, "\n", buf.String())
	buf.Reset()
}

func TestStyle_Fprint_WithOpts(t *testing.T) {
	ccolor.ForceEnableColor()
	defer ccolor.RevertColorSupport()

	buf := &bytes.Buffer{}

	ccolor.Success.Fprint(buf, "text")
	assert.Equal(t, "\x1b[32;1mtext\x1b[0m", buf.String())
	buf.Reset()

	ccolor.Error.Fprint(buf, "error")
	assert.Equal(t, "\x1b[97;41merror\x1b[0m", buf.String())
	buf.Reset()
}
