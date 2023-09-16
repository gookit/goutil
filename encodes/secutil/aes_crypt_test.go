package secutil_test

import (
	"testing"

	"github.com/gookit/goutil/byteutil"
	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/encodes/secutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestNewCrypt(t *testing.T) {
	p := secutil.NewAesCrypt()
	p.Config(func(c *secutil.CryptConfig) {
		c.Key = "abcd12345678abcd"
		c.IV = "1234567812345678"
	})

	srcStr := "hi123456"
	encStr, err := p.EncryptString(srcStr)
	assert.NoErr(t, err)
	assert.NotEmpty(t, encStr)
	// "07QaVXqF5Q8Ko9OtE6U9cg==" len=24
	dump.P(encStr)

	decStr, err := p.DecryptString(encStr)
	assert.NoErr(t, err)
	assert.Eq(t, srcStr, decStr)

	srcStr = "a long long long long long ooo ooo ooo ooo ...... string"
	encStr, err = p.EncryptString(srcStr)
	assert.NoErr(t, err)
	assert.NotEmpty(t, encStr)
	// string("fBVcqKLpwMl7J+nk2ncROHwGKlCUk0h8mL+KZr1mxA9qvIfv73Ed3+s/VPCtBJHZchsoA76Pl/MLSCEYpRfnHA=="), #len=88
	dump.P(encStr)

	decStr, err = p.DecryptString(encStr)
	assert.NoErr(t, err)
	assert.Eq(t, srcStr, decStr)
}

func TestNewCrypt_hexEnc(t *testing.T) {
	p := secutil.NewAesCrypt()
	p.Config(func(c *secutil.CryptConfig) {
		c.Key = "abcd12345678abcd"
		c.IV = "1234567812345678"
		c.Encoder = byteutil.HexEncoder
	})

	srcStr := "hi123456"
	encStr, err := p.EncryptString(srcStr)
	assert.NoErr(t, err)
	assert.NotEmpty(t, encStr)
	// "d3b41a557a85e50f0aa3d3ad13a53d72" len=32
	dump.P(encStr)

	decStr, err := p.DecryptString(encStr)
	assert.NoErr(t, err)
	assert.Eq(t, srcStr, decStr)

	srcStr = "a long long long long long ooo ooo ooo ooo ...... string"
	encStr, err = p.EncryptString(srcStr)
	assert.NoErr(t, err)
	assert.NotEmpty(t, encStr)
	// string("7c155ca8a2e9c0c97b27e9e4da7711387c062a509493487c98bf8a66bd66c40f6abc87efef711ddfeb3f54f0ad0491d9721b2803be8f97f30b482118a517e71c"), #len=128
	dump.P(encStr)

	decStr, err = p.DecryptString(encStr)
	assert.NoErr(t, err)
	assert.Eq(t, srcStr, decStr)
}
