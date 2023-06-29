package cflag_test

import (
	"flag"
	"testing"

	"github.com/gookit/goutil/cflag"
	"github.com/gookit/goutil/testutil/assert"
)

func TestAddPrefix(t *testing.T) {
	assert.Eq(t, "-a", cflag.AddPrefix("a"))
	assert.Eq(t, "--long", cflag.AddPrefix("long"))

	assert.Eq(t, "--long", cflag.AddPrefixes("long", nil))
	assert.Eq(t, "--long, -l", cflag.AddPrefixes("long", []string{"l"}))
	assert.Eq(t, "-l, --long", cflag.AddPrefixes2("long", []string{"l"}, true))
}

func TestIsFlagHelpErr(t *testing.T) {
	assert.False(t, cflag.IsFlagHelpErr(nil))
	assert.True(t, cflag.IsFlagHelpErr(flag.ErrHelp))

	// IsGoodName
	assert.True(t, cflag.IsGoodName("name"))

	// WrapColorForCode
	assert.Eq(t, "hello <mga>keywords</>", cflag.WrapColorForCode("hello `keywords`"))
}

func TestSplitShortcut(t *testing.T) {
	assert.Eq(t, []string{"a", "b"}, cflag.SplitShortcut("a,-b"))
	assert.Eq(t, []string{"a", "b"}, cflag.SplitShortcut("a, ,-b"))
	assert.Eq(t, []string{"ab", "cd"}, cflag.SplitShortcut("-- ab,,-cd"))
}

func TestReplaceShorts(t *testing.T) {
	assert.Len(t, cflag.ReplaceShorts([]string{}, map[string]string{
		"f": "file",
	}), 0)

	assert.Eq(t,
		[]string{"--file", "./config.ini", "-e"},
		cflag.ReplaceShorts([]string{"-f", "./config.ini", "-e"}, map[string]string{
			"f": "file",
		}),
	)
	assert.Eq(t,
		[]string{"--file", "./config.ini", "-e", "--number", "23"},
		cflag.ReplaceShorts([]string{"-f", "./config.ini", "-e", "--number", "23"}, map[string]string{
			"f": "file",
			"n": "number",
		}),
	)
	assert.Eq(t,
		[]string{"--file", "./config.ini", "-e", "--", "-n", "23"},
		cflag.ReplaceShorts([]string{"-f", "./config.ini", "-e", "--", "-n", "23"}, map[string]string{
			"f": "file",
			"n": "number",
		}),
	)
	assert.Eq(t,
		[]string{"--file=./config.ini", "-e", "--", "-n", "23"},
		cflag.ReplaceShorts([]string{"-f=./config.ini", "-e", "--", "-n", "23"}, map[string]string{
			"f": "file",
			"n": "number",
		}),
	)
}
