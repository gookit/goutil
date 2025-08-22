package ccolor_test

import (
	"testing"

	"github.com/gookit/goutil/testutil/assert"
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
