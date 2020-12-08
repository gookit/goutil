package strutil_test

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/gookit/goutil/strutil"
	"github.com/stretchr/testify/assert"
)

func TestMd5(t *testing.T) {
	assert.Equal(t, "e10adc3949ba59abbe56e057f20f883e", strutil.Md5("123456"))
	assert.Equal(t, "a906449d5769fa7361d7ecc6aa3f6d28", strutil.GenMd5("123abc"))
	assert.Equal(t, "289dff07669d7a23de0ef88d2f7129e7", strutil.GenMd5(234))
}

func TestRandomChars(t *testing.T) {
	for i := 0; i < 10; i++ {
		str := strutil.RandomChars(4)
		fmt.Println(str)

		assert.Len(t, str, 4)
	}
}

func TestRandomCharsV2(t *testing.T) {
	for i := 0; i < 10; i++ {
		str := strutil.RandomCharsV2(4)
		fmt.Println(str)

		assert.Len(t, str, 4)
	}
}

func TestRandomCharsV3(t *testing.T) {
	for i := 0; i < 10; i++ {
		str := strutil.RandomCharsV3(4)
		fmt.Println(str)

		assert.Len(t, str, 4)
	}
}

func TestRandomBytes(t *testing.T) {
	b, err := strutil.RandomBytes(3)

	// 1607400451937462000
	tsn := time.Now().UnixNano()
	rand.Seed(tsn)

	fmt.Println(tsn)
	fmt.Println(rand.Intn(12))
	fmt.Println(rand.Intn(12))

	fmt.Println(string(b))
	fmt.Println(base64.URLEncoding.EncodeToString(b))
	fmt.Println(base64.StdEncoding.EncodeToString(b))
	fmt.Println(hex.EncodeToString(b))
	assert.NoError(t, err)
}

func TestRandomString(t *testing.T) {
	s, err := strutil.RandomString(3)

	fmt.Println(s)
	assert.NoError(t, err)
	assert.True(t, len(s) > 3)
}
