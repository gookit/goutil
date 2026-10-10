package ccolor_test

import (
	"testing"

	"github.com/gookit/goutil/x/assert"
	"github.com/gookit/goutil/x/ccolor"
)

func TestTagCommon(t *testing.T) {
	assert.NotEmpty(t, ccolor.ColorTags())
	assert.True(t, ccolor.IsDefinedTag("info"))
}

func TestApplyTag(t *testing.T) {
	ccolor.ForceEnableColor()
	defer ccolor.RevertColorSupport()

	assert.Equal(t, "\x1b[0;32mMSG\x1b[0m", ccolor.ApplyTag("info", "MSG"))
}

func TestColorTag(t *testing.T) {
	// force open color render for testing
	ccolor.ForceEnableColor()
	defer ccolor.RevertColorSupport()

	is := assert.New(t)

	// sample 1
	r := ccolor.ReplaceTag("<err>text</>")
	is.NotContains(r, "<")
	is.NotContains(r, ">")

	// sample 2
	s := "abc <err>err-text</> def <info>info text</>"
	r = ccolor.ReplaceTag(s)
	is.NotContains(r, "<")
	is.NotContains(r, ">")

	// sample 3
	s = `abc <err>err-text</>
def <info>info text
</>`
	r = ccolor.ReplaceTag(s)
	is.NotContains(r, "<")
	is.NotContains(r, ">")

	// sample 4
	s = "abc <err>err-text</> def <err>err-text</> "
	r = ccolor.ReplaceTag(s)
	is.NotContains(r, "<")
	is.NotContains(r, ">")

	// sample 5
	s = "abc <err>err-text</> def <d>"
	r = ccolor.ReplaceTag(s)
	is.NotContains(r, "<err>")
	is.Contains(r, "<d>")

	// sample 6
	// s = "custom tag: <fg=yellow;bg=black;op=underscore;>hello, welcome</>"
	// r = ccolor.ReplaceTag(s)
	// is.NotContains(r, "<")
	// is.NotContains(r, ">")

	// no tags
	r = ccolor.Render("not color tags")
	is.Equal("not color tags", r)

	// empty
	s = ccolor.Render()
	is.Equal("", s)
}

func TestParseTag_unknownTagKeepsLaterTags(t *testing.T) {
	ccolor.ForceEnableColor()
	defer ccolor.RevertColorSupport()

	s := "  <info>--model</> Model, <provider>/<model> or id\n  <info>--no-project-rules</> Skip"
	got := ccolor.ParseTag(s)
	assert.Eq(t, "  \x1b[0;32m--model\x1b[0m Model, <provider>/<model> or id\n  \x1b[0;32m--no-project-rules\x1b[0m Skip", got)

	// unknown tag wrapping real content is left untouched
	assert.Eq(t, "<name>text</>", ccolor.ParseTag("<name>text</>"))
}

func TestClearTag_keepsPlainAngleText(t *testing.T) {
	s := "<info>--model</> Model, <provider>/<model> or id <fg=red;op=bold>x</> <>y</info>"
	assert.Eq(t, "--model Model, <provider>/<model> or id x y", ccolor.ClearTag(s))
}
