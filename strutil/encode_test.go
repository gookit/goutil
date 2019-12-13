package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/stretchr/testify/assert"
)

func TestEscape(t *testing.T) {
	tests := struct{ give, want string }{
		"<p>some text</p>",
		"&lt;p&gt;some text&lt;/p&gt;",
	}

	assert.Equal(t, tests.want, strutil.EscapeHTML(tests.give))

	tests = struct{ give, want string }{
		"<script>var a = 23;</script>",
		`\x3Cscript\x3Evar a = 23;\x3C/script\x3E`,
	}
	assert.Equal(t, tests.want, strutil.EscapeJS(tests.give))
}

func TestMd5(t *testing.T) {
	assert.Equal(t, "e10adc3949ba59abbe56e057f20f883e", strutil.Md5("123456"))
	assert.Equal(t, "a906449d5769fa7361d7ecc6aa3f6d28", strutil.GenMd5("123abc"))
	assert.Equal(t, "289dff07669d7a23de0ef88d2f7129e7", strutil.GenMd5(234))
}

func TestBase64(t *testing.T) {

}
