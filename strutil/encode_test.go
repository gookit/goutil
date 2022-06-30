package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/stretchr/testify/assert"
)

func TestMd5(t *testing.T) {
	assert.Equal(t, "e10adc3949ba59abbe56e057f20f883e", strutil.Md5("123456"))
	assert.Equal(t, "e10adc3949ba59abbe56e057f20f883e", strutil.MD5("123456"))
	assert.Equal(t, "a906449d5769fa7361d7ecc6aa3f6d28", strutil.GenMd5("123abc"))
	assert.Equal(t, "289dff07669d7a23de0ef88d2f7129e7", strutil.GenMd5(234))
}

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

func TestAddSlashes(t *testing.T) {
	assert.Equal(t, "", strutil.AddSlashes(""))
	assert.Equal(t, "", strutil.StripSlashes(""))

	assert.Equal(t, `{\"key\": 123}`, strutil.AddSlashes(`{"key": 123}`))
	assert.Equal(t, `{"key": 123}`, strutil.StripSlashes(`{\"key\": 123}`))
}

func TestURLEnDecode(t *testing.T) {
	is := assert.New(t)

	is.Equal("a.com/?name%3D%E4%BD%A0%E5%A5%BD", strutil.URLEncode("a.com/?name=你好"))
	is.Equal("a.com/?name=你好", strutil.URLDecode("a.com/?name%3D%E4%BD%A0%E5%A5%BD"))
	is.Equal("a.com", strutil.URLEncode("a.com"))
	is.Equal("a.com", strutil.URLDecode("a.com"))
}
