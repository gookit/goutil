package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestMd5(t *testing.T) {
	assert.Eq(t, "e10adc3949ba59abbe56e057f20f883e", strutil.Md5("123456"))
	assert.Eq(t, "e10adc3949ba59abbe56e057f20f883e", strutil.MD5("123456"))
	assert.Eq(t, "e10adc3949ba59abbe56e057f20f883e", strutil.MD5([]byte("123456")))
	assert.Eq(t, "a906449d5769fa7361d7ecc6aa3f6d28", strutil.GenMd5("123abc"))
	assert.Eq(t, "289dff07669d7a23de0ef88d2f7129e7", strutil.GenMd5(234))

	// short md5
	assert.Eq(t, "ac59075b964b0715", strutil.ShortMd5(123))
	assert.Eq(t, "3cd24fb0d6963f7d", strutil.ShortMd5("abc"))
}

func TestHashPasswd(t *testing.T) {
	key := "ot54c"
	pwd := "abc123456"

	msgMac := strutil.HashPasswd(pwd, key)
	dump.P(msgMac)
	assert.NotEmpty(t, msgMac)
	assert.True(t, strutil.VerifyPasswd(msgMac, pwd, key))
}
