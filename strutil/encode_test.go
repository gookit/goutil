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

	ret := strutil.EscapeJS("<script>var a = 23;</script>")
	assert.NotContains(t, ret, "<script>")
	assert.NotContains(t, ret, "</script>")
}

func TestBase64(t *testing.T) {

}

func TestURLEnDecode(t *testing.T) {
	is := assert.New(t)

	is.Equal("a.com/?name%3D%E4%BD%A0%E5%A5%BD", strutil.URLEncode("a.com/?name=你好"))
	is.Equal("a.com/?name=你好", strutil.URLDecode("a.com/?name%3D%E4%BD%A0%E5%A5%BD"))
	is.Equal("a.com", strutil.URLEncode("a.com"))
	is.Equal("a.com", strutil.URLDecode("a.com"))
}
