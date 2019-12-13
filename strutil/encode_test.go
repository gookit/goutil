package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/stretchr/testify/assert"
)

func TestEscape(t *testing.T)  {
	tests := struct {give, want string} {
		"<p>some text</p>",
		"&lt;p&gt;some text&lt;/p&gt;",
	}

	assert.Equal(t, tests.want, strutil.EscapeHTML(tests.give))

	tests = struct {give, want string} {
		 "<script>var a = 23;</script>",
		 `\x3Cscript\x3Evar a = 23;\x3C/script\x3E`,
	}
	assert.Equal(t, tests.want, strutil.EscapeJS(tests.give))
}
