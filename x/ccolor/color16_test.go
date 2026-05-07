package ccolor_test

import (
	"bytes"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/x/ccolor"
)

func TestColor_Fprint(t *testing.T) {
	ccolor.ForceEnableColor()
	defer ccolor.RevertColorSupport()

	buf := &bytes.Buffer{}

	ccolor.Green.Fprint(buf, "text")
	assert.Equal(t, "\x1b[32mtext\x1b[0m", buf.String())
	buf.Reset()

	ccolor.Red.Fprint(buf, "text", "more")
	assert.Equal(t, "\x1b[31mtextmore\x1b[0m", buf.String())
	buf.Reset()

	ccolor.Cyan.Fprint(buf)
	assert.Equal(t, "", buf.String())
	buf.Reset()
}

func TestColor_Fprintf(t *testing.T) {
	ccolor.ForceEnableColor()
	defer ccolor.RevertColorSupport()

	buf := &bytes.Buffer{}

	ccolor.Cyan.Fprintf(buf, "string %s", "arg0")
	assert.Equal(t, "\x1b[36mstring arg0\x1b[0m", buf.String())
	buf.Reset()

	ccolor.Yellow.Fprintf(buf, "number %d", 42)
	assert.Equal(t, "\x1b[33mnumber 42\x1b[0m", buf.String())
	buf.Reset()
}

func TestColor_Fprintln(t *testing.T) {
	ccolor.ForceEnableColor()
	defer ccolor.RevertColorSupport()

	buf := &bytes.Buffer{}

	ccolor.Green.Fprintln(buf, "text")
	assert.Equal(t, "\x1b[32mtext\x1b[0m\n", buf.String())
	buf.Reset()

	ccolor.Red.Fprintln(buf, "text", "more")
	assert.Equal(t, "\x1b[31mtext more\x1b[0m\n", buf.String())
	buf.Reset()

	ccolor.Cyan.Fprintln(buf)
	assert.Equal(t, "\n", buf.String())
	buf.Reset()
}

func TestColor_Fprint_AllColors(t *testing.T) {
	ccolor.ForceEnableColor()
	defer ccolor.RevertColorSupport()

	tests := []struct {
		name  string
		color ccolor.Color
		code  string
	}{
		{"FgBlack", ccolor.FgBlack, "30"},
		{"FgRed", ccolor.FgRed, "31"},
		{"FgGreen", ccolor.FgGreen, "32"},
		{"FgYellow", ccolor.FgYellow, "33"},
		{"FgBlue", ccolor.FgBlue, "34"},
		{"FgMagenta", ccolor.FgMagenta, "35"},
		{"FgCyan", ccolor.FgCyan, "36"},
		{"FgWhite", ccolor.FgWhite, "37"},
		{"FgLightRed", ccolor.FgLightRed, "91"},
		{"FgLightGreen", ccolor.FgLightGreen, "92"},
		{"BgRed", ccolor.BgRed, "41"},
		{"BgGreen", ccolor.BgGreen, "42"},
		{"OpBold", ccolor.OpBold, "1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			tt.color.Fprint(buf, "msg")
			assert.Equal(t, "\x1b["+tt.code+"mmsg\x1b[0m", buf.String())
		})
	}
}
