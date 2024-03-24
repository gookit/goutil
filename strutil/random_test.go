package strutil_test

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestRandomChars(t *testing.T) {
	for i := 0; i < 10; i++ {
		str := strutil.RandomChars(4)
		fmt.Println(str)
		assert.Len(t, str, 4)
	}
}

func TestRandomCharsV2(t *testing.T) {
	keyMp := make(map[string]bool)
	for i := 0; i < 10; i++ {
		str := strutil.RandomCharsV2(6)
		fmt.Println(str)
		assert.Len(t, str, 6)
		keyMp[str] = true
		// time.Sleep(time.Microsecond * 10)
	}

	assert.Len(t, keyMp, 10)
}

// https://github.com/gookit/goutil/issues/121
func TestRandomCharsV2_issues121(t *testing.T) {
	keyMp := make(map[string]bool)

	for i := 0; i < 10; i++ {
		str := strutil.RandomCharsV2(32)
		fmt.Println(str)
		assert.Len(t, str, 32)
		keyMp[str] = true
		// time.Sleep(time.Microsecond * 10)
	}

	assert.Len(t, keyMp, 10)
}

func TestRandomCharsV3(t *testing.T) {
	size := 10
	keyMp := make(map[string]bool)

	for i := 0; i < size; i++ {
		str := strutil.RandomCharsV3(4)
		fmt.Println(str)
		assert.Len(t, str, 4)
		keyMp[str] = true
		// time.Sleep(time.Microsecond * 10)
	}

	assert.Len(t, keyMp, size)
}

func TestRandWithTpl(t *testing.T) {
	size := 10
	keyMp := make(map[string]bool)

	for i := 0; i < size; i++ {
		str := strutil.RandWithTpl(4, strutil.AlphaBet1)
		fmt.Println(str)
		assert.Len(t, str, 4)
		keyMp[str] = true
		// time.Sleep(time.Microsecond * 10)
	}

	assert.NotEmpty(t, strutil.RandWithTpl(8, ""))
	assert.NotEmpty(t, strutil.RandWithTpl(8, strutil.AlphaBet1))

	assert.Len(t, keyMp, size)
}

func TestRandomBytes(t *testing.T) {
	b, err := strutil.RandomBytes(3)

	// 1607400451937462000
	tsn := time.Now().UnixNano()
	rn := rand.New(rand.NewSource(tsn))

	fmt.Println(tsn)
	fmt.Println(rn.Intn(12), rn.Intn(12))

	fmt.Println(string(b))
	fmt.Println(hex.EncodeToString(b))
	assert.NoErr(t, err)
}

func TestRandomString(t *testing.T) {
	s, err := strutil.RandomString(16)

	fmt.Println(s)
	assert.NoErr(t, err)
	assert.True(t, len(s) > 3)
}
